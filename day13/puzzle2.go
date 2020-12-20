package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
    //"sort"
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


    buses := strings.Split(text[1], ",")
    busesCount := 0

    for i := 0; i < len(buses); i++{
      if buses[i] != "x" {
        busesCount = busesCount + 1
      }
    }

    busesInteger := make([]int, busesCount)
    busesPosition := make([]int, busesCount)
    checkedPosition := make([]bool, busesCount)
    checkedPosition[0] = true

    used := 0
    for i := 0; i < len(buses); i++{
      if buses[i] != "x" {
        busesInteger[used], _ = strconv.Atoi(buses[i])
        busesPosition[used] = i
        used = used + 1
      }
    }

    firstDeparture := 0
    currentDeparture := 0
    found := false
    failed := false
    answerTimestamp := 0
    jump := busesInteger[0]

    for i := busesInteger[0]; !found; i+=jump{
      fmt.Println(i)
      for j:=0 ; j < len(busesInteger) && !failed; j++{
        if firstDeparture == 0 {
          firstDeparture = i
        }
        currentDeparture  = busesInteger[j]
        if (firstDeparture + busesPosition[j])%currentDeparture != 0{
          failed = true
        } else if !checkedPosition[j]{
          checkedPosition[j] = true
          jump = jump*busesInteger[j]
        }
      }
      if !failed {
        found = true
        answerTimestamp = firstDeparture
      } else{
        firstDeparture = 0
        failed = false
      }
    }


    fmt.Println("Answer: ", answerTimestamp)

}
