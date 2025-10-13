package tareas

import (
	"fmt"
	"strconv"
	"time"
)

func (g *GestorTareas) Agregar(descripcion string, fechaLimiteStr string) {
	tarea := Tarea{
		ID:            g.SiguienteID,
		Descripcion:   descripcion,
		Completada:    false,
		FechaCreacion: time.Now(),
	}

	// Procesar fecha límite si se proporciona
	if fechaLimiteStr != "" {
		fechaLimite, err := time.Parse("02/01/2006", fechaLimiteStr)
		if err == nil {
			tarea.FechaLimite = &fechaLimite
		} else {
			fmt.Printf("Formato de fecha inválido. Usa DD/MM/AAAA\n")
		}
	}

	g.Tareas = append(g.Tareas, tarea)
	g.SiguienteID++

	if tarea.FechaLimite != nil {
		fmt.Printf("Tarea agregada: %s (ID: %d) - Límite: %s\n",
			descripcion, tarea.ID, tarea.FechaLimite.Format("02/01/2006"))
	} else {
		fmt.Printf("Tarea agregada: %s (ID: %d)\n", descripcion, tarea.ID)
	}
}

func (g *GestorTareas) Listar() {
	if len(g.Tareas) == 0 {
		fmt.Println("No hay tareas pendientes")
		return
	}

	fmt.Println("\nLista de Tareas:")
	fmt.Println("ID | Estado    | Fecha Creación | Fecha Límite | Descripción")
	fmt.Println("---|-----------|----------------|--------------|-------------")

	for _, tarea := range g.Tareas {
		estado := "Pendiente"
		if tarea.Completada {
			estado = "Completada"
		}

		fechaCreacion := tarea.FechaCreacion.Format("02/01/2006")
		fechaLimite := "          "
		if tarea.FechaLimite != nil {
			fechaLimite = tarea.FechaLimite.Format("02/01/2006")

			// Verificar si la tarea está vencida
			if !tarea.Completada && tarea.FechaLimite.Before(time.Now()) {
				estado = "Vencida"
			}
		}

		fmt.Printf("%2d | %-9s | %-14s | %-12s | %s\n",
			tarea.ID, estado, fechaCreacion, fechaLimite, tarea.Descripcion)
	}
	fmt.Printf("\nTotal: %d tarea(s)\n", len(g.Tareas))
}

func (g *GestorTareas) Editar(idStr string, nuevaDescripcion string, nuevaFechaLimiteStr string) error {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("el ID debe ser un número")
	}

	for i := range g.Tareas {
		if g.Tareas[i].ID == id {
			// Guardar la descripción anterior para el mensaje
			descripcionAnterior := g.Tareas[i].Descripcion

			// Actualizar descripción si se proporciona
			if nuevaDescripcion != "" {
				g.Tareas[i].Descripcion = nuevaDescripcion
			}

			// Actualizar fecha límite si se proporciona
			if nuevaFechaLimiteStr != "" {
				if nuevaFechaLimiteStr == "quitar" {
					g.Tareas[i].FechaLimite = nil
					fmt.Printf("Tarea editada: '%s' → '%s' (fecha límite quitada)\n",
						descripcionAnterior, g.Tareas[i].Descripcion)
				} else {
					fechaLimite, err := time.Parse("02/01/2006", nuevaFechaLimiteStr)
					if err == nil {
						g.Tareas[i].FechaLimite = &fechaLimite
						fmt.Printf("Tarea editada: '%s' → '%s' (límite: %s)\n",
							descripcionAnterior, g.Tareas[i].Descripcion, fechaLimite.Format("02/01/2006"))
					} else {
						return fmt.Errorf("formato de fecha inválido. Usa DD/MM/AAAA")
					}
				}
			} else {
				fmt.Printf("Tarea editada: '%s' → '%s'\n",
					descripcionAnterior, g.Tareas[i].Descripcion)
			}

			return nil
		}
	}
	return fmt.Errorf("no se encontró la tarea con ID %d", id)
}

func (g *GestorTareas) Completar(idStr string) error {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("el ID debe ser un número")
	}

	for i := range g.Tareas {
		if g.Tareas[i].ID == id {
			if g.Tareas[i].Completada {
				return fmt.Errorf("la tarea %d ya estaba completada", id)
			}
			g.Tareas[i].Completada = true
			fmt.Printf("Tarea %d marcada como completada: %s\n", id, g.Tareas[i].Descripcion)
			return nil
		}
	}
	return fmt.Errorf("no se encontró la tarea con ID %d", id)
}

func (g *GestorTareas) Borrar(idStr string) error {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("el ID debe ser un número")
	}

	for i, tarea := range g.Tareas {
		if tarea.ID == id {
			g.Tareas = append(g.Tareas[:i], g.Tareas[i+1:]...)
			fmt.Printf("Tarea eliminada: %s (ID: %d)\n", tarea.Descripcion, id)
			return nil
		}
	}
	return fmt.Errorf("no se encontró la tarea con ID %d", id)
}
