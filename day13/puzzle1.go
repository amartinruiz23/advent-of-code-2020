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

    arrivalTime, _ := strconv.Atoi(text[0])
    earliestBusID := 0
    earliestBusTime := 999999
    elapsedTime := 0
    currentBusID := 0

    buses := strings.Split(text[1], ",")

    for i := 0; i < len(buses); i++{
      if buses[i] != "x"{
        currentBusID, _ = strconv.Atoi(buses[i])
        elapsedTime = ((arrivalTime/currentBusID)*currentBusID + currentBusID) - arrivalTime
        if elapsedTime < earliestBusTime{
          earliestBusTime = elapsedTime
          earliestBusID = currentBusID
        }
      }
    }

    fmt.Println("Answer: ", earliestBusID*earliestBusTime)

}
