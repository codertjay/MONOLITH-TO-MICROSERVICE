package price

import "errors"

var (
	ErrPriceTooLow     = errors.New("price must be greater than zero")
	ErrInvalidCurrency = errors.New("invalid currency")
)

type Price struct {
	cents    uint
	currency string
}

func NewPrice(cents uint, currency string) (Price, error) {
	if cents <= 0 {
		return Price{}, ErrPriceTooLow
	}
	if len(currency) != 3 {
		return Price{}, ErrInvalidCurrency
	}
	return Price{cents: cents, currency: currency}, nil
}

func NewPriceP(cents uint, currency string) Price {
	p, err := NewPrice(cents, currency)
	if err != nil {
		panic(err)
	}
	return p
}

func (p Price) Cents() uint {
	return p.cents
}

func (p Price) Currency() string {
	return p.currency
}
