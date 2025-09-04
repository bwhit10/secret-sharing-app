package secret

import (
	"math/rand"
	"sync"
	"time"
)

// Store holds all the secrets in memory
type Store struct {
	Secrets map[string]Secret
	Mu      sync.Mutex
}

// NewStore initializes a new secret store
func NewStore() *Store {
	return &Store{
		Secrets: make(map[string]Secret),
	}
}

// CleanupExpiredSecrets runs in the background to remove expired secrets
func (s *Store) CleanupExpiredSecrets() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now().Unix()
		s.Mu.Lock()
		for id, secret := range s.Secrets {
			if secret.Viewed || (now-secret.CreatedAt > secret.TTL) {
				delete(s.Secrets, id)
			}
		}
		s.Mu.Unlock()
	}
}

// randomID generates a random alphanumeric string of length n
func randomID(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

