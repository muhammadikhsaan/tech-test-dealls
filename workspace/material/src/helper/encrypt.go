package helper

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"dealls.test/material/src/core"
	"dealls.test/material/src/static"
	"golang.org/x/crypto/bcrypt"
)

var (
	salt   = static.HASH_SALT
	cost   = static.HASH_COST
	format = "%s-zee-salt-%s"
)

func Hash(plain string) (*string, *core.Error) {
	text := fmt.Sprintf(format, plain, salt)

	crypt_cost, err := strconv.Atoi(cost)

	if err != nil {
		return nil, &core.Error{
			StatusCode: http.StatusInternalServerError,
			Origin:     err,
			Message:    "failed to convert hast cost",
		}
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(text), crypt_cost)

	if err != nil {
		return nil, &core.Error{
			StatusCode: http.StatusInternalServerError,
			Origin:     err,
			Message:    "failed to hash password",
		}
	}

	return ToPointer(string(bytes)), nil
}

func Compare(hash, plain string) (bool, *core.Error) {
	text := fmt.Sprintf(format, plain, salt)

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(text))

	if err != nil {
		return false, &core.Error{
			StatusCode: http.StatusBadRequest,
			Origin:     err,
			Message:    "password don't match",
		}
	}

	return true, nil
}

func HashSha256(plain string, b ...string) string {
	alg := sha256.New()

	text := fmt.Sprintf(format, plain, salt)

	alg.Write([]byte(text))

	sumText := strings.Join(b, "-")

	encrypt := alg.Sum([]byte(sumText))

	return fmt.Sprintf("%x", encrypt)
}

func HashMD5(plain string, b ...string) string {
	alg := md5.New()

	text := fmt.Sprintf(format, plain, salt)

	alg.Write([]byte(text))

	sumText := strings.Join(b, "-")

	encrypt := alg.Sum([]byte(sumText))

	return fmt.Sprintf("%x", encrypt)
}
