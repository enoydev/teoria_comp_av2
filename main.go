package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

// --- START: Standalone constraints package implementation ---
// Ordered is a constraint that permits any type that supports the <, <=, >, >= operators.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// --- END: Standalone constraints package implementation ---

// --- START: Standalone math/min package implementation ---
// Min returns the smaller of x or y.
func Min[T Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// --- END: Standalone math/min package implementation ---

// --- START: Your Merge Sort Implementations ---
func merge[T Ordered](a []T, b []T) []T {
	var r = make([]T, len(a)+len(b))
	var i = 0
	var j = 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}
	}

	for i < len(a) {
		r[i+j] = a[i]
		i++
	}
	for j < len(b) {
		r[i+j] = b[j]
		j++
	}

	return r
}

func Merge[T Ordered](items []T) []T {
	if len(items) < 2 {
		return items
	}

	var middle = len(items) / 2
	var a = Merge(items[:middle])
	var b = Merge(items[middle:])
	return merge(a, b)
}

func MergeIter[T Ordered](items []T) []T {
	for step := 1; step < len(items); step += step {
		for i := 0; i+step < len(items); i += 2 * step {
			endOfSecondHalf := Min(i+2*step, len(items))
			tmp := merge(items[i:i+step], items[i+step:endOfSecondHalf])
			copy(items[i:], tmp)
		}
	}
	return items
}

func ParallelMerge[T Ordered](items []T) []T {
	if len(items) < 2 {
		return items
	}

	if len(items) < 2048 {
		return Merge(items)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	var middle = len(items) / 2
	var a []T
	go func() {
		defer wg.Done()
		a = ParallelMerge(items[:middle])
	}()
	var b = ParallelMerge(items[middle:])

	wg.Wait()
	return merge(a, b)
}

// --- END: Your Merge Sort Implementations ---

// --- Data Generation Functions for Go ---
func generateRandomSlice(size int) []int {
	s := make([]int, size)
	for i := 0; i < size; i++ {
		s[i] = rand.Intn(size * 10)
	}
	return s
}

func generateSortedSlice(size int) []int {
	s := make([]int, size)
	for i := 0; i < size; i++ {
		s[i] = i
	}
	return s
}

func generateReverseSortedSlice(size int) []int {
	s := make([]int, size)
	for i := 0; i < size; i++ {
		s[i] = size - 1 - i
	}
	return s
}

func isSliceSorted[T Ordered](s []T) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}
	}
	return true
}

// calculateMeanAndStdDev calcula a média e o desvio padrão de um slice de int64 (nanosegundos)
func calculateMeanAndStdDev(times []int64) (float64, float64) {
	if len(times) == 0 {
		return 0, 0
	}

	var sum int64
	for _, t := range times {
		sum += t
	}
	mean := float64(sum) / float64(len(times))

	if len(times) < 2 { // Desvio padrão não é significativo para 0 ou 1 amostra
		return mean, 0
	}

	var sumSqDiff float64
	for _, t := range times {
		diff := float64(t) - mean
		sumSqDiff += diff * diff
	}
	stdDev := math.Sqrt(sumSqDiff / float64(len(times)-1)) // Sample standard deviation

	return mean, stdDev
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Initialize random number generator

	N_VALUES := []int{10, 100, 1000, 5000, 10000, 50000, 100000, 200000, 500000} // Ajustei para incluir 'pequenas'
	INPUT_TYPES := map[string]func(int) []int{
		"aleatorio": generateRandomSlice,
		"ordenado":  generateSortedSlice,
		"inverso":   generateReverseSortedSlice,
	}
	NUM_TRIALS := 30 // Número de execuções para cada entrada

	fmt.Println("### Go Merge Sort (Recursivo) - Tempos de Execução (nanosegundos)")
	fmt.Printf("%-10s | %-15s | %-15s | %-15s | %-10s\n", "Tamanho N", "Tipo Entrada", "Média (ns)", "Desvio Padrão (ns)", "Ordenado?")
	fmt.Println("------------------------------------------------------------------")

	for _, n := range N_VALUES {
		for inputName, generator := range INPUT_TYPES {
			var executionTimes []int64
			var sortedFlag bool

			for i := 0; i < NUM_TRIALS; i++ {
				data := generator(n)
				// Faça uma cópia para a função de sorteio, para garantir que cada teste seja com dados "originais"
				dataCopy := make([]int, len(data))
				copy(dataCopy, data)

				start := time.Now()
				sortedData := Merge(dataCopy) // Usando a função Merge recursiva
				duration := time.Since(start)
				executionTimes = append(executionTimes, duration.Nanoseconds())

				if i == 0 { // Verifique a ordenação apenas uma vez por N e tipo de entrada
					sortedFlag = isSliceSorted(sortedData)
				}
			}

			mean, stdDev := calculateMeanAndStdDev(executionTimes)
			fmt.Printf("%-10d | %-15s | %-15.2f | %-19.2f | %-10t\n", n, inputName, mean, stdDev, sortedFlag)
		}
	}
	fmt.Println("------------------------------------------------------------------")
}
