package abstract_factory

import "fmt"

type BenzCar struct {
}

func (car *BenzCar) CreateSportCar() SportCarFactory {
	return new(BenzSportCar)
}

func (car *BenzCar) CreateBusinessCar() BusinessCarFactory {
	return new(BenzBusinessCar)
}

type BenzSportCar struct {
}

func (BenzSportCar) Drive() {
	fmt.Println("drive benz sport car")
}

type BenzBusinessCar struct {
}

func (BenzBusinessCar) Drive() {
	fmt.Println("drive benz business car")
}
