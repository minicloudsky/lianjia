package service

import (
	v1 "github.com/minicloudsky/lianjia/api/lianjia/v1"
	"github.com/minicloudsky/lianjia/internal/biz"
)

// Service is a lianjia service.
type Service struct {
	v1.UnimplementedLianjiaServer

	uc   *biz.LianjiaUsecase
	repo biz.LianjiaRepo
}

func NewService(uc *biz.LianjiaUsecase, repo biz.LianjiaRepo) *Service {
	return &Service{uc: uc, repo: repo}
}
