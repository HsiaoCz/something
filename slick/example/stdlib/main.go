package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("POST /user", HandleUser)

	log.Fatal(http.ListenAndServe(":9001", router))
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func HandleUser(w http.ResponseWriter, r *http.Request) {
	user := User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]any{
			"message": "Bad Request please check the input",
			"status":  http.StatusBadRequest,
		})
	}
	WriteJSON(w, http.StatusOK, &user)
}

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}
