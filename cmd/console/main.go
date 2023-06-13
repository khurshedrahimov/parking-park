package main

import (
	"fmt"
	"parking-park/pkg/models"
)

func main() {
	var command int
	var number, color, brand string
	parking := models.Parking{}

	fmt.Println("How many parking spaces does your parking have?")
	fmt.Scan(&parking.NumberOfParkingSpaces)

	for {
		fmt.Println("\nHere are the list of commands I can process:")
		fmt.Println("\t0 - Exit")
		fmt.Println("\t1 - Park a car into an empty slot")
		fmt.Println("\t2 - Remove a car from a slot")
		fmt.Println("\t3 - Show the list of slots with their cars")

		fmt.Print("Command: ")
		fmt.Scan(&command)

		if command == 0 {
			if parking.IsEmpty() {
				fmt.Println("\nTotal cars served:", parking.TotalCarsServed)
				fmt.Println("Exiting... Bye!")
				break
			} else {
				fmt.Println("\nWe cannot exit the program, there are cars in parking. Remove them before exiting.")
			}
		} else {
			switch command {
			case 1:
				if !parking.HasAvailableSpace() {
					fmt.Println("\nThere is no available space in parking!")
					break
				}
				fmt.Println("\nAdding car to parking. Enter car's info.")

				fmt.Println("Number:")
				fmt.Scan(&number)
				fmt.Println("Color:")
				fmt.Scan(&color)
				fmt.Println("Brand:")
				fmt.Scan(&brand)

				parking.AddCar(number, color, brand)

				fmt.Println("\nAdded car to parking")
			case 2:
				fmt.Println("Removing car from parking...")
				fmt.Print("Enter car's number: ")
				var removeNum string
				fmt.Scan(&removeNum)
				existence := parking.RemoveCar(removeNum)
				if existence {
					fmt.Println("\nCar removed from parking")
				} else {
					fmt.Println("\nThere is no such car in parking!")
				}
			case 3:
				parking.PrintListOfslots()
			default:
				fmt.Println("\nI don't know this command")
			}
		}
	}
}
