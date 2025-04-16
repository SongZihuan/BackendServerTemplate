// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package logpanic

import "time"

type PanicData struct {
	time time.Time
	msg  string
}

func Panic(t time.Time, msg string) *PanicData {
	panic(NewPanicData(t, msg))
}

func NewPanicData(t time.Time, msg string) *PanicData {
	return &PanicData{
		time: t,
		msg:  msg,
	}
}

func (p *PanicData) Time() time.Time {
	return p.time
}

func (p *PanicData) Msg() string {
	return p.msg
}
