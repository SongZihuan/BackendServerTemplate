// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package logformat

import (
	"github.com/SongZihuan/BackendServerTemplate/src/global"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"github.com/SongZihuan/BackendServerTemplate/utils/osutils"
	"os"
	"os/user"
	"strings"
	"time"
)

type LogData struct {
	Level         loglevel.LoggerLevel `json:"level"`
	Now           time.Time            `json:"-"`
	Date          string               `json:"date"`
	Zone          string               `json:"zone"`
	Timestamp     int64                `json:"timestamp"`
	Exec          string               `json:"exec"`
	Name          string               `json:"name"`
	Version       string               `json:"version"`
	User          *user.User           `json:"-"`
	Uid           string               `json:"uid"`
	Gid           string               `json:"gid"`
	UserName      string               `json:"username"`
	WorkDirectory string               `json:"work-directory"`
	Pid           int                  `json:"pid"`
	ParentPid     int                  `json:"ppid"`
	Msg           string               `json:"msg"`
}

func GetLogData(level loglevel.LoggerLevel, msg string, now time.Time) *LogData {
	var res = new(LogData)

	now = now.In(global.Location)

	res.Level = level
	res.Now = now
	res.Date = now.Format(time.DateTime)
	res.Zone = global.Location.String()
	res.Timestamp = now.Unix()
	res.Name = global.Name
	res.Version = global.Version

	u := getUser()
	if u != nil {
		res.User = u
		res.Uid = u.Uid
		res.Gid = u.Gid
		res.UserName = getUserName(u)
	} else {
		res.User = nil
		res.Uid = ""
		res.Gid = ""
		res.UserName = ""
	}

	wd := getWorkDir()
	if wd != "" {
		res.WorkDirectory = wd
	} else {
		res.WorkDirectory = ""
	}

	res.Pid = os.Getpid()
	res.ParentPid = os.Getppid()
	res.Exec = osutils.GetArgs0Name()

	res.Msg = strings.Replace(msg, "\"", "'", -1)

	return res
}

func getUserName(u *user.User) (name string) {
	if name = u.Name; name != "" {
		return name
	}

	if name = u.Username; name != "" {
		return name
	}

	return "-"
}

func getUser() *user.User {
	currentUser, err := user.Current()
	if err != nil {
		return nil
	}

	return currentUser
}

func getWorkDir() string {
	dir, err := os.Getwd()
	if err != nil {
		return "未知"
	}

	return dir
}
