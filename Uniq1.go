package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	usage = `Specify a command to execute and be careful options -u and -d are mutually exclusive:
  - uniq -u -d -c -s <num> +f <num> input_file
   `
)

func main() {
	//namePtr := flag.String("command", "uniq", "The name of command")
	uniquePtr := flag.Bool("u", false, "Only output unique lines")
	duplicatePtr := flag.Bool("d", false, "Only output duplicate lines")
	countPtr := flag.Bool("c", false, "Prefix lines by the number of occurrences")
	numPtrChar := flag.Int("s", 0, "Skip the first num characters of each line")
	numPtrField := flag.Int("f", 0, "Skip the first num fields of each line")

	flag.Parse()
	/*
		fmt.Println("The name of command ", *namePtr)
		fmt.Println("The value of u is ", *uniquePtr)
		fmt.Println("The value of d is ", *duplicatePtr)
		fmt.Println("The value of c is ", *countPtr)
		fmt.Println("The value of -num is ", *numPtrChar)
		fmt.Println("The value of +num is ", *numPtrField)
	*/

	if len(flag.Args()) < 1 || *uniquePtr && *duplicatePtr {
		fmt.Println(usage)
		os.Exit(1)
	}

	mode := "u"
	if *countPtr {
		mode = "c"
	} else if *duplicatePtr {
		mode = "d"
	}
	fileName := flag.Args()[0]
	file, err := os.Open(fileName)
	check_(err)

	lines := []string{}
	defultInt := []int{}
	index := []int{}
	result := make(map[string]int)
	resultTirmField := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if *numPtrChar == 0 && *numPtrField == 0 {
		result, index = standard(lines)
	} else if *numPtrChar > 0 && *numPtrField > 0 {
		resultTirmField, index = skipf(lines, *numPtrField)
		dataSkipField := make([]string, 0, len(resultTirmField))
		for k := range resultTirmField {
			dataSkipField = append(dataSkipField, k)
		}
		result, index = skips(dataSkipField, *numPtrChar, index)
	} else {
		if *numPtrField > 0 {
			result, index = skipf(lines, *numPtrField)
		} else if *numPtrChar > 0 {
			result, index = skips(lines, *numPtrChar, defultInt)
		}

	}
	switch mode {
	case "u":
		printUniqResultu(lines, result, index)
	case "d":
		printUniqResultd(lines, result, index)
	case "c":
		printUniqResultc(lines, result, index)

	}
}

func check_(e error) {
	if e != nil {
		panic(e)
	}

}
func standard(data []string) (map[string]int, []int) {
	lines := data
	index := make([]int, 0)
	duplicatedData := make(map[string]int)
	for i, item := range lines {
		_, item1 := duplicatedData[item]

		if item1 {
			duplicatedData[item] += 1
		} else {

			duplicatedData[item] = 1
			index = append(index, i)
		}
	}
	return duplicatedData, index
}

func skipf(data []string, num int) (map[string]int, []int) {
	lines := data
	duplicatedDataTirm := make(map[string]int)
	index := make([]int, 0)
	for i, item := range lines {
		if len(item) == 0 || len(item) < num {
			continue
		}
		tempString := strings.Fields(item)
		tempSlice := tempString[num:]
		tempStringJoin := strings.Join(tempSlice, " ")

		_, item1 := duplicatedDataTirm[tempStringJoin]

		if item1 {
			duplicatedDataTirm[tempStringJoin] += 1
		} else {
			duplicatedDataTirm[tempStringJoin] = 1
			index = append(index, i)
		}
	}

	return duplicatedDataTirm, index
}

func skips(data []string, num int, index []int) (map[string]int, []int) {
	lines := data
	duplicatedDataTirm := make(map[string]int)
	indexs := make([]int, 0)

	for i, item := range lines {
		if len(item) == 0 || len(item) < num {
			continue
		}
		tempChar := []rune(item)
		tempTirmChar := tempChar[num:]
		tempString := string(tempTirmChar)
		_, item1 := duplicatedDataTirm[tempString]

		if item1 {
			duplicatedDataTirm[tempString] += 1
		} else {
			duplicatedDataTirm[tempString] = 1
			if len(data) == len(index) {
				indexs = append(indexs, index[i])
			} else {
				indexs = append(indexs, i)
			}

		}
	}
	return duplicatedDataTirm, indexs
}

func printUniqResultc(data []string, datauniqResult map[string]int, index []int) {
	i := 0
	for _, v := range datauniqResult {
		fmt.Println(data[index[i]], "	", v)
		i++
	}

}
func printUniqResultu(data []string, datauniqResult map[string]int, index []int) {
	i := 0
	for _, v := range datauniqResult {
		if v == 1 {
			fmt.Println(data[index[i]])
			i++
		}

	}

}
func printUniqResultd(data []string, datauniqResult map[string]int, index []int) {
	i := 0
	for _, v := range datauniqResult {
		if v != 1 {
			fmt.Println(data[index[i]])
			i++
		}

	}

}
