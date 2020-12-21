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

func setBit(n int, pos int) int {
    n |= (1 << pos)
    return n
}

func clearBit(n int, pos int) int {
    mask := ^(1 << pos)
    n &= mask
    return n
}

func applyMask(n int, mask string) int{
  number := n

  for i := len(mask) -1 ; i >= 0; i--{
    if mask[i] == '0'{
      number = clearBit(number, len(mask) -1 - i)
    }
    if mask[i] == '1'{
      number = setBit(number, len(mask) -1 - i)
    }
  }

  return number
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

    mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
    var memory map[int]int
    memory = make(map[int]int)
    var number, position int

    for i := 0; i < len(text); i++{
      if strings.HasPrefix(text[i], "mask"){
        chunks := strings.Split(text[i], " ")
        mask = chunks[2]
      }
      if strings.HasPrefix(text[i], "mem"){
        chunks := strings.Split(text[i], " ")
        chunks2 := strings.Split(chunks[0], "[")
        chunks2[1] = strings.TrimSuffix(chunks2[1], "]")
        position, _ = strconv.Atoi(chunks2[1])
        number, _ = strconv.Atoi(chunks[2])
        number = applyMask(number, mask)
        memory[position] = number
      }
    }

    sum := 0
    for _, s := range memory{
      sum = sum + s
    }

    fmt.Println("Answer: ", sum)

}
