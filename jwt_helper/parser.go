package jwthelper

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/the-goat-jp/yabai-pkg/model"
)

const (
	refreshTime   time.Duration = 30 * time.Minute
	certsEndpoint string        = "/user/v1/certs"
)

type JwtParser struct {
	jwkCache *jwk.Cache
	endpoint string
}

func NewJWTParser(monoServiceHost string) (*JwtParser, error) {
	parser := JwtParser{
		jwkCache: jwk.NewCache(context.Background()),
		endpoint: monoServiceHost + certsEndpoint,
	}

	// register jwk cache
	parser.jwkCache.Register(monoServiceHost+certsEndpoint, jwk.WithMinRefreshInterval(refreshTime))

	// make sure certificate exist
	_, err := parser.jwkCache.Refresh(context.Background(), parser.endpoint)
	if err != nil {
		return nil, err
	}
	return &parser, nil
}

// Parse will decode jwt payload data, validate the timestamp, and verify the signature
func (jp *JwtParser) Parse(ctx context.Context, jwtInput string) (model.JWTData, error) {
	// get jwk set
	set, err := jp.jwkCache.Get(ctx, jp.endpoint)
	if err != nil {
		return model.JWTData{}, err
	}

	jwtToken, err := jwt.Parse([]byte(jwtInput), jwt.WithKeySet(set))
	if err != nil {
		return model.JWTData{}, err
	}

	dataByte, err := json.Marshal(jwtToken.PrivateClaims())
	if err != nil {
		return model.JWTData{}, err
	}

	var jwtData model.JWTData
	err = json.Unmarshal(dataByte, &jwtData)
	if err != nil {
		return model.JWTData{}, err
	}

	subID, err := strconv.ParseInt(jwtToken.Subject(), 10, 64)
	if err != nil {
		return model.JWTData{}, err
	}

	jwtData.ID = subID

	email, ok := jwtToken.Get("name")
	if ok {
		jwtData.Email = email.(string)
	}

	pic, ok := jwtToken.Get("url")
	if ok {
		jwtData.ProfileImage = pic.(string)
	}

	return jwtData, nil
}
