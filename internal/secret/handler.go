package secret

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func CreateSecretHandler(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var value string
		var ttl int64

		if r.Header.Get("Content-Type") == "application/json" {
			var body struct {
				Value string `json:"value"`
				TTL   int64  `json:"ttl"`
			}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.Value == "" {
				http.Error(w, "Invalid JSON input", http.StatusBadRequest)
				return
			}
			value = body.Value
			ttl = body.TTL
		} else {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Invalid form submission", http.StatusBadRequest)
				return
			}
			value = r.FormValue("value")
			fmt.Sscanf(r.FormValue("ttl"), "%d", &ttl)
		}

		enc, err := Encrypt(value)
		if err != nil {
			http.Error(w, "Failed to encrypt secret", http.StatusInternalServerError)
			return
		}

		id := randomID(12)
		secret := Secret{
			ID:        id,
			Encrypted: enc,
			CreatedAt: time.Now().Unix(),
			TTL:       ttl,
			Viewed:    false,
		}

		store.Mu.Lock()
		store.Secrets[id] = secret
		store.Mu.Unlock()

		if r.Header.Get("Content-Type") == "application/json" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"url": fmt.Sprintf("/secret/%s", id),
			})
		} else {
			http.Redirect(w, r, fmt.Sprintf("/secret/%s", id), http.StatusSeeOther)
		}
	}
}

func GetSecretHandler(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/secret/"):]
		
		store.Mu.Lock()
		secret, ok := store.Secrets[id]
		if !ok || secret.Viewed || (time.Now().Unix()-secret.CreatedAt > secret.TTL) {
			store.Mu.Unlock()
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Secret not found or expired",
			})
			return
		}
		secret.Viewed = true
		store.Secrets[id] = secret
		store.Mu.Unlock()

		decrypted, err := Decrypt(secret.Encrypted)
		if err != nil {
			http.Error(w, "Failed to decrypt secret", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"secret": decrypted,
		})
	}
}

