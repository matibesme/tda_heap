package cola_prioridad_test

import (
	TDAHeap "tdas/heap"
	"testing"
	"github.com/stretchr/testify/require"
)

//CONSTANTES
var(
	NUMEROS  = []int{5, 4, 3, 2, 1}
	PALABRAS = []string{"B", "C", "J", "E", "A"}
)

const(
	MSG_PANIC = "La cola esta vacia"
	TAMANO_VOLUMEN = 1000
)

//Comparación de strings
func compararStrings(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

//Comparación de números
func compararNumerosEnteros(a, b int) int {
	return a - b
}

//TESTS
func TestHeapVacio(t *testing.T) {
	t.Log("Vemos el comportamiento de una pila recien creada")
	heap := TDAHeap.CrearHeap[string](compararStrings)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, MSG_PANIC, func() { heap.VerMax() })
	require.PanicsWithValue(t, MSG_PANIC, func() { heap.Desencolar() })
}

func TestEncolarElementos(t *testing.T) {
	t.Log("Probamos encolar elementos, y que al desencolarlos se mantenga la prioridad")
	heap := TDAHeap.CrearHeap[int](compararNumerosEnteros)
	for _, numero := range NUMEROS {
		heap.Encolar(numero)
	}
	require.EqualValues(t, 5, heap.VerMax())
	
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
	valorDesencolado := heap.Desencolar()
	require.EqualValues(t, "J", valorDesencolado)
	require.EqualValues(t, "E", heap.VerMax())
}

