package main

import (
    "bufio"
    "log"
    "os"
    "strings"
    "strconv"
    "fmt"
)

type Pair struct {
    row int
    column int
}

func main() {
    file, err := os.Open("board.txt")
    defer file.Close()

    if err != nil {
        log.Fatal("Error opening file")
    }

    table := make([]string, 0)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if len(line) != 9 {
            log.Fatal("Row is not 9 numbers long")
        }
        table = append(table, line)
    }

    if len(table) != 9 {
        log.Fatal("Columns is not 9 numbers long")
    }

    for {
        if checkSolved(table) {
            break
        }

        filledsomething := false

        for i := 1; i < 10; i++ {

            emptyCoords := findEmpty(table)
            potentialCoords := make([]Pair, 0)
            
            fmt.Printf("now on number %d\n", i)
            for _, emptyelem := range emptyCoords {
                if !isInRow(rune('0' + i), emptyelem.row, table) &&
                    !isInColumn(byte('0' + i), emptyelem.column, table) &&
                    !isInBlock(byte('0' + i), getBlockNumber(emptyelem.row, emptyelem.column), table) {

                    //fmt.Printf("er=%d ec=%d\n", emptyelem.row, emptyelem.column)

                    potentialCoords = append(potentialCoords, emptyelem)
                }
            }

            for index, potentialelem := range potentialCoords {
                if !containsRow(potentialelem.row, potentialCoords, index, table) ||
                    !containsColumn(potentialelem.column, potentialCoords, index, table) ||
                    !containsBlock(getBlockNumber(potentialelem.row, potentialelem.column), potentialCoords, index, table) {

                        /*fmt.Printf("pr=%d pc=%d\n", potentialelem.row, potentialelem.column)
                        var input string
                        fmt.Scanln(&input)*/

                    replacestring := table[potentialelem.row]
                    table[potentialelem.row] = replacestring[:potentialelem.column] + strconv.Itoa(i) + replacestring[potentialelem.column+1:]
                    filledsomething = true
                }
            }
        }

        if !filledsomething {
            break
        }
    }

    printTable(table)
}

func printTable(table []string) {
    for _, elem := range table {
        fmt.Println(elem)
    }
}

func containsRow(rowtest int, potential []Pair, currentindex int, table []string) bool {
    for index, elem := range potential {
        if elem.row == rowtest && index != currentindex {
            return true
        }
    }
    return false
}

func containsColumn(columntest int, potential []Pair, currentindex int, table []string) bool {
    for index, elem := range potential {
        if elem.column == columntest && index != currentindex {
            return true
        }
    }
    return false
}

func containsBlock(blocknumber int, potential []Pair, currentindex int, table []string) bool {
    initialrow := (blocknumber / 3) * 3
    initialcolumn := (blocknumber % 3) * 3

    for index, elem := range potential {
        if elem.row >= initialrow && elem.row < initialrow+3 && elem.column >= initialcolumn && elem.column < initialcolumn+3 && index != currentindex {
            return true
        }
    }
    return false
}

func findEmpty(table []string) []Pair {
    emptyCoords := make([]Pair, 0)

    for i := 0; i < len(table); i++ {
        for j := 0; j < len(table[i]); j++ {
            if table[i][j] == '0' {
                emptyCoords = append(emptyCoords, Pair{i, j})
            }
        }
    }

    return emptyCoords
}

func isInRow(num rune, index int, table []string) bool {
    return strings.ContainsRune(table[index], num)
}

func isInColumn(num byte, index int, table []string) bool {
    for i := 0; i < len(table); i++ {
        if table[i][index] == num {
            return true
        }
    }

    return false
}

func getBlockNumber(row int, column int) int {
    return (row / 3) * 3 + (column / 3)
}

func isInBlock(num byte, blocknumber int, table []string) bool {
    initialrow := (blocknumber / 3) * 3
    initialcolumn := (blocknumber % 3) * 3
    
    for i := initialrow; i < initialrow+3; i++ {
        for j := initialcolumn; j < initialcolumn+3; j++ {
            if table[i][j] == num {
                return true
            }
        }
    }

    return false
}

func checkSolved(table []string) bool {
    for i := 0; i < len(table); i++ {
        for j := 0; j < len(table[i]); j++ {
            if table[i][j] == '0' {
                return false
            }
        }
    }

    return true
}
