package tareas

import "fmt"

func (g *GestorTareas) MostrarUso() {
	fmt.Println("")
	fmt.Println("                 _                   _        _                           ")
	fmt.Println("                | |                 | |      | |                           ")
	fmt.Println("  __ _  ___  ___| |_ ___  _ __    __| | ___  | |_ __ _ _ __ ___  __ _ ___ ")
	fmt.Println(" / _` |/ _ \\/ __| __/ _ \\| '__|  / _` |/ _ \\ | __/ _` | '__/ _ \\/ _` / __|")
	fmt.Println("| (_| |  __/\\__ \\ || (_) | |    | (_| |  __/ | || (_| | | |  __/ (_| \\__ \\")
	fmt.Println(" \\__, |\\___||___/\\__\\___/|_|     \\__,_|\\___|  \\__\\__,_|_|  \\___|\\__,_|___/")
	fmt.Println("  __/ |                                                                 ")
	fmt.Println(" |___/                                                                  ")
	fmt.Println("---------------------------------------------------------")
	fmt.Println("Andres Rodriguez C.I: 31.219.248")
	fmt.Println("---------------------------------------------------------")
	fmt.Println("Gestor de Tareas - Uso:")
	fmt.Println("   agregar  <tarea> [fecha]    - Agregar tarea (fecha: DD/MM/AAAA)")
	fmt.Println("   listar                      - Ver todas las tareas")
	fmt.Println("   editar   <id> <nueva_desc> [fecha] - Editar tarea")
	fmt.Println("   completar <id>              - Marcar tarea como completada")
	fmt.Println("   borrar    <id>              - Eliminar tarea")
	fmt.Println("   ayuda                       - Mostrar esta ayuda")
	fmt.Println("")
	fmt.Println("💡 EJEMPLOS PRÁCTICOS:")
	fmt.Println("  gestor agregar \"Estudiar matemáticas\" 25/12/2025")
	fmt.Println("  gestor editar 1 \"Estudiar cálculo\" 30/12/2025")
	fmt.Println("  gestor editar 1 \"Nueva descripción\" quitar  (quita fecha límite)")
	fmt.Println("  gestor listar")
	fmt.Println("  gestor completar 1")
	fmt.Println("  gestor borrar 2")
	fmt.Println("")
}
