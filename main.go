package main

import "fmt"

const Capacidade = 5

type FilaArray struct {
	dados  [Capacidade]int
	cabeca int
	cauda  int
	tam    int
}

func NovaFilaArray() *FilaArray {
	return &FilaArray{cabeca: 0, cauda: 0, tam: 0}
}

func (f *FilaArray) Enfileirar(v int) bool {
	if f.Cheia() {
		return false
	}
	f.dados[f.cauda] = v
	f.cauda = (f.cauda + 1) % Capacidade
	f.tam++
	return true
}

func (f *FilaArray) Desenfileirar() (int, bool) {
	if f.Vazia() {
		return 0, false
	}
	v := f.dados[f.cabeca]
	f.cabeca = (f.cabeca + 1) % Capacidade
	f.tam--
	return v, true
}

func (f *FilaArray) Frente() (int, bool) {
	if f.Vazia() {
		return 0, false
	}
	return f.dados[f.cabeca], true
}

func (f *FilaArray) Vazia() bool {
	return f.tam == 0
}

func (f *FilaArray) Cheia() bool {
	return f.tam == Capacidade
}

func main() {
	f := NovaFilaArray()

	valores := []int{10, 20, 30, 40, 50, 60}
	for _, v := range valores {
		ok := f.Enfileirar(v)
		fmt.Printf("Enfileirar(%d) → %v\n", v, ok)
	}

	v1, _ := f.Desenfileirar()
	v2, _ := f.Desenfileirar()
	fmt.Printf("Desenfileirou: %d %d\n", v1, v2)

	f.Enfileirar(60)
	f.Enfileirar(70)
	fmt.Println("Enfileirou 60 e 70 (wrap-around)")

	fmt.Print("Desenfileirados: ")
	for !f.Vazia() {
		v, ok := f.Desenfileirar()
		if ok {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Println()

	_, ok := f.Desenfileirar()
	fmt.Printf("Desenfileirar em fila vazia → ok == %v\n", ok)
}
