package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/segmentio/kafka-go"
	"sync"
)

type KafkaSender struct {
	Writer *kafka.Writer
}

type KafkaReader struct {
	Reader *kafka.Reader
}

type KafkaManager struct {
	Brokers   map[string][]string
	Senders   map[string]*KafkaSender
	Readers   map[string]*KafkaReader
	SendersMu sync.Mutex
	Logger    log.Logger
}

func (km *KafkaManager) Close() error {
	km.SendersMu.Lock()
	defer km.SendersMu.Unlock()
	kLogger := log.NewHelper(km.Logger)
	for _, sender := range km.Senders {
		err := sender.Writer.Close()
		if err != nil {
			kLogger.Errorf("fail to close kafka writer ! err: %v", err)
			return err
		}
	}
	return nil
}
