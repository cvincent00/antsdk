package open

import (
  "github.com/cvincent00/antsdk/api"
)

type AlipayOpenPublicAccountQueryResponse struct {
  api.AlipayResponse
  PublicBindAccounts []StdPublicBindAccount `json:"public_bind_accounts"` // 绑定账户列表
}
