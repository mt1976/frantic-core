package application

import (
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"unicode"

	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dockerHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/rivo/uniseg"
)

var name = "Application"

const (
	WINDOWS = "Windows"
	NIX     = "*nix"
)

func OS() string {
	if IsRunningOnWindows() {
		return WINDOWS
	} else {
		return NIX
	}
}

// Deprecated: Use dockerHelpers.IsDockerContainer()
func RunningInDockerContainer() bool {
	return dockerHelpers.IsDockerContainer()
}

func IsRunningOnWindows() bool {
	return runtime.GOOS == "windows"
}

func HostName() string {
	if IsRunningOnWindows() {
		rtn := hostname_windows()
		strings.Replace(rtn, "\n", "", -1)
		strings.Replace(rtn, "\r", "", -1)
		return rtn
	}
	hn, err := os.Hostname()
	if err != nil {
		logHandler.ErrorLogger.Printf("[%v] Error=[%v]", strings.ToUpper(name), err.Error())
		panic(commonErrors.ErrOSWrapper(err))
	}
	return strings.ToLower(hn)
}

func hostname_windows() string {

	cmd := exec.Command("hostname")

	hostname, err := cmd.Output()

	if err != nil {
		logHandler.ErrorLogger.Printf("[%v] Error=[%v]", strings.ToUpper(name), err.Error())
		panic(commonErrors.ErrOSWrapper(err))
	}
	rtn := string(hostname)
	rtn = strings.ToLower(strings.TrimSuffix(rtn, "\n"))
	rtn = strings.ToLower(strings.TrimSuffix(rtn, "\r"))

	return rtn
}

func HostIP() string {

	netInterfaceAddresses, err := net.InterfaceAddrs()

	if err != nil {
		return ""
	}

	for _, netInterfaceAddress := range netInterfaceAddresses {

		networkIp, ok := netInterfaceAddress.(*net.IPNet)

		if ok && !networkIp.IP.IsLoopback() && networkIp.IP.To4() != nil {

			ip := networkIp.IP.String()

			//	fmt.Println("Resolved Host IP: " + ip)

			return ip
		}
	}
	return ""

}

func get_IP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		logHandler.ErrorLogger.Printf("[%v] Error=[%v]", strings.ToUpper(name), err.Error())
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", nil
}

func SystemIdentity() string {
	id := strings.ToLower(HostName())
	return cleanID(id)
}

func cleanID(id string) string {
	id = strings.Replace(id, ".local", "", -1)
	r := strings.NewReplacer("\n", "", "\r", "", "\t", "")
	id = r.Replace(id)
	id = stripSpecial(id)
	return id
}

func stripSpecial(s string) string {
	var b strings.Builder
	gr := uniseg.NewGraphemes(s)
	for gr.Next() {
		r := gr.Runes()[0]
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteString(gr.Str())
		}
	}
	return b.String()
}
