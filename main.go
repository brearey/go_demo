package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getWeekDayIndex(weekDay string) int {
	var weekDays = []string{"Monday", "Tuesday", "Wednesday", "Thursday" ,"Friday" ,"Saturday" ,"Sunday"}
	var ind int
	for i := 0; i < len(weekDays); i++ {
		if weekDays[i] == weekDay {
			ind = i
			break
		}
	}
	return ind
}

func main() {
	// https://coderun.yandex.ru/problem/calendar-formatting?compiler=go

	// input
	reader := bufio.NewReaderSize(os.Stdin, 1<<4)
  writer := bufio.NewWriterSize(os.Stdout, 1<<4)
  defer writer.Flush()

  line, _ := reader.ReadString('\n')
  input := strings.Fields(line)

	nDays, _ := strconv.Atoi(input[0])
	weekday := input[1]

	// print input for debug
  writer.WriteString(strconv.Itoa(nDays * 2))
  writer.WriteByte(' ')
  writer.WriteString(weekday)
  writer.WriteByte('\n')

	// weekdays
	weekDayIndex := getWeekDayIndex(weekday)
	
	// main cycle
	weeksCount := nDays / 7
	dayNumber := 1
	fmt.Printf("weeksCount = %d\n", weeksCount)
	for weekIndex := 0; weekIndex < weeksCount; weekIndex++ {
		// first week
		for i := 0; i < weekDayIndex; i++ {
			fmt.Print(".. ")
		}
		fmt.Print(dayNumber)
		fmt.Println()
	}
}
