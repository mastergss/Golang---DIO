// Criar uma API Rest com todos os verbos HTTP
// Utilizar o pacote gorilla/mux 'go get -u github.com/gorilla/mux'

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	
	"github.com/gorilla/mux"
)

// estrutura para armazenar dados do usuário, e objeto 'Localização'
type Usuario struct {
	ID			int			`json:"id"`
	Nome		string		`json:"nome"`
	Idade		int			`json:"idade"`
	Localizacao	Localizacao	`json:"localizacao"`
}

// estrutura para armazenar localização do usuário
type Localizacao struct {
	Cidade	string	`json:"cidade"`
	Estado	string	`json:"estado"`
}
// variável global para armazenar array usuário, simulando um banco de dados
var usuarios = []Usuario{}
// variável global para gerar ID's de forma automática
var proximoID = 1

// função GET/usuario - listar todos os usuários
func listarUsuarios(w http.ResponseWriter, r *http.Request) {
	// definindo que a resposta será um arquivo JSON
	w.Header().Set("Content-Type", "application/json")

	// desempacotar o array 'usuarios' e converter para JSON e retornar usuários
	// recebe JSON e converte o arquivo para struct
	json.NewEncoder(w).Encode(usuarios)
}

// função GET/usuario/{id} - buscar usuário pelo id
func buscarUsuario(w http.ResponseWriter, r *http.Request) {
	// obter parâmetro da URL - (exemplo: /usuarios/1)
	params := mux.Vars(r)

	// converter o parâmetro 'ID' de string p/ inteiro
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		// se for um número 'id' inválido, retorna erro 400
		http.Error(w, "ID inválido!", http.StatusBadRequest)
		return
	}

	// percorrer array usuarios e validar se o 'id' é localizado
	for _, user := range usuarios {
		// validar se o 'id' é localizado
		if user.ID == id {
			// definindo que a resposta será um arquivo JSON
			w.Header().Set("Content-Type", "application/json")
			// retornar o usuario encontrado
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	// caso não localize o usuário, retornar status
	http.Error(w, "Usuário/ID não localizado!", http.StatusNotFound)
}

// função POST/usuario - criar novo usuário
func criarUsuario(w http.ResponseWriter, r *http.Request) {
	// após executar a função fechar o Body
	defer r.Body.Close()
	// criar variável local para criar um novo usuário
	var newUser Usuario
	
	// converter JSON recebido no Body para struct Usuario
	// envia arquivo em JSON
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		// se for um JSON inválido, retorna erro 400
		http.Error(w, "Arquivo JSON inválido!", http.StatusBadRequest)
		return
	}

	// gerar próximo ID de forma automática
	newUser.ID = proximoID
	// incrementar 'proximoID' para o novo usuário
	proximoID++
	// adicionar o novo usuário ao array 'Usuario'
	usuarios = append(usuarios, newUser)

	// definindo que a resposta será um arquivo JSON
	w.Header().Set("Content-Type", "application/json")

	// retornar novo usuário criado
	json.NewEncoder(w).Encode(newUser)
}

// função DELETE/usuario{id} - buscar usuário pelo id e realizar 'delete'
func deletarUsuario(w http.ResponseWriter, r * http.Request) {
	// obter parâmetro da URL - (exemplo: /usuarios/1)
	params := mux.Vars(r)

	// converter o parâmetro 'ID' de string p/ inteiro
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		// // se for um número 'id' inválido, retorna erro 400
		http.Error(w, "ID inválido!", http.StatusBadRequest)
		return
	}
	
	// percorrer array usuarios, com indice 'index', e validar se o 'id' é localizado
	for index, user := range usuarios {
		// verificar se o usuário foi encontrado
		if user.ID == id {
			// remover usuario/id do array usuarios
			usuarios = append(usuarios[:index], usuarios[index+1:]...)

			// retornar status 204 - conteúdo removido
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	// caso não localize o usuário, retornar status
	http.Error(w, "Usuário não localizado!", http.StatusNotFound)
}

func main() {
	// criar um novo roteador
	router := mux.NewRouter()

	// definir a rota GET para listar todos os usuários
	router.HandleFunc("/usuarios", listarUsuarios).Methods("GET")

	// definir a rota GET para buscar usuário pelo {id}
	router.HandleFunc("/usuarios/{id}", buscarUsuario).Methods("GET")

	// definir a rota POST para criar novo usuário
	router.HandleFunc("/usuarios", criarUsuario).Methods("POST")

	// definir a rota DELETE para deletar usuário pelo {id}
	router.HandleFunc("/usuarios/{id}", deletarUsuario).Methods("DELETE")

	// adicionar usuário
	usuarios = append(usuarios, Usuario{ID: 1, Nome: "Gerson", Idade: 44, Localizacao: Localizacao{Cidade: "São Paulo", Estado: "SP"}})

	// exibir mensagem no console que a API foi iniciada e está rodando na porta :8080
	log.Println("🚀 API rodando em http://localhost:8080")

	// iniciar o servidor na porta :8080
	log.Fatal(http.ListenAndServe(":8080", router))
}