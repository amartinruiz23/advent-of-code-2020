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

  correct := 0
  byr := false
  iyr := false
  eyr := false
  hgt := false
  hcl := false
  ecl := false
  pid := false

  for i := 0; i < len(text); i++ {

    chunks := strings.Split(text[i], " ")

    for j := 0; j < len(chunks); j++ {

      if len(text[i]) != 0 {

        if strings.Contains(chunks[j], "byr"){
          chunks2 := strings.Split(chunks[j], ":")
          number, _ := strconv.Atoi(chunks2[1])
          if number >= 1920 && number <= 2002{
            byr = true
          }
        }

        if strings.Contains(chunks[j], "iyr"){
          chunks2 := strings.Split(chunks[j], ":")
          number, _ := strconv.Atoi(chunks2[1])
          if  number >= 2010 && number <= 2020{
            iyr = true
          }
        }

        if strings.Contains(chunks[j], "eyr"){
          chunks2 := strings.Split(chunks[j], ":")
          number, _ := strconv.Atoi(chunks2[1])
          if  number >= 2020 && number <= 2030{
            eyr = true
          }
        }

        if strings.Contains(chunks[j], "hgt"){
          chunks2 := strings.Split(chunks[j], ":")
          if strings.HasSuffix(chunks2[1], "cm"){
            numberStr := strings.Replace(chunks2[1], "cm", "", 1)
            number, _ := strconv.Atoi(numberStr)
            if number >= 150 && number <= 193{
              hgt = true
            }
          }
          if strings.HasSuffix(chunks2[1], "in"){
            numberStr := strings.Replace(chunks2[1], "in", "", 1)
            number, _ := strconv.Atoi(numberStr)
            if number >= 59 && number <= 76{
              hgt = true
            }
          }
        }

        if strings.Contains(chunks[j], "hcl"){
          chunks2 := strings.Split(chunks[j], ":")
          correct := true
          if chunks2[1][0] != '#'{
            correct = false
          }
          for k := 1; k < len(chunks2[1]); k++{
            letter := chunks2[1][k] >= 'a' && chunks2[1][k] <= 'f'
            number := chunks2[1][k] >= '0' && chunks2[1][k] <= '9'
            if !letter && !number {
              correct = false
            }
          }
          if correct{
            hcl = true
          }
        }

        if strings.Contains(chunks[j], "ecl"){
          chunks2 := strings.Split(chunks[j], ":")
          if chunks2[1] == "amb" || chunks2[1] == "blu" ||
             chunks2[1] == "brn" || chunks2[1] == "gry" ||
             chunks2[1] == "hzl" || chunks2[1] == "oth" ||
             chunks2[1] == "grn" {
               ecl = true
          }
        }

        if strings.Contains(chunks[j], "pid"){
          chunks2 := strings.Split(chunks[j], ":")
          correct := false
          if len(chunks2[1]) == 9 {
            correct = true
            for k := 0; k < 9; k++{
              if chunks2[1][k] < '0' || chunks2[1][k] > '9'{
                correct = false
              }
            }
          }
          if correct{
            pid = true
          }
        }
      }

      if len(text[i]) == 0 || (i == len(text) - 1 && j == len(chunks) - 1) {
        if (byr && iyr && eyr && hgt && hcl && ecl && pid){
          correct = correct + 1
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

  }

  fmt.Println("Answer:", correct)
}
