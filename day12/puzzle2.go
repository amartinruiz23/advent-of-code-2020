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
    waypointEast := 10
    waypointNorth := 1
    var advancedNorth, advancedEast int
    for i := 0; i < len(text); i++ {
      if text[i][0] == 'N'{
        number, _ = strconv.Atoi(strings.TrimPrefix(text[i], "N"))
        waypointNorth = waypointNorth + number
      } else if text[i][0] == 'S'{
        number, _ = strconv.Atoi(strings.TrimPrefix(text[i], "S"))
        waypointNorth = waypointNorth - number
      } else if text[i][0] == 'E'{
        number, _ = strconv.Atoi(strings.TrimPrefix(text[i], "E"))
        waypointEast = waypointEast + number
      } else if text[i][0] == 'W'{
        number, _ = strconv.Atoi(strings.TrimPrefix(text[i], "W"))
        waypointEast = waypointEast - number
      } else if text[i][0] == 'L'{
        number, _ = strconv.Atoi(strings.TrimPrefix(text[i], "L"))
        for ;number >= 90; number = number-90 {
          aux := waypointEast
          waypointEast = -1*waypointNorth
          waypointNorth = aux
        }
      } else if text[i][0] == 'R'{
        number, _ = strconv.Atoi(strings.TrimPrefix(text[i], "R"))
        for ;number >= 90; number = number-90 {
          aux := waypointEast
          waypointEast = waypointNorth
          waypointNorth = -1*aux
        }
      } else if text[i][0] == 'F'{
        number, _ = strconv.Atoi(strings.TrimPrefix(text[i], "F"))
        advancedEast = advancedEast + waypointEast*number
        advancedNorth = advancedNorth + waypointNorth*number
      }
    }

    if advancedEast < 0 {
      advancedEast = -1*advancedEast
    }
    if advancedNorth < 0 {
      advancedNorth = -1*advancedNorth
    }

    fmt.Println("Answer: ", advancedEast + advancedNorth)

}
