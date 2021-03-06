package marketing

import (
  "github.com/cvincent00/antsdk/api"
)

type KoubeiMarketingCampaignActivityModifyResponse struct {
  api.AlipayResponse
  CampStatus  string `json:"camp_status"`   // 活动状态
  AuditStatus string `json:"audit_status"`  // 活动子状态，如审核中
}
