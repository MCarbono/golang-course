package usecase

import "cleanArch/internal/entity"

type GetOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetOrdersUseCase(OrderRepository entity.OrderRepositoryInterface) *GetOrdersUseCase {
	return &GetOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (g *GetOrdersUseCase) Execute() ([]entity.Order, error) {
	return g.OrderRepository.GetAll()
}
