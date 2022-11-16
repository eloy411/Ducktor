package routes

import "net/http"

func IniGame(w http.ResponseWriter, r *http.Request) {

	/**CONFIG HEADERS*/
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	/**REQUEST DE FRASES PARA INICIAR A JUGAR */

	w.Write([]byte("Iniciando el juego"))
}

func DataGame(w http.ResponseWriter, r *http.Request) {
	/**CONFIG HEADERS*/
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	/**SALVAR LOS DATOS (TIEMPO,PUNTUACIÓN) */
	w.Write([]byte("Datos del juego registrados"))
}