package util

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"time"
)

type Util struct {
	Valid *validator.Validate
}

func NewUtils() *Util {
	return &Util{Valid: validator.New()}
}

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

func (u *Util) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := u.Valid.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}
	return validationErrors
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

func (u *Util) Response(Status int, Data interface{}) ResponseData {
	return ResponseData{
		Data:   Data,
		Status: Status,
	}
}

func (u *Util) RandNumber(start int, end int) int {
	code := rand.Intn(end-start) + start
	return code
}

func (u *Util) GetUserID(c *fiber.Ctx) uint {
	return c.Locals("user").(*jwt.Token).Claims.(*JwtCustomClaims).UserID
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
