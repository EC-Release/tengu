/*
 * Copyright (c) 2016 General Electric Company. All rights reserved.
 *
 * The copyright to the computer software herein is the property of
 * General Electric Company. The software may be used and/or copied only
 * with the written permission of General Electric Company or in accordance
 * with the terms and conditions stipulated in the agreement/contract
 * under which the software has been supplied.
 *
 * author: apolo.yasuda@ge.com
 */

package main

import (
	"strings"
	"net/http"
	api "github.com/wzlib/wzapi"
	config "github.com/wzlib/wzconf"
	util "github.com/wzlib/wzutil"
	model "github.com/wzlib/wzschema"
	"encoding/json"
	"strconv"
	"net/url"
)

var (
	//tmplt map[string]interface{}
)
const (
	
)

// InitOAuth2 instantiate a oauth object and start its basic ops.
func InitOAuth2(ac *util.ConfigSet) {

	t:=ac.AgtConf
	oagt:=&OAuth2{}
	
	oa2:=api.NewOAuth2Mode(oagt, ac.AuthGrpConf, config.GetRev())
	oa2.Start(*t["_ap"].(*string),*t[util.AGENT_TLS_ENABLED].(*bool))

	//leave an internal reference
	oagt.oauthMode = oa2
}

// OAuth2 contains the mutex locker to protect the access to oauth instance
type OAuth2 struct{
	oauthMode *api.OAuth2Mode
}

//GetUser give the user detail as part of callback thats embedded in the x509 certificate
func (o *OAuth2) GetUser(w http.ResponseWriter, r *http.Request, cd *model.UserDetail){

	util.GetCLog().InfoLog("fetched user detail:",cd.FullName)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	_cfg, _ := json.Marshal(cd)
	_, err:=w.Write(_cfg)
	if err!=nil{
		util.ErrResponse(w, 500, err, err.Error())
		return
	}
	
	return
}

// CheckToken the function implemented as a callback once a token passed
//the validation https://tools.ietf.org/html/rfc6749
func (o *OAuth2) CheckToken(w http.ResponseWriter, r *http.Request, ts *model.TokenStatus){

	util.GetCLog().InfoLog("token validated for client id:",ts.DevId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	_cfg, _ := json.Marshal(ts)
	_, err:=w.Write(_cfg)
	if err!=nil{
		util.ErrResponse(w, 500, err, err.Error())
		return
	}
	return
}

// CreateToken implement the OAuth2 callback interface and will be invoke if, and only if the requestor passed the authentication
func (o *OAuth2) CreateToken(w http.ResponseWriter, r *http.Request, tkn *model.Token){
	
	util.GetCLog().InfoLog("token fetched.")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_cfg, _ := json.Marshal(tkn)
	_, err:=w.Write(_cfg)
	if err!=nil{
		util.ErrResponse(w, 500, err, err.Error())
		return
	}
	return
}

// CreateAuthCode implement the OAuth2 callback interface and will be invoke if, and only if the requestor passed the authentication.
func (o *OAuth2) CreateAuthCode(w http.ResponseWriter, r *http.Request, auth_code, callbackURL string){

	util.GetCLog().DbgLog("callbackURL",callbackURL)
	u, _ := url.ParseRequestURI(callbackURL)
	
	q := u.Query()

	q.Set("code", auth_code)

	u.RawQuery = q.Encode()
	urlStr := u.String()

	util.GetCLog().DbgLog(urlStr)
	//client := &http.Client{}
	r2, _ := http.NewRequest("GET", urlStr, nil) // URL-encoded payload
	r2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r2.Header.Add("Content-Length", strconv.Itoa(len(q.Encode())))
	r2.Header.Add("Cache-Control", "must-revalidate")
	r2.Header.Add("Cache-Control", "no-cache")

	http.Redirect(w, r2, urlStr, 302)	
}

// Authorization implement the OAuth2 callback interface and will be invoke if, and only if the requestor passed the authentication. This process is to give the user (@@uid) opportunity to decide, whether not to allow the client app (@cid) to access the information based on the scope.
func (o *OAuth2) Authorization(w http.ResponseWriter, r *http.Request, cid, uid, callbackURL string, scope []string) {
	util.GetCLog().InfoLog("authorised.")
	util.GetCLog().DbgLog("authorised info:cid,uid,scope,callbackURL",cid,uid,strings.Join(scope,":"),callbackURL)
	r.Header.Add("Cache-Control", "must-revalidate")
	r.Header.Add("Cache-Control", "no-cache")
	http.Redirect(w, r, callbackURL, 302)
}
