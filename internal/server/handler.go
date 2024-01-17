package server

import (
	pb "service/examples/pkg/pb/v1"

	"github.com/starfork/stargo"
	"github.com/starfork/stargo/logger"
)

type handler struct {
	logger logger.Logger
	pb.UnimplementedExamplesServer
}

func New(app *stargo.App) *handler {

	return &handler{
		logger: app.GetLogger(),
	}
}
