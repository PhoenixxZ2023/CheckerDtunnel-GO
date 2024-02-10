package device_handler

import (
	"context"

	device_use_case "github.com/PhoenixxZ2023/CheckerDtunnel-GO/src/domain/usecase/device"
	"github.com/PhoenixxZ2023/CheckerDtunnel-GO/src/infra/handler"
)

type listDevicesHandler struct {
	listDevicesUseCase *device_use_case.ListDevicesUseCase
}

func NewListDevicesHandler(listDevices *device_use_case.ListDevicesUseCase) handler.Handler {
	return &listDevicesHandler{listDevices}
}

func (h *listDevicesHandler) Handle(ctx context.Context, request *handler.HttpRequest) (*handler.HttpResponse, error) {
	devices, err := h.listDevicesUseCase.Execute(ctx)
	if err != nil {
		return nil, err
	}

	return &handler.HttpResponse{
		Status: 200,
		Body:   devices,
	}, nil
}
