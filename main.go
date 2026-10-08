package main

import (
	"fmt"
	"strconv"
	"strings"
)

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

func AvaliarRPN(expr string) (int, error) {
	p := NovaPilhaEncadeada()
	tokens := strings.Fields(expr)

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			b, okB := p.Desempilhar()
			a, okA := p.Desempilhar()
			if !okA || !okB {
				return 0, fmt.Errorf("operandos insuficientes")
			}
			var resultado int
			switch token {
			case "+":
				resultado = a + b
			case "-":
				resultado = a - b
			case "*":
				resultado = a * b
			case "/":
				if b == 0 {
					return 0, fmt.Errorf("divisão por zero")
				}
				resultado = a / b
			}
			p.Empilhar(resultado)
		default:
			n, err := strconv.Atoi(token)
			if err != nil {
				return 0, fmt.Errorf("token inválido: %s", token)
			}
			p.Empilhar(n)
		}
	}

	if p.Tamanho() == 0 {
		return 0, fmt.Errorf("pilha vazia")
	}
	if p.Tamanho() > 1 {
		return 0, fmt.Errorf("mais de um valor na pilha")
	}
	resultado, _ := p.Desempilhar()
	return resultado, nil
}

func main() {
	casos := []string{
		"3 4 + 2 *",
		"5 1 2 + 4 * + 3 -",
		"3 +",
		"1 2",
		"4 0 /",
		"2 x +",
	}

	for _, expr := range casos {
		resultado, err := AvaliarRPN(expr)
		if err != nil {
			fmt.Printf("%s → erro: %v\n", expr, err)
		} else {
			fmt.Printf("%s → %d\n", expr, resultado)
		}
	}
}
