package main

import (
	"encoding/gob" // codificacion de go object
	"fmt"
	"math" //se usa para hallar el trapecio
	"net"  //se usa este paquete para crear el TCP
)

func Trapecio(f func(float64) float64, a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	sum := 0.5 * (f(a) + f(b))
	for i := 1; i < n; i++ {
		sum += f(a + float64(i)*h)
	}
	return sum * h
}

func proceso(nTrapecios <-chan int, resultados chan<- float64) {
	f := func(x float64) float64 {
		return ((math.Pow(x, 2) + 1) / 2)
	}
	for n := range nTrapecios {
		resultados <- Trapecio(f, 5, 20, n)
	}
}

func servidor() {
	// se le asigna el puerto para abrirlo
	//conexión,error ||| protocolo, puerto
	s, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	//Si no ocurre el error se crea un ciclo que escucha
	//las peticiones de los clientes
	for {
		// el s.Accept() espera al cliente
		// c es el cliente y err el error
		c, err := s.Accept()
		c2, err2 := s.Accept()
		c3, err3 := s.Accept()
		if err != nil && (err2 != nil) && (err3 != nil) {
			fmt.Println(err)
			continue
		}
		//acepta la conexíon entrante y se acepta
		//lo que el cliente envia
		go handleClientA(c)
		go handleClientB(c2)
		go handleClient(c3)
	}
}

func handleClientA(c net.Conn) {
	var a2 int
	//Decode se hace un puntero
	err := gob.NewDecoder(c).Decode(&a2)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Mensaje: ", a2)
	}
}

func handleClientB(c2 net.Conn) {
	var b2 int
	//Decode se hace un puntero
	err2 := gob.NewDecoder(c2).Decode(&b2)
	if err2 != nil {
		fmt.Println(err2)
		return
	} else {
		fmt.Println("Mensaje: ", b2)
	}
}

func handleClient(c3 net.Conn) {
	var n int
	nTrapecios := make(chan int, n)
	resultados := make(chan float64, n)
	//Decode se hace un puntero
	err3 := gob.NewDecoder(c3).Decode(&n)
	if err3 != nil {
		fmt.Println(err3)
		return
	} else {

		for i := 0; i < n; i++ {
			go proceso(nTrapecios, resultados)
			go proceso(nTrapecios, resultados)
			go proceso(nTrapecios, resultados)
			go proceso(nTrapecios, resultados)
			go proceso(nTrapecios, resultados)
		}

		for i := 0; i < n; i++ {
			nTrapecios <- i
		}
		close(nTrapecios)

		for i := 0; i < n; i++ {
			fmt.Println(<-resultados)
		}
	}
}

func main() {
	//Se queda concurrente esperando a llamada del cliente
	go servidor()
	//hasta que no se escriba algo en la terminal el servidor va a seguir en un hilo
	var input string
	fmt.Scanln(&input)
}
