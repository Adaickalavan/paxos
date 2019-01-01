# Find Items

## Instructions

1. **Source code**
    + Unzip `paxos.zip` into `$GOPATH/src/paxos` folder in your computer

2. **Executable**
    + To create an executable, navigate to `$GOPATH/src/paxos/FindItems` directory and execute
        ```go
        go install FindItems
        ```

3. **Functional test**
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

4. **Running**
    + Launch by executing
        ```go
        $GOPATH/bin/FindItems.exe $GOPATH/src/paxos/FindItems/prices.txt maxPrice maxItems
        ```
        Here,
        + `$GOPATH/src/paxos/FindItems/prices.txt` refers to the input file with complete path
        + `maxPrice` refers to the balance in the gift card
        + `maxItems` refers to the desired number of items to be selected
    + Several example commands
        ```go
        $GOPATH/bin/FindItems.exe $GOPATH/src/paxos/FindItems/prices.txt 2300 //When maxItems is omitted, it defaults to 2
        ```
        ```go
        $GOPATH/bin/FindItems.exe $GOPATH/src/paxos/FindItems/prices.txt 2100 2
        ```
        ```go
        $GOPATH/bin/FindItems.exe $GOPATH/src/paxos/FindItems/prices.txt 2200 3 //maxItems can be any positive integer
        ```

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

1. **Complexity**
    + O(itemNum * maxPrice * maxItems).
    + itemNum is the total number of items available.
    + maxPrice is the balance of the gift card.
    + maxItems is the desired number of items to be selected.

1. **Algorithm and data structure**
    + In classic Knapsack problem, there is only one constraint and thus a 2 dimensional matrix is used to track the unconstrained variable. Since there are 2 constrains (total price and number of items) in our problem, the traditional solution is extended to a 3 dimensional matrix, where the second and third dimension models the constraint on price and number of items, respectively.
    + Only the last 2 rows of the dynamic programming matrix is needed and thus kept in memory, to reduce memory consumption.
    + The solution will work for any positive integer `maxItems` of items to be selected.