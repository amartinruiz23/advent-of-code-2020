
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
      appears := 0
      if letter[0] == password[first_number-1]{
        appears = appears + 1
      }
      if letter[0] == password[second_number-1]{
        appears = appears + 1
      }
      if appears == 1{
        correct = correct + 1
      }
    }

    fmt.Println("Number of correct passwords:", correct)
}
