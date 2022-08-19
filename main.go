package main

import (
	"container/heap"
	"fmt"
)

type Empleado struct {
	nombre    string
	direccion string
	fecha     string
	salario   int
	categoria string
	dni       string
	prioridad int

	index int
}

type PriorityQueue []*Empleado

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {

	return pq[i].prioridad > pq[j].prioridad
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	empleado := x.(*Empleado)
	empleado.index = n
	*pq = append(*pq, empleado)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	empleado := old[n-1]
	empleado.index = -1 //
	*pq = old[0 : n-1]
	return empleado
}

// actualiza y modifica la prioridad el valor de empleado en el queue

func (pq *PriorityQueue) update(empleado *Empleado, nombre string, direccion string, fecha string, salario int, categoria string, dni string, prioridad int) {
	empleado.nombre = nombre
	empleado.direccion = direccion
	empleado.fecha = fecha
	empleado.salario = salario
	empleado.categoria = categoria
	empleado.dni = dni
	empleado.prioridad = prioridad
	heap.Fix(pq, empleado.index)

}

func main() {
	ListaEmpleado := []*Empleado{{nombre: "Maria", direccion: "Cerro Brujo Honduras", fecha: "01-enero-1965", salario: 5000, categoria: "mantenimiento", dni: "0101-honduras-1965"},
		{nombre: "Juan", direccion: "Long Beach FLorida USA", fecha: "08-febrero-1965", salario: 4000, categoria: "conserje", dni: "0802-USA-1965"},
		{nombre: "Soraya", direccion: "Amapala Honduras", fecha: "07-marzo-1965", salario: 15000, categoria: "CEO", dni: "0703-honduras-1965"}}

	pq := make(PriorityQueue, len(ListaEmpleado))

	for i, empleado := range ListaEmpleado {

		pq[i] = empleado
		pq[i].index = i

	}

	heap.Init(&pq)

	for pq.Len() > 0 {
		empleado := heap.Pop(&pq).(*Empleado)
		fmt.Println("NOMBRE: ", empleado.nombre, "DIRECCION:", empleado.direccion, " FECHA DE NACIMIMENTO: ", empleado.fecha, "SALARIO: ", empleado.salario, " CATEGORIA: ", empleado.categoria, "DNI: ", empleado.dni)
	}
}
