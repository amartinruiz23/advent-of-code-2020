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

  highest := 0
  for i := 0; i < len(text) && text[i] != ""; i++ {

    text[i] = strings.ReplaceAll(text[i], "F", "0")
    text[i] = strings.ReplaceAll(text[i], "B", "1")
    text[i] = strings.ReplaceAll(text[i], "L", "0")
    text[i] = strings.ReplaceAll(text[i], "R", "1")

    integer, _ := strconv.ParseInt(text[i], 2, 0)
    seatID := int(integer)

    if seatID > highest{
      highest = seatID
    }
  }
  fmt.Println("Highest ID:", highest)
}
