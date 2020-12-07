package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    //"strconv"
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

  correct := 0
  byr := false
  iyr := false
  eyr := false
  hgt := false
  hcl := false
  ecl := false
  pid := false

  for i := 0; i < len(text); i++ {

    if len(text[i]) != 0 {
      if strings.Contains(text[i], "byr"){
        byr = true
      }
      if strings.Contains(text[i], "iyr"){
        iyr = true
      }
      if strings.Contains(text[i], "eyr"){
        eyr = true
      }
      if strings.Contains(text[i], "hgt"){
        hgt = true
      }
      if strings.Contains(text[i], "hcl"){
        hcl = true
      }
      if strings.Contains(text[i], "ecl"){
        ecl = true
      }
      if strings.Contains(text[i], "pid"){
        pid = true
      }
    }

    if len(text[i]) == 0 || i == len(text) - 1 {
      if (byr && iyr && eyr && hgt && hcl && ecl && pid){
        correct = correct + 1
        //fmt.Println("Line", i)
      }
      byr = false
      iyr = false
      eyr = false
      hgt = false
      hcl = false
      ecl = false
      pid = false
    }

  }

  fmt.Println("Answer:", correct)
}
