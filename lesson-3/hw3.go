package main

import "fmt"

type Car struct {
	brand        string
	Year         int
	TrunkVolume  int
	Engine       bool
	Window       bool
	TrunkPrecent int
}
type InformPerson struct {
	Id    int
	Name  string
	Phone int
}
type Truck struct {
	brand       string
	Year        int
	TrunkVolume int
	Person      InformPerson
}

var Car1 = Car{
	brand:        "Hundai",
	Year:         2012,
	TrunkVolume:  40,
	Engine:       true,
	Window:       false,
	TrunkPrecent: 50,
}
var Car2 = Truck{
	brand:       "Toyota",
	Year:        2016,
	TrunkVolume: 60,
	Person:      InformPerson{567456, "Ivan", 89996785432},
}

func main() {
	fmt.Println(Car1)
	fmt.Printf("%v\n", Car2)
}
