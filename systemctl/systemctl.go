package systemctl

import (
	"errors"
	"fmt"
	"github.com/codeskyblue/go-sh"
	"github.com/halidaltuner/go-sysmanager/messages"
	"os"
	"strings"
)

const (
	systemctlPath = "/bin/systemctl"
	libPath = "/lib/systemd/system/"
)

func GetServiceParams(serviceName string) map[string]string {
	showCommand, err := sh.Command(systemctlPath, "show", serviceName, "--no-page").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	commandTrim := strings.TrimSpace(string(showCommand))
	splitLines := strings.Split(commandTrim, "\n")
	serviceDetails := make(map[string]string)

	for i :=0; i<len(splitLines); i++ {
		kv := strings.Split(splitLines[i], "=")
		serviceDetails[kv[0]] = kv[1]
	}
	return serviceDetails
}

func GetServiceParam(key, serviceName string)string {
	serviceDetail := GetServiceParams(serviceName)
	return serviceDetail[key]
}

func ServiceExist(serviceName string) bool{
	if _, err := os.Stat(libPath+serviceName+".service"); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	return false
}

func StartService(serviceName string)(string, error) {
	if ServiceExist(serviceName) == true {
		_, err := sh.Command(systemctlPath, "start", serviceName).Output()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(messages.ServiceStarted, serviceName), err
	} else {
		return "", errors.New(messages.ServiceDoesNotExist)
	}
}

func RestartService(serviceName string)(string, error) {
	if ServiceExist(serviceName) == true {
		_, err := sh.Command(systemctlPath, "restart", serviceName).Output()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(messages.ServiceStarted, serviceName), err
	} else {
		return "", errors.New(messages.ServiceDoesNotExist)
	}
}

func StopService(serviceName string)(string, error) {
	if ServiceExist(serviceName) == true {
		_, err := sh.Command(systemctlPath, "stop", serviceName).Output()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(messages.ServiceStarted, serviceName), err
	} else {
		return "", errors.New(messages.ServiceDoesNotExist)
	}
}