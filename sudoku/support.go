package sudoku

import (
        "fmt"
        "bufio"
        "io"
        "os"
)

type Game struct {
    Map    [81]byte
}

func (game *Game) Print() {
    for i:=0;i<9;i++ {
        for j:=0;j<9;j++ {
            fmt.Printf("%d",game.Map[i*9+j])
        }
        fmt.Printf("\n")
    }
}


//Returns true if there are no more moves to make
func (game *Game) IsFilled() (solved bool) {
    for i:=0;i<len(game.Map);i++ {
        if game.Map[i] == 0 {
            return false
        }
    }
    return true
}

//Returns true if none of the three 1-9 number rules are broken
func (game *Game) IsValid() (valid bool) {
    for row:=0;row<9;row++ {
        if !game.TestRow(row) {
            return false
        }
    }
    for col:=0;col<9;col++ {
        if !game.TestColumn(col) {
            return false
        }
    }
    for sec:=0;sec<9;sec++ {
        if !game.TestSector(sec) {
            return false
        }
    }
    return true
}

func (game *Game) TestRow(row int) (valid bool) {
    used:=0
    for i:=0;i<9;i++ {
        test := game.Map[row*9+i]
        if (used &(1<<test))>>test == 1  {
            return false
        } else {
            //Add the number to the used list, except for 0 which is allowed to be repeated
            if test!=0 {
                used = used+1<<test
            }
        }
    }
    return true
}


func (game *Game) TestColumn(col int) (valid bool) {
    used:=0
    for i:=0;i<9;i++ {
        test := game.Map[i*9+col]
        if (used &(1<<test))>>test == 1  {
            return false
        } else {
            //Add the number to the used list, except for 0 which is allowed to be repeated
            if test!=0 {
                used = used+1<<test
            }
        }
    }
    return true
}


func (game *Game) TestSector(sec int) (valid bool) {
    used:=0
    first:=(sec/3)*27+(sec-(sec/3)*3)*3
    fmt.Printf("First = %d\n",first)
    for i:=0;i<3;i++ {
        for j:=0;j<3;j++ {
            test := game.Map[first+i*9+j]
            if (used &(1<<test))>>test == 1  {
                return false
            } else {
                //Add the number to the used list, except for 0 which is allowed to be repeated
                if test!=0 {
                    used = used+1<<test
                }
            }
        }
    }
    return true
}
func (game *Game) FromFile(name string) (err error) {
    file, err := os.Open(name)
    if err != nil {
        return err
    }
    fileinfo, err := file.Stat()

    r := bufio.NewReaderSize(file, int(fileinfo.Size()))

    game.Load(r)

    return err
}

func (game *Game) Load(r *bufio.Reader) (err error) {
    //data := make(Map, 0,80)

    i := 0
    for ; ; i++ {
        line, err := ReadLine(r)
        if err == io.EOF {
            break
        } else if err != nil {
            fmt.Printf("Error: %s\n", err)
        }

        for j:=0;j<len(line);j++ {
            if i*9+j<len(game.Map) {
            game.Map[i*9+j]=line[j]-48
            } else {
                fmt.Println("Map Overrun")
            }
        }
    }
    return nil
}


// Inspired by Alex Ray
func ReadLine(r *bufio.Reader) ([]byte, error) {
    l := make([]byte, 0, 4096)
    
    for {
        line, isPrefix, err := r.ReadLine()

        if err != nil && err != io.EOF {
            return nil, err
        }

        l = append(l, line...)

        if err == io.EOF {
            return l, err
        }
        if !isPrefix {
            break
        }
    }
    return l, nil
}
