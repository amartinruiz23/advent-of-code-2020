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

var m map[int]int
var m1 map[int]int

func countConexions(number int) int{

  value, exists := m1[number]

  if exists{
    return value
  }else{

    conexions := 0
    var exists1, exists2, exists3 bool

    if number == 1 || number == 2 || number == 3{
      conexions = conexions + 1
    }

    _, exists1 = m[number - 1]
    _, exists2 = m[number - 2]
    _, exists3 = m[number - 3]

    if exists1{
      conexions = conexions + countConexions(number-1)
    }
    if exists2{
      conexions = conexions + countConexions(number-2)
    }
    if exists3{
      conexions = conexions + countConexions(number-3)
    }

    m1[number] = conexions
    return conexions
  }
}

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

    allNumbers := make([]int, len(text))
    var number int
    for i := 0; i < len(text); i++{
      number, _ = strconv.Atoi(text[i])
      allNumbers[i] = number
    }

    sort.Ints(allNumbers[:])

    m = make(map[int]int)
    m1 = make(map[int]int)

    for  i := 0; i < len(allNumbers); i++{
      m[allNumbers[i]] = allNumbers[i]
    }

    fmt.Println("Answer: ", countConexions(allNumbers[len(allNumbers)-1]))
}
