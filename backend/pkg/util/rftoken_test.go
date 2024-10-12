package util

import (
	"testing"
	"time"
)

func TestRefresh(t *testing.T) {
	userId := uint(123)
	token, err := GenerateRefreshToken(userId)
	if err != nil {
		t.Fatalf("GenerateRefreshToken failed: %v", err)
	}

	claims, err := ParseRefreshToken(token)
	if err != nil {
		t.Fatalf("ParseRefreshToken failed: %v", err)
	}

	if claims.UserId != userId {
		t.Errorf("Expected userId %d, got %d", userId, claims.UserId)
	}

	if claims.Issuer != "gin-discord-clone" {
		t.Errorf("Expected issuer 'gin-discord-clone', got '%s'", claims.Issuer)
	}

	now := time.Now().Unix()
	if claims.ExpiresAt < now {
		t.Errorf("Token expired")
	}
}
