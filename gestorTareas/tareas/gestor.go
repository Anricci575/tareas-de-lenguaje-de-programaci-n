package tareas

type GestorTareas struct {
	Tareas      []Tarea `json:"tareas"`
	SiguienteID int     `json:"siguiente_id"`
}

func NuevoGestor() *GestorTareas {
	return &GestorTareas{
		Tareas:      []Tarea{},
		SiguienteID: 1,
	}
}
