package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func main() {

		file, err := os.Open("input")

		if err != nil {
			log.Fatalf("failed to open")
		}

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var input []int

		for scanner.Scan() {
			newNumber, _ := strconv.Atoi(scanner.Text())
			input = append(input, newNumber)
		}

		file.Close()

    var m map[int]int
    m = make(map[int]int)

    for _, s := range input {
      m[s] = s
    }

    var found bool
    for i := 0; i < len(input) && !found; i++ {
      s := input[i]
      value, exists := m[2020-s]
      if exists{
        found = true
        fmt.Println(s*value)
      }
    }

    if !found{
      fmt.Println("No answer found")
    }
}
