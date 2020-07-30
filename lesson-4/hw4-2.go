package main

import (
	"fmt"
	"sort"
)

type Phonebook struct {
	FirstName string
	Phone     int
	Adress    string
}
type PhoneSort []Phonebook

func (p PhoneSort) Len() int {
	return len(p)
}
func (p PhoneSort) Less(i, j int) bool {
	return p[i].Phone < p[j].Phone
}
func (p PhoneSort) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func main() {
	Adressbook := []Phonebook{
		{"Danny", 89992767892, "1street"},
		{"Mike", 89156547656, "2street"},
		{"Tim", 8564738, "3street"},
	}
	sort.Sort(PhoneSort(Adressbook))
	fmt.Println(Adressbook)
}

// type Phonebook2 []int

// func (p Phonebook2) ViewList() {
// 	for i, phone := range p {
// 		fmt.Printf("\t %v)%v\n", i, phone)
// 	}
// }

// func main() {
// 	Adressbook := make(map[string]Phonebook2)
// 	Adressbook["Danny"] = Phonebook2{89993546382}
// 	Adressbook["Viktor"] = Phonebook2{89996574382}
// 	for name, ph := range Adressbook {
// 		fmt.Println(name)
// 		ph.ViewList()
// 	}
// }
