# 🚗 Parking Lot System Program

This Go program simulates a parking lot system that allows users to park and remove cars from the lot, display a list of occupied slots with their car information, and exit the program. 

## 📖 How to Use

1. Run the program in the command line by typing `go run ./cmd/console/main.go` in the project directory.
2. Enter the number of parking spaces the lot has when prompted by the program.
3. The program enters into an infinite loop where it prompts the user to enter a command.
4. Available commands are:
    * 🅿️ `park` - Park a car in an empty slot.
    * 🚙 `remove` - Remove a car from a slot.
    * 📋 `list` - Show the list of slots with their cars.
    * 🛑 `exit` - Exit the program.
5. If the user chooses to park a car, the program will ask for the car's information such as number, color, and brand, and then add the car to the lot.
6. If the user wants to remove a car, the program will ask for the car's number and if the car exists in the lot, it will be removed.
7. If the user chooses to show the list of slots with their cars, the program will print out the list of occupied slots with the car's information.
8. If the user tries to exit the program while there are still cars parked in the lot, the program will not allow it and will prompt the user to remove the cars first.
9. If the user inputs a command that is not recognized by the program, the program will display an error message.

## ⚙️ How to Install

1. Clone the repository to your local machine.
2. Ensure you have Go installed on your system.
3. Run `go build` command in the project directory to build the executable.
4. Execute the program by running the generated executable file.


