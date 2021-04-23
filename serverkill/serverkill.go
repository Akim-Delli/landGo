package serverkill

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func KillServer(pidFile string) error {
	content, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return errors.Wrap(err, "Can't open pid file (is server running?)")
	}

	if err := os.Remove(pidFile); err != nil {
		log.Printf("warning: can't remove pid File - %s", err)
	}

	strPID := strings.TrimSpace(string(content))
	pid, err := strconv.Atoi(strPID)
	if err != nil {
		return errors.Wrap(err, "bad process ID")

	}

	fmt.Printf("killing server with pid=%d\n", pid)

	return nil
}
