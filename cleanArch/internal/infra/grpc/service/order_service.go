package service

import (
	"cleanArch/internal/infra/grpc/pb"
	"cleanArch/internal/usecase"
	"context"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	GetOrdersUseCase   usecase.GetOrdersUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, getOrderUseCase usecase.GetOrdersUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		GetOrdersUseCase:   getOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(context.Context, *pb.Blank) (*pb.OrdersList, error) {
	orders, err := s.GetOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}
	var output = make([]*pb.Orders, len(orders))
	for i := range output {
		output[i] = &pb.Orders{
			Id:         orders[i].ID,
			Price:      float32(orders[i].Price),
			Tax:        float32(orders[i].Tax),
			FinalPrice: float32(orders[i].FinalPrice),
		}
	}
	return &pb.OrdersList{Orders: output}, nil
}
