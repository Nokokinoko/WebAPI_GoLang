package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func requestGet(context echo.Context, uri string, logError string) error {
	request, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return err
	}

	request.Header.Add("Content-Type", "application/json")

	client := new(http.Client)
	response, err := client.Do(request)
	if err != nil {
		context.Echo().Logger.Error(logError)
		return err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var message map[string]interface{}
	if err = json.Unmarshal(body, &message); err != nil {
		return err
	}

	return context.JSON(response.StatusCode, message)
}

func requestPost(context echo.Context, uri string, fields map[string]interface{}, logError string) error {
	encoded, err := json.Marshal(fields)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(encoded))
	if err != nil {
		return err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Auth-Key", "app_auth") // TODO: parameter store

	client := new(http.Client)
	response, err := client.Do(request)
	if err != nil {
		context.Echo().Logger.Error(logError)
		return err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var message map[string]interface{}
	if err = json.Unmarshal(body, &message); err != nil {
		return err
	}

	return context.JSON(response.StatusCode, message)
}

func (ctrl Controller) urlBase(prm JsonParameter) string {
	if prm.isStaging() || IsDocker() {
		return urlStaging
	}
	return urlProduction
}

func (ctrl Controller) fqgn() string {
	return ""
}

func (ctrl Controller) identification() string {
	return ""
}

// ------ Balance ------
func (ctrl Controller) Balance(context echo.Context) error {
	prm, err := bindJson(context)
	if err != nil {
		context.Echo().Logger.Error("Error Balance parameter uri:" + context.Path() + ", parameter:" + prm.convString())
		return err
	}

	uri := ctrl.urlBase(prm) + "/user/" + prm.json["address"] + "/balance"
	if prm.isDisconnect() || IsDocker() {
		return debugResponse(context, uri, nil, debugResponseBalance())
	}

	return requestGet(context, uri, "Error Balance request uri:"+context.Path()+", parameter:"+prm.convString())
}

// ------ List Collections ------
func (ctrl Controller) ListCollections(context echo.Context) error {
	prm, err := bindJson(context)
	if err != nil {
		context.Echo().Logger.Error("Error ListCollections parameter uri:" + context.Path() + ", parameter:" + prm.convString())
		return err
	}

	uri := ctrl.urlBase(prm) + "/collections/" + ctrl.fqgn() + "/"
	if prm.isDisconnect() || IsDocker() {
		return debugResponse(context, uri, nil, debugResponseListCollections())
	}

	return requestGet(context, uri, "Error ListCollections request uri:"+context.Path()+", parameter:"+prm.convString())
}

// ------ Collection Detail ------
func (ctrl Controller) CollectionDetail(context echo.Context) error {
	prm, err := bindJson(context)
	if err != nil {
		context.Echo().Logger.Error("Error CollectionDetail parameter uri:" + context.Path() + ", parameter:" + prm.convString())
		return err
	}

	uri := ctrl.urlBase(prm) + "/collections/" + ctrl.fqgn() + "/" + prm.json["qCn"] + "/"
	if prm.isDisconnect() || IsDocker() {
		return debugResponse(context, uri, nil, debugResponseCollectionDetail())
	}

	return requestGet(context, uri, "Error CollectionDetail request uri:"+context.Path()+", parameter:"+prm.convString())
}

// ------ Clone ------
func (ctrl Controller) Clone(context echo.Context) error {
	prm, err := bindJson(context)
	if err != nil {
		context.Echo().Logger.Error("Error Clone parameter uri:" + context.Path() + ", parameter:" + prm.convString())
		return err
	}

	uri := ctrl.urlBase(prm) + "/tokens/manage/" + prm.json["fqTn"] + "/clone"
	fields := map[string]interface{}{
		"to":                     prm.json["to"],
		"mutableProperties":      prm.json["mutableProperties"],
		"cloneMutableProperties": prm.json["cloneMutableProperties"],
		"displayInDiscovery":     prm.json["displayInDiscovery"],
	}
	if prm.isDisconnect() || IsDocker() {
		return debugResponse(context, uri, fields, debugResponseClone())
	}

	return requestPost(context, uri, fields, "Error Clone request uri:"+context.Path()+", parameter:"+prm.convString())
}

// ------ Mint ------
func (ctrl Controller) Mint(context echo.Context) error {
	prm, err := bindJson(context)
	if err != nil {
		context.Echo().Logger.Error("Error Mint parameter uri:" + context.Path() + ", parameter:" + prm.convString())
		return err
	}

	uri := ctrl.urlBase(prm) + "/tokens/manage/" + prm.json["fqTn"] + "/mint"
	fields := map[string]interface{}{
		"to":     prm.json["to"],
		"amount": prm.json["amount"],
	}
	if prm.isDisconnect() || IsDocker() {
		return debugResponse(context, uri, fields, debugResponseMint())
	}

	return requestPost(context, uri, fields, "Error Mint request uri:"+context.Path()+", parameter:"+prm.convString())
}

// ------ Burn ------
func (ctrl Controller) Burn(context echo.Context) error {
	prm, err := bindJson(context)
	if err != nil {
		context.Echo().Logger.Error("Error Burn parameter uri:" + context.Path() + ", parameter:" + prm.convString())
		return err
	}

	uri := ctrl.urlBase(prm) + "/tokens/manage/burn"
	fields := map[string]interface{}{
		"from":     prm.json["from"],
		"fqTn":     prm.json["fqTn"],
		"quantity": prm.json["quantity"],
	}
	if prm.isDisconnect() || IsDocker() {
		return debugResponse(context, uri, fields, debugResponseBurn())
	}

	return requestPost(context, uri, fields, "Error Burn request uri:"+context.Path()+", parameter:"+prm.convString())
}

// ------ Burn to Mint ------
func (ctrl Controller) BurnToMint(context echo.Context) error {
	prm, err := bindJson(context)
	if err != nil {
		context.Echo().Logger.Error("Error BurnToMint parameter uri:" + context.Path() + ", parameter:" + prm.convString())
		return err
	}

	uri := ctrl.urlBase(prm) + "/tokens/manage/burn-to-mint"
	fields := map[string]interface{}{
		"address":    prm.json["address"],
		"fqCn":       prm.json["fqCn"],
		"burnTokens": prm.json["burnTokens"],
		"mintTokens": prm.json["mintTokens"],
	}
	if prm.isDisconnect() || IsDocker() {
		return debugResponse(context, uri, fields, debugResponseBurnToMint())
	}

	return requestPost(context, uri, fields, "Error BurnToMint request uri:"+context.Path()+", parameter:"+prm.convString())
}

// ------ Transfer ------
func (ctrl Controller) Transfer(context echo.Context) error {
	prm, err := bindJson(context)
	if err != nil {
		context.Echo().Logger.Error("Error Transfer parameter uri:" + context.Path() + ", parameter:" + prm.convString())
		return err
	}

	uri := ctrl.urlBase(prm) + "/tokens/manage/transfer"
	fields := map[string]interface{}{
		"to":     prm.json["to"],
		"fqCn":   prm.json["fqCn"],
		"tokens": prm.json["tokens"],
	}
	if prm.isDisconnect() || IsDocker() {
		return debugResponse(context, uri, fields, debugResponseTransfer())
	}

	return requestPost(context, uri, fields, "Error Transfer request uri:"+context.Path()+", parameter:"+prm.convString())
}

// ------ List Tokens ------
func (ctrl Controller) ListTokens(context echo.Context) error {
	prm, err := bindJson(context)
	if err != nil {
		context.Echo().Logger.Error("Error ListTokens parameter uri:" + context.Path() + ", parameter:" + prm.convString())
		return err
	}

	uri := ctrl.urlBase(prm) + "/tokens/" + ctrl.fqgn() + "/"
	if prm.isDisconnect() || IsDocker() {
		return debugResponse(context, uri, nil, debugResponseListTokens())
	}

	return requestGet(context, uri, "Error ListTokens request uri:"+context.Path()+", parameter:"+prm.convString())
}

// ------ List Tokens by Collection ------
func (ctrl Controller) ListTokensByCollection(context echo.Context) error {
	prm, err := bindJson(context)
	if err != nil {
		context.Echo().Logger.Error("Error ListTokensByCollection parameter uri:" + context.Path() + ", parameter:" + prm.convString())
		return err
	}

	uri := ctrl.urlBase(prm) + "/tokens/" + ctrl.fqgn() + "/" + prm.json["qCn"] + "/"
	if prm.isDisconnect() || IsDocker() {
		return debugResponse(context, uri, nil, debugResponseListTokensByCollection())
	}

	return requestGet(context, uri, "Error ListTokensByCollection request uri:"+context.Path()+", parameter:"+prm.convString())
}

// ------ Token Detail ------
func (ctrl Controller) TokenDetail(context echo.Context) error {
	prm, err := bindJson(context)
	if err != nil {
		context.Echo().Logger.Error("Error TokenDetail parameter uri:" + context.Path() + ", parameter:" + prm.convString())
		return err
	}

	uri := ctrl.urlBase(prm) + "/tokens/" + ctrl.fqgn() + "/" + prm.json["qCn"] + "/" + prm.json["qTn"]
	if prm.isDisconnect() || IsDocker() {
		return debugResponse(context, uri, nil, debugResponseTokenDetail())
	}

	return requestGet(context, uri, "Error TokenDetail request uri:"+context.Path()+", parameter:"+prm.convString())
}

// ------ List User Tokens ------
func (ctrl Controller) ListUserTokens(context echo.Context) error {
	prm, err := bindJson(context)
	if err != nil {
		context.Echo().Logger.Error("Error ListUserTokens parameter uri:" + context.Path() + ", parameter:" + prm.convString())
		return err
	}

	uri := ctrl.urlBase(prm) + "/user/" + prm.json["address"] + "/tokens/" + ctrl.fqgn() + "/"
	if prm.isDisconnect() || IsDocker() {
		return debugResponse(context, uri, nil, debugResponseListUserTokens())
	}

	return requestGet(context, uri, "Error ListUserTokens request uri:"+context.Path()+", parameter:"+prm.convString())
}

// ------ List User Tokens by Collection ------
func (ctrl Controller) ListUserTokensByCollection(context echo.Context) error {
	prm, err := bindJson(context)
	if err != nil {
		context.Echo().Logger.Error("Error ListUserTokensByCollection parameter uri:" + context.Path() + ", parameter:" + prm.convString())
		return err
	}

	uri := ctrl.urlBase(prm) + "/user/" + prm.json["address"] + "/tokens/" + ctrl.fqgn() + "/" + prm.json["qCn"] + "/"
	if prm.isDisconnect() || IsDocker() {
		return debugResponse(context, uri, nil, debugResponseListUserTokensByCollection())
	}

	return requestGet(context, uri, "Error ListUserTokensByCollection request uri:"+context.Path()+", parameter:"+prm.convString())
}

// ------ User Token Detail ------
func (ctrl Controller) UserTokenDetail(context echo.Context) error {
	prm, err := bindJson(context)
	if err != nil {
		context.Echo().Logger.Error("Error UserTokenDetail parameter uri:" + context.Path() + ", parameter:" + prm.convString())
		return err
	}

	uri := ctrl.urlBase(prm) + "/user/" + prm.json["address"] + "/tokens/" + ctrl.fqgn() + "/" + prm.json["qCn"] + "/" + prm.json["qTn"]
	if prm.isDisconnect() || IsDocker() {
		return debugResponse(context, uri, nil, debugResponseUserTokenDetail())
	}

	return requestGet(context, uri, "Error UserTokenDetail request uri:"+context.Path()+", parameter:"+prm.convString())
}

// ------ List Users by Token ------
func (ctrl Controller) ListUsersByToken(context echo.Context) error {
	prm, err := bindJson(context)
	if err != nil {
		context.Echo().Logger.Error("Error ListUsersByToken parameter uri:" + context.Path() + ", parameter:" + prm.convString())
		return err
	}

	uri := ctrl.urlBase(prm) + "/users/" + ctrl.fqgn() + "/" + prm.json["qCn"] + "/" + prm.json["qTn"] + "/"
	if prm.isDisconnect() || IsDocker() {
		return debugResponse(context, uri, nil, debugResponseListUsersByToken())
	}

	return requestGet(context, uri, "Error ListUsersByToken request uri:"+context.Path()+", parameter:"+prm.convString())
}
