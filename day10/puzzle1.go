package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    //"strings"
    "strconv"
    "sort"
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

    allNumbers := make([]int, len(text)+1)
    var number int
    allNumbers[0] = 0
    for i := 0; i < len(text); i++{
      number, _ = strconv.Atoi(text[i])
      allNumbers[i+1] = number
    }

    sort.Ints(allNumbers[:])

    oneDifference := 0
    threeDifference := 0
    var dif int

    for i:= 0; i < len(allNumbers); i++{
      if i == len(allNumbers) - 1 {
        threeDifference = threeDifference + 1
      }else{
        dif = allNumbers[i+1] - allNumbers[i]
        if dif == 3{
          threeDifference = threeDifference + 1
        }
        if dif == 1{
          oneDifference = oneDifference + 1
        }
      }
    }

    fmt.Println("Answer: ", oneDifference * threeDifference)

}
