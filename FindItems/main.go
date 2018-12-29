package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var outStream io.Writer = os.Stdout
var outFile io.Writer

type item struct {
	name  string
	price int
}

func main() {
	//Log errors
	mw := io.MultiWriter(outFile, outStream)
	log.SetOutput(mw)
	outFile, err := os.OpenFile("error.log", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	//Verify command line argument
	ii := len(os.Args)
	var maxPrice int
	var numTotItems int
	switch {
	case ii == 4:
		maxPrice, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal("Unknown maximum price of items. ", err)
		}
		numTotItems, err = strconv.Atoi(os.Args[3])
		if err != nil {
			log.Fatal("Unknown number of total items to select. ", err)
		}
	default:
		log.Fatal("Unknown command line input. ", err)
	}

	//Prepare to read file
	var scanner *bufio.Scanner
	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()
	scanner = bufio.NewScanner(inputFile)

	//Find items
	indices, err := findItems(scanner, maxPrice, numTotItems)

	//Display result
	if err != nil {
		fmt.Fprintln(outStream, err.Error())
		return
	}
	fmt.Println(indices)

}

//findItems finds combination of items satisfying the item-count and maximum-price and constraints
func findItems(scanner *bufio.Scanner, maxPrice int, numTotItems int) ([]int, error) {

	var indices []int
	priceArray1 := make([]int, maxPrice+1)
	priceArray2 := make([]int, maxPrice+1)
	itemChosen1 := make([]map[int]int, maxPrice+1)
	itemChosen2 := make([]map[int]int, maxPrice+1)
	var list []item

	newlineStr := getNewlineStr()
	rowNum := -1
	for scanner.Scan() {
		//Read-in items
		input := scanner.Text()
		input = strings.TrimRight(input, newlineStr)
		s := parse(input)
		curItemPrice, err := strconv.Atoi(s[1])
		if err != nil {
			log.Fatalf("Unknown item format in input file: %s", input)
		}
		rowNum++

		//Store items in list
		list = append(list, item{name: s[0], price: curItemPrice})
		fmt.Println(list)

		for ii := 0; ii <= maxPrice; ii++ {
			if curItemPrice <= ii {
				diffPrice := ii - curItemPrice
				prevTot := priceArray1[diffPrice]
				newTot := prevTot + curItemPrice
				switch {
				case newTot > priceArray1[ii]:
					priceArray2[ii] = newTot
					itemChosen2[ii] = computeMap(itemChosen1[diffPrice], rowNum, numTotItems)

				case newTot == priceArray1[ii]:
					priceArray2[ii] = newTot
					itemChosen2[ii] = computeMap(itemChosen1[diffPrice], rowNum, numTotItems)
					mergeMap(itemChosen2[ii], itemChosen1[ii])

				default:
					priceArray2[ii] = priceArray1[ii]
					itemChosen2[ii] = itemChosen1[ii]
				}
			} else {
				priceArray2[ii] = priceArray1[ii]
				itemChosen2[ii] = itemChosen1[ii]
			}
		}
		//Swap the arrays
		priceArray1, priceArray2 = priceArray2, priceArray1
		itemChosen1, itemChosen2 = itemChosen2, itemChosen1

		fmt.Println(priceArray1)
		fmt.Println(priceArray2)
		fmt.Println(itemChosen1)
		fmt.Println(itemChosen2)
		fmt.Printf("Input: %s \n", input)
		fmt.Println("------------------------")
	}

	col := maxPrice
	for ii := numTotItems; ii > 0; ii-- {
		if elem, ok := itemChosen1[col][ii]; ok {
			col = priceArray1[col] - list[elem].price
			indices = append(indices, elem)
		} else {
			return indices, errors.New("No combination of items available")
		}
	}

	fmt.Println(list)
	return indices, nil

}

//getNewlineStr identifies operating system and returns newline character used
func getNewlineStr() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}
	return "\n"
}

//parse string using given separator
func parse(input string) []string {
	s := strings.Split(input, ", ")
	return s
}

//mergeMap concatenates in2 to in1. Overlapping key values in in1 are overwritten with key values from in2.
func mergeMap(in1 map[int]int, in2 map[int]int) {
	for key, val := range in2 {
		in1[key] = val
	}
}

//computeMap (i) returns a new map[int]int{1:val} if in1 is nil map, or
//(ii) returns in1 where each key is incremented by 1 and assigned a value val, provided the key is smaller than maxKey.
func computeMap(in1 map[int]int, val int, maxKey int) map[int]int {
	out1 := make(map[int]int)
	if in1 == nil {
		out1[1] = val
	} else {
		for key := range in1 {
			if key+1 <= maxKey {
				out1[key+1] = val
			}
		}
	}
	return out1
}
