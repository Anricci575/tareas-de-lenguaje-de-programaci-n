package tareas

import "time"

type Tarea struct {
	ID            int        `json:"id"`
	Descripcion   string     `json:"descripcion"`
	Completada    bool       `json:"completada"`
	FechaCreacion time.Time  `json:"fecha_creacion"`
	FechaLimite   *time.Time `json:"fecha_limite,omitempty"` // Opcional
}
