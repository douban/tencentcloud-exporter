package collector

import (
	"strings"

	"github.com/go-kit/log"
	apiCommon "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"

	"github.com/tencentyun/tencentcloud-exporter/pkg/metric"
	"github.com/tencentyun/tencentcloud-exporter/pkg/util"
)

const (
	DcxNamespace     = "QCE/DCX"
	DcxInstanceidKey = "directConnectConnId"
)

var (
	DcxInvalidMetricNames = []string{"rxbytes", "txbytes"}
)

func init() {
	registerHandler(DcxNamespace, defaultHandlerEnabled, NewDcxHandler)
}

type dcxHandler struct {
	baseProductHandler
}

func (h *dcxHandler) IsMetricMetaValid(meta *metric.TcmMeta) bool {
	if util.IsStrInList(DcxInvalidMetricNames, strings.ToLower(meta.MetricName)) {
		return false
	}
	return true
}

func (h *dcxHandler) GetNamespace() string {
	return DcxNamespace
}

func (h *dcxHandler) IsMetricValid(m *metric.TcmMetric) bool {
	return true
}

func NewDcxHandler(cred apiCommon.CredentialIface, c *TcProductCollector, logger log.Logger) (handler ProductHandler, err error) {
	handler = &dcxHandler{
		baseProductHandler{
			monitorQueryKey: DcxInstanceidKey,
			collector:       c,
			logger:          logger,
		},
	}
	return

}
