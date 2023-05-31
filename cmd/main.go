package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "{1, 2, 5}"
	str = strings.Trim(str, "{}")
	//str = strings.Trim(str, " ")
	//fmt.Println(str)
	strValues := strings.Split(str, ",") // Разделяем элементы по запятой

	var intSlice []int
	for _, val := range strValues {
		num, err := strconv.Atoi(strings.TrimSpace(val))
		if err != nil {
			fmt.Printf("Ошибка преобразования числа: %v\n", err)
			continue
		}
		intSlice = append(intSlice, num)
	}

	fmt.Println(intSlice)
}
