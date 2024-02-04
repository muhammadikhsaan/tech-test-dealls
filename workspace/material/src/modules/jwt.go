package modules

import (
	"encoding/json"
	"net/http"
	"time"

	"dealls.test/material/src/contract"
	"dealls.test/material/src/core"
	"dealls.test/material/src/helper"
	"dealls.test/material/src/secret"
	"dealls.test/material/src/static"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

var (
	name = "user-claims"
)

func GenerateJWT(subject string, user *contract.UserFormToken) (*string, *core.Error) {
	key, err := jwk.ParseKey(secret.JWT_PRIVATEKEY, jwk.WithPEM(true))

	if err != nil {
		return nil, &core.Error{
			StatusCode: http.StatusInternalServerError,
			Origin:     err,
			Message:    "failed to get jwt key",
		}
	}

	b, err := jwt.NewBuilder().
		Subject(subject).
		Issuer(static.JWT_ISSUER).
		IssuedAt(time.Now()).
		Expiration(time.Now().AddDate(0, static.JWT_EXPIRED_TOKEN, 0)).
		Claim(name, user).Build()

	if err != nil {
		return nil, &core.Error{
			StatusCode: http.StatusInternalServerError,
			Origin:     err,
			Message:    "failed to generate jwt token",
		}
	}

	signed, err := jwt.Sign(b, jwt.WithKey(jwa.RS512, key))

	if err != nil {
		return nil, &core.Error{
			StatusCode: http.StatusInternalServerError,
			Origin:     err,
			Message:    "failed to sign jwt token",
		}
	}

	return helper.ToPointer(string(signed)), nil
}

func VerifyJWT(subject, token string) (*contract.UserFormToken, *core.Error) {
	user := new(contract.UserFormToken)

	key, err := jwk.ParseKey(secret.JWT_PUBLICKEY, jwk.WithPEM(true))

	if err != nil {
		return nil, &core.Error{
			StatusCode: http.StatusInternalServerError,
			Origin:     err,
			Message:    "failed to get jwt key",
		}
	}

	t, err := jwt.Parse([]byte(token), jwt.WithKey(jwa.RS512, key))

	if err != nil {
		return nil, &core.Error{
			StatusCode: http.StatusUnauthorized,
			Origin:     err,
			Message:    "failed to parse jwt token",
		}
	}

	iss := t.Issuer()

	if iss != static.JWT_ISSUER {
		return nil, &core.Error{
			StatusCode: http.StatusUnauthorized,
			Message:    "Issuer is not valid",
		}
	}

	sub := t.Subject()

	if sub != subject {
		return nil, &core.Error{
			StatusCode: http.StatusUnauthorized,
			Message:    "subject is not valid",
		}
	}

	claims := t.PrivateClaims()

	js, err := json.Marshal(claims[name])

	if err != nil {
		return nil, &core.Error{
			StatusCode: http.StatusInternalServerError,
			Origin:     err,
			Message:    "failed to marshal user claims",
		}
	}

	err = json.Unmarshal(js, &user)

	if err != nil {
		return nil, &core.Error{
			StatusCode: http.StatusInternalServerError,
			Origin:     err,
			Message:    "failed to unmarshal user claims",
		}
	}

	return user, nil
}
