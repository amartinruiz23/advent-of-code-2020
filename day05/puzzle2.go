package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
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

  var m map[int]int
  m = make(map[int]int)

  for i := 0; i < len(text) && text[i] != ""; i++ {

    text[i] = strings.ReplaceAll(text[i], "F", "0")
    text[i] = strings.ReplaceAll(text[i], "B", "1")
    text[i] = strings.ReplaceAll(text[i], "L", "0")
    text[i] = strings.ReplaceAll(text[i], "R", "1")

    integer, _ := strconv.ParseInt(text[i], 2, 0)
    seatID := int(integer)
    m[seatID] = seatID
  }

  for i := 0; i < 1024; i++ {
    _, exists1 := m[i - 1]
    _, exists2 := m[i + 1]
    _, exists3 := m[i]

    if exists1 && exists2 && !exists3 {
      fmt.Println("Your seat has the ID", i)
    }
  }


}
