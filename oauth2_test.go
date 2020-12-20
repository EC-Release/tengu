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
	"net/http"
	"testing"

	api "github.com/wzlib/wzapi"
	model "github.com/wzlib/wzschema"
	util "github.com/wzlib/wzutil"
)

func TestInitOAuth2(t *testing.T) {
	type args struct {
		ac *util.ConfigSet
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitOAuth2(tt.args.ac)
		})
	}
}

func TestOAuth2_GetUser(t *testing.T) {
	type fields struct {
		oauthMode *api.OAuth2Mode
	}
	type args struct {
		w  http.ResponseWriter
		r  *http.Request
		cd *model.UserDetail
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OAuth2{
				oauthMode: tt.fields.oauthMode,
			}
			o.GetUser(tt.args.w, tt.args.r, tt.args.cd)
		})
	}
}

func TestOAuth2_CheckToken(t *testing.T) {
	type fields struct {
		oauthMode *api.OAuth2Mode
	}
	type args struct {
		w  http.ResponseWriter
		r  *http.Request
		ts *model.TokenStatus
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OAuth2{
				oauthMode: tt.fields.oauthMode,
			}
			o.CheckToken(tt.args.w, tt.args.r, tt.args.ts)
		})
	}
}

func TestOAuth2_CreateToken(t *testing.T) {
	type fields struct {
		oauthMode *api.OAuth2Mode
	}
	type args struct {
		w   http.ResponseWriter
		r   *http.Request
		tkn *model.Token
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OAuth2{
				oauthMode: tt.fields.oauthMode,
			}
			o.CreateToken(tt.args.w, tt.args.r, tt.args.tkn)
		})
	}
}

func TestOAuth2_CreateAuthCode(t *testing.T) {
	type fields struct {
		oauthMode *api.OAuth2Mode
	}
	type args struct {
		w           http.ResponseWriter
		r           *http.Request
		auth_code   string
		callbackURL string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OAuth2{
				oauthMode: tt.fields.oauthMode,
			}
			o.CreateAuthCode(tt.args.w, tt.args.r, tt.args.auth_code, tt.args.callbackURL)
		})
	}
}

func TestOAuth2_Authorization(t *testing.T) {
	type fields struct {
		oauthMode *api.OAuth2Mode
	}
	type args struct {
		w           http.ResponseWriter
		r           *http.Request
		cid         string
		uid         string
		callbackURL string
		scope       []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OAuth2{
				oauthMode: tt.fields.oauthMode,
			}
			o.Authorization(tt.args.w, tt.args.r, tt.args.cid, tt.args.uid, tt.args.callbackURL, tt.args.scope)
		})
	}
}
