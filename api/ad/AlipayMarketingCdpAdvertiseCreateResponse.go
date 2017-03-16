package ad

import (
  "github.com/cvincent00/antsdk/api"
)

type AlipayMarketingCdpAdvertiseCreateResponse struct {
  api.AlipayResponse
  AdId string `json:"ad_id"`  // 创建广告唯一标识
}
