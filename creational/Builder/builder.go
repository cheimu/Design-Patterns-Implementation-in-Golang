package builder

type BuildProcess interface {
	setWheels() BuildProcess
	setSeats() BuildProcess
	setStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

//Director
type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	f.builder.setWheels()
	f.builder.setSeats()
	f.builder.setStructure()
	//f.builder.SetSeats().SetStructure().SetWheels()
}

func (f *ManufacturingDirector) SetBuilder(builder BuildProcess) {
	f.builder = builder
}

// CarBuilder is object Car
type CarBuilder struct {
	v VehicleProduct
}

func (cb *CarBuilder) setWheels() BuildProcess {
	cb.v.Wheels = 4
	return cb
}

func (cb *CarBuilder) setSeats() BuildProcess {
	cb.v.Seats = 5
	return cb
}

func (cb *CarBuilder) setStructure() BuildProcess {
	cb.v.Structure = "Car"
	return cb
}

func (cb *CarBuilder) GetVehicle() VehicleProduct {
	return cb.v
}

// BikeBuilder is object Bike
type BikeBuilder struct {
	v VehicleProduct
}

func (bb *BikeBuilder) setWheels() BuildProcess {
	bb.v.Wheels = 2
	return bb
}

func (bb *BikeBuilder) setSeats() BuildProcess {
	bb.v.Seats = 2
	return bb
}

func (bb *BikeBuilder) setStructure() BuildProcess {
	bb.v.Structure = "Bike"
	return bb
}

func (bb *BikeBuilder) GetVehicle() VehicleProduct {
	return bb.v
}

// MotorbikeBuilder is object Motorbike
type MotorbikeBuilder struct {
	v VehicleProduct
}

func (mb *MotorbikeBuilder) setWheels() BuildProcess {
	mb.v.Wheels = 2
	return mb
}

func (mb *MotorbikeBuilder) setSeats() BuildProcess {
	mb.v.Seats = 2
	return mb
}

func (mb *MotorbikeBuilder) setStructure() BuildProcess {
	mb.v.Structure = "Motorbike"
	return mb
}

func (mb *MotorbikeBuilder) GetVehicle() VehicleProduct {
	return mb.v
}

// BusBuilder is object Bus
type BusBuilder struct {
	v VehicleProduct
}

func (b *BusBuilder) setWheels() BuildProcess {
	b.v.Wheels = 8
	return b
}

func (b *BusBuilder) setSeats() BuildProcess {
	b.v.Seats = 30
	return b
}

func (b *BusBuilder) setStructure() BuildProcess {
	b.v.Structure = "Bus"
	return b
}

func (b *BusBuilder) GetVehicle() VehicleProduct {
	return b.v
}
