package hash_test

import (
	"gin/pkg/hash"
	"testing"
)

func TestHashService(t *testing.T) {
	service := hash.NewHashService()

	t.Run("Hash and Check password", func(t *testing.T) {
		password := "pass"

		hashed, err := service.HashPassword(password)
		if err != nil {
			t.Fatalf("Failed to hash password: %v", err)
		}

		if !service.CheckPasswordHash(password, hashed) {
			t.Error("Expected password to match the hash, but it didn't")
		}
	})
}
