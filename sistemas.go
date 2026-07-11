package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lector := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n1. Agregar")
		fmt.Println("2. Listar")
		fmt.Println("3. Buscar")
		fmt.Println("4. Cambiar Estado")
		fmt.Println("5. Eliminar")
		fmt.Println("6. Salir")
		fmt.Print("Opción: ")

		opcion, _ := lector.ReadString('\n')
		opcion = string(opcion[0])

		switch opcion {
		case "1":
			fmt.Println("Función Agregar - en construcción")
		case "2":
			fmt.Println("Función Listar - en construcción")
		case "3":
			fmt.Println("Función Buscar - en construcción")
		case "4":
			fmt.Println("Función Cambiar Estado - en construcción")
		case "5":
			fmt.Println("Función Eliminar - en construcción")
		case "6":
			fmt.Println("Adiós")
			return
		default:
			fmt.Println("Opción no válida")
		}
	}
}
