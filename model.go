package main

import (
	"errors"
)

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *product) getProduct() error {
	return errors.New("Not Implemented")
}

func (p *product) updateProduct() error {
	return errors.New("Not Implemented")
}

func (p *product) deleteProduct() error {
	return errors.New("Not Implemented")
}

func (p *product) createProduct() error {
	return errors.New("Not Implemented")
}

func getProducts(start, total int) ([]product, error) {
	return nil, errors.New("Not Implemented")
}
