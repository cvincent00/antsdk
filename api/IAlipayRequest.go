package api

import (
  "github.com/cvincent00/antsdk/utils"
)

type IAlipayRequest interface {
  GetApiMethodName() string
  GetApiVersion() string
  GetTerminalType() string
  GetTerminalInfo() string
  GetNotifyUrl() string
  GetReturnUrl() string
  GetProdCode() string
  GetTextParams() *utils.AlipayHashMap
  IsNeedEncrypt() bool
}
