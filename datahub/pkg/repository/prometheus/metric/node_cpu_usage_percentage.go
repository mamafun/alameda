package metric

import (
	"fmt"
	EntityPromthMetric "github.com/containers-ai/alameda/datahub/pkg/entity/prometheus/metric"
	DBCommon "github.com/containers-ai/alameda/internal/pkg/database/common"
	InternalPromth "github.com/containers-ai/alameda/internal/pkg/database/prometheus"
	"github.com/pkg/errors"
	"time"
)

// NodeCPUUsagePercentageRepository Repository to access metric node:node_cpu_utilisation:avg1m from prometheus
type NodeCpuUsagePercentageRepository struct {
	PrometheusConfig InternalPromth.Config
}

// NewNodeCPUUsagePercentageRepositoryWithConfig New node cpu usage percentage repository with prometheus configuration
func NewNodeCpuUsagePercentageRepositoryWithConfig(cfg InternalPromth.Config) NodeCpuUsagePercentageRepository {
	return NodeCpuUsagePercentageRepository{PrometheusConfig: cfg}
}

// ListMetricsByPodNamespacedName Provide metrics from response of querying request contain namespace, pod_name and default labels
func (n NodeCpuUsagePercentageRepository) ListMetricsByNodeName(nodeName string, options ...DBCommon.Option) ([]InternalPromth.Entity, error) {

	var (
		err error

		prometheusClient *InternalPromth.Prometheus

		queryLabelsString string

		queryExpressionSum string
		queryExpressionAvg string

		response InternalPromth.Response

		entities []InternalPromth.Entity
	)

	prometheusClient, err = InternalPromth.NewClient(&n.PrometheusConfig)
	if err != nil {
		return entities, errors.Wrap(err, "list node cpu usage metrics by node name failed")
	}

	opt := DBCommon.NewDefaultOptions()
	for _, option := range options {
		option(&opt)
	}

	//metricName = EntityPromthNodeCpu.MetricName
	metricNameSum := EntityPromthMetric.NodeCpuUsagePercentageMetricNameSum
	metricNameAvg := EntityPromthMetric.NodeCpuUsagePercentageMetricNameAvg

	queryLabelsString = n.buildQueryLabelsStringByNodeName(nodeName)

	if queryLabelsString != "" {
		queryExpressionSum = fmt.Sprintf("%s{%s}", metricNameSum, queryLabelsString)
		queryExpressionAvg = fmt.Sprintf("%s{%s}", metricNameAvg, queryLabelsString)
	} else {
		queryExpressionSum = fmt.Sprintf("%s", metricNameSum)
		queryExpressionAvg = fmt.Sprintf("%s", metricNameAvg)
	}

	stepTimeInSeconds := int64(opt.StepTime.Nanoseconds() / int64(time.Second))
	queryExpressionSum, err = InternalPromth.WrapQueryExpression(queryExpressionSum, opt.AggregateOverTimeFunc, stepTimeInSeconds)
	if err != nil {
		return entities, errors.Wrap(err, "list node cpu usage metrics by node name failed")
	}
	queryExpressionAvg, err = InternalPromth.WrapQueryExpression(queryExpressionAvg, opt.AggregateOverTimeFunc, stepTimeInSeconds)
	if err != nil {
		return entities, errors.Wrap(err, "list node cpu usage metrics by node name failed")
	}

	queryExpression := fmt.Sprintf("1000 * %s * %s", queryExpressionSum, queryExpressionAvg)

	response, err = prometheusClient.QueryRange(queryExpression, opt.StartTime, opt.EndTime, opt.StepTime)
	if err != nil {
		return entities, errors.Wrap(err, "list node cpu usage metrics by node name failed")
	} else if response.Status != InternalPromth.StatusSuccess {
		return entities, errors.Errorf("list node cpu usage metrics by node name failed: receive error response from prometheus: %s", response.Error)
	}

	entities, err = response.GetEntities()
	if err != nil {
		return entities, errors.Wrap(err, "list node cpu usage metrics by node name failed")
	}

	return entities, nil
}

func (n NodeCpuUsagePercentageRepository) buildQueryLabelsStringByNodeName(nodeName string) string {

	var (
		queryLabelsString = ""
	)

	if nodeName != "" {
		queryLabelsString += fmt.Sprintf(`%s = "%s"`, EntityPromthMetric.NodeCpuUsagePercentageLabelNode, nodeName)
	}

	return queryLabelsString
}
