package main

import (
	"fmt"
	"strings"
)

func AnalyzeText(text string) {
	separators := " .,!?!"
	words := strings.FieldsFunc(text, func(r rune) bool {
		return strings.ContainsRune(separators, r)
	})
	fmt.Println("Количество слов:", len(words))
	var unique_words = make(map[string]int)
	for _, value := range words {

		unique_words[strings.ToLower(value)]++
	}
	fmt.Println("Количество уникальных слов:", len(unique_words))
	var max = 0
	var max_str = ""
	for key, value := range unique_words {
		if value > max {
			max = value
			max_str = key
		}
	}
	fmt.Printf("Самое часто встречающееся слово: \"%s\" (встречается %d раз)\n", max_str, max)

	var max1 = 0
	var str1 = ""
	var str2 = ""
	var str3 = ""
	var str4 = ""
	var str5 = ""
	var max2 = 0
	var max3 = 0
	var max4 = 0
	var max5 = 0
	for key, value := range unique_words {
		if value > max1 {
			max1 = value
			str1 = key
		}
	}
	for key, value := range unique_words {
		if value > max2 && value < max1 {
			max2 = value
			str2 = key
		}
	}
	for key, value := range unique_words {
		if value > max3 && value < max2 {
			max3 = value
			str3 = key
		}
	}
	for key, value := range unique_words {
		if value > max4 && value < max3 {
			max4 = value
			str4 = key
		}
	}
	for key, value := range unique_words {
		if value > max5 && value < max4 {
			max5 = value
			str5 = key
		}
	}
	fmt.Println("Топ-5 самых часто встречающихся слов:")
	fmt.Printf("\"%s\": %d раз\n", str1, max1)
	fmt.Printf("\"%s\": %d раз\n", str2, max2)
	fmt.Printf("\"%s\": %d раз\n", str3, max3)
	fmt.Printf("\"%s\": %d раз\n", str4, max4)
	fmt.Printf("\"%s\": %d раз\n", str5, max5)
}

func getTopWords(wordMap map[string]int, n int) []string {
	var max1 = 0
	var str1 = ""
	var str2 = ""
	var str3 = ""
	var str4 = ""
	var str5 = ""
	var max2 = 0
	var max3 = 0
	var max4 = 0
	var max5 = 0
	for key, value := range wordMap {
		if value > max1 {
			max1 = value
			str1 = key
		}
	}
	for key, value := range wordMap {
		if value > max2 && value < max1 {
			max2 = value
			str2 = key
		}
	}
	for key, value := range wordMap {
		if value > max3 && value < max2 {
			max3 = value
			str3 = key
		}
	}
	for key, value := range wordMap {
		if value > max4 && value < max3 {
			max4 = value
			str4 = key
		}
	}
	for key, value := range wordMap {
		if value > max5 && value < max4 {
			max5 = value
			str5 = key
		}
	}
	var answer = make([]string, n)
	var words = []string{str1, str2, str3, str4, str5}
	for i := 0; i < n; i++ {
		answer[i] = words[i]
	}
	return answer
}

func main() {
	text := "Your sample text here for testing"
	AnalyzeText(text)
}
