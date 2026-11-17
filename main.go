package catalogo

type Contenido struct {
	ID          int
	Titulo      string
	Genero      string
	EdadMinima  int
	RequiereVIP bool
	DuracionMin int
}

type CriterioFiltro func(Contenido) bool

func FiltrarContenido(lista []Contenido, criterio CriterioFiltro) []Contenido {
	var resultado []Contenido
	for _, item := range lista {
		if criterio(item) {
			resultado = append(resultado, item)
		}
	}
	return resultado
}

func ObtenerCatalogoInicial() []Contenido {
	return []Contenido{
		{ID: 1, Titulo: "Mr. Robot", Genero: "Cyberpunk", EdadMinima: 16, RequiereVIP: true, DuracionMin: 55},
		{ID: 2, Titulo: "Interestelar", Genero: "Ciencia Ficcion", EdadMinima: 12, RequiereVIP: false, DuracionMin: 169},
		{ID: 3, Titulo: "El Conjuro", Genero: "Terror", EdadMinima: 18, RequiereVIP: true, DuracionMin: 112},
		{ID: 4, Titulo: "Toy Story", Genero: "Animacion", EdadMinima: 0, RequiereVIP: false, DuracionMin: 81},
	}
}
package usuarios

type Usuario struct {
	Nombre    string
	Edad      int
	EsPremium bool
}

// CrearUsuario actúa como constructor puro
func CrearUsuario(nombre string, edad int, esPremium bool) Usuario {
	return Usuario{
		Nombre:    nombre,
		Edad:      edad,
		EsPremium: esPremium,
	}
}
package reproductor

import (
	"fmt"
	"sistema-streaming/catalogo"
	"sistema-streaming/usuarios"
)

// IniciarStream evalúa de forma pura si el usuario cumple las condiciones para ver el título
func IniciarStream(u usuarios.Usuario, c catalogo.Contenido) string {
	if u.Edad < c.EdadMinima {
		return fmt.Sprintf("[BLOQUEADO] %s no cumple la edad mínima (%d+) para ver '%s'.", u.Nombre, c.EdadMinima, c.Titulo)
	}
	if c.RequiereVIP && !u.EsPremium {
		return fmt.Sprintf("[SUSCRIPCIÓN REQUERIDA] '%s' es exclusivo para usuarios Premium.", c.Titulo)
	}
	return fmt.Sprintf("[REPRODUCIENDO] %s está viendo '%s' (%d min) en calidad HD...", u.Nombre, c.Titulo, c.DuracionMin)
}
package main

import (
	"fmt"
	"sistema-streaming/catalogo"
	"sistema-streaming/reproductor"
	"sistema-streaming/usuarios"
)

func main() {
	// 1. Carga del catálogo base
	listaMaster := catalogo.ObtenerCatalogoInicial()

	// 2. Creación de usuarios
	usuarioFree := usuarios.CrearUsuario("Carlos", 15, false)
	usuarioVIP := usuarios.CrearUsuario("Ana", 22, true)

	// 3. Uso de Programación Funcional: pasamos funciones anónimas (closures) como filtro
	soloParaMenoresDe16 := catalogo.FiltrarContenido(listaMaster, func(c catalogo.Contenido) bool {
		return c.EdadMinima <= 16
	})
	soloCyberpunk := catalogo.FiltrarContenido(listaMaster, func(c catalogo.Contenido) bool {
		return c.Genero == "Cyberpunk"
	})

	fmt.Println("=== CATÁLOGO APTO PARA <= 16 AÑOS ===")
	for _, item := range soloParaMenoresDe16 {
		fmt.Printf("- %s (%s)\n", item.Titulo, item.Genero)
	}

	fmt.Println("\n=== PRUEBAS DE STREAMING ===")
	// Carlos intenta ver Mr. Robot (requiere 16+ y Premium)
	fmt.Println(reproductor.IniciarStream(usuarioFree, soloCyberpunk[0]))
	// Ana intenta ver Mr. Robot
	fmt.Println(reproductor.IniciarStream(usuarioVIP, soloCyberpunk[0]))
}