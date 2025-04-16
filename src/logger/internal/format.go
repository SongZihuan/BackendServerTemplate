// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package internal

import (
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/global"
	"github.com/SongZihuan/BackendServerTemplate/src/logger/loglevel"
	"os"
	"os/user"
	"strings"
	"time"
)

func (l *Logger) format(_level loglevel.LoggerLevel, msg string) string {
	var res = new(strings.Builder)

	level := string(_level)

	now := time.Now().In(global.Location)
	zone := global.Location.String()
	if strings.ToLower(zone) == "local" {
		zone, _ = now.Zone()
	}
	date := now.Format(time.DateTime)
	msg = strings.Replace(msg, "\"", "'", -1)
	level = strings.ToUpper(level)

	res.WriteString(fmt.Sprintf("%s %s | %s | unix-timestamp=\"%ds\" | app=\"%s\" | version=\"%s\"", date, zone, level, now.Unix(), global.Name, global.Version))

	u := getUser()
	if u != nil {
		res.WriteString(fmt.Sprintf(" | uid=\"%s\" | gid=\"%s\" | user=\"%s\"", u.Uid, u.Gid, u.Name))
	} else {
		res.WriteString(" | uid=without | gid=without | user=without")
	}

	wd := getWorkDir()
	if wd != "" {
		res.WriteString(fmt.Sprintf(" | work-directory=\"%s\"", wd))
	} else {
		res.WriteString(" | work-directory=without")
	}

	res.WriteString(fmt.Sprintf(" | msg=\"%s\"\n", msg))

	return res.String()
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
