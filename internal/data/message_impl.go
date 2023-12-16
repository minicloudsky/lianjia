package data

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

// KafkaService 接口定义
type KafkaService interface {
	SendMessage(topic, key string, value []byte) error
	ReceiveMessages(ctx context.Context, topic, groupID string, handler func(key, value []byte) error)
}

// KafkaServiceImpl 结构体实现 KafkaService 接口
type KafkaServiceImpl struct {
	brokers   []string
	kafkaConn *kafka.Conn
}

// NewKafkaService 创建 KafkaService 的实例
func NewKafkaService(brokers []string) KafkaService {
	return &KafkaServiceImpl{brokers: brokers}
}

// SendMessage 实现 KafkaService 接口的 SendMessage 方法
func (ks *KafkaServiceImpl) SendMessage(topic, key string, value []byte) error {
	_, err := ks.kafkaConn.WriteMessages(kafka.Message{
		Key:   []byte(key),
		Value: value,
	})
	if err != nil {
		return err
	}
	return nil
}

// ReceiveMessages 实现 KafkaService 接口的 ReceiveMessages 方法
func (ks *KafkaServiceImpl) ReceiveMessages(ctx context.Context, topic, groupID string, handler func(key, value []byte) error) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  ks.brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			continue
		}

		if err := handler(m.Key, m.Value); err != nil {
			fmt.Printf("Error handling message: %v\n", err)
			continue
		}
	}
}
