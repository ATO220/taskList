package taskList

import "fmt"

type taskList struct {
	tasks []*task
}

func (tl *taskList) AgregarALista(t *task) {
	tl.tasks = append(tl.tasks, t)
}
func (tl *taskList) EliminarDeLista(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

func (tl *taskList) ImprimirLista() {
	for _, tarea := range tl.tasks {
		fmt.Println("Nombre:", tarea.nombre)
		fmt.Println("Descripcion:", tarea.descripcion)
	}
}
func (tl *taskList) ImprimirListaCompletadas() {
	for _, tarea := range tl.tasks {
		if tarea.completado {
			fmt.Println("Nombre:", tarea.nombre)
			fmt.Println("Descripcion:", tarea.descripcion)
		}
	}
}

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}
