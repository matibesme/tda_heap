package cola_prioridad

const(
	TAM_INICIAL = 7
	MSG_COLA_VACIA = "La cola esta vacia"
	FACTOR_REDIMENSION = 2
)

type heap[T any] struct {
	datos []T
	cantidad int
	cmp funcCmp[T]
}

type funcCmp[T any] func(T, T) int

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &heap{
		datos: make([]T, TAM_INICIAL),
		cantidad: 0,
		cmp: funcion_cmp,
	}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &heap{
		datos: heapify(arreglo),
		cantidad: len(arreglo),
		cmp: funcion_cmp,
	}
}

//PRIMITIVAS
func (heap *heap[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap *heap[T]) Encolar(dato T) {
	if heap.cantidad == cap(heap.datos) {
		heap.redimensionar(cap(heap.datos) * FACTOR_REDIMENSION)
	}
	heap.datos[heap.cantidad] = dato
	heap.cantidad++
	heap.upHeap(cantidad-1)
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
	if heap.cantidad * 4 <= cap(heap.datos) && cap(heap.datos) > TAM_INICIAL {
		heap.redimensionar(cap(heap.datos) / FACTOR_REDIMENSION)
	}
	heap.datos[0], heap.datos[heap.cantidad-1] = heap.datos[heap.cantidad-1], heap.datos[0]
	desencolado := heap.datos[heap.cantidad-1]
	heap.cantidad--
	heap.downHeap(heap.datos[0])
	return desencolado
}

func (heap *heap[T]) Cantidad() int {
	return heap.cantidad
}

//UPHEAP Y DOWNHEAP
func  upHeap[T any](datos []T, pos int, cmp func(T, T) int) {
	pos_padre := buscarPadre(pos)
	if cmp(datos[pos_padre], datos[pos]) > 0 {
		return 
	}
	datos[pos_padre], datos[pos] = datos[pos], datos[pos_padre]
	upHeap(datos,pos_padre,cmp)
}

func buscarPadre(pos int) int {
	pos_padre := (pos - 1) / 2
	return pos_padre
}

func downHeap[T any](datos []T, pos int, cmp func(T, T) int) {
	pos_hijo_izq := buscarHijoIzq(pos)
	pos_hijo_der := buscarHijoDer(pos)
	if pos_hijo_izq > (heap.cantidad - 1) && pos_hijo_der > (heap.cantidad - 1) {
		return
	} else if pos_hijo_der > (heap.cantidad - 1) {
		if cmp(heap.datos[pos_hijo_izq], datos[pos]) < 0 {
			return
		} else {
			datos[pos_hijo_izq], datos[pos] = datos[pos], datos[pos_hijo_izq]
		}
	} else {
		if cmp(datos[pos_hijo_izq], datos[pos_hijo_der]) >= 0 {
			datos[pos_hijo_izq],heap.datos[pos] = datos[pos], datos[pos_hijo_izq]
			downHeap(datos,pos_hijo_izq,cmp)
		} else {
			datos[pos_hijo_der],datos[pos] = datos[pos], datos[pos_hijo_der]
			downHeap(datos,pos_hijo_der,cmp)
		}
	}
}

func buscarHijoIzq(pos int) int {
	pos_hijo_izq := pos * 2 + 1
	return pos_hijo_izq
}

func buscarHijoDer(pos int) int {
	pos_hijo_der := pos * 2 + 2
	return pos_hijo_der
}

//HEAPIFY
func heapify[T any](arr []T, cmp func(T, T) int) []T {
	for i := len(arr)- 1; i >= 0; i-- {
		downHeap(arr,i,cmp)
	}
}

