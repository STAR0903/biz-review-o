package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// OperationRepo 是操作仓库接口
type OperationRepo interface {
	// AuditReview 审核评论
	AuditReview(ctx context.Context, param *AuditReviewParam) (int64, int32, error)
	// AuditAppeal 审核申诉
	AuditAppeal(ctx context.Context, param *AuditAppealParam) (int64, int32, error)
}

// OperationUsecase 是操作用例
type OperationUsecase struct {
	repo OperationRepo
	log  *log.Helper
}

// NewOperationUsecase new a Operation usecase.
func NewOperationUsecase(repo OperationRepo, logger log.Logger) *OperationUsecase {
	return &OperationUsecase{repo: repo, log: log.NewHelper(logger)}
}

// AuditReview 审核评论
func (uc *OperationUsecase) AuditReview(ctx context.Context, param *AuditReviewParam) (int64, int32, error) {
	uc.log.WithContext(ctx).Debugf("AuditReview: %v", param)
	return uc.repo.AuditReview(ctx, param)

}

// AuditAppeal 审核申诉
func (uc *OperationUsecase) AuditAppeal(ctx context.Context, param *AuditAppealParam) (int64, int32, error) {
	uc.log.WithContext(ctx).Debugf("AuditAppeal: %v", param)
	return uc.repo.AuditAppeal(ctx, param)
}
