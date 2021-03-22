package abstract_factory

import "testing"

func TestMotorbikeFactory(t *testing.T) {
	motorbikeF, err := GetVehicleFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}
	motorbikeVehicle, err := motorbikeF.GetVehicle(SportMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Motorbike vehicle has %d wheels\n", motorbikeVehicle.GetWheels())
	t.Logf("Motorbike vehicle has %d seats\n", motorbikeVehicle.GetSeats())
	sportBike, ok := motorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Sport motorbike has type %d\n", sportBike.GetType())

	motorbikeVehicle, err = motorbikeF.GetVehicle(CruiseMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Motorbike vehicle has %d wheels\n", motorbikeVehicle.GetWheels())
	t.Logf("Motorbike vehicle has %d seats\n", motorbikeVehicle.GetSeats())
	cruiseMotorBike, ok := motorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Sport motorbike has type %d\n", cruiseMotorBike.GetType())
}

func TestCarFactory(t *testing.T) {
	carF, err := GetVehicleFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}
	carVehicle, err := carF.GetVehicle(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Car vehicle has %d wheels\n", carVehicle.GetWheels())
	t.Logf("Car vehicle has %d seats\n", carVehicle.GetSeats())
	luxuryCar, ok := carVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Luxury car has %d doors.\n", luxuryCar.GetDoors())

	carVehicle, err = carF.GetVehicle(FamilyCarType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Car vehicle has %d wheels\n", carVehicle.GetWheels())
	t.Logf("Car vehicle has %d seats\n", carVehicle.GetSeats())
	familyCar, ok := carVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Luxury car has %d doors.\n", familyCar.GetDoors())
}

func TestCarFactoryNotExists(t *testing.T) {
	_, err := GetVehicleFactory(10)
	if err == nil {
		t.Error("GetVehicleFactory with ID 10 must return an error")
	}
	t.Log("LOG:", err)

	carf, _ := GetVehicleFactory(CarFactoryType)
	_, err = carf.GetVehicle(10)
	if err == nil {
		t.Error("carFactory with ID 10 must return an error")
	}

	motorf, _ := GetVehicleFactory(MotorbikeFactoryType)
	_, err = motorf.GetVehicle(10)
	if err == nil {
		t.Error("MotorbikeFactory with ID 10 must return an error")
	}

}
