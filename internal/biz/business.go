package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// BusinessRepo 商家repo
type BusinessRepo interface {
	CreateReply(context.Context, *ReplyParam) (int64, error)   // B端 回复评价
	CreateAppeal(context.Context, *AppealParam) (int64, error) // B端 申诉评价
}

// BusinessUsecase 商家usecase
type BusinessUsecase struct {
	repo BusinessRepo
	log  *log.Helper
}

// NewBusinessUsecase 构造函数
func NewBusinessUsecase(repo BusinessRepo, logger log.Logger) *BusinessUsecase {
	return &BusinessUsecase{repo: repo, log: log.NewHelper(logger)}
}

// ReplyReview 回复评价
func (uc *BusinessUsecase) ReplyReview(ctx context.Context, param *ReplyParam) (int64, error) {
	//log
	uc.log.WithContext(ctx).Infof("ReplyReview - reviewID: %v", param.ReviewID)
	return uc.repo.CreateReply(ctx, param)
}

// AppealReview 申诉评价
func (uc *BusinessUsecase) AppealReview(ctx context.Context, param *AppealParam) (int64, error) {
	//log
	uc.log.WithContext(ctx).Infof("AppealReview - reviewID: %v", param.ReviewID)
	return uc.repo.CreateAppeal(ctx, param)
}
