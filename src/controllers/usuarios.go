package controllers

import "net/http"

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário..."))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários..."))
}

func BuscandoUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuário específico..."))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando dados do usuário..."))
}

func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário..."))
}