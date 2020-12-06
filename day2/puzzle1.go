
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

    file, err := os.Open("input2")

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

    correct := 0

    for _, each_line := range text {
      chunks := strings.Split(each_line, " ")
      numbers := strings.Split(chunks[0], "-")
      first_number, _ := strconv.Atoi(numbers[0])
      second_number, _ := strconv.Atoi(numbers[1])
      letter := strings.ReplaceAll(chunks[1], ":", "")
      password := chunks[2]
      count := strings.Count(password, letter)
      if first_number <= count && second_number >= count{
        correct = correct + 1
      }
    }

    fmt.Println("Number of correct passwords:", correct)
}
