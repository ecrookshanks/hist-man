package shell

import (
	"bufio"
	"errors"
	"io"
	"os"
	"os/user"
	"runtime"
	"strings"
)

const pwd_file = "/etc/passwd"

func IsWinodwsOS() bool {

	OS := runtime.GOOS

	if OS == "darwin" || OS == "linux" {
		return false
	}

	return true

}

func GetCurrentUserDefaultShell() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	if !IsWinodwsOS() {

		f, err := os.Open(pwd_file)
		if err != nil {
			return "", err
		}

		defer f.Close()

		r := bufio.NewReader(f)

		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				return "", err
			}
			sections := strings.Split(line, ":")
			if sections[0] == usr.Username {
				shell := strings.TrimSpace(sections[6])

				return shell, nil

			}
		}
	} else {
		return "bash-windows", nil
		// return "", errors.New("no default shell exists on Windows")
	}
	// If here, then the user was not found in the pwd file.  This means they're part
	// of a corp group or other 3rd party access control.  Use default shell based on OS
	OS := runtime.GOOS
	if OS == "darwin" {
		return "zsh", nil
	}
	if OS == "linux" {
		return "bash", nil
	}
	// Should be impossible - what could other error be?
	return "", errors.New("no user name found matching the current user")
}
