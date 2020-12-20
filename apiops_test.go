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

func TestInitAPI(t *testing.T) {
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
			InitAPI(tt.args.ac)
		})
	}
}

func TestAgentAPI_Ext(t *testing.T) {
	type fields struct {
		AgentOpsList *ConfLocker
		a            *api.APIMode
	}
	type args struct {
		w         http.ResponseWriter
		r         *http.Request
		tknStatus *model.TokenStatus
		db        *util.DB
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
			i := &AgentAPI{
				AgentOpsList: tt.fields.AgentOpsList,
				a:            tt.fields.a,
			}
			i.Ext(tt.args.w, tt.args.r, tt.args.tknStatus, tt.args.db)
		})
	}
}

func TestAgentAPI_SessionBegin(t *testing.T) {
	type fields struct {
		AgentOpsList *ConfLocker
		a            *api.APIMode
	}
	type args struct {
		r          *http.Request
		userDetail *model.UserDetail
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &AgentAPI{
				AgentOpsList: tt.fields.AgentOpsList,
				a:            tt.fields.a,
			}
			if err := i.SessionBegin(tt.args.r, tt.args.userDetail); (err != nil) != tt.wantErr {
				t.Errorf("AgentAPI.SessionBegin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAgentAPI_SessionEnd(t *testing.T) {
	type fields struct {
		AgentOpsList *ConfLocker
		a            *api.APIMode
	}
	type args struct {
		r         *http.Request
		tknStatus *model.TokenStatus
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &AgentAPI{
				AgentOpsList: tt.fields.AgentOpsList,
				a:            tt.fields.a,
			}
			if err := i.SessionEnd(tt.args.r, tt.args.tknStatus); (err != nil) != tt.wantErr {
				t.Errorf("AgentAPI.SessionEnd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
