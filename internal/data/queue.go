package data

import (
	"context"
	"github.com/minicloudsky/lianjia/internal/biz"
)

type Queue interface {
	Send(ctx context.Context, msg []biz.Message, houseType biz.HoueseType) error
	Receive(ctx context.Context) error
}

func (repo lianjiaRepo) Send(ctx context.Context, msg []biz.Message, houseType biz.HoueseType) error {
	return repo.data.queue.Send(ctx, msg, houseType)
}

func (repo lianjiaRepo) Receive(ctx context.Context) error {
	err := repo.data.queue.Receive(ctx)
	if err != nil {
		return err
	}
	return nil
}
