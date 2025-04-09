package resource

import (
	_ "embed"
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/reutils"
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

func init() {
	initCleanFile()
	initName()
	initBuildDate()
	initVersion()
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
	if Name == "" {
		Name = "Backend-Server"
	}
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
	SemanticVersioning = strings.TrimPrefix(strings.ToLower(version), "v")
	if SemanticVersioning == "" {
		SemanticVersioning = strings.TrimPrefix(strings.ToLower(GitTag), "v")
		if SemanticVersioning == "" {
			if GitCommitHash != "" {
				SemanticVersioning = fmt.Sprintf("0.0.0+dev-%d-%s", BuildTime.Unix(), GitCommitHash)
			} else {
				SemanticVersioning = fmt.Sprintf("0.0.0+dev-%d", BuildTime.Unix())
			}
			Version = "v" + SemanticVersioning
		} else if reutils.IsSemanticVersion(SemanticVersioning) {
			if GitCommitHash != "" && GitTagCommitHash != "" && GitCommitHash != GitTagCommitHash && !strings.Contains(SemanticVersioning, "dev") {
				SemanticVersioning = SemanticVersioning + fmt.Sprintf("+dev-%s", GitTagCommitHash)
			} else if strings.HasPrefix(SemanticVersioning, "0.") {
				SemanticVersioning = SemanticVersioning + "-dev"
			}
			Version = "v" + SemanticVersioning
		} else {
			panic(fmt.Sprintf("%s is not a semantic versioning.", SemanticVersioning))
		}
	} else if reutils.IsSemanticVersion(SemanticVersioning) {
		Version = "v" + SemanticVersioning
	} else {
		panic(fmt.Sprintf("%s is not a semantic versioning.", SemanticVersioning))
	}
}

func utilsClenFileData(data string) (res string) {
	res = utilsCheckAndRemoveBOM(data)
	res = strings.Replace(res, "\r", "", -1)
	res = strings.Split(res, "\n")[0]
	res = strings.TrimSpace(res)
	return res
}

func utilsClenFileDataMoreLine(data string) (res string) {
	res = utilsCheckAndRemoveBOM(data)
	res = strings.Replace(res, "\r", "", -1)
	return res
}

func utilsCheckAndRemoveBOM(s string) string {
	// UTF-8 BOM 的字节序列为 0xEF, 0xBB, 0xBF
	bom := []byte{0xEF, 0xBB, 0xBF}

	// 将字符串转换为字节切片
	bytes := []byte(s)

	// 检查前三个字节是否是 BOM
	if len(bytes) >= 3 && bytes[0] == bom[0] && bytes[1] == bom[1] && bytes[2] == bom[2] {
		// 如果存在 BOM，则删除它
		return string(bytes[3:])
	}

	// 如果不存在 BOM，则返回原始字符串
	return s
}
