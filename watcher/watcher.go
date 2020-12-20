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
	"net"
	//"fmt"
	"context"
	util "github.com/wzlib/wzutil"
	//"sync"
	//"time"
	//"gopkg.in/yaml.v2"
	//"errors"
	//"os/exec"
	//"runtime"
	//"encoding/json"
	"os"
	"os/signal"
	//"strconv"
	"google.golang.org/grpc"
	config "github.com/wzlib/wzconf"
	//agt "github.com/Enterprise-Connect/agent"
	//"github.com/opencontainers/runc/libcontainer"
	//_ "github.com/opencontainers/runc/libcontainer/nsenter"
)
const (
	HEALTH_DURATION = 600
)

type Watcher struct {
	Config *config.WatcherConfig
	killSwitch chan bool
}

type WatcherService struct {
	UnimplementedGatewayServer
}

var (
	wtrConf *config.WatcherConfig
)

func (ws *WatcherService)GetList(ctx context.Context, in *BadGatewayList) (*GoodGatewayList, error) {
	util.GetCLog().DbgLog("Received:",in.String())
	return &GoodGatewayList{List: []*GatewayInfo{
		&GatewayInfo{
			GtwId: "who cares",
			Active: true,
		},
	},}, nil
}

func (w *Watcher)Proxy(mod string) {

	switch mod {
	case "client":
		fallthrough
	case "gw:client":
		fallthrough
	case "x:client":
		go w.Config.ProxyTCPReq()
		go w.Config.ProxyHTTPReq()		
	default:
		go w.Config.ProxyHTTPReq()		
	}	
	
	w.Monitor()
}

type KillObjIntr interface {
	Stop()
}

func (w *Watcher)Monitor(kobjs ...KillObjIntr) {	

	w.killSwitch = make(chan bool, 1)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-interrupt:
			//_=w.Config.KillProc()
			for _,k:=range kobjs {
				k.Stop()
			}
			panic("system interrupted.")
		
		case <-w.killSwitch:
			util.GetCLog().InfoLog("exited monitoring.")
			return
		}
	}	
}

func (w *Watcher) DeployGRPC(){

	//port:=config.WATCHER_FIELD_DAEMON_PORT_VAL	
	lis, err := net.Listen("tcp", config.DAEMON_FIELD_DAEMON_PORT_VAL)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	RegisterGatewayServer(s, &WatcherService{})

	util.GetCLog().InfoLog("watcher grpc will listen on",config.DAEMON_FIELD_DAEMON_PORT_VAL)

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func (w *Watcher) Scale() {
	
	defer func(){
		if r:=recover();r!=nil{
			util.PanicRecovery(r)
		}
		w.killSwitch<-true
	}()

	wb := func(isWatcher bool) ([]string) {

		var op []string
		if isWatcher {	
			op=[]string{"./agent","-cfg","./conf.yml","-pxy","-wtr"}
		} else {
			op=[]string{"./agent","-cfg","./conf.yml"}
		}
		return op
		
	}

	if err:=w.Config.DeployWatcher(wb);err!=nil {
		panic(err)
	}
}

func (w *Watcher) Scheduler() {	
	
	go w.DeployGRPC()

	if err:=w.Config.InitOpsScheduler();err!=nil{
		panic(err)
	}		

	w.Monitor()
}

func InitWatcherD() (*Watcher) {

	wtr:=InitBareWatcher()
	
	go func(){
		defer func(){	
			if r:=recover();r!=nil{
				util.PanicRecovery(r)
			}
			
		}()
		if err:=wtr.Config.InitDaemonContainer();err!=nil{
			util.GetCLog().DbgLog("daemon failed initialisd.",err)
		}
	}()

	return wtr
}

func InitBareWatcher() (*Watcher) {

	return &Watcher{
		Config: config.NewWatcherConfig(),
	}

}

//intialise watcher internal RPC system
//@@cpath config file path
//func InitWatcher(watrURL string, cpath string) (*Watcher,error) {
func InitWatcher(ac *util.ConfigSet, isProxy bool) (*Watcher,error) {
	
	_cff:=util.CURRENT_PATH()+"/"+(*ac.Template[util.CONFIG_FILE_YML].(*string))
	util.GetCLog().DbgLog("_cff",_cff)
	wc:=config.NewWatcherConfig()
	wc.ExternalCfgFile=_cff
	
	if isProxy {
		err:=wc.InitProxy(ac.WatcherConf,ac.CredConf)
		if err!=nil {
			return nil, err
		}
	} else {
		err:=wc.Init(ac.WatcherConf,ac.CredConf)
		if err!=nil {
			return nil, err
		}
	}
		
	wtr:=&Watcher{
		Config: wc,
	}

	return wtr, nil	
}
