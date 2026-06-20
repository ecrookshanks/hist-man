package shell

import (
	"os"
	"runtime"
	"strings"
	"testing"
)

const pwd_file = "/etc/passwd"

func Test_openShellFile(t *testing.T) {

	OS := runtime.GOOS
	if OS == "darwin" || OS == "linux" {

		f, err := os.Open(pwd_file)
		if err != nil {
			t.Error("Count not open passwd file.")
		}

		defer f.Close()
		info, err := f.Stat()
		if err != nil {
			t.Error("Cound not get password file stats!")
		}
		size := info.Size()
		if size <= 0 {
			t.Errorf("passswd file is zero size!!")
		}
	}
}

func Test_findShellValueInPasswdLine(t *testing.T) {
	sim_etc_passwd := "ecrooksh:x:1000:1000:,,,:/home/ecrooksh:/bin/bash"

	sections := strings.Split(sim_etc_passwd, ":")

	if sections[6] != "/bin/bash" {
		t.Error("Cound not fid the correct shell entry!")
	}

}
