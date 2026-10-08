package main

import "fmt"

type No struct {
	valor   int
	proximo *No
}

type FilaEncadeada struct {
	cabeca  *No
	cauda   *No
	tamanho int
}

func NovaFilaEncadeada() *FilaEncadeada {
	return &FilaEncadeada{cabeca: nil, cauda: nil, tamanho: 0}
}

func (f *FilaEncadeada) Enfileirar(v int) {
	novo := &No{valor: v, proximo: nil}
	if f.Vazia() {
		f.cabeca = novo
		f.cauda = novo
	} else {
		f.cauda.proximo = novo
		f.cauda = novo
	}
	f.tamanho++
}

func (f *FilaEncadeada) Desenfileirar() (int, bool) {
	if f.Vazia() {
		return 0, false
	}
	v := f.cabeca.valor
	f.cabeca = f.cabeca.proximo
	if f.cabeca == nil {
		f.cauda = nil
	}
	f.tamanho--
	return v, true
}

func (f *FilaEncadeada) Frente() (int, bool) {
	if f.Vazia() {
		return 0, false
	}
	return f.cabeca.valor, true
}

func (f *FilaEncadeada) Vazia() bool {
	return f.cabeca == nil
}

func (f *FilaEncadeada) Tamanho() int {
	return f.tamanho
}

func main() {
	f := NovaFilaEncadeada()

	f.Enfileirar(1)
	f.Enfileirar(2)
	fmt.Printf("Após enfileirar 1 e 2 → Tamanho: %d\n", f.Tamanho())

	v1, _ := f.Desenfileirar()
	v2, _ := f.Desenfileirar()
	fmt.Printf("Desenfileirou: %d %d\n", v1, v2)
	fmt.Printf("Fila vazia → Tamanho: %d, Vazia: %v\n", f.Tamanho(), f.Vazia())

	f.Enfileirar(3)
	frente, _ := f.Frente()
	fmt.Printf("Após enfileirar 3 → Frente: %d, Tamanho: %d\n", frente, f.Tamanho())
}
