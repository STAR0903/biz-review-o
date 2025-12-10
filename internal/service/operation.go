package service

import (
	"context"
	v1 "review-o/api/operation/v1"
	"review-o/internal/biz"
)

// OperationService is a greeter service.
type OperationService struct {
	v1.UnimplementedOperationServer
	uc *biz.OperationUsecase
}

// NewOperationService
func NewOperationService(uc *biz.OperationUsecase) *OperationService {
	return &OperationService{uc: uc}
}

// AuditReview 审核评价
func (s *OperationService) AuditReview(ctx context.Context, req *v1.AuditReviewRequest) (*v1.AuditReviewReply, error) {
	param := &biz.AuditReviewParam{
		ReviewID: req.ReviewID,
		Status:   req.Status,
		OpUser:   req.OpUser,
		OpReason: req.OpReason,
	}
	reviewID, status, err := s.uc.AuditReview(ctx, param)
	if err != nil {
		return nil, err
	}
	return &v1.AuditReviewReply{
		ReviewID: reviewID,
		Status:   status,
	}, nil
}

// AuditAppeal 审核申诉
func (s *OperationService) AuditAppeal(ctx context.Context, req *v1.AuditAppealRequest) (*v1.AuditAppealReply, error) {
	param := &biz.AuditAppealParam{
		AppealID:  req.AppealID,
		Status:    req.Status,
		OpUser:    req.OpUser,
		OpRemarks: req.OpRemarks,
	}
	appealID, status, err := s.uc.AuditAppeal(ctx, param)
	if err != nil {
		return nil, err
	}
	return &v1.AuditAppealReply{
		AppealID: appealID,
		Status:   status,
	}, nil
}
