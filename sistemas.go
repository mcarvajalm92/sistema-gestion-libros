package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ============================================================
// CLASE (STRUCT) LIBRO
// ============================================================

type Libro struct {
	ID     int
	Titulo string
	Autor  string
	Estado string // "Disponible" o "Prestado"
}

// ============================================================
// CLASE (STRUCT) USUARIO (nuevo)
// ============================================================

type Usuario struct {
	ID        int
	Nombre    string
	Prestamos []Libro // IDs de libros prestados (por ahora solo IDs)
}

// ============================================================
// VARIABLES GLOBALES
// ============================================================

var libros []Libro
var usuarios []Usuario
var contadorLibros = 1
var contadorUsuarios = 1

// ============================================================
// FUNCIONES DEL SISTEMA
// ============================================================

func main() {
	lector := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n=====================================")
		fmt.Println("   SISTEMA DE GESTIÓN DE LIBROS (POO)")
		fmt.Println("=====================================")
		fmt.Println("1. Agregar Libro")
		fmt.Println("2. Listar Libros")
		fmt.Println("3. Buscar Libro")
		fmt.Println("4. Cambiar Estado")
		fmt.Println("5. Eliminar Libro")
		fmt.Println("6. Salir")
		fmt.Println("7. Agregar Usuario") // NUEVO
		fmt.Println("8. Listar Usuarios") // NUEVO
		fmt.Print("Opción: ")

		opcion, _ := lector.ReadString('\n')
		opcion = strings.TrimSpace(opcion)

		switch opcion {
		case "1":
			agregarLibro(lector)
		case "2":
			listarLibros()
		case "3":
			buscarLibro(lector)
		case "4":
			cambiarEstado(lector)
		case "5":
			eliminarLibro(lector)
		case "6":
			fmt.Println("Adiós")
			return
		case "7":
			agregarUsuario(lector)
		case "8":
			listarUsuarios()
		default:
			fmt.Println("Opción no válida")
		}
	}
}

// ============================================================
// FUNCIONES DE LIBROS
// ============================================================

func agregarLibro(lector *bufio.Reader) {
	fmt.Print("Título: ")
	titulo, _ := lector.ReadString('\n')
	fmt.Print("Autor: ")
	autor, _ := lector.ReadString('\n')
	libros = append(libros, Libro{contadorLibros, strings.TrimSpace(titulo), strings.TrimSpace(autor), "Disponible"})
	fmt.Println("Agregado ID:", contadorLibros)
	contadorLibros++
}

func listarLibros() {
	if len(libros) == 0 {
		fmt.Println("No hay libros")
		return
	}
	fmt.Println("ID | Título | Autor | Estado")
	fmt.Println("-------------------------------------")
	for _, l := range libros {
		fmt.Printf("%d | %s | %s | %s\n", l.ID, l.Titulo, l.Autor, l.Estado)
	}
}

func buscarLibro(lector *bufio.Reader) {
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

func eliminarLibro(lector *bufio.Reader) {
	fmt.Print("ID: ")
	idStr, _ := lector.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))
	for i, l := range libros {
		if l.ID == id {
			libros = append(libros[:i], libros[i+1:]...)
			fmt.Println("Eliminado")
			return
		}
	}
	fmt.Println("No encontrado")
}

// ============================================================
// FUNCIONES DE USUARIOS (NUEVAS)
// ============================================================

func agregarUsuario(lector *bufio.Reader) {
	fmt.Print("Nombre: ")
	nombre, _ := lector.ReadString('\n')
	nombre = strings.TrimSpace(nombre)
	usuarios = append(usuarios, Usuario{contadorUsuarios, nombre, []Libro{}})
	fmt.Println("Usuario agregado con ID:", contadorUsuarios)
	contadorUsuarios++
}

func listarUsuarios() {
	if len(usuarios) == 0 {
		fmt.Println("No hay usuarios registrados")
		return
	}
	fmt.Println("ID | Nombre | Libros Prestados")
	fmt.Println("-------------------------------------")
	for _, u := range usuarios {
		fmt.Printf("%d | %s | %d libros\n", u.ID, u.Nombre, len(u.Prestamos))
	}
}
