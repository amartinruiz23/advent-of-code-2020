package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    //"strconv"
)

func findParents(colour string, m map[string]map[string]string) map[string]string{
  parents := map[string]string{}
  m1, exists := m[colour]

  if !exists{
    return parents
  }else{
    for _, v := range m1{
      grandParents := findParents(v, m)
      parents[v] = v
      for _, v1 := range grandParents{
        parents[v1] = v1
      }
    }
  }
  return parents
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

  var m = map[string]map[string]string{}

  for i := 0; i < len(text); i++ {
    chunks1 := strings.Split(text[i], "contain")
    chunks2 := strings.Split(chunks1[0], " ")
    father :=  chunks2[0] + chunks2[1]
    if chunks1[1] != " no other bags."{
      chunks3 := strings.Split(chunks1[1], ",")
      for j:= 0; j < len(chunks3); j++{
        chunks4 := strings.Split(chunks3[j], " ")
        _, exists := m[chunks4[2] + chunks4[3]]
        if !exists{
          var m1 = map[string]string{}
          m1[father] = father
          m[chunks4[2] + chunks4[3]] = m1
        }else{
          m[chunks4[2] + chunks4[3]][father] = father
        }
      }
    }
  }

  howMany := findParents("shinygold", m)
  fmt.Println("Answer: ", len(howMany))

  /*// Prints the complete map
  for k, v := range m{
    fmt.Println("Son: ", k)
    for _, v1 := range v{
      fmt.Println(v1)
    }
  }*/

}
