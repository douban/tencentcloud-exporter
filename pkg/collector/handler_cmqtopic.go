package collector

import (
	"github.com/go-kit/log"
	apiCommon "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
)

const (
	CMQTopicNamespace     = "QCE/CMQTOPIC"
	CMQTopicInstanceIDKey = "topicId"
)

func init() {
	registerHandler(CMQTopicNamespace, defaultHandlerEnabled, NewCMQTopicHandler)
}

type cmqTopicHandler struct {
	baseProductHandler
}

func (h *cmqTopicHandler) GetNamespace() string {
	return CMQTopicNamespace
}
func NewCMQTopicHandler(cred apiCommon.CredentialIface, c *TcProductCollector, logger log.Logger) (handler ProductHandler, err error) {
	handler = &cmqTopicHandler{
		baseProductHandler{
			monitorQueryKey: CMQTopicInstanceIDKey,
			collector:       c,
			logger:          logger,
		},
	}
	return

}
