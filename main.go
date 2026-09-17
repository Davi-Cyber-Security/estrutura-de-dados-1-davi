package main

type No struct {
	valor int
	PROX  *No
	PREV  *No
}

type Lista struct {
	Head *No
	TAIL *No
}

func (lista *Lista) inserirNoInicio(valor int) {
	novo := &No{valor: valor}

	if lista.Head == nil {
		lista.Head = novo
		lista.TAIL = novo
	} else {
		novo.PROX = lista.Head
		lista.Head.PREV = novo
		lista.Head = novo
	}

}

func (lista *Lista) inserirNoFinal(valor int) {
	atual := lista.Head
	novo := &No{valor: valor}

	if atual == nil {
		lista.Head = novo
		lista.TAIL = novo
		return
	} else {
		for atual.PROX != nil {
			atual = atual.PROX
		}

		novo.PREV = atual
		atual.PROX = novo
	}
}

func (lista *Lista) removerNoInicio() {

	if lista.Head == nil {
		println("Não é possível remover do inicio. A lista está vazio!")
		return
	} else {
		lista.Head = lista.Head.PROX
		lista.Head.PREV = lista.Head
	}

}

func (lista *Lista) removerNoFinal() {
	atual := lista.Head

	if atual == nil {
		println("Não é possível remover do inicio. A lista está vazio!")
		return
	} else {
		for atual.PROX != nil {
			atual = atual.PROX
		}

		atual.PREV.PROX = nil
	}
}

func (lista *Lista) imprimir() {
	atual := lista.Head

	for atual != nil {
		print(atual.valor, " ")

		atual = atual.PROX
	}

	println()
}

func main() {
	lista := &Lista{}
	lista.inserirNoInicio(10)
	lista.inserirNoInicio(20)
	lista.inserirNoInicio(50)
	lista.inserirNoInicio(60)
	lista.inserirNoInicio(80)
	lista.inserirNoFinal(100)
	lista.removerNoInicio()
	lista.removerNoFinal()

	lista.imprimir()
}
