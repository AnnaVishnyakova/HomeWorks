package main

import (
	"fmt"
)

type Film struct {
	Name   string
	Year   string
	Rating string
}

func (a Film) String() string {
	//fmt.Println("Фильмы, которые стоит посмотреть:", a.Name, a.Year, a.Rating)
	return "Фильмы, которые стоит посмотреть:" + a.Name + " Год выхода:" + a.Year + " Рейтинг на кинопоиске: " + a.Rating
}

type Produser struct {
	YouName    string
	FilmCareer string
	BestFilm   string
}

func (b Produser) String() string {
	//fmt.Println(
	return "Знаменитые режисеры, чьи фильмы стоит посмотреть: " + b.YouName + " Фильмы режиссера: " + b.FilmCareer + " Лучший фильм: " + b.BestFilm
}

type Stringer interface {
	String() string
}

func main() {
	a := &Film{
		Name:   "Побег из Шоушенка",
		Year:   "1994",
		Rating: "9",
	}
	Print(a)
	b := &Produser{
		YouName:    "Фрэнк Дарабонт",
		FilmCareer: "Побег из Шоушенка, Зеленая миля, Мажестик,Мгла",
		BestFilm:   "Побег из Шоушенка",
	}
	Print(b)

}
func Print(s Stringer) {
	fmt.Println(s.String())
}
