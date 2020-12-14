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

    allNumbers := make([]int, len(text))

    for i := 0; i < len(text); i++{
      number, _ = strconv.Atoi(text[i])
      allNumbers[i] = number
    }

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

    found = false
    lookingFor := number
    smallest := 198620152477517 + 1
    biggest := 0
    counter := 0

    for i := 0; i < len(allNumbers) && !found; i++{
      counter = 0
      for j := 0; counter < lookingFor && i+j < len(allNumbers); j++{
        counter = counter + allNumbers[i+j]
        if allNumbers[i+j] > biggest{
          biggest = allNumbers[i+j]
        }
        if allNumbers[i+j] < smallest{
          smallest = allNumbers[i+j]
        }
      }
      if counter == lookingFor{
        found = true
      }else{
        biggest = 0
        smallest = 198620152477517 + 1
      }
    }

    fmt.Println("Answer: ", smallest + biggest)
}
