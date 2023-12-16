package biz

import (
	"context"
	"github.com/segmentio/kafka-go"
)

func (uc *LianjiaUsecase) HandleZufangMessage(ctx context.Context, msg []kafka.Message) error {
	return nil
}

func (uc *LianjiaUsecase) ListCityZufang(ctx context.Context, city *CityInfo) error {
	return nil
}
