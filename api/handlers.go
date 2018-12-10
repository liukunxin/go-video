package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/liukunxin/go-video/api/dbops"
	"github.com/liukunxin/go-video/api/defs"
	"github.com/liukunxin/go-video/api/session"
	"io"
	"io/ioutil"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}
	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
	}
	if err :=  dbops.AddUserCredential(ubody.Username, ubody.Pwd);err!=nil{
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	id := session.GenerateNewSessionId(ubody.Username)
	su := &defs.SignedUp{Success:true,SessionId:id}

	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	}else {
		sendNormalResponse(w, string(resp),201)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}
