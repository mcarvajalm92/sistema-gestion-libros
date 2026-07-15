package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Libro struct {
	ID     int
	Titulo string
	Autor  string
	Estado string
}

var libros []Libro
var contador = 1

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
		opcion = strings.TrimSpace(opcion)

		switch opcion {
		case "1":
			agregar(lector)
		case "2":
			listar()
		case "3":
			buscar(lector)
		case "4":
			cambiarEstado(lector)
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

func agregar(lector *bufio.Reader) {
	fmt.Print("Título: ")
	titulo, _ := lector.ReadString('\n')
	fmt.Print("Autor: ")
	autor, _ := lector.ReadString('\n')
	libros = append(libros, Libro{contador, strings.TrimSpace(titulo), strings.TrimSpace(autor), "Disponible"})
	fmt.Println("Agregado ID:", contador)
	contador++
}

func listar() {
	if len(libros) == 0 {
		fmt.Println("No hay libros")
		return
	}
	for _, l := range libros {
		fmt.Printf("%d | %s | %s | %s\n", l.ID, l.Titulo, l.Autor, l.Estado)
	}
}

func buscar(lector *bufio.Reader) {
	fmt.Print("Buscar: ")
	texto, _ := lector.ReadString('\n')
	texto = strings.TrimSpace(strings.ToLower(texto))
	for _, l := range libros {
		if strings.Contains(strings.ToLower(l.Titulo), texto) || strings.Contains(strings.ToLower(l.Autor), texto) {
			fmt.Printf("%d | %s | %s | %s\n", l.ID, l.Titulo, l.Autor, l.Estado)
		}
	}
}

func cambiarEstado(lector *bufio.Reader) {
	fmt.Print("ID: ")
	idStr, _ := lector.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))
	for i, l := range libros {
		if l.ID == id {
			if l.Estado == "Disponible" {
				libros[i].Estado = "Prestado"
			} else {
				libros[i].Estado = "Disponible"
			}
			fmt.Println("Estado actualizado")
			return
		}
	}
	fmt.Println("No encontrado")
}
