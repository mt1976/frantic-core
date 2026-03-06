package commonConfig

import (
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/mt1976/frantic-core/paths"
)

// var Data ConfigurationModel
var (
	name               = "common"
	filename           = ""
	commonSettingsFile = "common"
)

var (
	TRUE  = "true"
	FALSE = "false"
)

func Get() *Settings {
	var thisConfig Settings
	filename = paths.Application().String() + paths.Config().String() + string(os.PathSeparator) + commonSettingsFile + ".toml"
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("[%v] Error=[%v]", strings.ToUpper(name), err.Error())
		panic(err)
	}

	err = toml.Unmarshal(content, &thisConfig)
	if err != nil {
		panic(err)
	}

	return &thisConfig
}

func (s *Settings) Spew() {
	nm := strings.ToUpper(name)
	fmt.Printf("[%v] Config loaded from file: %v\n", nm, filename)
	fmt.Printf("[%v][APPLICATION]\n", nm)
	fmt.Printf("[%v] Name: %+v\n", nm, s)
}

func isTrueFalse(s string) bool {
	// We only disable the logging if the value is "true"/"t" or "yes"/"y"

	if s == "" {
		return false
	}

	logTrue := "true"
	if strings.EqualFold(s[:1], "y") {
		logTrue = "yes"
	}

	if strings.EqualFold(s, logTrue[:1]) {
		return true
	}
	if strings.EqualFold(s, logTrue) {
		return true
	}

	return false
}

var importedVersionInfo string

func init() {
	importedVersionInfo = ""
	// Import the db version from the version.no file stored in the root of the project
	// This is used to ensure that the database version is always in sync with the application version, and to prevent accidental changes to the database version in the common.toml file
	versionFile := paths.Application().String() + string(os.PathSeparator) + "version.no"
	content, err := os.ReadFile(versionFile)
	if err != nil {
		fmt.Printf("[%v] Error reading database version from file: %v\n", strings.ToUpper(name), err.Error())
	} else {
		importedVersionInfo = strings.TrimSpace(string(content))
		fmt.Printf("[%v] Imported database version from file: %v\n", strings.ToUpper(name), importedVersionInfo)
	}
	if importedVersionInfo == "" {
		fmt.Printf("[%v] No database version imported from file, using default from common.toml: %v\n", strings.ToUpper(name), importedVersionInfo)
		importedVersionInfo = Get().Application.Version
	}
}
