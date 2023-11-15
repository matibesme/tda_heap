package cola_prioridad

const (
	TAM_INICIAL              = 7
	MSG_COLA_VACIA           = "La cola esta vacia"
	FACTOR_REDIMENSION       = 2
	DOBLE_FACTOR_REDIMENSION = 4
)

type heap[T any] struct {
	datos    []T
	cantidad int
	cmp      funcCmp[T]
}

type funcCmp[T any] func(T, T) int

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{
		datos:    make([]T, TAM_INICIAL),
		cantidad: 0,
		cmp:      funcion_cmp,
	}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	if len(arreglo) == 0 {
		return CrearHeap[T](funcion_cmp)
	}
	capacidadInicial := TAM_INICIAL
	if len(arreglo) > TAM_INICIAL {
		capacidadInicial = len(arreglo)
	}
	copiaArreglo := make([]T, len(arreglo))
	copy(copiaArreglo, arreglo)
	heap := new(heap[T])
	heap.cantidad = len(copiaArreglo)
	heap.cmp = funcion_cmp
	heapify[T](copiaArreglo, funcion_cmp)
	heap.redimensionar(capacidadInicial)
	heap.datos = copiaArreglo
	return heap
}

// PRIMITIVAS
func (heap *heap[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap *heap[T]) Encolar(dato T) {
	if heap.cantidad == cap(heap.datos) {
		heap.redimensionar(cap(heap.datos) * FACTOR_REDIMENSION)
	}
	heap.datos[heap.cantidad] = dato
	heap.cantidad++
	upHeap(heap.datos, heap.cantidad-1, heap.cmp)

}

func (heap *heap[T]) VerMax() T {
	if heap.EstaVacia() {
		panic(MSG_COLA_VACIA)
	}
	return heap.datos[0]
}

func (heap *heap[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic(MSG_COLA_VACIA)
	}
	if heap.cantidad*DOBLE_FACTOR_REDIMENSION <= cap(heap.datos) && cap(heap.datos) > TAM_INICIAL {
		heap.redimensionar(cap(heap.datos) / FACTOR_REDIMENSION)
	}
	swap(heap.datos, 0, heap.cantidad-1)
	desencolado := heap.datos[heap.cantidad-1]
	heap.cantidad--
	downHeap(heap.datos, 0, heap.cantidad, heap.cmp)
	return desencolado
}

func (heap *heap[T]) Cantidad() int {
	return heap.cantidad
}

// SWAP
func swap[T any](datos []T, i, j int) {
	datos[i], datos[j] = datos[j], datos[i]
}

// UPHEAP Y DOWNHEAP
func upHeap[T any](datos []T, pos int, cmp func(T, T) int) {
	pos_padre := buscarPadre(pos)
	if cmp(datos[pos_padre], datos[pos]) < 0 {
		swap(datos, pos_padre, pos)
		upHeap(datos, pos_padre, cmp)
	}
}

func buscarPadre(pos int) int {
	pos_padre := (pos - 1) / 2
	return pos_padre
}

func downHeap[T any](datos []T, pos int, cantidad int, cmp func(T, T) int) {
	pos_hijo_izq := buscarHijoIzq(pos)
	pos_hijo_der := buscarHijoDer(pos)
	if pos_hijo_izq >= cantidad {
		return
	}
	if pos_hijo_der >= cantidad {
		if cmp(datos[pos_hijo_izq], datos[pos]) > 0 {
			swap(datos, pos_hijo_izq, pos)
		}
		return
	}
	pos_hijo_mayor := hijoMasGrande[T](datos, pos_hijo_izq, pos_hijo_der, cmp)
	if cmp(datos[pos_hijo_mayor], datos[pos]) > 0 {
		swap(datos, pos_hijo_mayor, pos)
		downHeap[T](datos, pos_hijo_mayor, cantidad, cmp)
	}
}

func hijoMasGrande[T any](datos []T, pos_hijo_izq int, pos_hijo_der int, cmp func(T, T) int) int {
	if cmp(datos[pos_hijo_izq], datos[pos_hijo_der]) >= 0 {
		return pos_hijo_izq
	}
	return pos_hijo_der
}

func buscarHijoIzq(pos int) int {
	pos_hijo_izq := pos*2 + 1
	return pos_hijo_izq
}

func buscarHijoDer(pos int) int {
	pos_hijo_der := pos*2 + 2
	return pos_hijo_der
}

// REDIMENSION
func (heap *heap[T]) redimensionar(nuevo_tam int) {
	nuevos_datos := make([]T, nuevo_tam)
	copy(nuevos_datos, heap.datos)
	heap.datos = nuevos_datos
}

// HEAPIFY
func heapify[T any](arreglo []T, cmp func(T, T) int) {
	for pos := len(arreglo) - 1; pos >= 0; pos-- {
		downHeap[T](arreglo, pos, len(arreglo), cmp)
	}
}

// HEAPSORT
func HeapSort[T any](lista []T, cmp func(T, T) int) {
	heapify[T](lista, cmp)
	for i := len(lista) - 1; i > 0; i-- {
		lista[0], lista[i] = lista[i], lista[0]
		downHeap(lista, 0, i, cmp)
	}
}
