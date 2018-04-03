//package tools provides tools required for moviems to create directories, create/write files, checking on environment
//variables and a standard json output for api handlers
package tools

import (
	"fmt"
	"strings"
	"os"
)

// EnvironmentManager is used to validate missing environment variables
type EnvironmentManager struct {
	Logs []string
}

// EnvironmentManager.Validate will check if tag exist from environment.
// if not found, tag is added to list as a record
func (self *EnvironmentManager) Validate(tag string) {
	if os.Getenv(tag) == "" {
		self.Logs = append(self.Logs, fmt.Sprint("environment variable required: ", tag))
	}
}

// EnvironmentManager.Result will check if list is empty. If not empty, a panic will occurred and
// output the list of missing environment variables required.
func (self *EnvironmentManager) Result() {
	if len(self.Logs) != 0 {
		panic(strings.Join(self.Logs, "\n"))
	}
}

//create new instance of EnvironmentManager
func NewEnvironmentManager() *EnvironmentManager {
	manager := EnvironmentManager{}
	return &manager
}
