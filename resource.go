package resource

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

//go:embed VERSION
var version string
var Version string
var SemanticVersioning string

//go:embed LICENSE
var License string

//go:embed REPORT
var Report string

//go:embed NAME
var Name string

//go:embed build_date.txt
var buildDateTxt string

var BuildTime time.Time

//go:embed commit_data.txt
var GitCommitHash string

//go:embed tag_data.txt
var GitTag string

//go:embed tag_commit_data.txt
var GitTagCommitHash string

//go:embed random_data.txt
var randomData string

func init() {
	initCleanFile() // 最先执行

	// 其他操作
	initName()
	initBuildDate()
	initVersion()
	initServiceConfig()
}

func initCleanFile() {
	License = utilsClenFileDataMoreLine(License)
	Report = utilsClenFileDataMoreLine(Report)

	version = utilsClenFileData(version)
	Name = utilsClenFileData(Name)
	buildDateTxt = utilsClenFileData(buildDateTxt)
	GitCommitHash = utilsClenFileData(GitCommitHash)
	GitTag = utilsClenFileData(GitTag)
	GitTagCommitHash = utilsClenFileData(GitTagCommitHash)
}

func initName() {
	if Name != "" {
		return
	}

	_args0, err := os.Executable()
	if err != nil {
		if len(os.Args) > 0 {
			_args0 = os.Args[0]
		} else {
			panic(fmt.Sprintf("name was empty: %s", err.Error()))
		}
	}

	if _args0 == "" {
		panic("name was empty")
	}

	_args0Name := filepath.Base(_args0)

	if _args0Name == "" {
		panic("name was empty")
	}

	Name = _args0Name
}

func initBuildDate() {
	if buildDateTxt == "" {
		BuildTime = time.Now()
		return
	}

	res, err := strconv.ParseInt(buildDateTxt, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("get build timestamp error: %s", err.Error()))
	}
	BuildTime = time.Unix(res, 0)
}

func initVersion() {
	ver := getDefaultVersion()
	if ver != "" {
		SemanticVersioning = ver
		Version = "v" + SemanticVersioning
		return
	}

	ver = getGitTagVersion()
	if ver != "" {
		SemanticVersioning = ver
		Version = "v" + SemanticVersioning
		return
	}

	ver = getRandomVersion()
	if ver != "" {
		SemanticVersioning = ver
		Version = "v" + SemanticVersioning
		return
	}

	panic("Get Version Failed")
}

func getDefaultVersion() (defVer string) {
	defVer = strings.TrimPrefix(strings.ToLower(version), "v")
	if defVer == "" || !utilsIsSemanticVersion(defVer) {
		return ""
	}
	return defVer
}

func getGitTagVersion() (gitVer string) {
	gitVer = strings.TrimPrefix(strings.ToLower(GitTag), "v")
	if GitCommitHash != "" && (GitTagCommitHash == "" || gitVer == "") {
		return fmt.Sprintf("0.0.0+dev.%d.%s", BuildTime.Unix(), GitCommitHash)
	} else if GitCommitHash != "" && GitTagCommitHash != "" && gitVer != "" && utilsIsSemanticVersion(gitVer) {
		if (GitCommitHash != GitTagCommitHash || strings.HasPrefix(gitVer, "0.")) && !strings.Contains(gitVer, "dev") {
			return gitVer + fmt.Sprintf("+dev.%d.%s", BuildTime.Unix(), GitCommitHash)
		}
		return gitVer
	} else {
		return ""
	}
}

func getRandomVersion() (randVer string) {
	return fmt.Sprintf("0.0.0+dev.%d.%s", BuildTime.Unix(), randomData)
}
