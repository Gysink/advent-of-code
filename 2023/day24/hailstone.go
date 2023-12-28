package main

import (
	"advent-of-code/pkg/util"
	"strings"
)

type Hailstone struct {
	sx, sy, sz float64
	vx, vy, vz float64
	a, b, c    float64
}

func NewHailstone(input string) Hailstone {
	sp := strings.Split(input, "@")
	sps := strings.Split(strings.ReplaceAll(sp[0], " ", ""), ",")
	spv := strings.Split(strings.ReplaceAll(sp[1], " ", ""), ",")

	h := Hailstone{
		sx: util.StringToFloat(sps[0]),
		sy: util.StringToFloat(sps[1]),
		sz: util.StringToFloat(sps[2]),
		vx: util.StringToFloat(spv[0]),
		vy: util.StringToFloat(spv[1]),
		vz: util.StringToFloat(spv[2]),
		a:  0,
		b:  0,
		c:  0,
	}

	h.a = h.vy
	h.b = -h.vx
	h.c = h.vy*h.sx - h.vx*h.sy
	return h
}
