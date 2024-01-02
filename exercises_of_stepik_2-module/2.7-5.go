package main

import (
	"fmt"
	"math"
)

type S struct {
	w, p, v, k, m, t float64
}

func main() {
	ex := &S{}
	fmt.Scan(&ex.k, &ex.p, &ex.v)
	ex.M()
	ex.W()
	ex.T()
	fmt.Println(ex.t)
}

func (s *S) T() {
	s.t = 6 / s.w
}

func (s *S) M() {
	s.m = s.p * s.v
}

func (s *S) W() {
	s.w = math.Sqrt(s.k / s.m)
}