package secret

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestExpiredSecret(t *testing.T) {
	store := NewStore()
	enc, _ := Encrypt("test")
	id := "expired123"
	store.Secrets[id] = Secret{
		ID: id, Encrypted: enc, CreatedAt: time.Now().Unix() - 500, TTL: 60, Viewed: false,
	}

	req := httptest.NewRequest("GET", "/secret/"+id, nil)
	w := httptest.NewRecorder()
	handler := GetSecretHandler(store)
	handler(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected 404, got %d", w.Code)
	}
}

func TestCreateAndRetrieveSecret(t *testing.T) {
	store := NewStore()

	// Create secret
	secretVal := "this is a test"
	enc, _ := Encrypt(secretVal)
	id := "onetimetest"
	store.Secrets[id] = Secret{
		ID:        id,
		Encrypted: enc,
		CreatedAt: time.Now().Unix(),
		TTL:       60,
		Viewed:    false,
	}

	// Retrieve secret
	req := httptest.NewRequest("GET", "/secret/"+id, nil)
	w := httptest.NewRecorder()
	handler := GetSecretHandler(store)
	handler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", w.Code)
	}

	// Try to retrieve again (should be one-time use)
	w2 := httptest.NewRecorder()
	handler(w2, req)
	if w2.Code != http.StatusNotFound {
		t.Errorf("Expected 404 on second attempt, got %d", w2.Code)
	}
}
