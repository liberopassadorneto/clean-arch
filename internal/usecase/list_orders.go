package usecase

import (
	"github.com/liberopassadorneto/clean-arch/internal/entity"
)

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

type ListOrdersOutputDTO struct {
	Orders []OrderOutputDTO `json:"orders"`
}

func NewListOrdersUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (listOrdersUseCase *ListOrdersUseCase) Execute() (ListOrdersOutputDTO, error) {
	orders, err := listOrdersUseCase.OrderRepository.FindMany()
	if err != nil {
		return ListOrdersOutputDTO{}, err
	}

	var ordersOutput []OrderOutputDTO
	for _, order := range orders {
		ordersOutput = append(ordersOutput, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return ListOrdersOutputDTO{
		Orders: ordersOutput,
	}, nil
}
