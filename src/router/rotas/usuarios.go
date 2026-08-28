package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:          "/usuarios",
		Metodo:       http.MethodPost,
		Funcao:       controllers.CriarUsuario,
		Autenticacao: false,
	},
	{
		URI:          "/usuarios",
		Metodo:       http.MethodGet,
		Funcao:       controllers.BuscarUsuarios,
		Autenticacao: false,
	},
	{
		URI:          "/usuarios/{usuarioId}",
		Metodo:       http.MethodGet,
		Funcao:       controllers.BuscandoUsuario,
		Autenticacao: false,
	},
	{
		URI:          "/usuarios/{usuarioId}",
		Metodo:       http.MethodPut,
		Funcao:       controllers.AtualizarUsuario,
		Autenticacao: false,
	},
	{
		URI:          "/usuarios/{usuarioId}",
		Metodo:       http.MethodDelete,
		Funcao:       controllers.DeleteUsuario,
		Autenticacao: false,
	},
}
