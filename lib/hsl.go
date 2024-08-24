package lib

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type HSL struct {
	H float64
	S float64
	L float64
}

func rmFraction(s string) string {
	return strings.Split(s, ".")[0]
}

func toInt(s string) (int, error) {
	i, err := strconv.Atoi(rmFraction(s))
	if err != nil {
		fmt.Println("error: ", err)
		return 0, err
	}
	return i, nil
}

func toFloat(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println("error: ", err)
		return 0, err
	}
	return f, nil
}

func format(s string) string {
	if len(s) > 0 && s[0] == ' ' {
		return format(s[1:])
	}
	return s
}

func ParseIntoHSL(s string) HSL {
	// example string " 210 40% 98%"
	str := format(s)
	split := strings.Split(str, " ")
	H, err := toFloat(split[0])
	if err != nil {
		fmt.Println("error: ", err)
	}
	S, err := toFloat(strings.Trim(split[1], "%"))
	if err != nil {
		fmt.Println("error: ", err)
	}
	L, err := toFloat(strings.Trim(split[2], "%"))
	if err != nil {
		fmt.Println("error: ", err)
	}
	return HSL{H, S, L}
}

func (hsl HSL) ToRGB() (r, g, b float64) {
	h := hsl.H
	s := hsl.S / 100
	l := hsl.L
	c := (1 - math.Abs(2*l-1)) * s
	x := c * (1 - math.Abs(math.Mod(h/60, 2)-1))
	m := l - c/2

	switch {
	case 0 <= h && h < 60:
		r, g, b = c, x, 0
	case 60 <= h && h < 120:
		r, g, b = x, c, 0
	case 120 <= h && h < 180:
		r, g, b = 0, c, x
	case 180 <= h && h < 240:
		r, g, b = 0, x, c
	case 240 <= h && h < 300:
		r, g, b = x, 0, c
	case 300 <= h && h < 360:
		r, g, b = c, 0, x
	}

	return r + m, g + m, b + m
}

func round_int(f float64) int {
	return int(math.Round(f))
}

func (hsl HSL) ToHex() string {
	r, g, b := hsl.ToRGB()

	rInt := round_int(r)
	gInt := round_int(g)
	bInt := round_int(b)
	return fmt.Sprintf("#%02x%02x%02x", rInt, gInt, bInt)
}
