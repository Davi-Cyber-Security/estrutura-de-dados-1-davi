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

/* Exercício 3 — Inserção em Posição Específica
Implemente adicionarPosicao(valor, posicao int) bool. Siga a regra de ouro vista em aula: primeiro novo.proximo = anterior.proximo, depois anterior.proximo = novo — nessa ordem, para nunca perder o restante da lista. Trate posição inválida (posicao < 0 ou maior que o tamanho da lista) retornando false.
Branch sugerida: 02-adicionar-posicao */

func (lista *Lista) adicionarPosicao(valor, posicao int) bool {
	novo := &no{valor: valor}
	atual := lista.inicio
	contador := 0

	if posicao < 0 {
		return false
	}

	if posicao == 0 {
		lista.adicionarInicio(valor)
	}

	for atual != nil && contador < posicao-1 {
		atual = atual.proximo
		contador++
	}

	if atual == nil {
		return false
	}

	novo.proximo = atual.proximo
	atual.proximo = novo

	return true
}

/*
Exercício 4 — Remoção no Início e no Fim
Implemente removerInicio() (int, bool) e removerFim() (int, bool), cada uma retornando o valor removido e um bool indicando sucesso. Trate o caso de lista vazia (retornando 0, false) e o caso de lista com um único elemento.
Branch sugerida: 03-remover-inicio-e-fim
*/

func (lista *Lista) removerInicio() (int, bool) {
	valorRemovido := lista.inicio

	if lista == nil {
		return 0, false
	}

	lista.inicio = lista.inicio.proximo

	return valorRemovido.valor, true
}

func (lista *Lista) removerFinal() (int, bool) {
	atual := lista.inicio

	if lista == nil {
		return 0, false
	}

	if lista.inicio.proximo == nil {
		valor := lista.inicio.valor
		lista.inicio = nil
		return valor, true
	}

	for atual.proximo.proximo != nil {

		atual = atual.proximo
	}

	valorRemovido := atual.proximo.valor
	atual.proximo = nil

	return valorRemovido, true
}

/* Exercício 5 — Remoção em Posição Específica
Implemente removerPosicao(posicao int) (int, bool), pulando o nó removido com anterior.proximo = anterior.proximo.proximo. Trate posição inválida e lista vazia.
Branch sugerida: 04-remover-posicao */

func (lista *Lista) removerPosicao(posicao int) bool {
	atual := lista.inicio
	contador := 0

	if posicao < 0 {
		return false
	}

	if posicao == 0 {
		lista.inicio = lista.inicio.proximo
		return true
	}

	for atual != nil && contador < posicao-1 {
		atual = atual.proximo
		contador++
	}

	if atual == nil || atual.proximo == nil {
		return false
	}

	atual.proximo = atual.proximo.proximo

	return true
}

func (lista *Lista) imprimir() {
	atual := lista.inicio

	if lista.inicio == nil {
		println("Não foi possível imprimir os resultados. A lista está vazia")
		return
	}
	for atual != nil {
		print(" ", atual.valor)

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

	println("\n\nInserindo na posição 3")
	ok := lista.adicionarPosicao(20, 3)
	println("Adicionado 20 na posição 3\n", ok)
	println("Imprimindo lista")
	lista.imprimir()

	valorRemovidoInicio, sucess := lista.removerInicio()
	println("\n\nValor removido do inicio", valorRemovidoInicio, sucess)
	println("\nLista:")
	lista.imprimir()

	valorRemovidoFinal, sucess := lista.removerFinal()
	println("\n\nValor removido do final", valorRemovidoFinal, sucess)
	println("\nLista:")
	lista.imprimir()

	removendoValorPosicao := lista.removerPosicao(3)
	println("\n\nValor removido da posição 3", removendoValorPosicao)
	println("\nLista:")
	lista.imprimir()
}
