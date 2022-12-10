package main

import (
	"encoding/gob"
	"fmt"
	"net" //se usa este paquete para crear el TCP
)

func cliente() {
	// El cliente se conecta al servidor que deberia de estar corriendo
	//conexión,error ||| protocolo, puerto
	c, err := net.Dial("tcp", ":9999")
	c2, err2 := net.Dial("tcp", ":9999")
	c3, err3 := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		//return es para que si termina la función la continue
		return
	}
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	a2 := 5
	b2 := 20
	n := 10000
	fmt.Println("Enviando valor de a2 -> ", a2)
	fmt.Println("Enviando valor de b2 -> ", b2)
	fmt.Println("Enviando número de trapecios -> ", n)
	//Códifica el mensaje a2, b2 ,n
	err = gob.NewEncoder(c).Encode(a2)
	err = gob.NewEncoder(c2).Encode(b2)
	err = gob.NewEncoder(c3).Encode(n)
	if err != nil {
		fmt.Println(err)
	}
	if err2 != nil {
		fmt.Println(err2)
	}
	if err3 != nil {
		fmt.Println(err3)
	}
	//Cerramos la conexión
	c.Close()
	c2.Close()
	c3.Close()
}

func main() {
	//El cliente va a estar corriendo concurrentemente
	go cliente()
	//La pausa
	var input string
	fmt.Scanln(&input)
}
