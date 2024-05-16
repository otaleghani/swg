package spg

import (
  "strconv"
  "math/rand"
)

func (g Generator) randomNumber(dice int) int {
  return rand.Intn(dice)
}

// Gives back a strigified number
func (g Generator) StringNumber(max int) string {
  number := g.randomNumber(max)
  return strconv.Itoa(number)
}

// Gives back x chars in numbers
func (g Generator) StringNumberFixed(chars int) string {
  result := ""
  for i := 0; i < chars; i++ {
    result = result + g.StringNumber(10)
  }
  return result
}
