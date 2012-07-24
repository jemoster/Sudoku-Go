package main

import (
        "os"
        "./sudoku"
        "fmt"
        "container/list"
)

func usage() {
    fmt.Fprintf(os.Stderr, "usage: %s [*.sudoku]\n", os.Args[0])
    os.Exit(2)
}

func main() {
    if len(os.Args) != 2 {
        usage()
    }

    fmt.Println("*********************");
    fmt.Println("*  Sudoku Solver 1  *");
    fmt.Println("*********************");

    game := new(sudoku.Game)
    err := game.FromFile(os.Args[1])

    if err != nil {
        fmt.Printf("Map failed to load, Error: %s\n", err)
    }

    game.Print()

    mapQ := list.New()
    mapQ.PushBack(game)

    var counter = 0
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

solved:
    tmpGame, ok := mapQ.Back().Value.(*sudoku.Game)
    if ok {
        fmt.Printf("Solution:\n")
        tmpGame.Print()
        fmt.Printf("Tried %d combinations\n",counter)
    }
}
