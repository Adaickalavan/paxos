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

func main() {
	// Log errors
	outFile, err := os.OpenFile("error.log", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	mw := io.MultiWriter(outFile, outStream)
	log.SetOutput(mw)

	// Read command line arguments
	args := len(os.Args)
	if args > 4 || args < 3 {
		log.Fatal("Unknown command line input. ", err)
	}
	maxPrice, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("Unknown maximum price of items. ", err)
	}
	var maxItems int
	if args == 4 {
		maxItems, err = strconv.Atoi(os.Args[3])
		if err != nil {
			log.Fatal("Unknown number of items to select. ", err)
		}
	} else {
		maxItems = 2
	}
	if maxPrice <= 0 || maxItems <= 0 {
		fmt.Fprint(outStream, "Not possible")
		return
	}

	// Prepare to read file
	var scanner *bufio.Scanner
	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()
	scanner = bufio.NewScanner(inputFile)

	// Get combination of items
	itemList, indices, err := getItems(scanner, maxPrice, maxItems)

	// Display result
	if err != nil {
		fmt.Fprint(outStream, err.Error())
		return
	}
	printer(itemList, indices, outStream)
}

type item struct {
	name  string
	price int
}

type path struct {
	total      int
	itemPicked int
}

// getItems finds combination of items satisfying the item-count and maximum-price constraints
func getItems(scanner *bufio.Scanner, maxPrice int, maxItems int) ([]item, []int, error) {
	// Initialize variables
	pathArray1 := make([][]path, maxPrice+1)
	pathArray2 := make([][]path, maxPrice+1)
	for price := 0; price <= maxPrice; price++ {
		pathArray1[price] = make([]path, maxItems+1)
		pathArray2[price] = make([]path, maxItems+1)
	}
	var itemList []item

	newlineStr := getNewlineStr()
	itemIndex := -1
	for scanner.Scan() {
		// Read items
		input := scanner.Text()
		input = strings.TrimRight(input, newlineStr)
		s := parse(input)
		itemPrice, err := strconv.Atoi(s[1])
		if err != nil {
			log.Fatalf("Unknown item format in input file: %s", input)
		}
		itemIndex++

		// Store items in itemList
		itemList = append(itemList, item{name: s[0], price: itemPrice})

		// Perform dynamic programming by iterating over each sub-price and sub-itemCount
		for price := 0; price <= maxPrice; price++ {
			for itemCount := 1; itemCount <= maxItems; itemCount++ {
				if itemPrice <= price {
					prevPath := pathArray1[price][itemCount]
					subPath := pathArray1[price-itemPrice][itemCount-1]
					newTot := itemPrice + subPath.total
					if newTot > prevPath.total {
						pathArray2[price][itemCount] = path{total: newTot, itemPicked: itemIndex}
					} else {
						pathArray2[price][itemCount] = pathArray1[price][itemCount]
					}
				} else {
					pathArray2[price][itemCount] = pathArray1[price][itemCount]
				}
			}
		}

		// Swap the arrays
		pathArray1, pathArray2 = pathArray2, pathArray1
	}

	fmt.Println("finished")

	// Backtrack to get chosen items' indices
	indices, err := backtrack(itemList, pathArray1, maxItems)

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

// backtrack returns indices of selected items from itemList. Number of selected items equals maxItems.
// Error is returned if no combination exists.
func backtrack(itemList []item, pathArray [][]path, maxItems int) ([]int, error) {
	col := len(pathArray) - 1
	var indices []int
	for itemCount := maxItems; itemCount > 0; itemCount-- {
		curPath := pathArray[col][itemCount]
		indices = append(indices, curPath.itemPicked)
		col = curPath.total - itemList[curPath.itemPicked].price
		if col < 0 {
			return indices, errors.New("Not possible")
		}
	}
	if col != 0 {
		return indices, errors.New("Not possible")
	}
	return indices, nil
}

// printer pretty prints, to out stream, the items from list as given by indices in increasing price order.
func printer(list []item, indices []int, out io.Writer) {
	for ii := len(indices) - 1; ii >= 1; ii-- {
		fmt.Fprintf(out, "%v %v, ", list[indices[ii]].name, list[indices[ii]].price)
	}
	fmt.Fprintf(out, "%v %v", list[indices[0]].name, list[indices[0]].price)
}
