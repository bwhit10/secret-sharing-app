package main

import (
    "fmt"
    "log"
    "net/http"
    "devex-challenge-bwhit10/internal/secret"
)


func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<title>Secret Sharing</title>
		</head>
		<body>
			<h2>🔐 Secret Sharing App</h2>
			<p>Use POST /secret to create a secret and GET /secret/{id} to retrieve it once.</p>
		</body>
		</html>
	`)
}


func main() {
	store := secret.NewStore()
	go store.CleanupExpiredSecrets()

	http.HandleFunc("/", index)
	http.HandleFunc("/secret", secret.CreateSecretHandler(store))
	http.HandleFunc("/secret/", secret.GetSecretHandler(store))

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}






