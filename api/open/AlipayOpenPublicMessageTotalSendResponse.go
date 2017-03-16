package open

import (
  "github.com/cvincent00/antsdk/api"
)

type AlipayOpenPublicMessageTotalSendResponse struct {
  api.AlipayResponse
  MessageId string `json:"message_id"`  // 消息ID
}
