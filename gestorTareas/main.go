package main

import (
	"fmt"
	"os"
	"strings"

	"gestorTareas/tareas"
)

func main() {

	// Crear el gestor de tareas
	gestor := tareas.NuevoGestor()

	// Cargar tareas existentes
	err := gestor.CargarDesdeArchivo()
	if err != nil {
		fmt.Printf("⚠️  %v\n", err)
	}

	// Verificar que se haya proporcionado un comando
	if len(os.Args) < 2 {
		gestor.MostrarUso()
		return
	}

	// Procesar el comando
	comando := os.Args[1]

	switch comando {
	case "agregar", "add", "a":
		if len(os.Args) < 3 {
			fmt.Println("❌ Error: Debes proporcionar una descripción para la tarea")
			fmt.Println("   Uso: gestor agregar \"Mi tarea\" [DD/MM/AAAA]")
			return
		}

		descripcion := os.Args[2]
		fechaLimite := ""

		// Verificar si el último argumento es una fecha
		if len(os.Args) >= 4 {
			// Intentar parsear la fecha
			_, err := parsearFecha(os.Args[3])
			if err == nil {
				fechaLimite = os.Args[3]
				// Si hay más argumentos, unirlos a la descripción
				if len(os.Args) > 4 {
					descripcion = strings.Join(os.Args[2:len(os.Args)-1], " ")
					fechaLimite = os.Args[len(os.Args)-1]
				}
			} else {
				// Si no es fecha, unir todo como descripción
				descripcion = strings.Join(os.Args[2:], " ")
			}
		} else {
			descripcion = strings.Join(os.Args[2:], " ")
		}

		gestor.Agregar(descripcion, fechaLimite)

	case "listar", "list", "l", "ls":
		gestor.Listar()

	case "editar", "edit", "e":
		if len(os.Args) < 4 {
			fmt.Println("❌ Error: Debes proporcionar ID y nueva descripción")
			fmt.Println("   Uso: gestor editar <id> \"Nueva descripción\" [DD/MM/AAAA|quitar]")
			return
		}

		id := os.Args[2]
		nuevaDescripcion := os.Args[3]
		nuevaFechaLimite := ""

		// Verificar si hay un cuarto argumento (fecha)
		if len(os.Args) >= 5 {
			nuevaFechaLimite = os.Args[4]
		}

		err := gestor.Editar(id, nuevaDescripcion, nuevaFechaLimite)
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			return
		}

	case "completar", "complete", "c":
		if len(os.Args) < 3 {
			fmt.Println("❌ Error: Debes proporcionar el ID de la tarea a completar")
			fmt.Println("   Uso: gestor completar 1")
			return
		}
		err := gestor.Completar(os.Args[2])
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			return
		}

	case "borrar", "delete", "remove", "d", "rm":
		if len(os.Args) < 3 {
			fmt.Println("❌ Error: Debes proporcionar el ID de la tarea a borrar")
			fmt.Println("   Uso: gestor borrar 1")
			return
		}
		err := gestor.Borrar(os.Args[2])
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			return
		}

	case "ayuda", "help", "h", "?":
		gestor.MostrarUso()

	default:
		fmt.Printf("❌ Comando no reconocido: %s\n", comando)
		gestor.MostrarUso()
		return
	}

	// Guardar los cambios después de cada comando
	err = gestor.GuardarEnArchivo()
	if err != nil {
		fmt.Printf("❌ Error guardando tareas: %v\n", err)
	}
}

func parsearFecha(fechaStr string) (string, error) {
	// Validar formato básico de fecha (DD/MM/AAAA)
	if len(fechaStr) != 10 || fechaStr[2] != '/' || fechaStr[5] != '/' {
		return "", fmt.Errorf("formato inválido")
	}
	return fechaStr, nil
}
