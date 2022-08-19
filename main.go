package main

import (
	"container/heap"
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

}
