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

    var number int
    direction := 'E'
    var advancedNorth, advancedEast int
    for i := 0; i < len(text); i++{
      if text[i][0] == 'N'{
        number, _ = strconv.Atoi(strings.TrimPrefix(text[i], "N"))
        advancedNorth = advancedNorth + number
      } else if text[i][0] == 'S'{
        number, _ = strconv.Atoi(strings.TrimPrefix(text[i], "S"))
        advancedNorth = advancedNorth - number
      } else if text[i][0] == 'E'{
        number, _ = strconv.Atoi(strings.TrimPrefix(text[i], "E"))
        advancedEast = advancedEast + number
      } else if text[i][0] == 'W'{
        number, _ = strconv.Atoi(strings.TrimPrefix(text[i], "W"))
        advancedEast = advancedEast - number
      } else if text[i][0] == 'L'{
        number, _ = strconv.Atoi(strings.TrimPrefix(text[i], "L"))
        for ;number >= 90; number = number-90 {
          if direction == 'E'{
            direction = 'N'
          } else if direction == 'N'{
            direction = 'W'
          } else if direction == 'W'{
            direction = 'S'
          } else if direction == 'S'{
            direction = 'E'
          }
        }
      } else if text[i][0] == 'R'{
        number, _ = strconv.Atoi(strings.TrimPrefix(text[i], "R"))
        for ;number >= 90; number = number-90 {
          if direction == 'E'{
            direction = 'S'
          } else if direction == 'N'{
            direction = 'E'
          } else if direction == 'W'{
            direction = 'N'
          } else if direction == 'S'{
            direction = 'W'
          }
        }
      } else if text[i][0] == 'F'{
        number, _ = strconv.Atoi(strings.TrimPrefix(text[i], "F"))
        if direction == 'E'{
          advancedEast = advancedEast + number
        } else if direction == 'N'{
          advancedNorth = advancedNorth + number
        } else if direction == 'W'{
          advancedEast = advancedEast - number
        } else if direction == 'S'{
          advancedNorth = advancedNorth - number
        }
      }
    }

    if advancedEast < 0{
      advancedEast = -1*advancedEast
    }
    if advancedNorth < 0{
      advancedNorth = -1*advancedNorth
    }

    fmt.Println("Answer: ", advancedEast + advancedNorth)

}
