package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/minicloudsky/lianjia/internal/biz"
	"github.com/segmentio/kafka-go"
	"time"
)

type KafkaQueue struct {
	km     *biz.KafkaManager
	logger log.Logger
}

func (kq *KafkaQueue) Send(ctx context.Context, msgs []biz.Message, houseType biz.HoueseType) error {
	helper := log.NewHelper(kq.logger)
	kafkaMessages := make([]kafka.Message, 0)
	for _, m := range msgs {
		kafkaMessages = append(kafkaMessages, kafka.Message{
			Topic:     houseType.Topic(),
			Partition: 1,
			Key:       []byte(houseType),
			Value:     m.Content,
			Time:      time.Now(),
		})
	}
	if sender, ok := kq.km.Senders[string(houseType)]; ok && sender.Writer != nil {
		err := sender.Writer.WriteMessages(ctx, kafkaMessages...)
		if err != nil {
			helper.Errorf("write message err! err: %v", err)
			return err
		}
	}
	return nil
}

func (kq *KafkaQueue) Receive(ctx context.Context) error {
	helper := log.NewHelper(kq.logger)
	helper.Infof("kafka start receiving messages...")
	batchSize := 100
	for topic, reader := range kq.km.Readers {
		r := reader.Reader
		go func(topic string) {
			for {
				kafkaMessages := make([]kafka.Message, 0)
				messages := make([]biz.Message, 0, batchSize)
				for i := 0; i < batchSize; i++ {
					m, err := r.ReadMessage(ctx)
					if err != nil {
						helper.Errorf("ReadMessage err! err: %v", err)
					}
					messages = append(messages, biz.Message{Content: m.Value})
					kafkaMessages = append(kafkaMessages, m)
				}
				switch topic {
				case string(biz.TopicErshoufang):
					biz.ErshoufangProcessChan <- messages
				case string(biz.TopicLoupan):
					biz.LoupanProcessChan <- messages
				case string(biz.TopicCommercial):
					biz.CommercialProcessChan <- messages
				case string(biz.TopicZufang):
					biz.ZufangProcessChan <- messages
				default:
					helper.Errorf("unknown message type! %s", topic)
				}
				// batch acknowledge kafka messages
				if err := r.CommitMessages(context.Background(), kafkaMessages...); err != nil {
					helper.Errorf("Error committing messages: %v", err)
				}
			}
		}(topic)
	}

	return nil
}
