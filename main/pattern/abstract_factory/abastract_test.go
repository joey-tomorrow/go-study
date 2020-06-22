package abstract_factory

import "testing"

func Test_abstract(t *testing.T) {
	var factory AbstractCarFactory
	factory = new(BenzCar)

	factory.CreateBusinessCar().Drive()
	factory.CreateSportCar().Drive()
}
