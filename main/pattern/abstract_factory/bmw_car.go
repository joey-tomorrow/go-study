package abstract_factory

import "fmt"

type BmwCar struct {
}

func (car *BmwCar) CreateSportCar() SportCarFactory {
	return new(BmwSportCar)
}

func (car *BmwCar) CreateBusinessCar() BusinessCarFactory {
	return new(BmwBusinessCar)
}

type BmwSportCar struct {
}

func (BmwSportCar) Drive() {
	fmt.Println("drive bmw sport car")
}

type BmwBusinessCar struct {
}

func (BmwBusinessCar) Drive() {
	fmt.Println("drive bmw business car")
}
