package v1alpha1

import (
	DaoRecommendation "github.com/containers-ai/alameda/datahub/pkg/dao/recommendation"
	DaoRecommendationImpl "github.com/containers-ai/alameda/datahub/pkg/dao/recommendation/impl"
	AutoScalingV1alpha1 "github.com/containers-ai/alameda/operator/pkg/apis/autoscaling/v1alpha1"
	ReconcilerAlamedaRecommendation "github.com/containers-ai/alameda/operator/pkg/reconciler/alamedarecommendation"
	AlamedaUtils "github.com/containers-ai/alameda/pkg/utils"
	DatahubV1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
	K8sErrors "k8s.io/apimachinery/pkg/api/errors"
	K8sTypes "k8s.io/apimachinery/pkg/types"
)

// CreatePodRecommendations add pod recommendations information to database
func (s *ServiceV1alpha1) CreatePodRecommendations(ctx context.Context, in *DatahubV1alpha1.CreatePodRecommendationsRequest) (*status.Status, error) {
	scope.Debug("Request received from CreatePodRecommendations grpc function: " + AlamedaUtils.InterfaceToString(in))

	var containerDAO DaoRecommendation.ContainerOperation = &DaoRecommendationImpl.Container{
		InfluxDBConfig: *s.Config.InfluxDB,
	}

	podRecommendations := in.GetPodRecommendations()
	for _, podRecommendation := range podRecommendations {
		podNS := podRecommendation.GetNamespacedName().Namespace
		podName := podRecommendation.GetNamespacedName().Name
		alamedaRecommendation := &AutoScalingV1alpha1.AlamedaRecommendation{}

		if err := s.K8SClient.Get(context.TODO(), K8sTypes.NamespacedName{
			Namespace: podNS,
			Name:      podName,
		}, alamedaRecommendation); err == nil {
			alamedarecommendationReconciler := ReconcilerAlamedaRecommendation.NewReconciler(s.K8SClient, alamedaRecommendation)
			if alamedaRecommendation, err = alamedarecommendationReconciler.UpdateResourceRecommendation(podRecommendation); err == nil {
				if err = s.K8SClient.Update(context.TODO(), alamedaRecommendation); err != nil {
					scope.Error(err.Error())
				}
			}
		} else if !K8sErrors.IsNotFound(err) {
			scope.Error(err.Error())
		}
	}

	if err := containerDAO.AddPodRecommendations(in); err != nil {
		scope.Error(err.Error())
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}, err
	}

	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}

// CreatePodRecommendations add controller recommendations information to database
func (s *ServiceV1alpha1) CreateControllerRecommendations(ctx context.Context, in *DatahubV1alpha1.CreateControllerRecommendationsRequest) (*status.Status, error) {
	scope.Debug("Request received from CreateControllerRecommendations grpc function: " + AlamedaUtils.InterfaceToString(in))

	controllerDAO := DaoRecommendationImpl.Controller{
		InfluxDBConfig: *s.Config.InfluxDB,
	}

	controllerRecommendationList := in.GetControllerRecommendations()
	err := controllerDAO.AddControllerRecommendations(controllerRecommendationList)

	if err != nil {
		scope.Error(err.Error())
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}, err
	}

	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}

// ListPodRecommendations list pod recommendations
func (s *ServiceV1alpha1) ListPodRecommendations(ctx context.Context, in *DatahubV1alpha1.ListPodRecommendationsRequest) (*DatahubV1alpha1.ListPodRecommendationsResponse, error) {
	scope.Debug("Request received from ListPodRecommendations grpc function: " + AlamedaUtils.InterfaceToString(in))

	var containerDAO DaoRecommendation.ContainerOperation = &DaoRecommendationImpl.Container{
		InfluxDBConfig: *s.Config.InfluxDB,
	}

	podRecommendations, err := containerDAO.ListPodRecommendations(in)
	if err != nil {
		scope.Error(err.Error())
		return &DatahubV1alpha1.ListPodRecommendationsResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INTERNAL),
				Message: err.Error(),
			},
		}, nil
	}

	res := &DatahubV1alpha1.ListPodRecommendationsResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		PodRecommendations: podRecommendations,
	}
	scope.Debug("Response sent from ListPodRecommendations grpc function: " + AlamedaUtils.InterfaceToString(res))
	return res, nil
}

// ListAvailablePodRecommendations list pod recommendations
func (s *ServiceV1alpha1) ListAvailablePodRecommendations(ctx context.Context, in *DatahubV1alpha1.ListPodRecommendationsRequest) (*DatahubV1alpha1.ListPodRecommendationsResponse, error) {
	scope.Debug("Request received from ListAvailablePodRecommendations grpc function: " + AlamedaUtils.InterfaceToString(in))

	containerDAO := &DaoRecommendationImpl.Container{
		InfluxDBConfig: *s.Config.InfluxDB,
	}

	podRecommendations, err := containerDAO.ListAvailablePodRecommendations(in)
	if err != nil {
		scope.Error(err.Error())
		return &DatahubV1alpha1.ListPodRecommendationsResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INTERNAL),
				Message: err.Error(),
			},
		}, nil
	}

	res := &DatahubV1alpha1.ListPodRecommendationsResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		PodRecommendations: podRecommendations,
	}
	scope.Debug("Response sent from ListPodRecommendations grpc function: " + AlamedaUtils.InterfaceToString(res))
	return res, nil
}

// ListControllerRecommendations list controller recommendations
func (s *ServiceV1alpha1) ListControllerRecommendations(ctx context.Context, in *DatahubV1alpha1.ListControllerRecommendationsRequest) (*DatahubV1alpha1.ListControllerRecommendationsResponse, error) {
	scope.Debug("Request received from ListControllerRecommendations grpc function: " + AlamedaUtils.InterfaceToString(in))

	controllerDAO := &DaoRecommendationImpl.Controller{
		InfluxDBConfig: *s.Config.InfluxDB,
	}

	controllerRecommendations, err := controllerDAO.ListControllerRecommendations(in)
	if err != nil {
		scope.Errorf("api ListControllerRecommendations failed: %v", err)
		response := &DatahubV1alpha1.ListControllerRecommendationsResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INTERNAL),
				Message: err.Error(),
			},
			ControllerRecommendations: controllerRecommendations,
		}
		return response, nil
	}

	response := &DatahubV1alpha1.ListControllerRecommendationsResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		ControllerRecommendations: controllerRecommendations,
	}

	scope.Debug("Response sent from ListControllerRecommendations grpc function: " + AlamedaUtils.InterfaceToString(response))
	return response, nil
}
