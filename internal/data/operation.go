package data

import (
	"context"
	v1 "review-o/api/review/v1"
	"review-o/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type operationRepo struct {
	data *Data
	log  *log.Helper
}

// NewOperationRepo .
func NewOperationRepo(data *Data, logger log.Logger) biz.OperationRepo {
	return &operationRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// AuditReview 审核评论
func (r *operationRepo) AuditReview(ctx context.Context, param *biz.AuditReviewParam) (int64, int32, error) {
	reply, err := r.data.rc.AuditReview(ctx, &v1.AuditReviewRequest{
		ReviewID: param.ReviewID,
		Status:   param.Status,
		OpUser:   param.OpUser,
		OpReason: param.OpReason,
	})
	if err != nil {
		return 0, 0, err
	}
	return reply.ReviewID, reply.Status, nil
}

// AuditAppeal 审核申诉
func (r *operationRepo) AuditAppeal(ctx context.Context, param *biz.AuditAppealParam) (int64, int32, error) {
	reply, err := r.data.rc.AuditAppeal(ctx, &v1.AuditAppealRequest{
		AppealID:  param.AppealID,
		Status:    param.Status,
		OpUser:    param.OpUser,
		OpRemarks: param.OpRemarks,
	})
	if err != nil {
		return 0, 0, err
	}
	return reply.AppealID, reply.Status, nil
}
