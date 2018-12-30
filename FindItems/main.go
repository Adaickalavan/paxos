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
	var numItems int
	switch {
	case ii == 4:
		maxPrice, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal("Unknown maximum price of items. ", err)
		}
		numItems, err = strconv.Atoi(os.Args[3])
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
	itemList, indices, err := findItems(scanner, maxPrice, numItems)

	//Display result
	if err != nil {
		fmt.Fprintln(outStream, err.Error())
		return
	}
	printer(itemList, indices, outStream)

}

//findItems finds combination of items satisfying the item-count and maximum-price and constraints
func findItems(scanner *bufio.Scanner, maxPrice int, numItems int) ([]item, []int, error) {

	priceArray1 := make([]map[int]int, maxPrice+1)
	priceArray2 := make([]map[int]int, maxPrice+1)
	itemChosen1 := make([]map[int]int, maxPrice+1)
	itemChosen2 := make([]map[int]int, maxPrice+1)
	var itemList []item

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

		//Store items in itemList
		itemList = append(itemList, item{name: s[0], price: curItemPrice})

		//Perform dynamic programming
		for ii := 0; ii <= maxPrice; ii++ {
			if curItemPrice <= ii {
				diffPrice := ii - curItemPrice
				prevTot := priceArray1[diffPrice]
				newTot := addMap(map[int]int{1: curItemPrice}, prevTot)
				newMap := computeMap(itemChosen1[diffPrice], rowNum, numItems)
				switch {
				case newTot > priceArray1[ii] && len(newMap) != 0:
					priceArray2[ii] = newTot
					itemChosen2[ii] = newMap

				case newTot > priceArray1[ii] && len(newMap) == 0:
					priceArray2[ii] = newTot
					itemChosen2[ii] = newMap

				case newTot == priceArray1[ii]:
					priceArray2[ii] = newTot
					itemChosen2[ii] = newMap
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
		// fmt.Println(priceArray2)
		fmt.Println(itemChosen1)
		// fmt.Println(itemChosen2)
		fmt.Printf("Input: %s \n", input)
		fmt.Println("------------------------")
	}

	indices, err := backtrack(itemList, priceArray1, itemChosen1, numItems)
	return itemList, indices, err
}

// getNewlineStr identifies operating system and returns newline character used
func getNewlineStr() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}
	return "\n"
}

// parse string using given separator
func parse(input string) []string {
	s := strings.Split(input, ", ")
	return s
}

// addMap adds values from in2 to in1 by matching keys.
func addMap(in1 map[int]int, in2 map[int]int) map[int]int {
	for key, val := range in2 {
		in1[key] = in1[key] + val
	}
	return in1
}

// mergeMap concatenates in2 to in1. Overlapping key values in in1 are overwritten with key values from in2.
func mergeMap(in1 map[int]int, in2 map[int]int) {
	for key, val := range in2 {
		in1[key] = val
	}
}

// computeMap (i) returns a new map[int]int{1:val} if in1 is nil map, or
// (ii) returns in1 where each key is incremented by 1 and assigned a value val, provided the key is smaller than maxKey.
func computeMap(in1 map[int]int, val int, maxKey int) map[int]int {
	out1 := make(map[int]int)
	if len(in1) == 0 {
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

/*
backtrack returns indices of selected items from itemList. Number of selected items equals numItems.
Error is returned if no combination exists.

itemList: lists the items.

priceArray: lists the maximum achievable price for each smaller pricepoint.

itemChosen: lists the last chosen items for each smaller pricepoint.

numItems: is the desired number of total items to be selected.
*/
func backtrack(itemList []item, priceArray []int, itemChosen []map[int]int, numItems int) ([]int, error) {
	col := len(itemChosen) - 1
	var indices []int
	for ii := numItems; ii > 0; ii-- {
		if elem, ok := itemChosen[col][ii]; ok {
			col = priceArray[col] - itemList[elem].price
			indices = append(indices, elem)
		} else {
			return indices, errors.New("No combination of items available")
		}
	}
	return indices, nil
}

// printer pretty prints, to out stream, the items from list as given by indices in reverse order.
func printer(list []item, indices []int, out io.Writer) {
	for ii := len(indices) - 1; ii >= 1; ii-- {
		fmt.Fprintf(out, "%v %v, ", list[indices[ii]].name, list[indices[ii]].price)
	}
	fmt.Fprintf(out, "%v %v", list[indices[0]].name, list[indices[0]].price)
}
