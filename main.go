package main

import "fmt"

type No struct {
	valor   int
	proximo *No
}

type PilhaEncadeada struct {
	topo    *No
	tamanho int
}

func NovaPilhaEncadeada() *PilhaEncadeada {
	return &PilhaEncadeada{topo: nil, tamanho: 0}
}

func (p *PilhaEncadeada) Empilhar(v int) {
	novo := &No{valor: v, proximo: p.topo}
	p.topo = novo
	p.tamanho++
}

func (p *PilhaEncadeada) Desempilhar() (int, bool) {
	if p.Vazia() {
		return 0, false
	}
	v := p.topo.valor
	p.topo = p.topo.proximo
	p.tamanho--
	return v, true
}

func (p *PilhaEncadeada) Topo() (int, bool) {
	if p.Vazia() {
		return 0, false
	}
	return p.topo.valor, true
}

func (p *PilhaEncadeada) Vazia() bool {
	return p.topo == nil
}

func (p *PilhaEncadeada) Tamanho() int {
	return p.tamanho
}

func main() {
	p := NovaPilhaEncadeada()

	p.Empilhar(1)
	p.Empilhar(2)
	p.Empilhar(3)
	topo, _ := p.Topo()
	fmt.Printf("Após empilhar 1, 2, 3 → Tamanho: %d, Topo: %d\n", p.Tamanho(), topo)

	for !p.Vazia() {
		v, _ := p.Desempilhar()
		fmt.Printf("Desempilhou: %d\n", v)
	}
	fmt.Printf("Após esvaziar → Tamanho: %d\n", p.Tamanho())

	fmt.Println("\n--- Sequência do Exercício 1 ---")
	valores := []int{10, 20, 30, 40, 50, 60}
	for _, v := range valores {
		p.Empilhar(v)
		fmt.Printf("Empilhar(%d)\n", v)
	}

	fmt.Print("Desempilhados: ")
	for !p.Vazia() {
		v, ok := p.Desempilhar()
		if ok {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Println()

	_, ok := p.Desempilhar()
	fmt.Printf("Desempilhar em pilha vazia → ok == %v\n", ok)
}
