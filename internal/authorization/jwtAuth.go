// Package authz provides authentication and authorization functionality, including JWT token handling.
package authz

import (
	"context"
	"crypto/sha256"
	"errors"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/wurt83ow/gophkeeper-server/internal/config"
	"github.com/wurt83ow/gophkeeper-server/internal/models"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/crypto/bcrypt"
)

// Storage is an interface representing methods for inserting user data.
type Storage interface {
}

// CustomClaims represents custom claims for JWT token.
type CustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// Log is an interface representing a logger with Info method.
type Log interface {
	Info(string, ...zapcore.Field)
}

// JWTAuthz provides JWT token creation, decoding, and middleware functionality for authentication and authorization.
type JWTAuthz struct {
	jwtSigningKey    []byte
	log              Log
	jwtSigningMethod *jwt.SigningMethodHMAC
	defaultCookie    http.Cookie
}

// NewJWTAuthz creates a new JWTAuthz instance with the provided signing key and logger.
func NewJWTAuthz(signingKey string, log Log) *JWTAuthz {
	return &JWTAuthz{
		jwtSigningKey:    []byte(config.GetAsString("JWT_SIGNING_KEY", signingKey)),
		log:              log,
		jwtSigningMethod: jwt.SigningMethodHS256,

		defaultCookie: http.Cookie{
			HttpOnly: true,
		},
	}
}

func (j *JWTAuthz) JWTAuthzMiddleware(storage Storage, log Log) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var userID string
			var err error

			jwtToken := r.Header.Get("Authorization")

			if jwtToken != "" {
				userID, err = j.DecodeJWTToUser(jwtToken)

				if err != nil {
					userID = ""
					log.Info("Error occurred decoding JWT token", zap.Error(err))
				}
			}

			// If userID is still empty, return an authorization error
			if userID == "" {

				http.Error(w, "Authorization error", http.StatusUnauthorized)
				return
			}

			var keyUserID models.Key = "userID"
			ctx := r.Context()
			ctx = context.WithValue(ctx, keyUserID, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}

// CreateJWTTokenForUser creates a JWT token for the specified user ID.
func (j *JWTAuthz) CreateJWTTokenForUser(userid string) string {
	claims := CustomClaims{
		userid,
		jwt.StandardClaims{},
	}

	// Encode to token string
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.jwtSigningKey)
	if err != nil {
		log.Println("Error occurred generating JWT", err)
		return ""
	}

	return tokenString
}

// DecodeJWTToUser decodes a JWT token to retrieve the user ID.
func (j *JWTAuthz) DecodeJWTToUser(token string) (string, error) {
	// Decode
	decodeToken, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		if !(j.jwtSigningMethod == token.Method) {
			// Check our method hasn't changed since issuance
			return nil, errors.New("signing method mismatch")
		}
		return j.jwtSigningKey, nil
	})

	// There's two parts. We might decode it successfully but it might
	// be the case we aren't Valid so you must check both
	if decClaims, ok := decodeToken.Claims.(*CustomClaims); ok && decodeToken.Valid {
		return decClaims.Email, nil
	}

	return "", err
}

// GetHash computes the SHA-256 hash of the concatenation of email and password.
func (j *JWTAuthz) GetHash(email string, password string) []byte {
	src := []byte(email + password)

	// create a new hash.Hash that calculates the SHA-256 checksum
	h := sha256.New()
	// transfer bytes for hashing
	h.Write(src)
	// calculate the hash

	return h.Sum(nil)
}

// AuthCookie creates an http.Cookie with the specified name and token value.
func (j *JWTAuthz) AuthCookie(name string, token string) *http.Cookie {
	d := j.defaultCookie
	d.Name = name
	d.Value = token
	d.Path = "/"

	return &d
}

// IsBcryptHash проверяет, является ли данная строка хешем bcrypt.
func (j *JWTAuthz) IsBcryptHash(s string) bool {
	_, err := bcrypt.Cost([]byte(s))
	return err == nil
}

func (j *JWTAuthz) CompareHashAndPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
