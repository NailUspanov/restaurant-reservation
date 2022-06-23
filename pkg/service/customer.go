package service

import "restaurant-reservation/pkg/repository"

type CustomerService struct {
	repos repository.Customer
}

func NewCustomerService(repos repository.Customer) *CustomerService {
	return &CustomerService{repos: repos}
}

func (c *CustomerService) GetCustomerIdByPhone(phone string) (int, error) {
	return c.repos.GetCustomerIdByPhone(phone)
}

func (c *CustomerService) Create(name string, phone string) (int, error) {
	return c.repos.Create(name, phone)
}
