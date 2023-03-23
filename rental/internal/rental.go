package rental

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/auth"

	"go.uber.org/zap"
)

type Service struct {
	Logger *zap.Logger
}

func (this *Service) CreateTrip(ctx context.Context, req *rentalpb.CreateTripRequest) (*rentalpb.CreateTripResponse, error) {
	accountId, err := auth.AccountIDFromContext(ctx)

	if err != nil {
		return nil, err
	}

	this.Logger.Info("create trip", zap.String("start", req.Start), zap.String("accountId", accountId))

	return &rentalpb.CreateTripResponse{}, nil
}
