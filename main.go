package main

import "fmt"

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
		return true
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

func (lista *Lista) removerPosicao(posicao int) (int, bool) {
	atual := lista.inicio
	valorAnterior := 0
	contador := 0

	if posicao < 0 {
		return 0, false
	}

	if posicao == 0 {
		lista.inicio = lista.inicio.proximo
		return 0, true
	}

	for atual != nil && contador < posicao-1 {
		atual = atual.proximo
		contador++
	}

	if atual == nil || atual.proximo == nil {
		return 0, false
	}

	valorAnterior = atual.proximo.valor

	atual.proximo = atual.proximo.proximo

	return valorAnterior, true
}

/* Exercício 6 — Busca por Valor
Implemente posicao(valorProcurado int) (int, bool), que percorre a lista comparando atual.valor a cada passo e retorna o índice onde o valor foi encontrado (ou 0, false se não existir).
Branch sugerida: 05-encontrar-posicao */

func (lista *Lista) posicao(valorProcurado int) (int, bool) {
	contador := 0
	atual := lista.inicio

	if atual == nil || valorProcurado < 0 {
		return 0, false
	}

	for atual != nil && contador < valorProcurado-1 {

		if atual.valor == valorProcurado {
			return contador, true
		}

		atual = atual.proximo
		contador++
	}

	return 0, false
}

/* Exercício 7 — Busca por Posição
Implemente valorNaPosicao(posicaoProcurada int) (int, bool), que percorre a lista até o índice pedido e retorna o valor armazenado ali (ou 0, false se a posição não existir).
Branch sugerida: 06-valor-na-posicao */

func (lista *Lista) valorNaPosicao(posicaoProcurada int) (int, bool) {
	atual := lista.inicio
	contador := 0

	if atual == nil || posicaoProcurada < 0 {
		return 0, false
	}

	for atual != nil && contador < posicaoProcurada {
		atual = atual.proximo
		contador++
	}

	if atual == nil {
		return 0, false
	}

	return atual.valor, true
}

/* Exercício 8 — Tamanho, Impressão e Programa Integrador
Implemente tamanho() int, que percorre a lista e conta os nós, e imprimir(), que percorre a lista exibindo os valores no formato:
10 -> 20 -> 30 -> nil

Em seguida, una todas as operações dos Exercícios 2 a 7 em um único programa Go com um menu interativo no terminal (usando fmt.Scan). Teste cada operação com lista vazia, lista de 1 elemento e índices fora do intervalo — confirme que ela devolve false/0 como esperado nesses casos.
Branch sugerida: 07-tamanho-imprimir-integrador */

func (lista *Lista) tamanho() int {
	atual := lista.inicio
	contador := 0

	if atual == nil {
		return 0
	}

	for atual != nil {
		atual = atual.proximo
		contador++
	}

	return contador
}

func (lista *Lista) imprimir() {
	atual := lista.inicio

	if lista.inicio == nil {
		println("Não foi possível imprimir os resultados. A lista está vazia")
		return
	}
	for atual != nil {
		fmt.Printf("-> %d ", atual.valor)

		atual = atual.proximo

		if atual == nil {
			fmt.Printf("-> nil ")
		}
	}

}

func main() {
	var lista Lista
	var opcao int
	var valores int
	var aux int

	for {
		print("\n1 - Adicionar valor no inicio da lista\n2 - Adicionar no final da lista\n3 - Adicionar na posição\n4 - Remover valor do inicio\n5 - Remover valor do final\n6 - Remover da posição\n7 - Procurar valor na lista\n8 - Procurar valor pela posição\n9 - Listar tamanho da lista\n10 - Fechar programa\n\nEscolha a opção: ")
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			for {
				fmt.Print("\nAdicione o valor: ")
				fmt.Scan(&valores)
				lista.adicionarInicio(valores)
				println("Valor adicionado com sucesso!")

				println("\nLista:")
				lista.imprimir()
				println()

				print("\n1 - Sim\n2 - Não\nDeseja continuar? ")
				fmt.Scan(&opcao)

				if opcao == 2 {
					break
				}
			}
		case 2:
			for {
				fmt.Print("\nAdicione o valor: ")
				fmt.Scan(&valores)
				lista.adicionarFim(valores)
				println("Valor adicionado com sucesso!")

				println("\nLista:")
				lista.imprimir()
				println()

				print("\n01 - Sim\n02 - Não\nDeseja continuar? ")
				fmt.Scan(&opcao)

				if opcao == 2 {
					break
				}
			}
		case 3:
			for {
				fmt.Print("\nAdicione o valor na posicao: ")
				fmt.Scan(&valores)
				fmt.Print("\nDigite a posição para ser inserido: ")
				fmt.Scan(&aux)
				ok := lista.adicionarPosicao(valores, aux)
				fmt.Printf("Adicionado %d na posição %d status %t \n", valores, aux, ok)

				println("\nLista:")
				lista.imprimir()
				println()

				print("\n01 - Sim\n02 - Não\nDeseja continuar? ")
				fmt.Scan(&opcao)

				if opcao == 2 {
					break
				}
			}
		case 4:
			for {
				valorRemovidoInicio, sucess := lista.removerInicio()
				fmt.Printf("\n\nValor %d removido do inicio. Status: %t", valorRemovidoInicio, sucess)

				println("\nLista:")
				lista.imprimir()
				println()

				print("\n01 - Sim\n02 - Não\nDeseja continuar? ")
				fmt.Scan(&opcao)

				if opcao == 2 {
					break
				}
			}
		case 5:
			for {
				valorRemovidoInicio, sucess := lista.removerFinal()
				fmt.Printf("\n\nValor %d removido do final. Status: %t", valorRemovidoInicio, sucess)

				println("\nLista:")
				lista.imprimir()
				println()

				print("\n01 - Sim\n02 - Não\nDeseja continuar? ")
				fmt.Scan(&opcao)

				if opcao == 2 {
					break
				}
			}
		case 6:
			for {
				print("Digite a posição que deseja remover: ")
				fmt.Scan(&opcao)
				removendoValorPosicao, sucess := lista.removerPosicao(opcao)
				fmt.Printf("\n\nValor removido da posição %d valor removido: %d status: %t", opcao, removendoValorPosicao, sucess)

				println("\nLista:")
				lista.imprimir()
				println()

				print("\n01 - Sim\n02 - Não\nDeseja continuar? ")
				fmt.Scan(&opcao)

				if opcao == 2 {
					break
				}
			}
		case 7:
			for {
				print("Digite o valor que deseja buscar pelo valor: ")
				fmt.Scan(&opcao)
				buscandoValorPosicao, sucess := lista.posicao(opcao)
				fmt.Printf("\n\nValor procurado: %d índice: %d status: %t", opcao, buscandoValorPosicao, sucess)

				println("\nLista:")
				lista.imprimir()
				println()

				print("\n01 - Sim\n02 - Não\nDeseja continuar? ")
				fmt.Scan(&opcao)

				if opcao == 2 {
					break
				}
			}
		case 8:
			for {
				print("Digite a posição que deseja buscar pela posição: ")
				fmt.Scan(&opcao)
				buscandoValorNaPosicao, sucess := lista.valorNaPosicao(opcao)
				fmt.Printf("\n\nPosição procurado: %d valor: %d status: %t", opcao, buscandoValorNaPosicao, sucess)

				print("\n01 - Sim\n02 - Não\nDeseja continuar? ")
				fmt.Scan(&opcao)

				if opcao == 2 {
					break
				}
			}
		case 9:
			for {
				tamanho := lista.tamanho()
				fmt.Printf("Tamanho da lista: %d", tamanho)

				print("\n\n01 - Sim\n02 - Não\nDeseja continuar? ")
				fmt.Scan(&opcao)

				if opcao == 2 {
					break
				}
			}

		}

		if opcao == 10 {
			break
		}
	}
}
