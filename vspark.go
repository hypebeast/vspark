package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	characterList   = []string{"▏", "▎", "▍", "▌", "▋", "▊", "▉", "█"}
	GRAPH_SIZE      = 1
	CUSTOM_CHAR     = ""
	DISPLAY_NUMBERS = false
)

func init() {
	if s := os.Getenv("GRAPH_SIZE"); s != "" {
		GRAPH_SIZE, _ = strconv.Atoi(s)
	}

	if s := os.Getenv("CUSTOM_CHAR"); s != "" {
		CUSTOM_CHAR = s
	}

	if s := os.Getenv("DISPLAY_NUMBERS"); s != "" {
		DISPLAY_NUMBERS = true
	}
}

func main() {
	var args []string
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if fi.Mode()&os.ModeNamedPipe == 0 {
		// Print the help text if vspark was called with no argument
		if len(os.Args) <= 1 {
			usage()
		}

		// Print the help text if the user asks for help
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			usage()
		}

		args = os.Args[1:]
	} else {
		b, _ := ioutil.ReadAll(os.Stdin)
		args = strings.Split(string(b), "\n")
		if args[len(args)-1] == "" {
			args = args[:len(args)-1]
		}
	}

	var list []int
	for _, s := range args {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		list = append(list, i)
	}
	sort.Ints(list)

	for _, arg := range args {
		i, err := strconv.Atoi(arg)
		if err != nil {
			panic(err)
		}
		graphLine := getGraphLine(i, list[len(list)-1])
		fmt.Fprintf(os.Stdout, "%s\n", graphLine)
	}
}

func getGraphLine(number int, max int) string {
	var chars []string
	if CUSTOM_CHAR == "" {
		chars = characterList
	} else {
		chars = append(chars, CUSTOM_CHAR)
	}

	index := (number * len(chars)) / max

	var graphLine string
	if DISPLAY_NUMBERS {
		graphLine += fmt.Sprintf("%*d ", len(strconv.Itoa(max)), number)
	}

	graphLine += strings.Repeat(chars[len(chars)-1], int(index/len(chars)))
	graphLine += chars[int(index%len(chars))]
	return graphLine
}

var helpText = `USAGE:
    $vspark [-h|--help] VALUE1 VALUE2 ... VALUEX

    EXAMPLES:
      vspark 0 30 55 80 33 150
      ▏
      ▎
      ▍
      ▋
      ▎
      █▏
      seq 0 3 | vspark
      ▏
      ▍
      ▊
      █▏
`

func usage() {
	fmt.Fprintf(os.Stderr, "%s", helpText)
	os.Exit(1)
}
