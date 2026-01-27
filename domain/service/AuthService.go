package service

import (
	"BLOG/domain/model"
	. "BLOG/domain/model"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)


func NewAuthService(repo AuthRepo, secret string) *AuthService {
	return &AuthService{
		authRepo: repo,
		secret:   secret,
		expiry:   time.Hour * 24,
	}
}

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
	FindByEmail(Email string) (model.Login, error)
	GetProfile(ID int) (model.Profile, error)
	UpadateProfile(ID int, profile *model.Profile) error
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

//untuk login masuk dashboard blog

func (a *AuthService) LoginAllRoles(email, password string) (string, string, error) {
	user, err := a.authRepo.FindByEmail(email)
	if err != nil {
		return "", "", errors.New("user tidak ditemukan")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	); err != nil {
		return "", "", errors.New("password salah")
	}

	// VALIDASI ROLE
	if user.RoleName != "admin" && user.RoleName != "user" {
		return "", "", errors.New("role tidak diizinkan")
	}

	token, err := a.GenerateToken(user.ID, user.RoleName)
	if err != nil {
		return "", "", err
	}

	return token, user.RoleName, nil
}

//untuk login masuk dashbord admin saja

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

func (a *AuthService) LoginService(c *fiber.Ctx) error {
	var login Login
	if err := c.BodyParser(&login); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input tidak valid"})
	}

	// Panggil logika bisnis LoginAllRoles
	token, role, err := a.LoginAllRoles(login.Email, login.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"token": token,
		"role":  role,
	})
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

func (a *AuthService) RegisterService(c *fiber.Ctx) error {
	var user Register
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Format JSON salah"})
	}

	res, err := a.Register(user.Email, user.Password, user.Username)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Registrasi Berhasil",
		"data":    res,
	})
}

func (a *AuthService) UpdateProfile(ID int, profile *model.Profile) error {
	return a.authRepo.UpadateProfile(ID, profile)
}

//update profle
func (a *AuthService) UpdateProfileService(c *fiber.Ctx) error {
	var profile Profile
	if err := c.BodyParser(&profile); err !=nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "Format tidak benar",
		})
	}

	ID := c.Locals("userID").(int)

	if err := a.authRepo.UpadateProfile(ID, &profile); err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : "Gagal Update Profile",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"pesan" : "Berhasil Update Profile",
	})
}

func (a *AuthService) GetProfile(ID int) (model.Profile, error) {
	return a.authRepo.GetProfile(ID)
}

//tampilkan profile
func (a *AuthService) GetProfileService(c *fiber.Ctx) error {
	ID := c.Locals("userID").(int)

	profile, err := a.authRepo.GetProfile(ID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : "Gagal Mengambil profile",
		})
	}

	return c.Status(fiber.StatusOK).JSON(profile)
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
