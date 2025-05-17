// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package logformat

import (
	"encoding/json"
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"strings"
)

type FormatFunc func(data *LogData) string

func FormatJson(data *LogData) string {
	d, err := json.Marshal(data)
	if err != nil {
		return fmt.Sprintf("{\"errmgs\":\"%s\"}", err.Error())
	}

	return strings.TrimRight(string(d), "\n")
}

func FormatFile(data *LogData) string {
	var res = new(strings.Builder)

	res.WriteString(fmt.Sprintf("%s %s | %s | unix-timestamp=\"%ds\" | exec=\"%s\"  | app=\"%s\" | version=\"%s\"", data.Date, data.Zone, data.Level, data.Timestamp, data.Exec, data.Name, data.Version))

	if data.User != nil {
		res.WriteString(fmt.Sprintf(" | uid=\"%s\" | gid=\"%s\" | user=\"%s\"", data.Uid, data.Gid, data.UserName))
	} else {
		res.WriteString(" | uid=without | gid=without | user=without")
	}

	if data.WorkDirectory != "" {
		res.WriteString(fmt.Sprintf(" | work-directory=\"%s\"", data.WorkDirectory))
	} else {
		res.WriteString(" | work-directory=without")
	}

	res.WriteString(fmt.Sprintf(" | pid=\"%d\" | ppid=\"%d\"", data.Pid, data.ParentPid))

	res.WriteString(fmt.Sprintf(" | msg=\"%s\"", data.Msg))

	return res.String()
}

func FormatConsole(data *LogData) string {
	return fmt.Sprintf("%s %s | %s | %s | pid=%d | msg=\"%s\"", data.Date, data.Zone, data.Level, data.Name, data.Pid, data.Msg)
}

func FormatConsolePretty(data *LogData) string {
	if data.Level == loglevel.LevelWarn || data.Level == loglevel.LevelInfo || data.Level == loglevel.LevelDebug {
		return fmt.Sprintf("\u001B[1;3;37;44m %s %s \u001B[0m|\033[1;37;42m %s \033[0m| \u001B[1;3;4m%s\u001B[0m | \u001B[1;3;4m%d\u001B[0m |\u001B[1;7m %s \u001B[0m", data.Date, data.Zone, data.Level, data.Name, data.Pid, data.Msg)
	} else if data.Level == loglevel.PseudoLevelTag || data.Level == loglevel.LevelError || data.Level == loglevel.LevelPanic {
		return fmt.Sprintf("\u001B[1;3;37;44m %s %s \u001B[0m|\033[1;37;41m %s \033[0m| \u001B[1;3;4m%s\u001B[0m | \u001B[1;3;4m%d\u001B[0m |\u001B[1;7m %s \u001B[0m", data.Date, data.Zone, data.Level, data.Name, data.Pid, data.Msg)
	}
	panic(fmt.Sprintf("unknown loglevel: %s", data.Level))
}
