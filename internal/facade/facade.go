package facade

import (
	"context"
	"github.com/calebtracey/api-template/external/models"
	"github.com/calebtracey/api-template/external/models/request"
	psql2 "github.com/calebtracey/api-template/internal/facade/psql"
	"strings"
)

type ServiceI interface {
	FacadeResponse(ctx context.Context, req models.Request) (resp models.Response)
}

type Facade struct {
	PSQLDao psql2.FacadeI
}

func (s Facade) FacadeResponse(ctx context.Context, req models.Request) (resp models.Response) {
	//TODO add validation
	if strings.EqualFold(req.Type, "Insert") {
		_ = s.PSQLDao.AddNew(ctx, request.PSQLRequest{})
		//TODO add mappers
	}

	return resp
}
