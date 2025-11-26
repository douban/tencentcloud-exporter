package collector

import (
	"github.com/go-kit/log"
	apiCommon "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"

	"github.com/tencentyun/tencentcloud-exporter/pkg/metric"
)

const (
	SqlServerNamespace     = "QCE/SQLSERVER"
	SqlServerInstanceidKey = "resourceId"
)

func init() {
	registerHandler(SqlServerNamespace, defaultHandlerEnabled, NewSqlServerHandler)
}

type sqlServerHandler struct {
	baseProductHandler
}

func (h *sqlServerHandler) IsMetricMetaValid(meta *metric.TcmMeta) bool {
	return true
}

func (h *sqlServerHandler) GetNamespace() string {
	return SqlServerNamespace
}

func (h *sqlServerHandler) IsMetricValid(m *metric.TcmMetric) bool {
	return true
}

func NewSqlServerHandler(cred apiCommon.CredentialIface, c *TcProductCollector, logger log.Logger) (handler ProductHandler, err error) {
	handler = &sqlServerHandler{
		baseProductHandler{
			monitorQueryKey: SqlServerInstanceidKey,
			collector:       c,
			logger:          logger,
		},
	}
	return
}
