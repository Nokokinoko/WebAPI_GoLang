package handler

import (
	"encoding/json"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

const urlProduction string = "https://production"
const urlStaging string = "https://staging"

func IsDocker() bool {
	return os.Getenv("docker") != ""
}

type IController interface {
	fqgn() string
	identification() string

	Balance(echo.Context) error
	ListCollections(echo.Context) error
	CollectionDetail(echo.Context) error
	Clone(echo.Context) error
	Mint(echo.Context) error
	Burn(echo.Context) error
	BurnToMint(echo.Context) error
	Transfer(echo.Context) error
	ListTokens(echo.Context) error
	ListTokensByCollection(echo.Context) error
	TokenDetail(echo.Context) error
	ListUserTokens(echo.Context) error
	ListUserTokensByCollection(echo.Context) error
	UserTokenDetail(echo.Context) error
	ListUsersByToken(echo.Context) error

	Version(echo.Context) error
	Log(echo.Context) error
}

type Controller struct {
}

type JsonParameter struct {
	json map[string]string
}

func bindJson(context echo.Context) (JsonParameter, error) {
	var prm JsonParameter
	length, err := strconv.Atoi(context.Request().Header.Get("Content-Length"))
	if err != nil {
		return prm, err
	}

	body := make([]byte, length)
	length, err = context.Request().Body.Read(body)
	if err != nil && err != io.EOF {
		return prm, err
	}

	err = json.Unmarshal(body[:length], &prm.json)
	return prm, err // include error
}

func (prm JsonParameter) convString() string {
	json, _ := json.Marshal(prm.json)
	return string(json)
}

func (prm JsonParameter) isStaging() bool {
	_, exist := prm.json["stg"]
	return exist
}

func (prm JsonParameter) isDisconnect() bool {
	_, exist := prm.json["disconn"]
	return exist
}

func datetimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func datetimeYmd() string {
	return time.Now().Format("20060102")
}
