package abstract_factory

type AbstractCarFactory interface {
	CreateSportCar() SportCarFactory
	CreateBusinessCar() BusinessCarFactory
}
