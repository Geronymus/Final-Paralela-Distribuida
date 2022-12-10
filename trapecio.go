package main

import (
	"fmt"
	"math"
	"time"
)

func proceso(nTrapecios <-chan int, resultados chan<- float64) {
	f := func(x float64) float64 {
		return ((math.Pow(x, 2) + 1) / 2)
	}
	for n := range nTrapecios {
		resultados <- Trapecio(f, 5, 20, n)
	}
}

func Trapecio(f func(float64) float64, a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	sum := 0.5 * (f(a) + f(b))
	for i := 1; i < n; i++ {
		sum += f(a + float64(i)*h)
	}
	return sum * h
}

func main() {
	n := 100

	nTrapecios := make(chan int, n)
	resultados := make(chan float64, n)

	go proceso(nTrapecios, resultados)
	go proceso(nTrapecios, resultados)

	startMain := time.Now()
	for i := 0; i < n; i++ {
		nTrapecios <- i
	}
	endMain := time.Since(startMain).Nanoseconds()
	fmt.Println("Tiempo final:", endMain)

	close(nTrapecios)

	for i := 0; i < n; i++ {
		fmt.Println(<-resultados)
	}
}
