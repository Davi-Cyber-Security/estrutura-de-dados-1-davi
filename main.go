package main

import "fmt"

const Capacidade = 5

type PilhaArray struct {
	dados [Capacidade]int
	topo  int
}

func NovaPilhaArray() *PilhaArray {
	return &PilhaArray{topo: -1}
}

func (p *PilhaArray) Empilhar(v int) bool {
	if p.Cheia() {
		return false
	}
	p.topo++
	p.dados[p.topo] = v
	return true
}

func (p *PilhaArray) Desempilhar() (int, bool) {
	if p.Vazia() {
		return 0, false
	}

	v := p.dados[p.topo]
	p.topo--

	return v, true
}

func (p *PilhaArray) Topo() (int, bool) {
	if p.Vazia() {
		return 0, false
	}

	return p.dados[p.topo], true
}

func (p *PilhaArray) Vazia() bool {
	return p.topo == -1
}

func (p *PilhaArray) Cheia() bool {
	return p.topo == Capacidade-1
}

func main() {
	fmt.Println("Exercício 1")
	p := NovaPilhaArray()

	valores := []int{10, 20, 30, 40, 50, 60}
	for _, v := range valores {
		ok := p.Empilhar(v)
		fmt.Printf("Empilhar(%d) -> %v\n", v, ok)
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
	fmt.Printf("Desempilhar em pilha vazia -> ok == %v\n", ok)

	println()
	fmt.Println("Exercício 2")

}
