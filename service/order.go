package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/myself-demo/order-service-api/pb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"math/rand"
)

type OrderServiceProvider struct {
}

func (provider *OrderServiceProvider) QueryByID(ctx context.Context, request *wrapperspb.StringValue) (*pb.Order, error) {
	return &pb.Order{
		OrderID:     request.GetValue(),
		UserID:      uuid.New().String(),
		ProductName: "This order is provided by Golang",
		PayMoney:    rand.Float64(),
	}, nil
}

func NewOrderServiceProvider() *OrderServiceProvider {
	return &OrderServiceProvider{}
}
