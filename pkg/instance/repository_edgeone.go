package instance

import (
	"github.com/go-kit/log"
	apiCommon "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	sdk "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/teo/v20220901"

	"github.com/tencentyun/tencentcloud-exporter/pkg/client"
	"github.com/tencentyun/tencentcloud-exporter/pkg/config"
)

func init() {
	registerRepository("QCE/EDGEONE_L7", NewEdgeoneTcInstanceRepository)
}

type EdgeoneTcInstanceRepository struct {
	client *sdk.Client
	logger log.Logger
}

func (repo *EdgeoneTcInstanceRepository) GetInstanceKey() string {
	return "InstanceId"
}

func (repo *EdgeoneTcInstanceRepository) Get(id string) (instance TcInstance, err error) {
	return
}

func (repo *EdgeoneTcInstanceRepository) ListByIds(id []string) (instances []TcInstance, err error) {
	return
}

func (repo *EdgeoneTcInstanceRepository) ListByFilters(filters map[string]string) (instances []TcInstance, err error) {
	return
}

func NewEdgeoneTcInstanceRepository(cred apiCommon.CredentialIface, c *config.TencentConfig, logger log.Logger) (repo TcInstanceRepository, err error) {

	cli, err := client.NewEdgeOneClient(cred, c)
	if err != nil {
		return
	}
	repo = &EdgeoneTcInstanceRepository{
		client: cli,
		logger: logger,
	}
	return
}
