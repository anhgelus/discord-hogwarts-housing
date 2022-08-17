package xp

import (
	"github.com/anhgelus/discord-hogwarts-housing/src/util"
	"math"
)

func NewMessage(m string) float64 {
	return calc(len(m), util.GetNumberOfChar(m))
}

// l int - length of the message
// v int - Number of character in the message
func calc(l int, v int) float64 {
	// f(x)=((0.025 x^(1.25))/(50^(-0.5)))+1
	result := 0.025 * math.Pow(float64(l), 1.25)
	result = result / math.Pow(float64(v), -0.5)
	return result + 1
}
