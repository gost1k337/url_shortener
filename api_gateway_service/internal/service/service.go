package service

import "github.com/gost1k337/url_shortener/api_gateway_service/pkg/logging"

type Services struct {
	logger logging.Logger
}

type ServicesDependencies struct {
}

func NewServices(deps *ServicesDependencies, logger logging.Logger) *Services {
	return &Services{
		logger: logger,
	}
}
