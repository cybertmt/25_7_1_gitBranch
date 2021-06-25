package main

import (
	"fmt"
	"strings"
)
// Структура с данными человека
type Man struct {
	Name string
	LastName string
	Age int
	Gender string
	Crimes int
}

func main(){
// Мапа с персданными, ключ -имя
	peopleMap :=  map[string]Man {

		"John" : {"John","D.",20,"Male",0},
		"Doe" : {"Doe","J.",21,"Male",0},
		"Mary" : {"Mary","M.",18,"Female",0},
		"Jane" : {"Jane","J.",30,"Female",1},
		"Patrick" : {"Patrick","K.",28,"Male",15},
		"Freddy" : {"Freddy","",40,"Male",55},
	}

	var suspects []string // Список подозреваемых
	var topCriminal string // Имя криминального лидера
	topCrimesNumber := 0 // Наибольшее количество преступлений

	for _, value := range peopleMap {

		if value.Crimes != 0 {
			suspects = append(suspects, value.Name)

			if value.Crimes > topCrimesNumber {
				topCrimesNumber = value.Crimes
				topCriminal = value.Name
			}
		}

	}

	switch topCrimesNumber {
	case 0: fmt.Println("City can sleep peacefully")
	default: {
		fmt.Println("Suspects:", strings.Join(suspects, ", "))
		fmt.Printf( "%s is the criminal with %d crimes\n", topCriminal, topCrimesNumber)
	}
	}
	fmt.Println("Version 1.1")
	fmt.Print("\nPress any key...")
	_, _ = fmt.Scanln()
}
