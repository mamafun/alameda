package v1alpha1

import (
	DaoWeaveScope "github.com/containers-ai/alameda/datahub/pkg/dao/weavescope"
	AlamedaUtils "github.com/containers-ai/alameda/pkg/utils"
	DatahubV1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
)

func (s *ServiceV1alpha1) GetWeaveScopeHostDetails(ctx context.Context, in *DatahubV1alpha1.ListWeaveScopeHostsRequest) (*DatahubV1alpha1.WeaveScopeResponse, error) {
	scope.Debug("Request received from GetWeaveScopeHostDetails grpc function: " + AlamedaUtils.InterfaceToString(in))

	response := &DatahubV1alpha1.WeaveScopeResponse{}

	weaveScopeDAO := DaoWeaveScope.WeaveScope{
		WeaveScopeConfig: s.Config.WeaveScope,
	}
	rawdata, err := weaveScopeDAO.GetWeaveScopeHostDetails(in)

	if err != nil {
		scope.Error(err.Error())
		return &DatahubV1alpha1.WeaveScopeResponse{
			Status: &status.Status{
				Code: int32(code.Code_OK),
			},
			Rawdata: rawdata,
		}, nil
	}

	response.Rawdata = rawdata
	return response, nil
}

func (s *ServiceV1alpha1) GetWeaveScopePodDetails(ctx context.Context, in *DatahubV1alpha1.ListWeaveScopePodsRequest) (*DatahubV1alpha1.WeaveScopeResponse, error) {
	scope.Debug("Request received from GetWeaveScopePodDetails grpc function: " + AlamedaUtils.InterfaceToString(in))

	response := &DatahubV1alpha1.WeaveScopeResponse{}

	weaveScopeDAO := DaoWeaveScope.WeaveScope{
		WeaveScopeConfig: s.Config.WeaveScope,
	}
	rawdata, err := weaveScopeDAO.GetWeaveScopePodDetails(in)

	if err != nil {
		scope.Error(err.Error())
		return &DatahubV1alpha1.WeaveScopeResponse{
			Status: &status.Status{
				Code: int32(code.Code_OK),
			},
			Rawdata: rawdata,
		}, nil
	}

	response.Rawdata = rawdata
	return response, nil
}

func (s *ServiceV1alpha1) GetWeaveScopeContainerDetails(ctx context.Context, in *DatahubV1alpha1.ListWeaveScopeContainersRequest) (*DatahubV1alpha1.WeaveScopeResponse, error) {
	scope.Debug("Request received from GetWeaveScopeContainerDetails grpc function: " + AlamedaUtils.InterfaceToString(in))

	response := &DatahubV1alpha1.WeaveScopeResponse{}

	weaveScopeDAO := DaoWeaveScope.WeaveScope{
		WeaveScopeConfig: s.Config.WeaveScope,
	}
	rawdata, err := weaveScopeDAO.GetWeaveScopeContainerDetails(in)

	if err != nil {
		scope.Error(err.Error())
		return &DatahubV1alpha1.WeaveScopeResponse{
			Status: &status.Status{
				Code: int32(code.Code_OK),
			},
			Rawdata: rawdata,
		}, nil
	}

	response.Rawdata = rawdata
	return response, nil
}

func (s *ServiceV1alpha1) ListWeaveScopeHosts(ctx context.Context, in *DatahubV1alpha1.ListWeaveScopeHostsRequest) (*DatahubV1alpha1.WeaveScopeResponse, error) {
	scope.Debug("Request received from ListWeaveScopeHosts grpc function: " + AlamedaUtils.InterfaceToString(in))

	response := &DatahubV1alpha1.WeaveScopeResponse{}

	weaveScopeDAO := DaoWeaveScope.WeaveScope{
		WeaveScopeConfig: s.Config.WeaveScope,
	}

	rawdata, err := weaveScopeDAO.ListWeaveScopeHosts(in)

	if err != nil {
		scope.Error(err.Error())
		return &DatahubV1alpha1.WeaveScopeResponse{
			Status: &status.Status{
				Code: int32(code.Code_OK),
			},
			Rawdata: rawdata,
		}, nil
	}

	response.Rawdata = rawdata
	return response, nil
}

func (s *ServiceV1alpha1) ListWeaveScopePods(ctx context.Context, in *DatahubV1alpha1.ListWeaveScopePodsRequest) (*DatahubV1alpha1.WeaveScopeResponse, error) {
	scope.Debug("Request received from ListWeaveScopePods grpc function: " + AlamedaUtils.InterfaceToString(in))

	response := &DatahubV1alpha1.WeaveScopeResponse{}

	weaveScopeDAO := DaoWeaveScope.WeaveScope{
		WeaveScopeConfig: s.Config.WeaveScope,
	}
	rawdata, err := weaveScopeDAO.ListWeaveScopePods(in)

	if err != nil {
		scope.Error(err.Error())
		return &DatahubV1alpha1.WeaveScopeResponse{
			Status: &status.Status{
				Code: int32(code.Code_OK),
			},
			Rawdata: rawdata,
		}, nil
	}

	response.Rawdata = rawdata
	return response, nil
}

func (s *ServiceV1alpha1) ListWeaveScopeContainers(ctx context.Context, in *DatahubV1alpha1.ListWeaveScopeContainersRequest) (*DatahubV1alpha1.WeaveScopeResponse, error) {
	scope.Debug("Request received from ListWeaveScopeContainers grpc function: " + AlamedaUtils.InterfaceToString(in))

	response := &DatahubV1alpha1.WeaveScopeResponse{}

	weaveScopeDAO := DaoWeaveScope.WeaveScope{
		WeaveScopeConfig: s.Config.WeaveScope,
	}
	rawdata, err := weaveScopeDAO.ListWeaveScopeContainers(in)

	if err != nil {
		scope.Error(err.Error())
		return &DatahubV1alpha1.WeaveScopeResponse{
			Status: &status.Status{
				Code: int32(code.Code_OK),
			},
			Rawdata: rawdata,
		}, nil
	}

	response.Rawdata = rawdata
	return response, nil
}

func (s *ServiceV1alpha1) ListWeaveScopeContainersByHostname(ctx context.Context, in *DatahubV1alpha1.ListWeaveScopeContainersRequest) (*DatahubV1alpha1.WeaveScopeResponse, error) {
	scope.Debug("Request received from ListWeaveScopeContainersByHostname grpc function: " + AlamedaUtils.InterfaceToString(in))

	response := &DatahubV1alpha1.WeaveScopeResponse{}

	weaveScopeDAO := DaoWeaveScope.WeaveScope{
		WeaveScopeConfig: s.Config.WeaveScope,
	}
	rawdata, err := weaveScopeDAO.ListWeaveScopeContainersByHostname(in)

	if err != nil {
		scope.Error(err.Error())
		return &DatahubV1alpha1.WeaveScopeResponse{
			Status: &status.Status{
				Code: int32(code.Code_OK),
			},
			Rawdata: rawdata,
		}, nil
	}

	response.Rawdata = rawdata
	return response, nil
}

func (s *ServiceV1alpha1) ListWeaveScopeContainersByImage(ctx context.Context, in *DatahubV1alpha1.ListWeaveScopeContainersRequest) (*DatahubV1alpha1.WeaveScopeResponse, error) {
	scope.Debug("Request received from ListWeaveScopeContainersByImage grpc function: " + AlamedaUtils.InterfaceToString(in))

	response := &DatahubV1alpha1.WeaveScopeResponse{}

	weaveScopeDAO := DaoWeaveScope.WeaveScope{
		WeaveScopeConfig: s.Config.WeaveScope,
	}
	rawdata, err := weaveScopeDAO.ListWeaveScopeContainersByImage(in)

	if err != nil {
		scope.Error(err.Error())
		return &DatahubV1alpha1.WeaveScopeResponse{
			Status: &status.Status{
				Code: int32(code.Code_OK),
			},
			Rawdata: rawdata,
		}, nil
	}

	response.Rawdata = rawdata
	return response, nil
}
