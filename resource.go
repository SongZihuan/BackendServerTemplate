package resource

import (
	_ "embed"
	"github.com/SongZihuan/BackendServerTemplate/utils/cleanstringutils"
	"sync"
)

//go:embed VERSION
var Version string

//go:embed LICENSE
var License string

//go:embed REPORT
var Report string

var once sync.Once

func InitVar() {
	once.Do(func() {
		License = cleanstringutils.GetString(cleanstringutils.CheckAndRemoveBOM(License))
		Report = cleanstringutils.GetString(cleanstringutils.CheckAndRemoveBOM(Report))
		Version = cleanstringutils.GetString(cleanstringutils.CheckAndRemoveBOM(Version))
	})
}
