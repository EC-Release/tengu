/*
 * Copyright (c) 2016 General Electric Company. All rights reserved.
 *
 * The copyright to the computer software herein is the property of
 * General Electric Company. The software may be used and/or copied only
 * with the written permission of General Electric Company or in accordance
 * with the terms and conditions stipulated in the agreement/contract
 * underch the software has been supplied.
 *
 * author: apolo.yasuda@ge.com
 */

package main

import (
	"strings"
	"flag"
	//"errors"
	"os"
	"fmt"
	util "github.com/wzlib/wzutil"
	model "github.com/wzlib/wzschema"
	watcher "github.com/EC-Release/tengu/watcher"
	conf "github.com/wzlib/wzconf"
	//remove the pprof dependency to simplify agent binary
	//_ "net/http/pprof"
)

const (
	// EC_LOGO serves as the team's trademark
	EC_LOGO = `
           ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄
          ▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌
          ▐░█▀▀▀▀▀▀▀▀▀ ▐░█▀▀▀▀▀▀▀▀▀
          ▐░▌          ▐░▌   
          ▐░█▄▄▄▄▄▄▄▄▄ ▐░▌
          ▐░░░░░░░░░░░▌▐░▌
          ▐░█▀▀▀▀▀▀▀▀▀ ▐░▌
          ▐░▌          ▐░▌
          ▐░█▄▄▄▄▄▄▄▄▄ ▐░█▄▄▄▄▄▄▄▄▄ 
          ▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌
           ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀  @Enterprise-Connect 
`
	COPY_RIGHT = "Enterprise-Connect,  @General Electric"
	ISSUE_TRACKER = "https://github.com/EC-Release/sdk/issues"

	// AUTH_HEADER agent authorization header
	AUTH_HEADER = "Authorization"
	
	// EC_SUB_HEADER ec service header in predix
	EC_SUB_HEADER  = "Predix-Zone-Id"

	// CF_INS_IDX_EV app index available in a cf environment
	CF_INS_IDX_EV  = "CF_INSTANCE_INDEX"
	// CF_INS_HEADER forwarding header targeting a cf environment
	CF_INS_HEADER  = "X-CF-APP-INSTANCE"

	// EC_INS_IDX_EV app index available in a watcher environment
	EC_INS_IDX_EV  = "EC_INSTANCE_INDEX"
	// EC_INS_HEADER forwarding header targeting a watcher environement
	EC_INS_HEADER  = "X-EC-APP-INSTANCE"

	CA_URL = "https://github.com/EC-Release/certifactory"
	
	FILE_TO_UPLOAD = "_fu"
	FILE_TO_DOWNLOAD = "_fd"
	
	AGENT_MODE = "_md"
	AGENT_REVISION = "_v"
	AGENT_LOCAL_PORT = "_lp"
	AGENT_LOCAL_PORT_VAL = ":7990"
	AGENT_HEALTH_PORT = "_hc"
	AGENT_HEALTH_PORT_VAL = ":7991"
	AGENT_DEBUG = "_dg"

	AGENT_GW_PORT = "_gt"
	//AGENT_INSTANCE_ID = "_is"
	//WATCHER_ENV = "_ev"
	//WATCHER_INITIAL = "_in"
	//WATCHER_LAUNCH_CMD = "_ee"

	WATCHER_MODE = "_wt"
	WATCHER_PROXY = "_px"
	WATCHER_DAEMON = "_dm"

	TOKEN_EXPIRATION = "_ex"

)

var (
	// CFM is a collection of flags setting
	CFM = make(map[string]interface{})
	CERT_BASE64 string = ""

	// CLI_FLAGS bootstrap the cli
	CLI_FLAGS = map[string]interface{}{
		"fup":[]interface{}{FILE_TO_UPLOAD,"",`Specify a file to upload to the server agent.`},
		"fdw":[]interface{}{FILE_TO_DOWNLOAD,"",`Specify a file to download from the client agent.`},
		"cfg":[]interface{}{util.CONFIG_FILE_YML,"",`Specify the config file to launch the agent.`},
		"mod":[]interface{}{AGENT_MODE,"agent",`Specify the EC Agent Mode in "client", "server", or "gateway".`},
		"sed":[]interface{}{"_sd","",`Specify upstream seed node URL to for joining a BC cluster.`},
		"ver":[]interface{}{AGENT_REVISION, false, "Show EC Agent's version."},
		
		"hca":[]interface{}{AGENT_HEALTH_PORT,AGENT_HEALTH_PORT_VAL,`Specify a port# to turn on the Healthcheck API. This flag is always on when in the "gateway mode" with the provisioned local port. Upon provisioned, the api is available at <agent_uri>/health.`},
		"lpt":[]interface{}{AGENT_LOCAL_PORT,AGENT_LOCAL_PORT_VAL,`Specify the default EC port#.`},
		"gpt":[]interface{}{AGENT_GW_PORT,AGENT_LOCAL_PORT_VAL,`Specify the gateway port# in fuse-mode. (gw:server|gw:client)`},
		//deprecated. pprof imported in main pkg
		//"911":[]interface{}{"_91",false,`Internal system profiling`},
		"apt":[]interface{}{"_ap","17990",`Specify the EC http endpoint port# of the agent when -api is set.`},
		"cps":[]interface{}{"_cp",0,`Specify the Websocket compression-ratio for agent's inter-communication purpose. E.g. [0-9]. "0" is no compression whereas "9" is the best. "-1" indicates system default. "-2" is HuffmanOnly. See RFC 1951 for more detail.`},
		"tid":[]interface{}{"_ts","",`Specify the Target EC Server Id if the "client" mode is set`},
		//"ins":[]interface{}{util.AGENT_INSTANCE_ID,"0","Agent instance ID."},
		"hst":[]interface{}{"_gh","","Specify a local http endpoint/gateway URL which's accessible to public. E.g. wss://<somedomain>:8989"},
		"sst":[]interface{}{"_sh","","Specify the EC Service URI. E.g. https://<service.of.predix.io>"},
		"pth":[]interface{}{"_ph","","Specify the directory to the certificate/key."},
		"dat":[]interface{}{"_da","","Specify the string to support the cryto usage."},
		"sgn":[]interface{}{"_sg",false,"Start a CA x509 Cert-Signing process."},
		"ivk":[]interface{}{"_ik",false,"Invoke a http request. The flag requires endpoint url (-url), bearer token (-tkn), and json data (-dat) to complete."},
		"mtd":[]interface{}{"_mt","GET","Indicate a http method (POST, GET, PUT, DELETE) for the HTTP request (-ivk)."},
		"log":[]interface{}{"_lg",false,"Enable remote log session for a URL (-url)"},
		"url":[]interface{}{"_ul","","Specify the http endpoint to invoke."},
		"agt":[]interface{}{"_at","","Shorthand indicate the agent mode used when fusing the api and an gateway mode"},
		//deprecated
		//"sgf":[]interface{}{"_sf",false,"Sign a GPG-to-EC-specific file and generate its signature, keypair."},
		"sig":[]interface{}{util.SIG_FILE_FLD,"","a GPG-to-EC-specific signature file."},
		
		//bgn ca api
		"enc":[]interface{}{"_ep",false,"Generate a base64 encrypted string with base64 cert string (-pbk), based on the input string (-dat)."},
		"dec":[]interface{}{"_dp",false,"Decript an encrypted string with the private key (-pvk), based on the given encrypted string (-dat), and output the base64-encoded original string."},

		"gsg":[]interface{}{"_gs",false,"Generate the signature based on the given private key (-pvk), and the message (-dat)."},
		"vsg":[]interface{}{"_vs",false,"Verify the signature part (-dat) of the agent token based on the validatee's x509 cert (-pbk), and original message (-osg) part of the token, all in base64 encoded. When paired w/ -smp, the function will return the owner id of this signature."},
		"vst":[]interface{}{"_vt",false,"Verify the signature (-dat) based on the validatee's x509 cert (-pbk), and original message (-osg), all in base64 encoded."},
		"vfy":[]interface{}{"_vf",false,"Verify the compatibility of the base64 encoded x509 certificate (-pbk)."},
		"exp":[]interface{}{TOKEN_EXPIRATION,"60","Evaluate the token expiration in minutes since the token creation."},
		"osg":[]interface{}{"_od","","The signed message in base64 encoded format."},
		"hsh":[]interface{}{"_hs",false,"Generate the agent hashed string based on the agent mode value (-mod)"},
		"pvk":[]interface{}{util.BASE64PRVK_FLD,"","Specify a RSA/DSA private key in base64 encoded format to support crypto usage."},
		//deprecated for enhance security
		//"psp":[]interface{}{"_pw","","Specify the passphrase of the private key, combined with the flags (-dec -pth <dir-pvtkey> -psp <passphrase> -dat <encypted-data>) for one-step decryption."},
		"rnw":[]interface{}{"_rn",false,"Renew a previous-issued x509 certificate."},
		"vld":[]interface{}{util.VALIDATE_FILE_ENC_FLD,false,"validate the authenticity of the file-encryption keypair, and signature."},
		"pvd":[]interface{}{"_pd",false,`Print the decrypted private key with -pvk <base64 private pem key>`},
		"pbk":[]interface{}{util.BASE64CERT_FLD,"","Base64 encoded x509 certificate or DSA public string to support crypto usage."},
		//"tlp":[]interface{}{"_pk","","Specify the relative path to a TLS key when -tls is set. E.g. ./path/to/key.pem."},
		//"tlb":[]interface{}{"_pc","","Specify the relative path to a TLS cert when -tls is set. E.g. ./path/to/cert.pem."},
		"tls":[]interface{}{util.AGENT_TLS_ENABLED,false,"When -tls is set, both tlp (privatekey) tlb (public) are required."},
		
		//end ca api
		"rht":[]interface{}{"_rh","",`Specify the Resource Host if "x:server" mode is set. E.g. <someip>, <somedomain>. value will be discard when TLS is specified.`},
		"rpt":[]interface{}{"_rp","0",`Specify the Resource Port# if the "server" mode is set. E.g. 8989, 38989`},
		"aid":[]interface{}{"_id","","Specify the agent Id assigned by the EC Service. You may find it in the Cloud Foundry VCAP_SERVICE"},
		"tkn":[]interface{}{"_tk","","Specify the OAuth Token. The token may expire depending on your OAuth provisioner. This flag is ignored if OAuth2 Auto-Refresh were set."},
		"gtk":[]interface{}{"_gk",false,`Output the bearer token from instance "-oa2". Client ID "-cid" must be specified.`},
		
		"pxy":[]interface{}{WATCHER_PROXY,false,"when specified, watcher will serve as a network node to proxy requests, and/or to load-balance the traffic for child-containers."},
		"cid":[]interface{}{"_ci","","Specify the client Id to auto-refresh the OAuth2 token."},

		//deprecated in v1.1. introducing agent hash.
		//"csc":[]interface{}{"_cs","","Specify the client secret to auto-refresh the OAuth2 token."},

		"oa2":[]interface{}{"_oa","","Specify URL of the OAuth2 provisioner. E.g. https://<somedomain>/oauth/token"},
		//deprecated in x:<agent_mode>
		"dur":[]interface{}{"_du",0,"Specify the duration for the next token refresh in seconds. (default 100 years)"},
		//deprecated
		//"crt":[]interface{}{"_ct","","Specify the relative path of a digital certificate to operate the EC agent. (.pfx, .cer, .p7s, .der, .pem, .crt)"},
		
		"wtl":[]interface{}{"_wl","0.0.0.0/0,::/0","Specify the ip(s) whitelist in the cidr net format. Concatenate ips by comma. E.g. 89.24.9.0/24, 7.6.0.0/16"},
		"bkl":[]interface{}{"_bl","","Specify the ip(s) blocklist in the IPv4/IPv6 format. Concatenate ips by comma. E.g. 10.20.30.5, 2002:4559:1FE2::4559:1FE2"},
		"plg":[]interface{}{"_pg",false,`Enable EC plugin list. This requires the plugins.yml file presented in the agent path.`},
		"inf":[]interface{}{"_if",false,"The Product Information."},
		"dbg":[]interface{}{AGENT_DEBUG,false,"Turn on debug mode. This will introduce more error information. E.g. connection error."},
		
		//deprecated
		"zon":[]interface{}{"_zn","",`Specify the Zone/Service Inst. Id. required in the "gateway" mode.`},
		
		"gen":[]interface{}{"_gc",false,"Generate a x509 certificate request for the usage validation purpose."},
		"shc":[]interface{}{"_sc",false,"Health API requires basic authentication for Health APIs."},
		"vln":[]interface{}{"_vn",false,"Enable support for EC VLAN Network."},
		"wtr":[]interface{}{WATCHER_MODE,false,"Enable watcher process."},
		"dae":[]interface{}{WATCHER_DAEMON,false,"Enable daemon process."},
		"grp":[]interface{}{"_gp","","GroupID needed for Agent Client/Server."},
		
		//deprecated for security
		//"tse":[]interface{}{"_et",false,"Create a EC-compatible Token, with publickey (-pbk) and an optional 32-digits uuid (-dat). "},
		//deprecated for security
		//"tsd":[]interface{}{"_dt",false,"Check the detail of the EC token (-dat) "},
		//"exe":[]interface{}{WATCHER_LAUNCH_CMD,false,"Watcher launch command."},
		"out":[]interface{}{"_ot","","convert the current setting (command|yaml) to a .properties file."},
		"fil":[]interface{}{"_fl","","load the content of input file directory."},
	

		//deprecated. use "-mod api" instead
		//"api":[]interface{}{"_ht",false,"Operate agent in HTTP mode."},
		"app":[]interface{}{"_an","webui","The context path for the entry of the Web UI when -api is set. This is secured by the OAuth authentication."},

		//deprecated. use "-mod oauth" instead
		//"oau":[]interface{}{"_ou",false,"Launch OAuth2 Provider Store. Combined command with private key (-pvk), cert (-pbk)"},
		"smp":[]interface{}{"_sm",false,"simplifying the output for integration purpose."},

		"env":[]interface{}{util.CONFIG_FROM_ENV,false,"agent configuration set via environment variables."},
		//watcher mode
		//deprecated
		//"ini":[]interface{}{WATCHER_INITIAL,false,"Inidicate whether this process belong to the watcher container initialisation."},
		//"rev":[]interface{}{"_rv","v1beta.fukuoka.1676","The default agent revision for watcher ops."},
		//"art":[]interface{}{"_at","ecagent_linux_sys","The target artifact resulting from the tar.gz file."},
	}
)

func init(){

	bc:=&model.BrandingConfig{
		CONFIG_MAIN: "/.ec",
		BRAND_CONFIG: "EC",
		// the agent will consume the passphrase extension as in the format of <BRAND_CONF>_<PASSPHRASE_EXT>
		PASSPHRASE_EXT: "PPS",
		ART_NAME: "agent",
		LOGO: EC_LOGO,
		COPY_RIGHT: COPY_RIGHT,
		HEADER_PLUGIN: "ec-plugin",
		HEADER_CONFIG: "ec-config",
		STREAM_PATH: "/agent",
		HEADER_AUTH: AUTH_HEADER,
		HEADER_SUB_ID: EC_SUB_HEADER,
		HEADER_CF_INST: CF_INS_HEADER,
		HEADER_INST: EC_INS_HEADER,
		ENV_CF_INST_IDX: CF_INS_IDX_EV,
		ENV_INST_IDX: EC_INS_IDX_EV,
		URL_CA: CA_URL,
		//URL_WATCHER_CONF: WATCHER_URL,
		//URL_WATCHER_REPO: WATCHER_CONTR_URL,
		URL_ISSUE_TRACKER: ISSUE_TRACKER,
	}
	
	util.Branding(bc)
	l:=util.NewAppLog("agent")
	util.SetCLog(*l)
}

func main() {	
	defer func(){
		if r:=recover();r!=nil{
			//util.PanicRecovery(r)
			fmt.Println(" [EC Agent] main pacakge exception:",r)
			os.Exit(1)
			return
		}
		os.Exit(0)
	}()

	//avoid i/o
	//fmt.Println(" [EC Agent] loading application parameters..")

	//dynamically assign flags
	for k, v := range CLI_FLAGS {
		_v:=v.([]interface{})
		//util.InfoLog(_v[2].(string))
		switch _v[1].(type) {
		case string:
			CFM[_v[0].(string)]=flag.String(k,_v[1].(string),_v[2].(string))
		case bool:
			CFM[_v[0].(string)]=flag.Bool(k,_v[1].(bool),_v[2].(string))
		case int:
			CFM[_v[0].(string)]=flag.Int(k,_v[1].(int),_v[2].(string))
		default:
			panic("flags "+k+" is not implemented.")
		}
	}
	
	flag.Parse()

	ac:=util.InitConfigSet(CFM, CLI_FLAGS)

	if *CFM[AGENT_DEBUG].(*bool) {
		util.GetCLog().SetDebug()
	}
	
	if *CFM[util.CONFIG_FILE_YML].(*string)!=""{
		
		if *CFM["_dp"].(*bool) {
			err:=ac.LoadFromEncryptedYAML(conf.GetRev())
			if err!=nil {
				panic(err)
			}
		} else {
			err:=ac.LoadFromYAML()
			if err!=nil {
				panic(err)
			}
		}
		
		if *ac.AgtConf[util.CONFIG_FILE_OUT].(*string)!=""{
	
			p:=*ac.AgtConf[util.CONFIG_FILE_OUT].(*string)
			err:=ac.OutputPropertyFile(p)
			if err!=nil {
				panic(err)
			}

			util.GetCLog().InfoLog("property file",p,"generated.")
			return
		}

	} else if (*CFM[util.CONFIG_FROM_ENV].(*bool)){
		if err:=ac.LoadFromEnv();err!=nil{
			panic(err)
		}
	} else {
		ac.AgtConf = CFM
	}

	if *ac.Template[AGENT_GW_PORT].(*string)!=AGENT_LOCAL_PORT_VAL {
		ac.AgtConf[AGENT_GW_PORT]=ac.Template[AGENT_GW_PORT]
	}

	
	//watcher cmd flag must be explicitly specified
	if *CFM[WATCHER_MODE].(*bool)==true {

		if *CFM[WATCHER_DAEMON].(*bool) {
			util.SetCLog(*util.NewAppLog("daemon"))
			wat:=watcher.InitBareWatcher()
			wat.Scheduler()
			return
		}
					
		if *CFM[WATCHER_PROXY].(*bool) {
			wtr,err:=watcher.InitWatcher(ac,true)
			if err!=nil {
				panic(err)
			}

			wtr.Proxy(*ac.AgtConf[AGENT_MODE].(*string))
			
			return
		}
		
		wtr,err:=watcher.InitWatcher(ac,false)
		if err!=nil {
			panic(err)
		}

		if ac.WatcherConf!=nil {
			go wtr.Scale()
			
		}
		
		_=watcher.InitWatcherD()
		wtr.Monitor()
		
		return
	}

	//deprecated
	//util.Init(*ac.AgtConf[AGENT_MODE].(*string),*ac.AgtConf[AGENT_DEBUG].(*bool))
	if *ac.AgtConf[AGENT_DEBUG].(*bool) {
		util.GetCLog().SetDebug()
	}
	
	if *ac.AgtConf[AGENT_MODE].(*string)=="api" {
		
		InitAPI(ac)

		wtr:=watcher.InitWatcherD()
		wtr.Monitor()

		return
	}

	if *ac.AgtConf[AGENT_MODE].(*string)=="oauth2" {
		
		//util.Init("oauth",*ac.AgtConf[AGENT_DEBUG].(*bool))
		InitOAuth2(ac)

		wtr:=watcher.InitWatcherD()
		wtr.Monitor()

		return
	}
	
	//CFM = ac.AgtConf
	agtOps,err:=NewAgentOps(ac)
	if err!=nil{
		panic(err)
	}

	//file ops example
	if *ac.AgtConf[FILE_TO_UPLOAD].(*string)!="" {
		for i:=0;i<10;i++ {

			f:=strings.Split(*ac.AgtConf[FILE_TO_UPLOAD].(*string),":")
			err=agtOps.FileUploadOps(f[0],f[1])
		}
	}
	//file ops example
	if *ac.AgtConf[FILE_TO_DOWNLOAD].(*string)!="" {
		for i:=0;i<10;i++ {
			f:=strings.Split(*ac.AgtConf[FILE_TO_UPLOAD].(*string),":")
			err=agtOps.FileDownloadOps(f[0],f[1])
		}
	}	

	agtOps.Start()
	agtOps.Operation()

	wtr:=watcher.InitWatcherD()
	wtr.Monitor(agtOps)
}
