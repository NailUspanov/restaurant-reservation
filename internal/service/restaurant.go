package service

import (
	"restaurant-reservation/internal/domain"
	"restaurant-reservation/internal/domain/dto"
	"restaurant-reservation/internal/repository"
)

type RestaurantService struct {
	repo repository.Repository
}

func NewRestaurantService(repo repository.Repository) *RestaurantService {
	return &RestaurantService{repo: repo}
}

func (r *RestaurantService) Create(restaurant domain.Restaurant) (int, error) {
	return r.repo.Restaurant.Create(restaurant)
}

func (r *RestaurantService) GetAvailable(peopleQuantity int, time string) ([]dto.AvailableRestaurantResponse, error) {

	reservations, err := r.repo.Reservation.GetAllByTime(time)
	if err != nil {
		return nil, err
	}

	unavailableRestaurantTables := make(map[int][]int)
	availableRestaurantTables := make(map[int][]domain.Table)
	availableRestaurantResponse := make([]dto.AvailableRestaurantResponse, 0, 5)

	// если брони в указанное время есть, проверяю, есть ли достаточное количество мест для новой брони
	if len(reservations) > 0 {

		for _, v := range reservations {
			// инициализирую пустыми слайсами
			if unavailableRestaurantTables[v.Restaurant] == nil {
				unavailableRestaurantTables[v.Restaurant] = make([]int, 0, 3)
				availableRestaurantTables[v.Restaurant] = make([]domain.Table, 0, 3)
			}
			// добавляю занятый стол в мапу, где key - id ресторана, value - слайс занятых столов
			unavailableRestaurantTables[v.Restaurant] = append(unavailableRestaurantTables[v.Restaurant], v.Table)
		}

		// заполняю мапу свободных столов, где key - id ресторана, value - слайс свободных столов
		for _, v := range unavailableRestaurantTables {
			tables, err := r.repo.Table.GetAllNotIn(v)
			if err != nil {
				return nil, err
			}
			for _, i2 := range tables {
				availableRestaurantTables[i2.Restaurant] = append(availableRestaurantTables[i2.Restaurant], i2)
			}
		}

		// из множества доступных столов вычисляю общее количество доступных мест
		for k, v := range availableRestaurantTables {
			seats := 0
			for _, table := range v {
				seats += table.Capacity
			}
			if peopleQuantity > seats {
				delete(availableRestaurantTables, k)
			}
		}

		// создаю и заполняю слайс с ресторанами(id), где есть достаточное количество свободных мест
		availableRestaurants := make([]int, len(availableRestaurantTables))
		i := 0
		for k := range availableRestaurantTables {
			availableRestaurants[i] = k
			i++
		}

		// получаю слайс всех доступных ресторанов по их id
		restaurants, err := r.repo.Restaurant.GetByIds(availableRestaurants)
		if err != nil {
			return nil, err
		}

		// формирую ответ
		for _, restaurant := range restaurants {
			availableRestaurantResponse = append(availableRestaurantResponse, dto.AvailableRestaurantResponse{
				Name:            restaurant.Name,
				Location:        restaurant.Location,
				AvgWaitingTime:  restaurant.AvgWaitingTime,
				AvgBillAmount:   restaurant.AvgBillAmount,
				AvailableTables: availableRestaurantTables[restaurant.Id],
			})
		}
	} else { // если броней в указанное время нет, проверяю вместимость доступных ресторанов для компании

		// получаю вместительность всех ресторанов в виде мапы, где key - id ресторана, value - вместимость
		restaurantsCapacity, err := r.repo.Table.GetAllRestaurantsCapacity()
		if err != nil {
			return nil, err
		}

		// проверяю, поместится ли компания людей в ресторан
		for k, v := range restaurantsCapacity {
			if peopleQuantity <= v {
				tables, err := r.repo.Table.GetAllByRestaurant(k)
				if err != nil {
					return nil, err
				}
				availableRestaurantTables[k] = tables
			}
		}

		// создаю и заполняю слайс доступных ресторанов
		availableRestaurants := make([]int, len(availableRestaurantTables))
		i := 0
		for k := range availableRestaurantTables {
			availableRestaurants[i] = k
			i++
		}

		// получаю слайс всех доступных ресторанов по их id
		restaurants, err := r.repo.Restaurant.GetByIds(availableRestaurants)
		if err != nil {
			return nil, err
		}

		// формирую ответ
		for _, restaurant := range restaurants {
			var availableRestaurant dto.AvailableRestaurantResponse

			availableRestaurant.Name = restaurant.Name
			availableRestaurant.Location = restaurant.Location
			availableRestaurant.AvgBillAmount = restaurant.AvgBillAmount
			availableRestaurant.AvgWaitingTime = restaurant.AvgWaitingTime
			availableRestaurant.AvailableTables = availableRestaurantTables[restaurant.Id]

			availableRestaurantResponse = append(availableRestaurantResponse, availableRestaurant)
		}

	}

	return availableRestaurantResponse, nil
}

func (r *RestaurantService) GetAll() ([]domain.Restaurant, error) {
	return r.repo.Restaurant.GetAll()
}

func (r *RestaurantService) GetById(restaurantId int) (domain.Restaurant, error) {
	return r.repo.Restaurant.GetById(restaurantId)
}

func (r *RestaurantService) Delete(restaurantId int) error {
	return r.repo.Restaurant.Delete(restaurantId)
}
