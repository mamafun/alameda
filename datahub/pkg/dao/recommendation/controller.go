package recommendation

import (
	datahub_v1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
)

// ContainerOperation defines container measurement operation of recommendation database
type ControllerOperation interface {
	AddControllerRecommendations([]*datahub_v1alpha1.ControllerRecommendation) error
	//	ListPodRecommendations(podNamespacedName *datahub_v1alpha1.NamespacedName,
	//		queryCondition *datahub_v1alpha1.QueryCondition,
	//		kind datahub_v1alpha1.Kind) ([]*datahub_v1alpha1.PodRecommendation, error)
}