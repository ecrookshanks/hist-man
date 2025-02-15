package hist

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/user"
	"slices"
	"strings"

	"github.com/ecrookshanks/hist-man/shell"
)

type HistResult struct {
	Lines      int
	Size       int64
	Unique     int
	Dups       int
	All        []string
	DupVals    []string
	UniqueVals []string
	DupCounts  map[string]int
}

const linux_bash_file = "/.bash_history"
const mac_bash_file = "/.zsh_history"

func GetBashFileStats() (*HistResult, error) {
	hr := HistResult{}

	file, err := constructCompleteFileName()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	readFileLineByLine(file, &hr)

	return &hr, nil
}

func TestHistory(fileName string) (*HistResult, error) {

	hr := HistResult{}
	readFileLineByLine(fileName, &hr)
	return &hr, nil
}

func constructCompleteFileName() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	dir := usr.HomeDir

	default_shell, err := shell.GetCurrentUserDefaultShell()
	if err != nil {
		return "", err
	}

	if strings.Contains(default_shell, "bash") {
		return dir + linux_bash_file, nil
	} else {
		return dir + mac_bash_file, nil
	}
}

func readFileLineByLine(file string, hr *HistResult) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		return err
	}
	hr.Size = info.Size()

	counts := make(map[string]int)
	lines := []string{}
	dups := []string{}
	uniques := []string{}
	r := bufio.NewReader(f)

	num_lines := 0
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Error reading file: %s\n", err)
		}
		line = strings.TrimSpace(line)

		counts[line]++
		num_lines++

		lines = append(lines, line)

		// Manage the dups and unique slices
		if counts[line] == 2 {
			dups = append(dups, line)
			// remove from uniques
			ndex := slices.Index(uniques, line)
			if ndex != -1 {
				uniques = slices.Delete(uniques, ndex, ndex+1)
			}
		}

		if counts[line] == 1 && !slices.Contains(uniques, line) {
			uniques = append(uniques, line)
		}

	}

	hr.Lines = num_lines
	hr.Unique = len(uniques)
	hr.Dups = len(dups)
	hr.All = lines
	hr.DupVals = dups
	hr.UniqueVals = uniques
	hr.DupCounts = counts

	return nil

}

func FindMaxDupValueAndName(dups map[string]int) (string, int) {
	max := 0
	maxKey := ""
	for k, v := range dups {
		if v > max {
			max = v
			maxKey = k
		}
	}

	return maxKey, max
}
