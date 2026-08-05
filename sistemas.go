package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ============================================================
// CLASE LIBRO (CON ENCAPSULACIÓN)
// ============================================================

type Libro struct {
	id     int
	titulo string
	autor  string
	estado string // "disponible" o "prestado"
}

// Constructores y métodos getter
func NewLibro(titulo, autor string) *Libro {
	return &Libro{
		id:     0,
		titulo: titulo,
		autor:  autor,
		estado: "disponible",
	}
}

func (l *Libro) GetID() int {
	return l.id
}

func (l *Libro) GetTitulo() string {
	return l.titulo
}

func (l *Libro) GetAutor() string {
	return l.autor
}

func (l *Libro) GetEstado() string {
	return l.estado
}

func (l *Libro) SetEstado(estado string) {
	l.estado = estado
}

func (l *Libro) SetID(id int) {
	l.id = id
}

// ============================================================
// CLASE USUARIO
// ============================================================

type Usuario struct {
	id        int
	nombre    string
	prestamos []int // IDs de libros prestados
}

func NewUsuario(nombre string) *Usuario {
	return &Usuario{
		id:        0,
		nombre:    nombre,
		prestamos: []int{},
	}
}

func (u *Usuario) GetID() int {
	return u.id
}

func (u *Usuario) GetNombre() string {
	return u.nombre
}

func (u *Usuario) GetPrestamos() []int {
	return u.prestamos
}

func (u *Usuario) SetID(id int) {
	u.id = id
}

func (u *Usuario) AgregarPrestamo(libroID int) {
	u.prestamos = append(u.prestamos, libroID)
}

func (u *Usuario) RemoverPrestamo(libroID int) bool {
	for i, id := range u.prestamos {
		if id == libroID {
			u.prestamos = append(u.prestamos[:i], u.prestamos[i+1:]...)
			return true
		}
	}
	return false
}

// ============================================================
// CLASE BIBLIOTECA (GESTOR CENTRAL)
// ============================================================

type Biblioteca struct {
	libros           []*Libro
	usuarios         []*Usuario
	contadorLibros   int
	contadorUsuarios int
}

func NewBiblioteca() *Biblioteca {
	return &Biblioteca{
		libros:           []*Libro{},
		usuarios:         []*Usuario{},
		contadorLibros:   1,
		contadorUsuarios: 1,
	}
}

// ---------- MÉTODOS DE LIBROS ----------

func (b *Biblioteca) AgregarLibro(titulo, autor string) int {
	libro := NewLibro(titulo, autor)
	libro.SetID(b.contadorLibros)
	b.libros = append(b.libros, libro)
	b.contadorLibros++
	return libro.GetID()
}

func (b *Biblioteca) ListarLibros() {
	if len(b.libros) == 0 {
		fmt.Println("📭 No hay libros en la biblioteca.")
		return
	}
	fmt.Println("ID | TÍTULO | AUTOR | ESTADO")
	fmt.Println("-------------------------------------")
	for _, l := range b.libros {
		fmt.Printf("%d | %s | %s | %s\n", l.GetID(), l.GetTitulo(), l.GetAutor(), l.GetEstado())
	}
}

func (b *Biblioteca) BuscarLibros(texto string) []*Libro {
	resultados := []*Libro{}
	texto = strings.ToLower(texto)
	for _, l := range b.libros {
		if strings.Contains(strings.ToLower(l.GetTitulo()), texto) || strings.Contains(strings.ToLower(l.GetAutor()), texto) {
			resultados = append(resultados, l)
		}
	}
	return resultados
}

func (b *Biblioteca) CambiarEstado(id int) error {
	for _, l := range b.libros {
		if l.GetID() == id {
			if l.GetEstado() == "disponible" {
				l.SetEstado("prestado")
			} else {
				l.SetEstado("disponible")
			}
			return nil
		}
	}
	return fmt.Errorf("libro con ID %d no encontrado", id)
}

func (b *Biblioteca) EliminarLibro(id int) error {
	for i, l := range b.libros {
		if l.GetID() == id {
			// Si está prestado, no se puede eliminar
			if l.GetEstado() == "prestado" {
				return fmt.Errorf("el libro '%s' está prestado y no se puede eliminar", l.GetTitulo())
			}
			b.libros = append(b.libros[:i], b.libros[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("libro con ID %d no encontrado", id)
}

func (b *Biblioteca) ObtenerLibroPorID(id int) *Libro {
	for _, l := range b.libros {
		if l.GetID() == id {
			return l
		}
	}
	return nil
}

// ---------- MÉTODOS DE USUARIOS ----------

func (b *Biblioteca) AgregarUsuario(nombre string) int {
	usuario := NewUsuario(nombre)
	usuario.SetID(b.contadorUsuarios)
	b.usuarios = append(b.usuarios, usuario)
	b.contadorUsuarios++
	return usuario.GetID()
}

func (b *Biblioteca) ListarUsuarios() {
	if len(b.usuarios) == 0 {
		fmt.Println("No hay usuarios registrados")
		return
	}
	fmt.Println("ID | Nombre | Libros Prestados")
	fmt.Println("-------------------------------------")
	for _, u := range b.usuarios {
		fmt.Printf("%d | %s | %d libros\n", u.GetID(), u.GetNombre(), len(u.GetPrestamos()))
	}
}

func (b *Biblioteca) ObtenerUsuarioPorID(id int) *Usuario {
	for _, u := range b.usuarios {
		if u.GetID() == id {
			return u
		}
	}
	return nil
}

// ---------- MÉTODOS DE PRÉSTAMOS ----------

func (b *Biblioteca) PrestarLibro(usuarioID, libroID int) error {
	usuario := b.ObtenerUsuarioPorID(usuarioID)
	if usuario == nil {
		return fmt.Errorf("usuario con ID %d no encontrado", usuarioID)
	}

	libro := b.ObtenerLibroPorID(libroID)
	if libro == nil {
		return fmt.Errorf("libro con ID %d no encontrado", libroID)
	}

	if libro.GetEstado() == "prestado" {
		return fmt.Errorf("el libro '%s' ya está prestado", libro.GetTitulo())
	}

	// Realizar préstamo
	libro.SetEstado("prestado")
	usuario.AgregarPrestamo(libroID)
	return nil
}

func (b *Biblioteca) DevolverLibro(usuarioID, libroID int) error {
	usuario := b.ObtenerUsuarioPorID(usuarioID)
	if usuario == nil {
		return fmt.Errorf("usuario con ID %d no encontrado", usuarioID)
	}

	libro := b.ObtenerLibroPorID(libroID)
	if libro == nil {
		return fmt.Errorf("libro con ID %d no encontrado", libroID)
	}

	if libro.GetEstado() == "disponible" {
		return fmt.Errorf("el libro '%s' no está prestado", libro.GetTitulo())
	}

	// Verificar que el usuario tiene este libro prestado
	if !usuario.RemoverPrestamo(libroID) {
		return fmt.Errorf("el usuario no tiene prestado el libro con ID %d", libroID)
	}

	libro.SetEstado("disponible")
	return nil
}

// ============================================================
// FUNCIÓN MAIN (MENÚ PRINCIPAL)
// ============================================================

func main() {
	lector := bufio.NewReader(os.Stdin)
	biblioteca := NewBiblioteca()

	for {
		fmt.Println("\n=====================================")
		fmt.Println("   SISTEMA DE GESTIÓN DE LIBROS (POO)")
		fmt.Println("=====================================")
		fmt.Println("1. Agregar Libro")
		fmt.Println("2. Listar Libros")
		fmt.Println("3. Buscar Libro")
		fmt.Println("4. Cambiar Estado")
		fmt.Println("5. Eliminar Libro")
		fmt.Println("6. Agregar Usuario")
		fmt.Println("7. Listar Usuarios")
		fmt.Println("8. Prestar Libro")  // NUEVO
		fmt.Println("9. Devolver Libro") // NUEVO
		fmt.Println("10. Salir")
		fmt.Print("Opción: ")

		opcionStr, _ := lector.ReadString('\n')
		opcionStr = strings.TrimSpace(opcionStr)
		opcion, err := strconv.Atoi(opcionStr)
		if err != nil {
			fmt.Println("❌ Error: Debes ingresar un número válido.")
			continue
		}

		switch opcion {
		case 1:
			fmt.Print("Título: ")
			titulo, _ := lector.ReadString('\n')
			titulo = strings.TrimSpace(titulo)
			fmt.Print("Autor: ")
			autor, _ := lector.ReadString('\n')
			autor = strings.TrimSpace(autor)
			id := biblioteca.AgregarLibro(titulo, autor)
			fmt.Printf("✅ Libro agregado con ID: %d\n", id)

		case 2:
			biblioteca.ListarLibros()

		case 3:
			fmt.Print("🔍 Ingresa el título o autor a buscar: ")
			texto, _ := lector.ReadString('\n')
			texto = strings.TrimSpace(texto)
			resultados := biblioteca.BuscarLibros(texto)
			if len(resultados) == 0 {
				fmt.Println("❌ No se encontraron libros.")
			} else {
				fmt.Println("\n--- RESULTADOS ---")
				for _, l := range resultados {
					fmt.Printf("ID: %d | Título: %s | Autor: %s | Estado: %s\n", l.GetID(), l.GetTitulo(), l.GetAutor(), l.GetEstado())
				}
			}

		case 4:
			fmt.Print("📌 Ingresa el ID del libro: ")
			idStr, _ := lector.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))
			err := biblioteca.CambiarEstado(id)
			if err != nil {
				fmt.Println("❌", err)
			} else {
				fmt.Println("✅ Estado actualizado correctamente.")
			}

		case 5:
			fmt.Print("🗑️ Ingresa el ID del libro a eliminar: ")
			idStr, _ := lector.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))
			err := biblioteca.EliminarLibro(id)
			if err != nil {
				fmt.Println("❌", err)
			} else {
				fmt.Println("✅ Libro eliminado permanentemente.")
			}

		case 6:
			fmt.Print("Nombre: ")
			nombre, _ := lector.ReadString('\n')
			nombre = strings.TrimSpace(nombre)
			id := biblioteca.AgregarUsuario(nombre)
			fmt.Printf("✅ Usuario agregado con ID: %d\n", id)

		case 7:
			biblioteca.ListarUsuarios()

		case 8:
			fmt.Print("📌 ID del usuario: ")
			usuarioIDStr, _ := lector.ReadString('\n')
			usuarioID, _ := strconv.Atoi(strings.TrimSpace(usuarioIDStr))
			fmt.Print("📌 ID del libro: ")
			libroIDStr, _ := lector.ReadString('\n')
			libroID, _ := strconv.Atoi(strings.TrimSpace(libroIDStr))
			err := biblioteca.PrestarLibro(usuarioID, libroID)
			if err != nil {
				fmt.Println("❌", err)
			} else {
				fmt.Println("✅ Préstamo exitoso.")
			}

		case 9:
			fmt.Print("📌 ID del usuario: ")
			usuarioIDStr, _ := lector.ReadString('\n')
			usuarioID, _ := strconv.Atoi(strings.TrimSpace(usuarioIDStr))
			fmt.Print("📌 ID del libro: ")
			libroIDStr, _ := lector.ReadString('\n')
			libroID, _ := strconv.Atoi(strings.TrimSpace(libroIDStr))
			err := biblioteca.DevolverLibro(usuarioID, libroID)
			if err != nil {
				fmt.Println("❌", err)
			} else {
				fmt.Println("✅ Devolución exitosa.")
			}

		case 10:
			fmt.Println("👋 ¡Hasta luego!")
			return

		default:
			fmt.Println("❌ Opción no válida.")
		}
	}
}
