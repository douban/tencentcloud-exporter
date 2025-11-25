package instance

import (
	"fmt"

	"reflect"

	sdk "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/teo/v20220901"
)

type EdgeoneTcInstance struct {
	baseTcInstance
	meta *sdk.AccelerationDomain
}

func (ins *EdgeoneTcInstance) GetMeta() interface{} {
	return ins.meta
}

func NewEdgeoneTcInstance(instanceId string, meta *sdk.AccelerationDomain) (ins *EdgeoneTcInstance, err error) {
	if instanceId == "" {
		return nil, fmt.Errorf("instanceId is empty ")
	}
	if meta == nil {
		return nil, fmt.Errorf("meta is empty ")
	}
	ins = &EdgeoneTcInstance{
		baseTcInstance: baseTcInstance{
			instanceId: instanceId,
			value:      reflect.ValueOf(*meta),
		},
		meta: meta,
	}
	return
}
