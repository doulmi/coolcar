package rental

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/auth"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type Service struct {
	Logger *zap.Logger
	DB     *gorm.DB
}

func (this *Service) CreateTrip(ctx context.Context, req *rentalpb.CreateTripRequest) (*rentalpb.TripEntity, error) {
	_, err := auth.AccountIDFromContext(ctx)

	// trip := rentalpb.Trip{
	// 	AccountId: accountId,
	// 	Start:     req.Start,
	// }
	// this.DB.Create()
	if err != nil {
		return nil, err
	}

	return nil, status.Error(codes.Unimplemented, "")
}

func (this *Service) GetTrip(context.Context, *rentalpb.GetTripRequest) (*rentalpb.Trip, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (this *Service) GetTrips(context.Context, *rentalpb.GetTripsRequest) (*rentalpb.GetTripsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (this *Service) UpdateTrip(context.Context, *rentalpb.UpdateTripRequest) (*rentalpb.Trip, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
