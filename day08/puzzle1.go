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

    accumulator := 0
    visited := make([]bool, len(text))

    for i := 0; i < len(text) && !visited[i]; i++{
      visited[i] = true
      chunks := strings.Split(text[i], " ")
      number, _ := strconv.Atoi(chunks[1])
      if chunks[0] == "acc"{
        accumulator = accumulator + number
      }
      if chunks[0] == "jmp"{
        i = i + number - 1
      }
    }

    fmt.Println("Answer: ", accumulator)
}
