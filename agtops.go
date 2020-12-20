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
	"time"
	"strings"
	"errors"
	//api "github.com/wzlib/wzapi"
	config "github.com/wzlib/wzconf"
	"os"
	"os/signal"
	util "github.com/wzlib/wzutil"
	core "github.com/wzlib/wzcore"
	"strconv"
)

const (
	REPORT_DUR=60
)

// AgentOps define Agent operations
type AgentOps struct {
	//cli-compatible vars
	//Config  map[string]interface{}
	Agent  core.AgentIntr
	killSwitch chan bool
}

// NewAgentOps returns an AgentOps instance and the agent interface
func NewAgentOps(ac *util.ConfigSet) (*AgentOps,error){
	//func NewAgentOps(c map[string]interface{}) (*AgentOps,error)
	//c:=ac.AgtConf

	//util.Init(*ac.AgtConf[AGENT_MODE].(*string),*ac.AgtConf[AGENT_DEBUG].(*bool))
	if *ac.AgtConf[AGENT_DEBUG].(*bool) {
		util.GetCLog().SetDebug()
	}
	
	_agt,err:=core.InitAgent(ac)
	if err!=nil{
		return nil,err
	}
	
	return &AgentOps{
		Agent: _agt,
	},nil
}

// Start invokes start in AgentOps
func (a *AgentOps) Start(){
	defer func(){
		if r:=recover();r!=nil{
			util.PanicRecovery(r)
		}
	}()
	
	a.Agent.Hire()	
}

// Stop invokes stop in AgentOps
func (a *AgentOps) Stop() {
	defer func(){
		if r:=recover();r!=nil{
			util.PanicRecovery(r)
		}
	}()

	a.Agent.Fire()
	a.killSwitch<-true
}

// FileUploadOps file upload supported in x:client modes only  
func (a *AgentOps) FileUploadOps(from,to string) error {
	defer func(){
		if r:=recover();r!=nil{
			util.PanicRecovery(r)
		}
	}()

	conf:=a.Agent.Config()
	from=strings.Trim(from," ")
	to=strings.Trim(to," ")
	conf.FileToUpload=from+":"+to
	f, ok := a.Agent.(*core.XClient)
	if !ok {
		return errors.New("incompatible agent mode")
	}

	if err:=f.UploadFile();err!=nil{
		return err
	}

	return nil
}

// FileDownloadOps file download supported in x:client modes only  
func (a *AgentOps) FileDownloadOps(from,to string) error {
	defer func(){
		if r:=recover();r!=nil{
			util.PanicRecovery(r)
		}
	}()

	conf:=a.Agent.Config()
	from=strings.Trim(from," ")
	to=strings.Trim(to," ")
	conf.FileToDownload=from+":"+to
	f, ok := a.Agent.(*core.XClient)
	if !ok {
		return errors.New("incompatible agent mode")
	}

	if err:=f.DownloadFile();err!=nil{
		return err
	}

	return nil
}

// Config return the agent config flags from main package
func (a *AgentOps) Config()(*config.Config){
	return a.Agent.Config()
}

func (a *AgentOps) Operation(){

	a.killSwitch = make(chan bool,1)
	go func(){
		defer func(){
			if r:=recover();r!=nil{
				util.PanicRecovery(r)
			}
			
			//util.DbgLog("agent exited all ops.")
			//os.Exit(0)
		}()

		conf:=a.Agent.Config()

		defer func(){
			switch conf.ClientType {
			case "gateway":
				fallthrough
			case "gw:server":
				fallthrough			
			case "gw:client":
				if conf.Report.LastUsage>0{
					_ref:=conf.UsageReporting()
					_=<-_ref
				}
			case "x:gateway":
				//placeholder
			}
			
			conf.Report.LastUsage=0
			conf.Report.LastReport=time.Now()
		}()

		interrupt := make(chan os.Signal, 1)

		signal.Notify(interrupt, os.Interrupt)

		//bearer token ticker
		ticker := time.NewTicker(time.Second*time.Duration(conf.Duration))
		
		
		// let all modes to exercise the _rt service time ticker instead
		/*if conf.ClientType!="gateway"{
		//preset to 100 year
		_ref3=config.AUTOREFRESH_DUR
	}*/

		//service ticker
		_ref3:=int64(REPORT_DUR)
		_rt := time.NewTicker(time.Second*time.Duration(_ref3))
		
		defer ticker.Stop()
		defer _rt.Stop()
		
		for {
			select {
				
			case <-_rt.C:
				go func(){
					switch conf.ClientType {
					case "gateway":
						fallthrough
					case "gw:server":
						fallthrough			
					case "gw:client":
						//report the usage 
						//if conf.Report.LastUsage>0 {
						util.GetCLog().InfoLog("reporting usage..")

						//for {
						_ref:=conf.UsageReporting()
						_resp:=<-_ref
						if _resp["status"]=="EC_GRANT" {
							//break
							conf.Report.LastUsage=0
							util.GetCLog().InfoLog("usage reported.")
							
						} else {
							util.GetCLog().ErrLog("service is invalid. Contact your EC Service administrator for more detail. agent reporting will re-try in next",REPORT_DUR,"sec")
						}
						//x:gateway only relevant reporting usage
					case "x:gateway":
						//placeholder
					}
					
				}()
			case <-ticker.C:
				go func(){
					switch conf.ClientType {
					case "client":
						fallthrough
					case "server":
						util.GetCLog().InfoLog("Triggering token Auto-Refresh..")
						_op:=util.TokenRefresh(conf.ClientId,conf.ClientSecret, conf.OAuthURL)
						_ref:=<-_op
						conf.CFToken=_ref.AccessToken
						util.GetCLog().InfoLog("Token refreshed. The token will be expired in "+strconv.FormatInt(_ref.ExpiresIn/60,10)+" minutes. Approx. "+strconv.FormatInt(conf.Duration/60,10)+" minutes to the next auto-refresh.")
					}
				}()
			case <-a.killSwitch:
				/*case <-interrupt:
			agt.Fire()
			p, err := os.FindProcess(os.Getpid())

				util.GetCLog().ErrLog("failed in find process.",err)

			if err=p.Signal(os.Interrupt);err!=nil{
				util.GetCLog().ErrLog("failed in signify os.interrupt.",err)				
			}*/
				panic("exited "+conf.ClientType+" mode")
			}
		}
	}()
	
	return
}
