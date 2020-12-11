package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    //"strings"
    //"strconv"
)

func main() {

  file, err := os.Open("input")

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

  answer := 0
  group := 0
  var m map[byte]int
  m = make(map[byte]int)

  for i := 0; i < len(text); i++ {

    if len(text[i]) != 0 {
      group = group + 1
      for j := 0; j < len(text[i]); j++{
        _, exists := m[text[i][j]]
        if !exists{
          m[text[i][j]] = 1
        }else{
          m[text[i][j]] = m[text[i][j]] + 1
        }
      }
    }

    if len(text[i]) == 0 || i == len(text) - 1 {
      selectedByAll := 0
      for _, number := range m {
        if number == group{
          selectedByAll = selectedByAll + 1
        }
      }
      answer = answer + selectedByAll
      m = make(map[byte]int)
      group = 0
    }

  }

  fmt.Println("Answer:", answer)
}
