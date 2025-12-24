package main

import (
	"fmt"
)

func main() {
	var first string = "-"
	var second string = "-"
	var third string = "-"
	var fourth string = "-"
	var fifth string = "-"
	available := 5

	for {
		var name string
		var num int

		n, _ := fmt.Scan(&name)

		if n == 0 {
			continue
		}

		if name == "конец" {
			fmt.Printf("1. %s\n2. %s\n3. %s\n4. %s\n5. %s\n",
				first, second, third, fourth, fifth)
			break
		}

		if name == "очередь" {
			fmt.Printf("1. %s\n2. %s\n3. %s\n4. %s\n5. %s\n",
				first, second, third, fourth, fifth)
			continue
		}

		if name == "количество" {
			fmt.Printf("Осталось свободных мест: %d\nВсего человек в очереди: %d\n",
				available, 5-available)
			continue
		}

		fmt.Scan(&num)
		if num < 1 || num > 5 {
			fmt.Printf("Запись на место номер %d невозможна: некорректный ввод\n", num)
			continue
		}

		if available == 0 {
			fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", num)
			continue
		}

		switch num {
		case 1:
			if first == "-" {
				first = name
				available--
			} else {
				fmt.Println("Запись на место номер 1 невозможна: место уже занято")
			}
		case 2:
			if second == "-" {
				second = name
				available--
			} else {
				fmt.Println("Запись на место номер 2 невозможна: место уже занято")
			}
		case 3:
			if third == "-" {
				third = name
				available--
			} else {
				fmt.Println("Запись на место номер 3 невозможна: место уже занято")
			}
		case 4:
			if fourth == "-" {
				fourth = name
				available--
			} else {
				fmt.Println("Запись на место номер 4 невозможна: место уже занято")
			}
		case 5:
			if fifth == "-" {
				fifth = name
				available--
			} else {
				fmt.Println("Запись на место номер 5 невозможна: место уже занято")
			}
		}
	}
}
