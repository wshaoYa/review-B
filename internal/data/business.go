package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "review-b/api/review/v1"
	"review-b/internal/biz"
)

type businessRepo struct {
	data *Data
	log  *log.Helper
}

// NewBusinessRepo . 构造函数
func NewBusinessRepo(data *Data, logger log.Logger) biz.BusinessRepo {
	return &businessRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateReply 商家回复评价
func (b businessRepo) CreateReply(ctx context.Context, param *biz.ReplyParam) (int64, error) {
	reply, err := b.data.rc.ReplyReview(ctx, &v1.ReplyReviewRequest{
		ReviewID:  param.ReviewID,
		StoreID:   param.StoreID,
		Content:   param.Content,
		PicInfo:   param.PicInfo,
		VideoInfo: param.VideoInfo,
	})
	if err != nil {
		b.log.Errorf("data business CreateReply failed, err:%v \n", err)
		return 0, err
	}
	return reply.ReplyID, nil
}

func (b businessRepo) CreateAppeal(ctx context.Context, param *biz.AppealParam) (int64, error) {
	return 0, nil
}
