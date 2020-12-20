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

package watcher

import (
	"context"
	"reflect"
	"testing"

	config "github.com/wzlib/wzconf"
	util "github.com/wzlib/wzutil"
)

func TestWatcherService_GetList(t *testing.T) {
	type fields struct {
		UnimplementedGatewayServer UnimplementedGatewayServer
	}
	type args struct {
		ctx context.Context
		in  *BadGatewayList
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GoodGatewayList
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := &WatcherService{
				UnimplementedGatewayServer: tt.fields.UnimplementedGatewayServer,
			}
			got, err := ws.GetList(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("WatcherService.GetList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WatcherService.GetList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWatcher_Proxy(t *testing.T) {
	type fields struct {
		Config     *config.WatcherConfig
		killSwitch chan bool
	}
	type args struct {
		mod string
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
			w := &Watcher{
				Config:     tt.fields.Config,
				killSwitch: tt.fields.killSwitch,
			}
			w.Proxy(tt.args.mod)
		})
	}
}

func TestWatcher_Monitor(t *testing.T) {
	type fields struct {
		Config     *config.WatcherConfig
		killSwitch chan bool
	}
	type args struct {
		kobjs []KillObjIntr
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
			w := &Watcher{
				Config:     tt.fields.Config,
				killSwitch: tt.fields.killSwitch,
			}
			w.Monitor(tt.args.kobjs...)
		})
	}
}

func TestWatcher_DeployGRPC(t *testing.T) {
	type fields struct {
		Config     *config.WatcherConfig
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
			w := &Watcher{
				Config:     tt.fields.Config,
				killSwitch: tt.fields.killSwitch,
			}
			w.DeployGRPC()
		})
	}
}

func TestWatcher_Scale(t *testing.T) {
	type fields struct {
		Config     *config.WatcherConfig
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
			w := &Watcher{
				Config:     tt.fields.Config,
				killSwitch: tt.fields.killSwitch,
			}
			w.Scale()
		})
	}
}

func TestWatcher_Scheduler(t *testing.T) {
	type fields struct {
		Config     *config.WatcherConfig
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
			w := &Watcher{
				Config:     tt.fields.Config,
				killSwitch: tt.fields.killSwitch,
			}
			w.Scheduler()
		})
	}
}

func TestInitWatcherD(t *testing.T) {
	tests := []struct {
		name string
		want *Watcher
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitWatcherD(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitWatcherD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitBareWatcher(t *testing.T) {
	tests := []struct {
		name string
		want *Watcher
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitBareWatcher(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitBareWatcher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitWatcher(t *testing.T) {
	type args struct {
		ac      *util.ConfigSet
		isProxy bool
	}
	tests := []struct {
		name    string
		args    args
		want    *Watcher
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InitWatcher(tt.args.ac, tt.args.isProxy)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitWatcher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitWatcher() = %v, want %v", got, tt.want)
			}
		})
	}
}
