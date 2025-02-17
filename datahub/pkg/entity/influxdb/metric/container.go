package metric

import (
	"github.com/containers-ai/alameda/datahub/pkg/utils"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb"
	"strconv"
	"time"
)

const (
	ContainerTime         influxdb.Tag = "time"
	ContainerPodNamespace influxdb.Tag = "pod_namespace"
	ContainerPodName      influxdb.Tag = "pod_name"
	ContainerName         influxdb.Tag = "name"
	ContainerRateRange    influxdb.Tag = "rate_range"

	ContainerValue influxdb.Field = "value"
)

var (
	ContainerTags    = []influxdb.Tag{ContainerPodNamespace, ContainerPodName, ContainerName}
	ContainerFields  = []influxdb.Field{ContainerValue}
	ContainerColumns = []string{string(ContainerPodNamespace), string(ContainerPodName), string(ContainerName), string(ContainerValue)}
)

type ContainerEntity struct {
	Time      time.Time
	Namespace *string
	PodName   *string
	Name      *string
	RateRange *string

	Value *float64
}

func NewContainerEntityFromMap(data map[string]string) ContainerEntity {
	tempTimestamp, _ := utils.ParseTime(data[string(ContainerTime)])

	entity := ContainerEntity{
		Time: tempTimestamp,
	}

	// InfluxDB tags
	if valueStr, exist := data[string(ContainerPodNamespace)]; exist {
		entity.Namespace = &valueStr
	}
	if valueStr, exist := data[string(ContainerPodName)]; exist {
		entity.PodName = &valueStr
	}
	if valueStr, exist := data[string(ContainerName)]; exist {
		entity.Name = &valueStr
	}
	if valueStr, exist := data[string(ContainerRateRange)]; exist {
		entity.RateRange = &valueStr
	}

	// InfluxDB fields
	if valueFloat, exist := data[string(ContainerValue)]; exist {
		value, _ := strconv.ParseFloat(valueFloat, 64)
		entity.Value = &value
	}

	return entity
}
