package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

func main() {
	// 创建HTTP服务器

	http.Handle("/new/health", http.HandlerFunc(Apihealth))
	http.Handle("/new/up", http.HandlerFunc(UpHandler))
	http.Handle("/new/down", http.HandlerFunc(DownHandler))
	http.Handle("/new/logout", http.HandlerFunc(LogoutHandler))
	http.Handle("/new/ping", http.HandlerFunc(PingHandleder))
	http.Handle("/new/status", http.HandlerFunc(StatusHandleder))
	result := make(map[string]interface{}, 1)
	result["status"] = true
	result["msg"] = "suc"
	marshal, _ := json.Marshal(result)
	fmt.Printf(string(marshal))
	// 启动HTTP服务器
	err := http.ListenAndServe(":65427", nil)
	if err != nil {
		panic(err)
	}
}

type BodyData struct {
	Path string `json:"path"`
}
type ResData struct {
	Data string `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func Apihealth(w http.ResponseWriter, r *http.Request) {

	resData := ResData{
		Data: "",
		Code: 0,
		Msg:  "ok",
	}
	json.NewEncoder(w).Encode(resData)
	return
}

func HandlerCheck(w http.ResponseWriter, r *http.Request) (bool, string) {

	resData := ResData{
		Data: "",
		Code: 0,
		Msg:  "",
	}
	if r.Method != "POST" {
		resData.Code = 400
		resData.Msg = "bad request"
		json.NewEncoder(w).Encode(resData)
		return false, ""
	}
	all, err := io.ReadAll(r.Body)
	if err != nil {
		resData.Code = 400
		resData.Msg = err.Error()
		json.NewEncoder(w).Encode(resData)
		return false, ""
	}
	var bodyData BodyData
	err = json.Unmarshal(all, &bodyData)
	if err != nil {
		resData.Code = 400
		resData.Msg = err.Error()
		json.NewEncoder(w).Encode(resData)
		return false, ""
	}
	return true, bodyData.Path
}
func UpHandler(w http.ResponseWriter, r *http.Request) {
	resData := ResData{
		Data: "",
		Code: 0,
		Msg:  "",
	}
	check, path := HandlerCheck(w, r)
	if check != true {
		return
	}
	//cmd := exec.Command(path, "up", "--register-server=http://solopace-controler.doaction.tech:8080", "--accept-routes=true", "--accept-dns=false", "--unattended=true")
	cmd := exec.Command(path, "up", "--register-server=https://solopace-controler.doaction.pro", "--accept-routes=true", "--accept-dns=false", "--unattended=true")
	stdout, err := cmd.Output()
	if err != nil {
		resData.Code = 400
		resData.Msg = err.Error()
		json.NewEncoder(w).Encode(resData)
		return
	}
	var res ResData
	err = json.Unmarshal(stdout, &res)
	if err != nil {
		resData.Code = 400
		resData.Msg = err.Error()
		json.NewEncoder(w).Encode(resData)
		return
	}
	json.NewEncoder(w).Encode(res)
}
func DownHandler(w http.ResponseWriter, r *http.Request) {
	resData := ResData{
		Data: "",
		Code: 0,
		Msg:  "",
	}
	check, path := HandlerCheck(w, r)
	if check != true {
		return
	}
	cmd := exec.Command(path, "down")
	stdout, err := cmd.Output()
	if err != nil {
		resData.Code = 400
		resData.Msg = err.Error()
		json.NewEncoder(w).Encode(resData)
		return
	}

	var res ResData
	err = json.Unmarshal(stdout, &res)
	if err != nil {
		resData.Code = 400
		resData.Msg = err.Error()
		json.NewEncoder(w).Encode(resData)
		return
	}
	json.NewEncoder(w).Encode(res)
}
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	resData := ResData{
		Data: "",
		Code: 0,
		Msg:  "",
	}
	check, path := HandlerCheck(w, r)
	if check != true {
		return
	}
	cmd := exec.Command(path, "logout")
	stdout, err := cmd.Output()
	if err != nil {
		resData.Code = 400
		resData.Msg = err.Error()
		json.NewEncoder(w).Encode(resData)
		return
	}

	var res ResData
	err = json.Unmarshal(stdout, &res)
	if err != nil {
		resData.Code = 400
		resData.Msg = err.Error()
		json.NewEncoder(w).Encode(resData)
		return
	}
	json.NewEncoder(w).Encode(res)
}
func PingHandleder(w http.ResponseWriter, r *http.Request) {
	resData := ResData{
		Data: "",
		Code: 0,
		Msg:  "",
	}
	check, path := HandlerCheck(w, r)
	if check != true {
		return
	}
	cmd := exec.Command(path, "ping")
	stdout, err := cmd.Output()
	if err != nil {
		resData.Code = 400
		resData.Msg = err.Error()
		json.NewEncoder(w).Encode(resData)
		return
	}

	var res ResData
	err = json.Unmarshal(stdout, &res)
	if err != nil {
		resData.Code = 400
		resData.Msg = err.Error()
		json.NewEncoder(w).Encode(resData)
		return
	}
	json.NewEncoder(w).Encode(res)
}
func StatusHandleder(w http.ResponseWriter, r *http.Request) {
	resData := ResData{
		Data: "",
		Code: 0,
		Msg:  "",
	}
	check, path := HandlerCheck(w, r)
	if check != true {
		return
	}
	cmd := exec.Command(path, "status", "--gem=true")
	stdout, err := cmd.Output()
	fmt.Printf(string(stdout))
	if err != nil {
		resData.Code = 400
		resData.Msg = err.Error()
		json.NewEncoder(w).Encode(resData)
		return
	}
	var res ResData
	err = json.Unmarshal(stdout, &res)
	if err != nil {
		resData.Code = 400
		resData.Msg = err.Error()
		json.NewEncoder(w).Encode(resData)
		return
	}
	json.NewEncoder(w).Encode(res)
}
