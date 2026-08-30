package main

type no struct {
	valor   int
	proximo *no
}

type Lista struct {
	inicio *no
}

/* Exercício 2 — Struct e Inserção no Início e no Fim
Defina as structs no e lista mostradas acima e implemente adicionarInicio(valor int) e adicionarFim(valor int) como métodos de *lista. Escreva um main() que insira ao menos 5 valores (misturando início e fim) e confira o resultado.
Branch sugerida: 01-adicionar-inicio-e-fim */

func (lista *Lista) adicionarInicio(valor int) {
	novo := &no{valor: valor}

	novo.proximo = lista.inicio
	lista.inicio = novo
}

func (lista *Lista) adicionarFim(valor int) {
	novo := &no{valor: valor}
	atual := lista.inicio

	if lista.inicio == nil {
		lista.inicio = novo
		return
	}

	for atual.proximo != nil {
		atual = atual.proximo
	}

	atual.proximo = novo
}

func (lista *Lista) imprimir() {
	atual := lista.inicio

	if lista.inicio == nil {
		println("Não foi possível imprimir os resultados. A lista está vazia")
		return
	}
	for atual != nil {
		println(atual.valor)

		atual = atual.proximo
	}
}

func main() {
	var lista Lista

	lista.adicionarInicio(5)
	lista.adicionarFim(10)
	lista.adicionarFim(18)
	lista.adicionarInicio(15)
	lista.adicionarInicio(12)

	lista.imprimir()
}
