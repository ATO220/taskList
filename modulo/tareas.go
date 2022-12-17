package taskList

import "fmt"

type TaskList struct {
	tasks []*Task
}

func (tl *TaskList) AgregarALista(t *Task) {
	tl.tasks = append(tl.tasks, t)
}
func (tl *TaskList) EliminarDeLista(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

func (tl *TaskList) ImprimirLista() {
	for _, tarea := range tl.tasks {
		fmt.Println("Nombre:", tarea.nombre)
		fmt.Println("Descripcion:", tarea.descripcion)
	}
}
func (tl *TaskList) ImprimirListaCompletadas() {
	for _, tarea := range tl.tasks {
		if tarea.completado {
			fmt.Println("Nombre:", tarea.nombre)
			fmt.Println("Descripcion:", tarea.descripcion)
		}
	}
}

type Task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *Task) MarcarCompleta() {
	t.completado = true
}

func (t *Task) ActualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *Task) ActualizarNombre(nombre string) {
	t.nombre = nombre
}
