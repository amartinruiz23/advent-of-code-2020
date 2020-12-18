package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    //"strings"
    //"strconv"
    //"sort"
)

func applyRules(text [][]byte) ([][]byte, bool){
  madeChanges := false
  newText := make([][]byte, len(text))
  for i := range newText {
    newText[i] = make([]byte, len(text))
  }
  for i := range text{
    for j := range text[i]{
      newText[i][j] = text[i][j]
    }
  }

  adjacentOccupied := 0
  right := false
  left := false
  up := false
  down := false

  for i := 0; i < len(text); i++{
    for j:= 0; j < len(text[i]); j++{
      if text[i][j] != '.'{
        if i == 0{
          up = true
        }
        if i == len(text) - 1{
          down = true
        }
        if j == 0{
          left = true
        }
        if j == len(text[i]) -1 {
          right = true
        }

        if !up{
          if text[i-1][j] == '#'{
            adjacentOccupied = adjacentOccupied + 1
          }
          if !left{
            if text[i-1][j-1] == '#'{
              adjacentOccupied = adjacentOccupied + 1
            }
          }
          if !right{
            if text[i-1][j+1] == '#'{
              adjacentOccupied = adjacentOccupied + 1
            }
          }
        }

        if !down{
          if text[i+1][j] == '#'{
            adjacentOccupied = adjacentOccupied + 1
          }
          if !left{
            if text[i+1][j-1] == '#'{
              adjacentOccupied = adjacentOccupied + 1
            }
          }
          if !right{
            if text[i+1][j+1] == '#'{
              adjacentOccupied = adjacentOccupied + 1
            }
          }
        }

        if !left{
          if text[i][j-1] == '#'{
            adjacentOccupied = adjacentOccupied + 1
          }
        }

        if !right{
          if text[i][j+1] == '#'{
            adjacentOccupied = adjacentOccupied +   1
          }
        }

        if (adjacentOccupied == 0 && text[i][j] == 'L'){
          madeChanges = true
          newText[i][j] = '#'
        }

        if (adjacentOccupied >= 4 && text[i][j] == '#'){
          madeChanges = true
          newText[i][j] = 'L'
        }
        adjacentOccupied = 0
        right = false
        left = false
        up = false
        down = false
      }
    }
  }

  /*fmt.Println("newText")
  for i:= 0; i < len(newText); i++{
    fmt.Println(newText[i])
  }*/
  return newText, madeChanges
}

func main() {

    file, err := os.Open("input2")

    if err != nil {
        log.Fatalf("failed to open")
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    var text []string

    for scanner.Scan() {
        text = append(text, scanner.Text())
    }
    file.Close()

    textMatrix := make([][]byte, len(text))
    for i := range textMatrix {
      textMatrix[i] = make([]byte, len(text))
    }

    for i := range textMatrix{
      for j := range textMatrix[i]{
        textMatrix[i][j] = text[i][j]
      }
    }

    //fmt.Println(textMatrix)

    newText, changes := applyRules(textMatrix)

    //fmt.Println(newText)

    for k := 0 ; changes; k++{
      newText, changes = applyRules(newText)
      //fmt.Println(k)
    }

    count := 0
    for i := range newText{
      for j := range newText[i]{
        if newText[i][j] == '#'{
          count = count + 1
        }
      }
    }

    fmt.Println(count)

}
