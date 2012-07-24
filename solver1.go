package main

import (
        "os"
        "./sudoku"
        "fmt"
)

func usage() {
    fmt.Fprintf(os.Stderr, "usage: %s [*.Sudoku]\n", os.Args[0])
    os.Exit(2)
}

func main() {
    if len(os.Args) != 2 {
        usage()
    }

    fmt.Println("********************");
    fmt.Println("* Solver 1         *");
    fmt.Println("********************");

    game := new(sudoku.Game)
    err := game.FromFile(os.Args[1])

    if err != nil {
        fmt.Printf("Map failed to load, Error: %s\n", err)
    }

    game.Print()
    testNum:=0
    fmt.Printf("Row %d:%t\n",testNum,game.TestRow(testNum))
    fmt.Printf("Column %d: %t\n",testNum,game.TestColumn(testNum))
    fmt.Printf("Sector %d: %t\n",testNum,game.TestSector(testNum))
    fmt.Printf("Is valid: %t\n",game.IsValid())
}
