package take

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func realIndex(i, l int) int {
	var res int

	if i < 0 {
		res = l + i
		if res < 0 {
			res = 0
		}
	} else {
		res = i
		if res >= l {
			res = l - 1
		}
	}

	return res
}

func Process(from, to int, delimiter string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Split(strings.TrimSpace(line), delimiter)
		l := len(columns)

		start := realIndex(from, l)
		end := realIndex(to, l)
		if start <= end {
			fmt.Println(strings.Join(columns[start:end+1], delimiter))
		} else {
			fmt.Println()
		}
	}
}
