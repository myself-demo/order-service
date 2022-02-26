package service

import (
	"context"
	"github.com/myself-demo/order-service-api/pb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"time"
)

type OrderServiceProvider struct {
}

// QueryById 通过ID查询订单
// 如果查询的订单ID不等于123则返回错误
func (provider *OrderServiceProvider) QueryById(ctx context.Context, param *wrapperspb.StringValue) (*pb.QueryByIdResult, error) {
	if param.GetValue() != "123" {
		return &pb.QueryByIdResult{
			HandleResult: &pb.HandleResult{Code: 1, Message: "the Order not found!"},
		}, nil
	}
	return &pb.QueryByIdResult{
		HandleResult: &pb.HandleResult{Code: 0},
		Order: &pb.Order{
			OrderId:     param.GetValue(),
			UserId:      "88888888",
			ProductName: "屠龙宝刀",
			PayMoney:    123.5,
			OrderTime:   uint64(time.Now().UnixMilli()),
		},
	}, nil
}

// ToOrder 下单
// 如果用户ID不是88888888则返回下单失败,用户不存在
// 如果下单的金额大于200返回用户余额不足
// 正常情况下返回订单ID
func (provider *OrderServiceProvider) ToOrder(ctx context.Context, param *pb.Order) (*pb.ToOrderResult, error) {
	if param.GetUserId() != "88888888" {
		return &pb.ToOrderResult{
			HandleResult: &pb.HandleResult{Code: 2, Message: "下单失败"},
			OrderId:      "",
		}, nil
	}
	if param.GetPayMoney() > 200 {
		return &pb.ToOrderResult{
			HandleResult: &pb.HandleResult{Code: 3, Message: "用户余额不足"},
			OrderId:      "",
		}, nil
	}
	return &pb.ToOrderResult{
		HandleResult: &pb.HandleResult{Code: 0},
		OrderId:      "123",
	}, nil
}

func NewOrderServiceProvider() *OrderServiceProvider {
	return &OrderServiceProvider{}
}
