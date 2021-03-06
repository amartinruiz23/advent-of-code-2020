package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    //"strings"
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
    found := false
    sums := false
    var number int

    for i := 25; i < len(text) && !found; i++{
      m = make(map[int]int)
      number, _ = strconv.Atoi(text[i])
      for j := i - 25; j < i; j++{
        otherNumber, _ := strconv.Atoi(text[j])
        m[otherNumber] = otherNumber
      }
      for _, v := range m{
        _, exists := m[number - v]
        if exists{
          sums = true
        }
      }
      if (!sums){
        found = true
      }else{
        sums = false
      }
    }

    if found{
      fmt.Println("Answer: ", number)
    }else{
      fmt.Println("No answer found")
    }
}
