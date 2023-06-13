package models

import "fmt"

type Parking struct {
	NumberOfParkingSpaces int
	TotalCarsServed       int
	slots                 map[int]Car
}

func (p *Parking) HasAvailableSpace() bool {

	return len(p.slots) != p.NumberOfParkingSpaces
}

func (p *Parking) IsEmpty() bool {

	return len(p.slots) == 0
}

func (p *Parking) AddCar(number, color, brand string) bool {
	if !p.HasAvailableSpace() {
		return false
	}

	if p.slots == nil {
		p.slots = make(map[int]Car)
	}
	parkingSlot := map[int]Car{}
	parkingSlot = p.slots

	var parkingCar Car
	parkingCar = Car{number, color, brand}

	var keys []int
	for i := 0; i < p.NumberOfParkingSpaces; i++ {
		keys = append(keys, i+1)
	}
	for _, k := range keys {
		_, ok := parkingSlot[k]

		if !ok {
			parkingSlot[k] = parkingCar
			p.slots = parkingSlot
			p.TotalCarsServed += 1
			break
		}
	}
	return true
}
func (p *Parking) RemoveCar(number string) bool {
	parkedCars := make(map[int]Car)
	parkedCars = p.slots

	for k, v := range parkedCars {
		if v.Number == number {
			delete(parkedCars, k)
			return true
		}
	}
	return false
}

func (p *Parking) PrintListOfslots() {
	parkedCar := make(map[int]Car)
	parkedCar = p.slots

	var keys []int
	for i := 0; i < p.NumberOfParkingSpaces; i++ {
		keys = append(keys, i+1)
	}
	empty := make(map[int]Car)
	for _, k := range keys {
		if empty[k] == parkedCar[k] {
			fmt.Print(k, ": -- Empty --\n")
		} else {
			fmt.Printf("%v: %s %s with numbers of %s \n", k, parkedCar[k].Color, parkedCar[k].Brand, parkedCar[k].Number)
		}
	}
}
