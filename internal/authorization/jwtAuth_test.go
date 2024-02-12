package authz

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
	"golang.org/x/crypto/bcrypt"
)

type MockLogger struct{}

func (m *MockLogger) Info(string, ...zapcore.Field) {}

func TestJWTAuthz_CreateJWTTokenForUser(t *testing.T) {
	jwtAuthz := NewJWTAuthz("secret", &MockLogger{})

	token := jwtAuthz.CreateJWTTokenForUser("user123")

	assert.NotEmpty(t, token)
}

func TestJWTAuthz_DecodeJWTToUser(t *testing.T) {
	jwtAuthz := NewJWTAuthz("secret", &MockLogger{})

	token := jwtAuthz.CreateJWTTokenForUser("user123")
	userID, err := jwtAuthz.DecodeJWTToUser(token)

	assert.NoError(t, err)
	assert.Equal(t, "user123", userID)
}

func TestJWTAuthz_Middleware(t *testing.T) {
	jwtAuthz := NewJWTAuthz("secret", &MockLogger{})
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", jwtAuthz.CreateJWTTokenForUser("user123"))
	rr := httptest.NewRecorder()

	middleware := jwtAuthz.JWTAuthzMiddleware(nil, &MockLogger{})(handler)
	middleware.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestJWTAuthz_Middleware_Unauthorized(t *testing.T) {
	jwtAuthz := NewJWTAuthz("secret", &MockLogger{})
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Handler should not have been called")
	})

	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	middleware := jwtAuthz.JWTAuthzMiddleware(nil, &MockLogger{})(handler)
	middleware.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}

func TestJWTAuthz_GetHash(t *testing.T) {
	jwtAuthz := NewJWTAuthz("secret", &MockLogger{})

	email := "test@example.com"
	password := "password123"
	expectedHash := jwtAuthz.GetHash(email, password)

	assert.NotNil(t, expectedHash)
}

func TestJWTAuthz_IsBcryptHash(t *testing.T) {
	jwtAuthz := NewJWTAuthz("secret", &MockLogger{})

	password := "password123"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	assert.NoError(t, err)

	assert.True(t, jwtAuthz.IsBcryptHash(string(hashedPassword)))
}

func TestJWTAuthz_AuthCookie(t *testing.T) {
	jwtAuthz := NewJWTAuthz("secret", &MockLogger{})

	cookieName := "test_cookie"
	tokenValue := "test_token"
	cookie := jwtAuthz.AuthCookie(cookieName, tokenValue)

	assert.Equal(t, cookieName, cookie.Name)
	assert.Equal(t, tokenValue, cookie.Value)
	assert.Equal(t, "/", cookie.Path)
	assert.True(t, cookie.HttpOnly)
}

func TestJWTAuthz_CompareHashAndPassword(t *testing.T) {
	jwtAuthz := NewJWTAuthz("secret", &MockLogger{})

	password := "password123"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	assert.NoError(t, err)

	assert.True(t, jwtAuthz.CompareHashAndPassword(string(hashedPassword), password))
}
