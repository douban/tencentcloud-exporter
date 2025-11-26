package collector

import (
	"fmt"

	apiCommon "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/tencentyun/tencentcloud-exporter/pkg/metric"
	"github.com/tencentyun/tencentcloud-exporter/pkg/util"
)

const (
	EdgeOneL7Namespace   = "QCE/EDGEONE_L7"
	EdgeOneInstanceidKey = "domain"
)

func init() {
	registerHandler(EdgeOneL7Namespace, defaultHandlerEnabled, NewEdgeOneHandler)
}

type edgeoneHandler struct {
	baseProductHandler
}

func (h *edgeoneHandler) IsMetricMetaValid(meta *metric.TcmMeta) bool {
	return true
}

func (h *edgeoneHandler) GetNamespace() string {
	return EdgeOneL7Namespace
}

func (h *edgeoneHandler) IsMetricValid(m *metric.TcmMetric) bool {
	_, ok := excludeMetricName[m.Meta.MetricName]
	if ok {
		return false
	}
	p, err := m.Meta.GetPeriod(m.Conf.StatPeriodSeconds)
	if err != nil {
		return false
	}
	if p != m.Conf.StatPeriodSeconds {
		return false
	}
	return true
}

func (h *edgeoneHandler) GetSeries(m *metric.TcmMetric) (slist []*metric.TcmSeries, err error) {
	if m.Conf.IsIncludeOnlyInstance() {
		return h.GetSeriesByOnly(m)
	}

	if m.Conf.IsIncludeAllInstance() {
		return h.GetSeriesByAll(m)
	}

	if m.Conf.IsCustomQueryDimensions() {
		return h.GetSeriesByCustom(m)
	}

	return nil, fmt.Errorf("must config all_instances or only_include_instances or custom_query_dimensions")
}

func (h *edgeoneHandler) GetSeriesByOnly(m *metric.TcmMetric) ([]*metric.TcmSeries, error) {
	var slist []*metric.TcmSeries
	for _, insId := range m.Conf.OnlyIncludeInstances {
		ins, err := h.collector.InstanceRepo.Get(insId)
		if err != nil {
			level.Error(h.logger).Log("msg", "Instance not found", "id", insId)
			continue
		}
		zoneId, err := ins.GetFieldValueByName("ZoneId")
		if err != nil {
			level.Error(h.logger).Log("msg", "ZoneId not found")
			continue
		}
		domain, err := ins.GetFieldValueByName("Domain")
		if err != nil {
			level.Error(h.logger).Log("msg", "domain not found")
			continue
		}
		ql := map[string]string{
			"domain": domain,
			"zoneid": zoneId,
		}
		s, err := metric.NewTcmSeries(m, ql, ins)
		if err != nil {
			level.Error(h.logger).Log("msg", "Create metric series fail",
				"metric", m.Meta.MetricName, "instance", insId)
			continue
		}
		slist = append(slist, s)
	}
	return slist, nil
}

func (h *edgeoneHandler) GetSeriesByAll(m *metric.TcmMetric) ([]*metric.TcmSeries, error) {
	var slist []*metric.TcmSeries
	insList, err := h.collector.InstanceRepo.ListByFilters(m.Conf.InstanceFilters)
	if err != nil {
		return nil, err
	}
	for _, ins := range insList {
		if len(m.Conf.ExcludeInstances) != 0 && util.IsStrInList(m.Conf.ExcludeInstances, ins.GetInstanceId()) {
			continue
		}
		zoneId, err := ins.GetFieldValueByName("ZoneId")
		domain, err := ins.GetFieldValueByName("Domain")
		if err != nil {
			level.Error(h.logger).Log("msg", "zoneId or domain not found")
			continue
		}
		ql := map[string]string{
			"domain": domain,
			"zoneId": zoneId,
		}
		s, err := metric.NewTcmSeries(m, ql, ins)
		if err != nil {
			level.Error(h.logger).Log("msg", "Create metric series fail",
				"metric", m.Meta.MetricName, "instance", ins.GetInstanceId())
			continue
		}
		slist = append(slist, s)
	}
	return slist, nil
}

func (h *edgeoneHandler) GetSeriesByCustom(m *metric.TcmMetric) ([]*metric.TcmSeries, error) {
	var slist []*metric.TcmSeries
	for _, ql := range m.Conf.CustomQueryDimensions {
		if !h.checkMonitorQueryKeys(m, ql) {
			continue
		}

		s, err := metric.NewTcmSeries(m, ql, nil)
		if err != nil {
			level.Error(h.logger).Log("msg", "Create metric series fail", "metric", m.Meta.MetricName,
				"ql", fmt.Sprintf("%v", ql))
			continue
		}
		slist = append(slist, s)
	}
	return slist, nil
}

func (h *edgeoneHandler) checkMonitorQueryKeys(m *metric.TcmMetric, ql map[string]string) bool {
	for k := range ql {
		if !util.IsStrInList(m.Meta.SupportDimensions, k) {
			level.Error(h.logger).Log("msg", fmt.Sprintf("not found %s in supportQueryDimensions", k),
				"ql", fmt.Sprintf("%v", ql))
			return false
		}
	}
	return true
}

func NewEdgeOneHandler(cred apiCommon.CredentialIface, c *TcProductCollector, logger log.Logger) (handler ProductHandler, err error) {
	handler = &edgeoneHandler{
		baseProductHandler{
			collector: c,
			logger:    logger,
		},
	}
	return
}
