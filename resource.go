package resource

import (
	_ "embed"
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/reutils"
	"strings"
)

//go:embed VERSION
var Version string

var SemanticVersioning string

//go:embed LICENSE
var License string

//go:embed REEPORT
var Report string

//go:embed NAME
var Name string

func init() {
	if Name == "" {
		Name = "Backend-Server"
	}

	Version = strings.ToLower(Version)
	ver := strings.TrimPrefix(Version, "v")
	if !reutils.IsSemanticVersion(ver) {
		panic(fmt.Sprintf("%s is not a semantic versioning.", Version))
	} else {
		SemanticVersioning = ver
	}
}
