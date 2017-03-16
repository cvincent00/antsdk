package marketing

import (
  "github.com/cvincent00/antsdk/api"
)

type KoubeiMarketingCampaignActivityQueryResponse struct {
  api.AlipayResponse
  CampDetail CampDetail `json:"camp_detail"`  // 活动详情
}
