package helper

import (
	"encoding/csv"
	"mime/multipart"
	"net/http"
	"net/mail"
	"os"
	"strings"
	"unicode"

	"github.com/google/uuid"
	"github.com/gorilla/securecookie"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/api/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/category/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/user/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/cart/model"
	"golang.org/x/crypto/bcrypt"
)

// Verify checks the user e-mail structure
func VerifyEMail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// VerifyPassword checks user password that match or not 1 lower, 1 upper, 1 number, 1 special chars and length of password
func VerifyPassword(s string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) >= 7 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

// HashPassword helps to hashing password before save the db
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
// CheckPasswordHash helps to decode hashed password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
// DecodeCookie checks cookie and returns token
func DecodeCookie(req *http.Request, user *User) (*Token, error) {
	var hashKey = []byte(os.Getenv("COOKIE_SECRET"))
	var s = securecookie.New(hashKey, nil)
	var value Token
	if cookie, err := req.Cookie(user.Id); err == nil {
		if err = s.Decode(os.Getenv("TOKEN_NAME"), cookie.Value, &value); err != nil {
			return nil, err
		}

	}
	return &value, nil
}

// ReadCSV helps to read file and format to list
func ReadCSV(file *multipart.File) (CategoryList, error) {
	csvReader := csv.NewReader(*file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var categorieslist CategoryList

	for _, line := range records[1:] {
		categorieslist = append(categorieslist, Category{
			CategoryID:   uuid.New().String(),
			CategoryName: line[0],
			IconURL:      line[1],
		})
	}
	return categorieslist, nil
}

// CompareCategories helps to compare db data and upload file data
func CompareCategories(db, uploaded *CategoryList) CategoryList {
	var out CategoryList
	up := *uploaded
	d := *db

	for i := 0; i < len(up); i++ {
		res := contains(d, up[i])
		if !res {
			out = append(out, up[i])
		}
	}
	return out
}

// contains checks data is created before
func contains(clist CategoryList, c Category) bool {
	for _, v := range clist {
		if strings.ToLower(v.CategoryName) == strings.ToLower(c.CategoryName) {
			return true
		}
	}
	return false
}

// SetCookie creates cookie depends on token
func SetCookie(tkn *Token, user *User) *http.Cookie{
	var hashKey = []byte(os.Getenv("COOKIE_SECRET"))
	var s = securecookie.New(hashKey, nil)
	encoded, err := s.Encode(os.Getenv("TOKEN_NAME"), tkn)
	if err==nil{
		cookie := &http.Cookie{
			Name:     user.Id,
			Value:    encoded,
			Path:     "/",
			Domain:   "127.0.0.1",
			Secure:   false,
			HttpOnly: true,
		}
		return cookie
	}
	return nil
}



func UpdateValues(cart Cart, cartitems []CartsItem) *Cart{

	cart.CartLength = len(cartitems)
	var val = 0
	for i := 0; i < len(cartitems); i++ {
		val += int(cartitems[i].TotalPrice)
	}
	cart.CartTotalPrice = float64(val)

	return &cart
}