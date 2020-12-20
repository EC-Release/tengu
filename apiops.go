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
	//"fmt"
	//util "github.com/wzlib/wzutil"
	"sync"
	"errors"
	//"io/ioutil"
	api "github.com/wzlib/wzapi"
	"net/http"
	util "github.com/wzlib/wzutil"
	model "github.com/wzlib/wzschema"
	"encoding/json"
	"strings"
	//"strconv"
)

var (
	tmplt map[string]interface{}
)
const (
	
)

// InitAPI instantiates agent in API mode and sart the api ops 
func InitAPI(ac *util.ConfigSet) {
	
	t:=ac.AgtConf
	tmplt=t
	aapi:= &AgentAPI{
		AgentOpsList: &ConfLocker{
			m: make(map[string]interface{}),
		},
	}
	
	am:=api.NewAPIMode(aapi,ac)

	if *t["_sd"].(*string)=="" {
		panic("missing seed URL (-sed)")
	}
	
	if *t["_gh"].(*string)=="" {
		panic("missing host URL (-hst). E.g. https://myhostname/v1.2beta/<the app context path (-app)>")
	}

	am.Start(*t["_tl"].(*bool), *t["_ap"].(*string), *t["_an"].(*string), *t["_sd"].(*string), *t["_gh"].(*string))
	aapi.a=am
}

// ConfLocker is a mutex structure protecting agent config map
type ConfLocker struct{
	sync.RWMutex
	m map[string]interface{}
}

// AgentAPI structure host a reference of config map
type AgentAPI struct{
	AgentOpsList *ConfLocker
	a *api.APIMode
}

// InitOps serves as callback of the HTTP request during ops initialisation.
// endpoint /<rev>/<app-name>/initops
/*func (i *AgentAPI) InitOps(w http.ResponseWriter, r *http.Request, opsID string){
	
	//api.ErrResponse(w, 500, errors.New("internal error."), "internal error.")

	//create a placeholder for this operation forward
	i.AgentOpsList.Lock()
	i.AgentOpsList.m[opsID]="*"
	i.AgentOpsList.Unlock()

	return
}*/

// POSTConfig post agent config. being deprecated and replaced by internal RPC call
// endpoint /<rev>/<app-name>/config
/*func (i *AgentAPI) POSTConfig(w http.ResponseWriter, r *http.Request, opsID string){

	defer func(){
		if r:=recover();r!=nil{
			util.PanicRecovery(r)
		}
	}()

	var ok bool
	i.AgentOpsList.RLock()
	_,ok = i.AgentOpsList.m[opsID]
	i.AgentOpsList.RUnlock()
	if !ok{
		
		api.ErrResponse(w, 500, errors.New("operation unauthorised"), "operation unauthorised.")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		api.ErrResponse(w, 500, errors.New("internal error"), err.Error())
		return
	}

	ac:=util.InitConfigSet(tmplt,CLI_FLAGS)
	cfm,err:=ac.LoadFromJSON(body)
	if err!=nil{
		api.ErrResponse(w, 500,err, err.Error())		
		return
	}
	
	s:="verified"
	cfm["status"]=&s
	cfm["opsId"]=&opsID

	ac.AgtConf = cfm
	//cfg needs to be in flags format
	agtOps,err:=NewAgentOps(ac)
	if err!=nil{
		api.ErrResponse(w, 500, err, err.Error())
		return
	}
	i.AgentOpsList.Lock()
	i.AgentOpsList.m[opsID]=agtOps
	i.AgentOpsList.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)	
	_cfg, _ := json.Marshal(agtOps.Config())
	_, err=w.Write(_cfg)
	if err!=nil{
		api.ErrResponse(w, 500, err, err.Error())
		return
	}
	
	return
}*/

// GETConfig get agent config. being deprecated and replaced by internal RPC call	
// endpoint /<rev>/<app-name>/config
/*func (i *AgentAPI) GETConfig(w http.ResponseWriter, r *http.Request, opsID string){
	var ok bool
	i.AgentOpsList.RLock()
	_,ok = i.AgentOpsList.m[opsID]
	i.AgentOpsList.RUnlock()
	if !ok{
		api.ErrResponse(w, 401, errors.New("operation unauthorised"), "operation unauthorised.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	var agtOps *AgentOps
	i.AgentOpsList.RLock()
	agtOps,k:=i.AgentOpsList.m[opsID].(*AgentOps)
	if !k {
		api.ErrResponse(w, 500, errors.New("agent operation not found"), "operation not found")
		return
	}
	i.AgentOpsList.RUnlock()

	_cfg, _ := json.Marshal(agtOps.Config())
	_, err:=w.Write(_cfg)
	if err!=nil{
		api.ErrResponse(w, 500, err, err.Error())
		return
	}

	return
}*/

// PUTConfig put update agent config. being deprecated and replaced by internal RPC call
// endpoint /<rev>/<app-name>/config
/*func (i *AgentAPI) PUTConfig(w http.ResponseWriter, r *http.Request, opsID string){
	var ok bool
	i.AgentOpsList.RLock()
	_,ok = i.AgentOpsList.m[opsID]
	i.AgentOpsList.RUnlock()
	if !ok{
		api.ErrResponse(w, 500, errors.New("operation unauthorised"), "operation unauthorised.")
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err:=w.Write([]byte(`{"status":"updated"}`))
	if err!=nil{
		api.ErrResponse(w, 500, err, err.Error())
		return
	}

	return
}*/

// Hire implements agent ops interface. being deprecated and replaced by internal RPC call
// endpoint /<rev>/<app-name>/hire
/*func (i *AgentAPI) Hire(w http.ResponseWriter, r *http.Request, opsID string){

	var ok bool
	i.AgentOpsList.RLock()
	_,ok = i.AgentOpsList.m[opsID]
	i.AgentOpsList.RUnlock()
	if !ok{
		api.ErrResponse(w, 500, errors.New("operation unauthorised"), "operation unauthorised.")
		return
	}
	
	var agtOps *AgentOps
	i.AgentOpsList.RLock()
	agtOps,k:=i.AgentOpsList.m[opsID].(*AgentOps)
	if !k {
		api.ErrResponse(w, 500, errors.New("agent operation not found"), "operation not found")
		return
	}
	i.AgentOpsList.RUnlock()

	go agtOps.Start()
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err:=w.Write([]byte(`{"status":"running"}`))
	if err!=nil{
		api.ErrResponse(w, 500, err, err.Error())
		return
	}
	return
}*/

// Resume implements agent ops interface. being deprecated and replaced by internal RPC call
// endpoint /<rev>/<app-name>/resume
/*func (i *AgentAPI) Resume(w http.ResponseWriter, r *http.Request, opsID string){

	var ok bool
	i.AgentOpsList.RLock()
	_,ok = i.AgentOpsList.m[opsID]
	i.AgentOpsList.RUnlock()
	if !ok{
		api.ErrResponse(w, 500, errors.New("operation unauthorised"), "operation unauthorised.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err:=w.Write([]byte(`{"status":"running"}`))
	if err!=nil{
		api.ErrResponse(w, 500, err, err.Error())
		return
	}
	return

}*/

// Suspend implements agent ops interface. being deprecated and replaced by internal RPC call
// endpoint /<rev>/<app-name>/suspend
/*func (i *AgentAPI) Suspend(w http.ResponseWriter, r *http.Request, opsID string){

	var ok bool
	i.AgentOpsList.RLock()
	_,ok = i.AgentOpsList.m[opsID]
	i.AgentOpsList.RUnlock()
	if !ok{
		api.ErrResponse(w, 500, errors.New("operation unauthorised"), "operation unauthorised.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err:=w.Write([]byte(`{"status":"stop"}`))
	if err!=nil{
		api.ErrResponse(w, 500, err, err.Error())
		return
	}
	return

}*/

// Status implements agent ops interface. being deprecated and replaced by internal RPC call
// endpoint /<rev>/<app-name>/status
/*func (i *AgentAPI) Status(w http.ResponseWriter, r *http.Request, opsID string){

	var ok bool
	i.AgentOpsList.RLock()
	_,ok = i.AgentOpsList.m[opsID]
	i.AgentOpsList.RUnlock()
	if !ok{
		api.ErrResponse(w, 500, errors.New("operation unauthorised"), "operation unauthorised.")
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err:=w.Write([]byte(`{"status":"ok"}`))
	if err!=nil{
		api.ErrResponse(w, 500, err, err.Error())
		return
	}
	return

}*/

// Fire implements agent ops interface. being deprecated and replaced by internal RPC call
// endpoint /<rev>/<app-name>/fire
/*func (i *AgentAPI) Fire(w http.ResponseWriter, r *http.Request, opsID string){

	var ok bool
	i.AgentOpsList.RLock()
	_,ok = i.AgentOpsList.m[opsID]
	i.AgentOpsList.RUnlock()
	if !ok{
		api.ErrResponse(w, 500, errors.New("operation unauthorised"), "operation unauthorised.")
		return
	}

	var agtOps *AgentOps

	i.AgentOpsList.RLock()
	agtOps,k:=i.AgentOpsList.m[opsID].(*AgentOps)
	i.AgentOpsList.RUnlock()
	if !k {
		api.ErrResponse(w, 500, errors.New("agent is invalid"), "invalid")
		return
	}
	go agtOps.Stop()
	
	i.AgentOpsList.Lock()
	delete(i.AgentOpsList.m,opsID)
	i.AgentOpsList.Unlock()
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err:=w.Write([]byte(`{"status":"removed"}`))
	if err!=nil{
		api.ErrResponse(w, 500, err, err.Error())
		return
	}

	return
}*/

// Ext extend the Agent API in the following format-
//<agent_url>/<rev. e,g v1.1|v1.1beta>/<app name from -app flag>/<api>/custom/path/name
//this api extension should assume that the user has been authenticated and
//promoted by the security chain.
func (i *AgentAPI) Ext(w http.ResponseWriter, r *http.Request, tknStatus *model.TokenStatus, db *util.DB) {
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	var jsonString []byte
	
	segs:=strings.Split(r.URL.Path,"/")
	
	if len(segs)>4 {
		switch r.Method {
		case http.MethodPost:
			fallthrough
		case http.MethodPatch:
			fallthrough
		case http.MethodPut:

			var t map[string]interface{}
			err := json.NewDecoder(r.Body).Decode(&t)
			if err != nil {
				util.ErrResponse(w, 500, err, err.Error())
				return		
			}

			//example getting the user detail
			userDetail,err:=i.a.FetchUserDetailByUserId(tknStatus.DevId)
			if err!=nil {
				util.ErrResponse(w, 500, err, err.Error())
				return		
			}
			
			t["status"]="ok"
			t["userId"]=tknStatus.DevId
			t["email"]=userDetail.Email

			db.Set(tknStatus.DevId,segs[4],t)
				
			jsonString,err=json.Marshal(t)
			if err!=nil{
				util.ErrResponse(w, 500, err, err.Error())
				return	
			}
			
		case http.MethodGet:
			val:=db.Get(segs[4])
			if val==nil{
				err:=errors.New("key "+segs[4]+" not found")
				util.ErrResponse(w, 500, err, err.Error())
				return
			}

			j,err:=json.Marshal(val)
			if err!=nil{
				util.ErrResponse(w, 500, err, err.Error())
				return	
			}
			jsonString=j
		case http.MethodDelete:
			val:=db.Del(segs[4])
			if val==nil{
				err:=errors.New("key "+segs[4]+" not found")
				util.ErrResponse(w, 500, err, err.Error())
				return
			}

			j,err:=json.Marshal(val)
			if err!=nil{
				util.ErrResponse(w, 500, err, err.Error())
				return	
			}
			jsonString=j
			
		default:
			util.ErrResponse(w, 401, errors.New("unsupported method"), "unsupported")
			return
		}
	}

	if len(segs)==4 {
		val:=db.GetKeys()
		if val==nil{
			err:=errors.New("key "+segs[4]+" not found")
			util.ErrResponse(w, 500, err, err.Error())
			return
		}

		j,err:=json.Marshal(val)
		if err!=nil{
			util.ErrResponse(w, 500, err, err.Error())
			return	
		}
		jsonString=j
	}

	_, err:=w.Write(jsonString)
	if err!=nil{
		util.ErrResponse(w, 500, err, err.Error())
	}
}

// SessionBegin invoked first when received a user authentication callback. In the event of OIDC integration, the userDetail refers to the user information in the targeted framework provided by the integration endpoint. Like SessionEnd, the interface SessionBegin will suspend the remaining process if the return error is not nil
func (i *AgentAPI) SessionBegin(r *http.Request, userDetail *model.UserDetail) error {
	util.GetCLog().InfoLog("user ",userDetail.FullName,"(",userDetail.Email,userDetail.UserId,") is now authenticated to the system.")
	return nil
	
}

// SessionEnd similar to GC invoked by the endpoint /<rev>/logout for the purpose of managing app resource. The Logout process will stop if the return error is not nil
// endpoint /<rev>/logout
func (i *AgentAPI) SessionEnd(r *http.Request, tknStatus *model.TokenStatus) error {

	util.GetCLog().InfoLog("request url",r.URL.String(),"user Id",tknStatus.DevId)
	return nil
}
