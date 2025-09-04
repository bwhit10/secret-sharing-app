package main

import "testing"

func TestEncryptionCycle(t *testing.T) {
    original := "hello-articulate"

    encrypted, err := Encrypt(original)
    if err != nil {
        t.Fatalf("Encryption failed: %v", err)
    }

    if encrypted == original {
        t.Fatalf("Encrypted text should not match original text")
    }

    decrypted, err := Decrypt(encrypted)
    if err != nil {
        t.Fatalf("Decryption failed: %v", err)
    }

    if decrypted != original {
        t.Fatalf("Expected %s but got %s", original, decrypted)
    }
}
