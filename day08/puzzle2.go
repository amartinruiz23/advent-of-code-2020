package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

func execute(code []string)(int, bool){
  accumulator := 0
  visited := make([]bool, len(code))
  finished := false

  i := 0
  for ; i < len(code) && !visited[i]; i++{
    visited[i] = true
    chunks := strings.Split(code[i], " ")
    number, _ := strconv.Atoi(chunks[1])
    if chunks[0] == "acc"{
      accumulator = accumulator + number
    }
    if chunks[0] == "jmp"{
      i = i + number - 1
    }
  }

  if i == len(code){
    finished = true
  }

  return accumulator, finished
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

    accumulator := 0
    finished := false

    textCopy := make([]string, len(text))
    for i := 0; i < len(text) && !finished; i++{
      copy(textCopy, text)
      chunks := strings.Split(text[i], " ")
      if chunks[0] == "nop"{
        textCopy[i] = "jmp " + chunks[1]
        accumulator, finished = execute(textCopy)
      }
      if chunks[0] == "jmp"{
        textCopy[i] = "nop " + chunks[1]
        accumulator, finished = execute(textCopy)
      }
    }

    fmt.Println("Answer: ", accumulator)
}
