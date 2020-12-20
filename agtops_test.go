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
	"reflect"
	"testing"

	config "github.com/wzlib/wzconf"
	core "github.com/wzlib/wzcore"
	util "github.com/wzlib/wzutil"
)

func TestNewAgentOps(t *testing.T) {
	type args struct {
		ac *util.ConfigSet
	}
	tests := []struct {
		name    string
		args    args
		want    *AgentOps
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAgentOps(tt.args.ac)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAgentOps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAgentOps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAgentOps_Start(t *testing.T) {
	type fields struct {
		Agent      core.AgentIntr
		killSwitch chan bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AgentOps{
				Agent:      tt.fields.Agent,
				killSwitch: tt.fields.killSwitch,
			}
			a.Start()
		})
	}
}

func TestAgentOps_Stop(t *testing.T) {
	type fields struct {
		Agent      core.AgentIntr
		killSwitch chan bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AgentOps{
				Agent:      tt.fields.Agent,
				killSwitch: tt.fields.killSwitch,
			}
			a.Stop()
		})
	}
}

func TestAgentOps_FileUploadOps(t *testing.T) {
	type fields struct {
		Agent      core.AgentIntr
		killSwitch chan bool
	}
	type args struct {
		from string
		to   string
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
			a := &AgentOps{
				Agent:      tt.fields.Agent,
				killSwitch: tt.fields.killSwitch,
			}
			if err := a.FileUploadOps(tt.args.from, tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("AgentOps.FileUploadOps() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAgentOps_FileDownloadOps(t *testing.T) {
	type fields struct {
		Agent      core.AgentIntr
		killSwitch chan bool
	}
	type args struct {
		from string
		to   string
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
			a := &AgentOps{
				Agent:      tt.fields.Agent,
				killSwitch: tt.fields.killSwitch,
			}
			if err := a.FileDownloadOps(tt.args.from, tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("AgentOps.FileDownloadOps() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAgentOps_Config(t *testing.T) {
	type fields struct {
		Agent      core.AgentIntr
		killSwitch chan bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *config.Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AgentOps{
				Agent:      tt.fields.Agent,
				killSwitch: tt.fields.killSwitch,
			}
			if got := a.Config(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AgentOps.Config() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAgentOps_Operation(t *testing.T) {
	type fields struct {
		Agent      core.AgentIntr
		killSwitch chan bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AgentOps{
				Agent:      tt.fields.Agent,
				killSwitch: tt.fields.killSwitch,
			}
			a.Operation()
		})
	}
}
