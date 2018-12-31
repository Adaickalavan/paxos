# Find Items

## Instructions
  
1. **Executable**
    + To create an executable in the `$GOPATH/bin/` directory, execute
        ```go
        go install FindItems
        ```
2. **Unit test and functional test**
    + To run complete test suite, run
        ```go
        go test -v FindItems
        ```
        Here, `-v` is the verbose command flag.
    + To run specific test, run
        ```go
        go test -v FindItems -run xxx
        ```
        Here, `xxx` is the name of test function.
    + Test coverage: 88.9% of statements
3. **Running**
    + Launch file input mode by executing
        ```go
        $GOPATH/bin/parking_lot.exe $GOPATH/src/parking_lot/inputFile.txt
        ```
        Here, `$GOPATH/src/parking_lot/inputFile.txt` refers to the input file with complete path.

## Project structure

The project structure is as follows:

```txt
FindItems             # main folder
├── error.log         # contains errors generated during runtime
├── main_test.go      # functional test of the main code
├── main.go           # main file of Go code
└── prices.txt        # sample input file for testing
```

## Notes on solution

1. **Assumptions**
   + All items have integer prices > 0 cents. In other words, no item is free.

2. **Complexity**
    + O(itemNum * maxPrice * maxItems).
    + itemNum is the total number of items available.
    + maxPrice is the balance of the gift card.
    + maxItem is the desired number of items to be selected.

3. **Algorithm and data structure**
    + A modified dynamic programming solution to Knapsack problem is used. In classic Knapsack problem, there is only one constraint and thus a 2 dimensional matrix is used to track the progression. In our problem, 
    + 