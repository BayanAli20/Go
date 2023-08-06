package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

var (
	usage = `Specify a command to execute and be careful options -u and -d are mutually exclusive:
	-u -d -c -s <num> +f <num> input_file
   `
)

func main() {

	lines := []string{}
	prevLine := ""
	result := make(map[string]int)
	//namePtr := flag.String("command", "uniq", "The name of command")
	uniquePtr := flag.Bool("u", false, "Only output unique lines")
	duplicatePtr := flag.Bool("d", false, "Only output duplicate lines")
	countPtr := flag.Bool("c", false, "Prefix lines by the number of occurrences")
	numPtrChar := flag.Int("s", 0, "Skip the first num characters of each line")
	numPtrField := flag.Int("f", 0, "Skip the first num fields of each line")

	flag.Parse()
	if len(flag.Args()) < 1 || *uniquePtr && *duplicatePtr {
		fmt.Println(usage)
		os.Exit(1)
	}
	fileName := flag.Args()[0]
	file, err := os.Open(fileName)
	check_(err)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	checkEmptyFile(lines)
	for _, item := range lines {
		fmt.Println(item)
	}
	if *countPtr || *duplicatePtr {
		sort.Strings(lines)
		fmt.Println("=========")
	}
	uniquList := []string{}
	cutList := make(map[int]string)
	for index, line := range lines {
		if len(line) == 0 {
			continue
		}
		prevOrginalLine := line
		newLine := line
		if *numPtrField > 0 {
			space := regexp.MustCompile(`\s+`)
			s := space.ReplaceAllString(line, " ")
			tempString := strings.Fields(s)
			if len(tempString) < *numPtrField {
				break
			}
			tempSlice := tempString[*numPtrField:]
			newLine = strings.Join(tempSlice, " ")
		}
		if *numPtrChar > 0 {
			tempChar := []rune(newLine)
			if len(tempChar) < *numPtrChar {
				break
			}
			tempTirmChar := tempChar[*numPtrChar:]
			newLine = string(tempTirmChar)

		}
		flag := false
		if newLine != prevLine {
			uniquList = append(uniquList, line)
			prevLine = newLine
			prevOrginalLine = line

		} else if *uniquePtr {
			result[prevOrginalLine] += 1
		}
		if *countPtr || *duplicatePtr {
			for i, item := range cutList {
				if item == newLine {
					result[lines[i]] += 1
					flag = true
					break
				}
			}
		}

		if !flag {
			result[line] = 1
			fmt.Println(newLine, " ", line)
			cutList[index] = newLine
			flag = false
		}
	}
	check_res(result)
	if *uniquePtr {
		for _, item := range uniquList {

			fmt.Println(item)

		}
	} else if *duplicatePtr && *countPtr {
		for item, iter := range result {
			if iter > 1 {
				fmt.Println(iter, " ", item)
			}

		}
	} else if *duplicatePtr {
		for item, iter := range result {
			if iter > 1 {
				fmt.Println(item)
			}
		}
	} else if *countPtr {
		for item, iter := range result {
			fmt.Println(iter, " ", item)
		}

	}

}

func checkEmptyFile(file []string) {
	if len(file) == 0 {
		fmt.Println("Empty file!")
		os.Exit(0)
	}
}
func check_(e error) {
	if e != nil {
		panic(e)
	}
}
func check_res(res map[string]int) {
	if len(res) == 0 {
		fmt.Println("Nothing duplicated")
		os.Exit(0)
	}
}
