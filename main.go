package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/getlantern/elevate"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

var args struct {
	Path string
}

func main() {
	flag.StringVar(&args.Path, "path", "", "Please enter the file path of the executable")
	flag.Parse()
	result := make(map[string]interface{}, 2)

	if strings.Contains(args.Path, "solopacedoored.exe") {
		cmd := elevate.Command(args.Path, "install-system-daemon")
		// 开始运行
		err := cmd.Run()
		if err != nil {
			result["status"] = false
			result["msg"] = err.Error()
			marshal, _ := json.Marshal(result)
			fmt.Printf(string(marshal))
		} else {
			result["status"] = true
			result["msg"] = "suc"
			marshal, _ := json.Marshal(result)
			fmt.Printf(string(marshal))
		}

	} else {
		cmd := exec.Command(args.Path)
		attr := &syscall.SysProcAttr{
			HideWindow:    true,
			CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP | 0x00000010,
		}

		cmd.SysProcAttr = attr

		if err := cmd.Start(); err != nil {
			result["status"] = false
			result["msg"] = err.Error()
			marshal, _ := json.Marshal(result)
			fmt.Printf(string(marshal))
		}

		process, err := os.FindProcess(cmd.Process.Pid)
		if err != nil {
			result["status"] = false
			result["msg"] = err.Error()
			marshal, _ := json.Marshal(result)
			fmt.Printf(string(marshal))
		}

		if err = process.Release(); err != nil {
			result["status"] = false
			result["msg"] = err.Error()
			marshal, _ := json.Marshal(result)
			fmt.Printf(string(marshal))
		}
		result["status"] = true
		result["msg"] = "suc"
		marshal, _ := json.Marshal(result)
		fmt.Printf(string(marshal))
	}
	os.Exit(0)
}
