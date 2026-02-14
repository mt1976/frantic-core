package netHandler

import (
	"os/exec"
	"strings"

	"github.com/mt1976/frantic-core/application"
	"github.com/mt1976/frantic-core/logHandler"
)

var sucess = "Host %v is reachable"
var failure = "Host %v is not reachable"

func CheckHostAvailability(addr string) (bool, error) {
	//
	// This is a simple ping function that will return true if the host is up and false if it is down
	//
	// addr: string: The IP address of the host to ping
	// windows: bool: If the host is running windows, then we need to use a different method to ping it
	logHandler.Service.Printf("Pinging %v", addr)
	var out []byte
	var err error

	if application.IsRunningOnWindows() {
		logHandler.Service.Printf("Running Windows Ping - [ping %v -n 5 -w 3000]", addr)
		out, err = exec.Command("ping", addr, "-n", "5", "-w", "3000").Output()
	} else {
		logHandler.Service.Printf("Running Linux Ping - [ping %v -c 5 -i 3 -W 10]", addr)
		out, err = exec.Command("ping", addr, "-c 5", "-i 3", "-W 10").Output()
	}
	if err != nil {
		if strings.Contains(err.Error(), "exit status 68") {
			logHandler.Warning.Printf(failure, addr)
			return false, nil
		}
	}

	if isHostReachable(out) {
		logHandler.Service.Printf(sucess, addr)
		return true, nil
	}

	logHandler.Warning.Printf(failure, addr)
	return false, nil
}

func isHostReachable(out []byte) bool {

	switch {
	case strings.Contains(string(out), "Destination Host Unreachable"):
		return false
	case strings.Contains(string(out), "Request timed out"):
		return false
	case strings.Contains(string(out), "100% packet loss"):
		return false
	case strings.Contains(string(out), "Request timeout"):
		return false
	case strings.Contains(string(out), "cannot resolve"):
		return false
	case strings.Contains(string(out), "unknown host"):
		return false
	case strings.Contains(string(out), "Name or service not known"):
		return false
	case strings.Contains(string(out), "Temporary failure in name resolution"):
		return false
	case strings.Contains(string(out), "Ping request could not find host"):
		return false
	}
	return true
}
