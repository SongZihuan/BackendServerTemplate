// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package consoleutils

const (
	CodePageUTF8 uint = 65001
	CodePageGBK  uint = 936
)

type Event interface {
	String() string
	ConsoleEvent()
	GetCode() uint
}

type EventData struct {
	Name string
	Code uint
}

func (e *EventData) String() string {
	return e.Name
}

func (e *EventData) GetCode() uint {
	return e.Code
}

func (*EventData) ConsoleEvent() {}

// 定义控制台事件类型
//
//goland:noinspection GoSnakeCaseUsage
var (
	CTRL_C_EVENT Event = &EventData{
		Name: "CTRL_C_EVENT",
		Code: 0,
	} // ctrl+c

	CTRL_BREAK_EVENT Event = &EventData{
		Name: "CTRL_BREAK_EVENT",
		Code: 1,
	} // ctrl+break

	CTRL_CLOSE_EVENT Event = &EventData{
		Name: "CTRL_CLOSE_EVENT",
		Code: 2,
	} // console关闭

	CTRL_LOGOFF_EVENT Event = &EventData{
		Name: "CTRL_LOGOFF_EVENT",
		Code: 5,
	} // 用户注销

	CTRL_SHUTDOWN_EVENT Event = &EventData{
		Name: "CTRL_SHUTDOWN_EVENT",
		Code: 6,
	} // 系统关机
)

var EventMap = map[uint]Event{
	0: CTRL_C_EVENT,
	1: CTRL_BREAK_EVENT,
	2: CTRL_CLOSE_EVENT,
	5: CTRL_LOGOFF_EVENT,
	6: CTRL_SHUTDOWN_EVENT,
}
