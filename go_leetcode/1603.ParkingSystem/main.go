package main

import "fmt"

func main() {
	ps := Constructor(1, 1, 0)
	fmt.Printf("ps.AddCar(1): %v\n", ps.AddCar(1))
	fmt.Printf("ps.AddCar(2): %v\n", ps.AddCar(2))
	fmt.Printf("ps.AddCar(3): %v\n", ps.AddCar(3))
	fmt.Printf("ps.AddCar(1): %v\n", ps.AddCar(1))
}

type ParkingSystem struct {
	Big    int
	Medium int
	Small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		Big:    big,
		Medium: medium,
		Small:  small,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if this.Small--; this.Small <= 0 {
			return false
		}
	case 2:
		if this.Medium--; this.Medium <= 0 {
			return false
		}
	case 3:
		if this.Big--; this.Big <= 0 {
			return false
		}
	}
	return true
}
