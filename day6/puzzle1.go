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
  var m map[byte]byte
  m = make(map[byte]byte)

  for i := 0; i < len(text); i++ {

    if len(text[i]) != 0 {
      for j := 0; j < len(text[i]); j++{
        _, exists := m[text[i][j]]
        if !exists{
          m[text[i][j]] = text[i][j]
        }
      }
    }

    if len(text[i]) == 0 || i == len(text) - 1 {
      answer = answer + len(m)
      m = make(map[byte]byte)
    }

  }

  fmt.Println("Answer:", answer)
}
