package biz

// AuditReviewParam 审核评价参数
type AuditReviewParam struct {
	ReviewID int64
	Status   int32
	OpUser   string
	OpReason string
}

// AuditAppealParam 审核申诉参数
type AuditAppealParam struct {
	AppealID  int64
	Status    int32
	OpUser    string
	OpRemarks *string
}
