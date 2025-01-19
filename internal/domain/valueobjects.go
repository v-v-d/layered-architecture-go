package domain

type Price struct {
	value int32
}

func NewPrice(value int32) (Price, error) {
	if value <= 0 {
		return Price{}, &InvalidPriceError{Value: value}
	}
	return Price{value: value}, nil
}

func (p Price) Value() int32 {
	return p.value
}

func (p Price) String() string {
	return string(p.Value())
}

type Quantity struct {
	value int32
}

func NewQuantity(value int32) (Quantity, error) {
	if value <= 0 {
		return Quantity{}, &InvalidQuantityError{Value: value}
	}
	return Quantity{value: value}, nil
}

func (q Quantity) Value() int32 {
	return q.value
}

func (q Quantity) String() string {
	return string(q.value)
}
