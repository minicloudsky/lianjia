package data

import (
	"context"
	"github.com/minicloudsky/lianjia/internal/biz"
)

func (repo lianjiaRepo) GetKafkaManager(ctx context.Context) (km *biz.KafkaManager, err error) {
	return repo.data.km, err
}
