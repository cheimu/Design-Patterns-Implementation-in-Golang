package prototype

import (
	"errors"
	"fmt"
)

type ShirtCloner interface {
	GetClone(m int) (ItemInfoGetter, error)
}

const (
	White = iota
	Black
	Blue
)

func GetShirtsCloner() ShirtCloner {
	return &shirtCache
}

type ShirtsCache struct{}

var shirtCache ShirtsCache = ShirtsCache{}

func (s *ShirtsCache) GetClone(m int) (ItemInfoGetter, error) {
	var newItem Shirt
	switch m {
	case White:
		newItem = *whitePrototype
	case Black:
		newItem = *blackPrototype
	case Blue:
		newItem = *bluePrototype
	default:
		return nil, errors.New("Shirt model not recognized")
	}
	return &newItem, nil
}

type ItemInfoGetter interface {
	GetInfo() string
	GetPrice() float32
}

type ShirtColor int

type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs %f\n", s.SKU, s.Color, s.Price)
}

func (s *Shirt) GetPrice() float32 {
	return s.Price
}

var whitePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}

var blackPrototype *Shirt = &Shirt{
	Price: 16.00,
	SKU:   "empty",
	Color: Black,
}

var bluePrototype *Shirt = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}
