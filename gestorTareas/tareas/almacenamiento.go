package tareas

import (
	"encoding/json"
	"fmt"
	"os"
)

const nombreArchivo = "tareas.json"

func (g *GestorTareas) CargarDesdeArchivo() error {
	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		// Si el archivo no existe, empezamos con tareas vacías
		return nil
	}
	defer archivo.Close()

	decoder := json.NewDecoder(archivo)
	err = decoder.Decode(g)
	if err != nil {
		return fmt.Errorf("error cargando tareas: %v", err)
	}

	return nil
}

func (g *GestorTareas) GuardarEnArchivo() error {
	archivo, err := os.Create(nombreArchivo)
	if err != nil {
		return fmt.Errorf("error creando archivo: %v", err)
	}
	defer archivo.Close()

	encoder := json.NewEncoder(archivo)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(g)
	if err != nil {
		return fmt.Errorf("error guardando tareas: %v", err)
	}

	return nil
}
