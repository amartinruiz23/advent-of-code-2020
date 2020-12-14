package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func treesEncountered(right, down int, input []string) int{
  trees := 0
  counter := 0

  for i := 0; i < len(input); i = i + down {
    if input[i][counter] == '#'{
      trees = trees + 1
    }
    counter = counter + right
    if counter >= len(input[i]){
      counter = counter - len(input[i])
    }
  }
  return trees
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

  answer := treesEncountered(1,1,text)*treesEncountered(3,1,text)*
            treesEncountered(5,1,text)*treesEncountered(7,1,text)*
            treesEncountered(1,2,text)
            
  fmt.Println("Answer:", answer)
}
