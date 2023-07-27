package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}

}

func uniq(data string) map[string]int {
	lines := strings.Split(data, "\n")
	duplicatedData := make(map[string]int)
	for _, item := range lines {
		_, item1 := duplicatedData[item]

		if item1 {
			duplicatedData[item] += 1
		} else {

			duplicatedData[item] = 1
		}
	}
	return duplicatedData
}

func uniq1(data string, ch string, num int) map[string]int {
	lines := strings.Split(data, "\n")
	duplicatedData := make(map[string]int)
	duplicatedDataTirm := make(map[string]int)
	switch ch {
	case "S":
		for i, item := range lines {
			if len(item) == 0 {
				continue
			}
			tempChar := []rune(item)
			tempTirmChar := tempChar[num:]
			if i == len(lines)-1 {
				tempTirmChar = tempChar[num:]
			} else {
				tempTirmChar = tempChar[num : len(tempChar)-1]
			}
			tempString := string(tempTirmChar)
			_, item1 := duplicatedDataTirm[tempString]

			if item1 {
				duplicatedDataTirm[tempString] += 1
				for k, _ := range duplicatedData {
					tempIfChar := []rune(k)
					tempIfTrimChar := tempIfChar[num : len(tempIfChar)-1]

					tempIfString := string(tempIfTrimChar)
					if tempIfString == tempString {
						duplicatedData[k] += 1
					}
				}
			} else {
				duplicatedDataTirm[tempString] = 1
				duplicatedData[item] = 1
			}
		}

	case "F":
		for _, item := range lines {
			if len(item) == 0 {
				continue
			}
			tempString := strings.Fields(item)
			tempSlice := tempString[num:]
			tempStringJoin := strings.Join(tempSlice, " ")
			_, item1 := duplicatedDataTirm[tempStringJoin]

			if item1 {
				duplicatedDataTirm[tempStringJoin] += 1
				for k := range duplicatedData {
					tempIfString := strings.Fields(k)
					tempIfSlice := tempIfString[num:]
					tempIfStringJoin := strings.Join(tempIfSlice, " ")
					if tempIfStringJoin == tempStringJoin {
						duplicatedData[k] += 1
					}
				}
			} else {
				duplicatedDataTirm[tempStringJoin] = 1
				duplicatedData[item] = 1
			}
		}

	default:
		fmt.Println("Invalid Option!!")

	}
	return duplicatedData
}

func main() {
	data, err := os.ReadFile("example.txt")
	check(err)
	dataStr := string(data)
	//fmt.Println(dataStr)

	uniqCom := uniq(dataStr)
	fmt.Println("======The First uniq=====")
	for k, v := range uniqCom {
		//fmt.Println(k)
		fmt.Println("The item is: ", k)
		fmt.Print("The frequency is: ", v)
		fmt.Println("")
	}
	fmt.Println("=================================")
	/////////////////////////////////////
	fmt.Println("====== uniq with skip characters=====")

	uniqCom1 := uniq1(dataStr, "S", 5)
	for k, v := range uniqCom1 {
		//fmt.Println(k)
		fmt.Println("The item is: ", k)
		fmt.Print("The frequency is: ", v)
		fmt.Println("")
	}

	fmt.Println("=================================")
	fmt.Println("=================================")
	//////////////////////////////////////////////////////////
	fmt.Println("======uniq with skip Fields=====")

	uniqCom2 := uniq1(dataStr, "F", 1)
	fmt.Println("=================================")
	for k, v := range uniqCom2 {
		//fmt.Println(k)
		fmt.Println("The item is: ", k)
		fmt.Print("The frequency is: ", v)
		fmt.Println("")
	}

}
