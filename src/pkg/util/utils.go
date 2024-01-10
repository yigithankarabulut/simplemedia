package util

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
)

type Util struct {
	Valid *validator.Validate
}

func NewUtils() *Util {
	return &Util{Valid: validator.New(validator.WithRequiredStructEnabled())}
}

func (u *Util) Validate(c *fiber.Ctx, data interface{}) error {
	err := c.BodyParser(data)
	if err != nil {
		err = c.QueryParser(data)
		if err != nil {
			return err
		}
	}
	err = u.Valid.Struct(data)
	if err != nil {
		return err
	}
	return nil
}

func (u *Util) ValidateImageSize(file *multipart.FileHeader) error {
	if file.Size > 5000000 { // 5MB max
		return errors.New("image size is too big")
	}
	return nil
}

func (u *Util) SavePicture(file *multipart.FileHeader, username, path string) (string, error) {
	if err := u.ValidateImageSize(file); err != nil {
		return "", err
	}
	split := strings.Split(file.Filename, ".")
	extention := split[len(split)-1]
	if extention != "jpg" && extention != "jpeg" && extention != "png" {
		return "", fmt.Errorf("invalid file type: %s", extention)
	}
	dirPath := fmt.Sprintf("./uploads/%s/%s", username, path)
	if err := os.MkdirAll(dirPath, 0777); err != nil {
		return "", err
	}
	filepath := fmt.Sprintf("%s/%s.%s", dirPath, time.Now().Format("2006-01-02-15-04-05"), extention)
	if _, err := os.Create(filepath); err != nil {
		return "", err
	}
	return filepath, nil
}

func (u *Util) Response(Status int, Data interface{}) ResponseData {
	return ResponseData{
		Data:   Data,
		Status: Status,
	}
}

func (u *Util) BasicError(d interface{}, status int) BasicErrorResponse {
	var errorResponse BasicErrorResponse
	switch d.(type) {
	case error:
		errorResponse.Error = d.(error).Error()
	case string:
		errorResponse.Error = d.(string)
	case nil:
		errorResponse.Error = "unknown error."
	default:
		errorResponse.Error = "unknown error."
	}
	errorResponse.Status = status
	return errorResponse
}

func (u *Util) GetUserID(c *fiber.Ctx) uint {
	user := c.Locals("user").(*jwt.Token).Claims.(*JwtCustomClaims)
	return user.UserID
}

func (u *Util) PasswordControl(hash, pass string) bool {
	passwordControl := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if passwordControl != nil {
		return false
	}
	return true
}

func (u *Util) GeneratePassword(password string) (string, error) {
	hasPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hasPassword), err
}

func (u *Util) JwtMiddleware(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token).Claims.(*JwtCustomClaims)
	if user.UserID == 0 {
		return c.JSON(u.BasicError("User not found", 404))
	}
	return c.Next()
}

func (u *Util) PaginatorPage(c *fiber.Ctx, key ...string) int {
	var queryKey string
	if len(key) > 0 {
		queryKey = key[0]
	} else {
		queryKey = "page"
	}

	if c == nil {
		return 0
	}

	p, err := strconv.Atoi(c.FormValue(queryKey))
	if err != nil {
		p = 0
	}

	return p
}

func (u *Util) ParseTime(t string) (time.Time, error) {

	tm, err := time.Parse("2006-01-02 15:04:05", t)
	return tm, err
}

func (u *Util) ParseDateTime(t string) (time.Time, error) {

	tm, err := time.Parse("15:04:05", t)
	return tm, err
}

func (u *Util) FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
