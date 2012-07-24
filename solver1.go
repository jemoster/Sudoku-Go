package main

import (
        "os"
        "./sudoku"
        "fmt"
        "container/list"
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
    fmt.Printf("First Zero: %d\n",game.GetFirstOpen())

    mapQ := list.New()
    mapQ.PushBack(game)

    var counter = 0
    for i:=1;i<20;i++ {
        fmt.Printf("Array Size: %d\n",mapQ.Len())
        for e:= mapQ.Front(); e!= nil; {
            tmpGame ,ok := e.Value.(*sudoku.Game)
            if ok {
                k:=tmpGame.GetFirstOpen()
                for j:=1;j<=9;j++ {
                    newGame := new(sudoku.Game)
                    *newGame = *tmpGame
                    newGame.Map[k]=byte(j)
                    if newGame.IsValid() {
                        counter++
                        mapQ.PushBack(newGame)
                        if newGame.IsFilled() {
                            goto solved
                        }
                    }
                }
                e = e.Next();
                if e!=nil {
                    mapQ.Remove(e.Prev())
                }
            }
        }
    }

solved:
    tmpGame, ok := mapQ.Back().Value.(*sudoku.Game)
    if ok {
        tmpGame.Print()
        fmt.Printf("Tried %d combinations\n",counter)
    }
}
