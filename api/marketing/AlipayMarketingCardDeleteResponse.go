package marketing

import (
  "github.com/cvincent00/antsdk/api"
)

type AlipayMarketingCardDeleteResponse struct {
  api.AlipayResponse
  BizSerialNo string `json:"biz_serial_no"` // 支付宝端删卡业务流水号
}
