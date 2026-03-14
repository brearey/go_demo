package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

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

	// save into the map
	var weekDays
}
