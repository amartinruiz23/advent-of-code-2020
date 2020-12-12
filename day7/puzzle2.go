package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

type pair struct{
  colour string
  number string
}

func countSons(colour string, m map[string]map[pair]pair) int{
  sons := 0
  m1, exists := m[colour]
  _, isempty := m1[pair{colour:"noother", number:"0"}]

  if !exists || isempty{
    return 0
  }else{
    for _, v := range m1{
      colourName := v.colour
      grandSons := countSons(colourName, m)
      convertedNumber, _ := strconv.Atoi(v.number)
      sons = sons + convertedNumber + convertedNumber*grandSons
    }
  }
  return sons
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

  var m = map[string]map[pair]pair{}

  for i := 0; i < len(text); i++ {
    chunks1 := strings.Split(text[i], "contain")
    chunks2 := strings.Split(chunks1[0], " ")
    father :=  chunks2[0] + chunks2[1]
    if chunks1[1] == " no other bags."{
      var m1 = map[pair]pair{}
      m1[pair{colour:"noother", number:"0"}] = pair{colour:"noother", number:"0"}
      m[father] = m1
    }else{
      chunks3 := strings.Split(chunks1[1], ",")
      var m1 = map[pair]pair{}
      for j:= 0; j < len(chunks3); j++{
        chunks4 := strings.Split(chunks3[j], " ")
        m1[pair{colour: chunks4[2] + chunks4[3], number:chunks4[1]}] = pair{colour: chunks4[2] + chunks4[3], number:chunks4[1]}
      }
      _, exists := m[father]
      if !exists{
        m[father] = m1
      } else {
        for k, v := range m1 {
          m[father][k] = v
        }
      }
    }
  }

  fmt.Println("Answer: ", countSons("shinygold", m))

}
