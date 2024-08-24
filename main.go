package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"mhirii/hexparse/lib"

	"github.com/crazy3lf/colorconv"
)

func main() {
	lines := readfile()
	out := []string{}
	for _, line := range lines {

		if isColorLine(line) {
			fmt.Println("line: ", line)
			hsl := getHsl(line)
			fmt.Println("hsl: ", hsl)
			hex := toHex(hsl)
			fmt.Println("hex: ", hex)
			line = strings.ReplaceAll(line, hsl, toHex(hsl))
		}

		out = append(out, line)
	}
	fmt.Println(strings.Join(out, "\n"))
}

func readfile() []string {
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Println("error: ", err)
			return lines
		}

		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func toHex(str string) string {
	hsl := lib.ParseIntoHSL(str)
	h, s, l := hsl.H, hsl.S, hsl.L
	color, err := colorconv.HSLToColor(h, s/100, l/100)
	if err != nil {
		fmt.Println(err)
	}
	hex := colorconv.ColorToHex(color)
	return strings.ReplaceAll(hex, "0x", "#")
}

func getHsl(s string) string {
	start := 1 + strings.IndexFunc(s, func(r rune) bool {
		return r == ':'
	})

	end := strings.IndexFunc(s, func(r rune) bool {
		return r == ';'
	})

	hex := s[start:end]
	return hex
}

func isColorLine(line string) bool {
	matchers := []string{":", "--", "%"}
	for _, m := range matchers {
		if !strings.Contains(line, m) {
			return false
		}
	}

	return true
}
