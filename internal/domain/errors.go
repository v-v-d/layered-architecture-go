package domain

import "fmt"

type InvalidQuantityError struct {
	Value int32
}

func (e *InvalidQuantityError) Error() string {
	return fmt.Sprintf("Error: Quantity %d is invalid.", e.Value)
}

type InvalidPriceError struct {
	Value int32
}

func (e *InvalidPriceError) Error() string {
	return fmt.Sprintf("Error: Price %d is invalid.", e.Value)
}
