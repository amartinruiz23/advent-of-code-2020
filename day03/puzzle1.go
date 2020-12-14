package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
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

  trees := 0
  counter := 0

  for _, eachLine := range text {
    if eachLine[counter] == '#'{
      trees = trees + 1
    }
    counter = counter + 3
    if counter >= len(eachLine){
      counter = counter - len(eachLine)
    }
  }

  fmt.Println("Number of trees:", trees)
}
