package main

import "fmt"

func main() {
	myflat := flat{
		area:       150,
		furniture:  []string{"Кровать", "Шкаф", "Стол", "Диван", "Кресло"},
		appliances: []string{"Холодильник", "Духовка", "Пылесос", "Чайник", "Миксер", "Фен"},
		family:     []string{"Мама", "Папа", "Бабушка", "Дедушка", "Сестра", "Собака", "Черепаха"},
	}
	myflat.showFlat()
}

type flat struct {
	area       float64
	furniture  []string
	appliances []string
	family     []string
}

type furnitureStruct struct {
	name string
}
type appliancesStruct struct {
	name string
}
type familyStruct struct {
	name string
}

func (f flat) showFlat() {
	fmt.Println("Площадь квартиры:", f.area, "кв.м")
	fmt.Println("Мебель: ")
	for index, furnitureStruct := range f.furniture {
		fmt.Println(index+1, furnitureStruct)
	}
	fmt.Print("Бытовая техника: \n")
	for index, appliancesStruct := range f.appliances {
		fmt.Println(index+1, appliancesStruct)
	}
	fmt.Println("Семья: ")
	for index, familyStruct := range f.family {
		fmt.Println(index+1, familyStruct)
	}
}
