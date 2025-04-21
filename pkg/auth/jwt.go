package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/hegervalesin/helplinego/internal/models"
)

// JWTClaim representa os dados armazenados no token JWT
type JWTClaim struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Name   string `json:"name,omitempty"` // Adicionado campo nome
	jwt.RegisteredClaims
}

// Erros comuns de autenticação
var (
	ErrInvalidToken  = errors.New("token inválido")
	ErrExpiredToken  = errors.New("token expirado")
	ErrParsingClaims = errors.New("erro ao processar claims do token")
)

// GenerateJWT cria um novo token JWT para o usuário
func GenerateJWT(user models.User) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	if len(secretKey) == 0 {
		return "", errors.New("chave secreta JWT não configurada")
	}

	// Tempo de expiração configurável
	expirationTime := 24 * time.Hour
	if expStr := os.Getenv("JWT_EXPIRATION_HOURS"); expStr != "" {
		if expHours, err := time.ParseDuration(expStr + "h"); err == nil {
			expirationTime = expHours
		}
	}

	claims := JWTClaim{
		UserID: user.ID.String(),
		Email:  user.Email,
		//Role:   user.Role,
		Name: user.Name, // Assumindo que o modelo User tem um campo Name
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "helplinego-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// ValidateToken verifica e extrai as informações de um token JWT
func ValidateToken(signedToken string) (*JWTClaim, error) {
	if signedToken == "" {
		return nil, ErrInvalidToken
	}

	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	if len(secretKey) == 0 {
		return nil, errors.New("chave secreta JWT não configurada")
	}

	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			// Verificação do método de assinatura
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("método de assinatura inválido")
			}
			return secretKey, nil
		},
	)

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok || !token.Valid {
		return nil, ErrParsingClaims
	}

	return claims, nil
}

// RefreshToken gera um novo token a partir de um token válido existente
func RefreshToken(signedToken string) (string, error) {
	claims, err := ValidateToken(signedToken)
	if err != nil {
		return "", err
	}

	// Cria um usuário temporário com os dados do token para gerar um novo
	user := models.User{
		ID:    uuid.MustParse(claims.UserID),
		Email: claims.Email,
		//Role:  claims.Role,
		Name: claims.Name,
	}

	return GenerateJWT(user)
}
