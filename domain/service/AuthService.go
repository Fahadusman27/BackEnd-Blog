package service

import (
	. "BLOG/domain/model"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Auth interface {
	Register(Email, Username, Password string) (*Register, error)
	Login(Email, Password string) (string, *Login, error)
	Profile(ID int) (*Profile, error)
	ParseToken(tokenString string) (*jwt.RegisteredClaims, error)
}

type AuthService struct {
	authRepo AuthRepo
	secret   string
	expiry   time.Duration
}

type AuthRepo interface {
	CreateUsers(Users *Register) error
	FindByEmail(Email string) (*Login, error)
	GetProfile(ID int) (*Profile, error)
}

func (a *AuthService) GenerateToken(id int, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": id,
        "role":    role,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.secret))
}

func (a *AuthService) LoginAllRoles(email, password string) (string, string, error) {
    user, err := a.authRepo.FindByEmail(email)
    if err != nil {
        return "", "", errors.New("user tidak ditemukan")
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return "", "", errors.New("password salah")
    }

    token, err := a.GenerateToken(user.ID, user.RoleName)
    if err != nil {
        return "", "", err
    }

    return token, user.RoleName, nil
}

func (a *AuthService) LoginAdminOnly(email, password string) (string, error) {
    user, err := a.authRepo.FindByEmail(email)
    if err != nil {
        return "", errors.New("akun admin tidak ditemukan")
    }

    if user.RoleName != "admin" {
        return "", errors.New("akses ditolak: anda bukan admin")
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return "", errors.New("password salah")
    }

    token, err := a.GenerateToken(user.ID, user.RoleName)
    if err != nil {
        return "", err
    }

    return token, nil
}

func (a *AuthService) Register(Email, Password, Username string) (*Register, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password),bcrypt.DefaultCost,)
	if err != nil {
		return nil, err
	}

	user := &Register{
		Username: Username,
		Email:    Email,
		Password: string(hashedPassword),
	}

	if err := a.authRepo.CreateUsers(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (a *AuthService) Profile(UserID int) (*Profile, error) {
	profile, err :=a.authRepo.GetProfile(UserID)
	if err != nil{
		return nil,err
	}

	return profile,err
}





func (s *AuthService) ParseToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token claims")
}
