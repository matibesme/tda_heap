package cola_prioridad_test

import (
	"github.com/stretchr/testify/require"
	TDAHeap "tdas/cola_prioridad"
	"testing"
)

// CONSTANTES
var (
	NUMEROS  = []int{5, 4, 3, 2, 1}
	PALABRAS = []string{"B", "C", "J", "E", "A"}
)

const (
	MSG_PANIC      = "La cola esta vacia"
	TAMANO_VOLUMEN = 1000
)

// Comparación de strings
func compararStrings(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// Comparación de números
func compararNumerosEnteros(a, b int) int {
	return a - b
}

// TESTS
func TestHeapVacio(t *testing.T) {
	t.Log("Vemos el comportamiento de una pila recien creada")
	heap := TDAHeap.CrearHeap[string](compararStrings)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, MSG_PANIC, func() { heap.VerMax() })
	require.PanicsWithValue(t, MSG_PANIC, func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())
}

func TestEncolarElementos(t *testing.T) {
	t.Log("Probamos encolar elementos, y que al desencolarlos se mantenga la prioridad")
	heap := TDAHeap.CrearHeap[int](compararNumerosEnteros)
	for _, numero := range NUMEROS {
		heap.Encolar(numero)
	}
	require.EqualValues(t, 5, heap.VerMax())
	require.EqualValues(t, len(NUMEROS), heap.Cantidad())

	for _, numero := range NUMEROS {
		require.EqualValues(t, numero, heap.VerMax())
		valorDesencolado := heap.Desencolar()
		require.EqualValues(t, numero, valorDesencolado)
	}
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, MSG_PANIC, func() { heap.VerMax() })
	require.PanicsWithValue(t, MSG_PANIC, func() { heap.Desencolar() })

}

func TestDeVolumen(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](compararNumerosEnteros)
	for n := 1; n <= TAMANO_VOLUMEN; n++ {
		heap.Encolar(n)
	}
	require.EqualValues(t, TAMANO_VOLUMEN, heap.VerMax())
	for n := TAMANO_VOLUMEN; n > 1; n-- {
		valorDesencolado := heap.Desencolar()
		require.EqualValues(t, n, valorDesencolado)
		require.EqualValues(t, n-1, heap.VerMax())
	}
}

func TestComportamientoPostVaciar(t *testing.T) {
	t.Log("Comprobamos que al desencolar hasta que está vacía hace que la cola se comporte como recién creada.")
	heap := TDAHeap.CrearHeap[int](compararNumerosEnteros)
	heap.Encolar(NUMEROS[0])
	require.EqualValues(t, NUMEROS[0], heap.VerMax())
	ValorDesencolado := heap.Desencolar()
	require.EqualValues(t, NUMEROS[0], ValorDesencolado)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, MSG_PANIC, func() { heap.VerMax() })
	require.PanicsWithValue(t, MSG_PANIC, func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())
}

func TestVerPimeroDesencolarNuevaCola(t *testing.T) {
	t.Log("Probamos que las acciones de desencolar y ver_max en una cola recién creada sean inválidas")
	heap := TDAHeap.CrearHeap[int](compararNumerosEnteros)
	require.PanicsWithValue(t, MSG_PANIC, func() { heap.VerMax() })
	require.PanicsWithValue(t, MSG_PANIC, func() { heap.Desencolar() })
}

func TestNuevaColaEstaVacia(t *testing.T) {
	t.Log("Probamos que la acción de esta_vacía en una cola recién creada es verdadero")
	heap := TDAHeap.CrearHeap[int](compararNumerosEnteros)
	require.True(t, heap.EstaVacia())
}

func TestColaStrings(t *testing.T) {
	t.Log("Probamos encolar diferentes tipos de datos. En este caso strings")
	heap := TDAHeap.CrearHeap[string](compararStrings)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, MSG_PANIC, func() { heap.VerMax() })
	require.PanicsWithValue(t, MSG_PANIC, func() { heap.Desencolar() })
	for _, palabra := range PALABRAS {
		heap.Encolar(palabra)
	}
	require.EqualValues(t, "J", heap.VerMax())
	valorDesencolado := heap.Desencolar()
	require.EqualValues(t, "J", valorDesencolado)
	require.EqualValues(t, "E", heap.VerMax())
}

func TestCrearHeapArreglo(t *testing.T) {
	t.Log("Probamos que a partir de un arreglo se cree el heap con su correcta prioridad")
	heap := TDAHeap.CrearHeapArr[string](PALABRAS, compararStrings)
	require.EqualValues(t, "J", heap.VerMax())
	require.EqualValues(t, len(PALABRAS), heap.Cantidad())
	valorDesencolado := heap.Desencolar()
	require.EqualValues(t, "J", valorDesencolado)
	require.EqualValues(t, "E", heap.VerMax())
}

func TestCrearHeapArregloVacio(t *testing.T) {
	t.Log("Probamos que al crear un heap desde un arreglo vacio no se rompa y se pueda utilizar con normalidad")
	heap := TDAHeap.CrearHeapArr[int](make([]int, 0), compararNumerosEnteros)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, MSG_PANIC, func() { heap.VerMax() })
	require.PanicsWithValue(t, MSG_PANIC, func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())
	for _, numero := range NUMEROS {
		heap.Encolar(numero)
	}
	require.EqualValues(t, 5, heap.VerMax())
	require.EqualValues(t, len(NUMEROS), heap.Cantidad())

	for _, numero := range NUMEROS {
		require.EqualValues(t, numero, heap.VerMax())
		valorDesencolado := heap.Desencolar()
		require.EqualValues(t, numero, valorDesencolado)
	}
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, MSG_PANIC, func() { heap.VerMax() })
	require.PanicsWithValue(t, MSG_PANIC, func() { heap.Desencolar() })
}

func TestHeapSortString(t *testing.T) {
	t.Log("Probamos la herramienta de heapsort en un arreglo de strings")
	arr_ordenado := []string{"A", "B", "C", "E", "J"}
	arr_desordenado := []string{"B", "C", "J", "E", "A"}

	TDAHeap.HeapSort[string](arr_desordenado, compararStrings)

	for i, letra := range arr_ordenado {
		require.EqualValues(t, letra, arr_desordenado[i])
	}

}

func TestHeapSortInt(t *testing.T) {
	t.Log("Probamos la herramienta de heapsort en un arreglo de enteros")
	arr_ordenado := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr_desordenado := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}

	TDAHeap.HeapSort[int](arr_desordenado, compararNumerosEnteros)

	for i, letra := range arr_ordenado {
		require.EqualValues(t, letra, arr_desordenado[i])
	}

}

func TestHeapSortArregloVacio(t *testing.T) {
	t.Log("Probamos la herramienta de heapsort en un arreglo de enteros")
	arr_ordenado := []string{""}
	arr_desordenado := []string{""}
	TDAHeap.HeapSort[string](arr_desordenado, compararStrings)
	require.EqualValues(t, arr_ordenado[0], arr_desordenado[0])
}
