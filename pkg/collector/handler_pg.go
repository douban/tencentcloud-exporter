package collector

import (
	"github.com/go-kit/log"
	apiCommon "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
)

const (
	PGNamespace       = "QCE/POSTGRES"
	PGDBInstanceIDKey = "resourceId"
)

func init() {
	registerHandler(PGNamespace, defaultHandlerEnabled, NewPGHandler)
}

type pgHandler struct {
	baseProductHandler
}

func (h *pgHandler) GetNamespace() string {
	return MariaDBNamespace
}
func NewPGHandler(cred apiCommon.CredentialIface, c *TcProductCollector, logger log.Logger) (handler ProductHandler, err error) {
	handler = &pgHandler{
		baseProductHandler{
			monitorQueryKey: PGDBInstanceIDKey,
			collector:       c,
			logger:          logger,
		},
	}
	return

}
