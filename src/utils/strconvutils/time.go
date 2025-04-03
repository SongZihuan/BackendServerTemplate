// Copyright 2025 BackendServerTemplate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strconvutils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ReadTimeDuration(str string) time.Duration {
	if str == "forever" || str == "none" {
		return -1
	}

	if strings.HasSuffix(strings.ToUpper(str), "Y") {
		numStr := str[:len(str)-1]
		num, _ := strconv.ParseUint(numStr, 10, 64)
		return time.Hour * 24 * 365 * time.Duration(num)
	} else if strings.HasSuffix(strings.ToLower(str), "year") {
		numStr := str[:len(str)-4]
		num, _ := strconv.ParseUint(numStr, 10, 64)
		return time.Hour * 24 * 365 * time.Duration(num)
	}

	if strings.HasSuffix(strings.ToUpper(str), "M") {
		numStr := str[:len(str)-1]
		num, _ := strconv.ParseUint(numStr, 10, 64)
		return time.Hour * 24 * 31 * time.Duration(num)
	} else if strings.HasSuffix(strings.ToLower(str), "month") {
		numStr := str[:len(str)-5]
		num, _ := strconv.ParseUint(numStr, 10, 64)
		return time.Hour * 24 * 31 * time.Duration(num)
	}

	if strings.HasSuffix(strings.ToUpper(str), "W") {
		numStr := str[:len(str)-1]
		num, _ := strconv.ParseUint(numStr, 10, 64)
		return time.Hour * 24 * 7 * time.Duration(num)
	} else if strings.HasSuffix(strings.ToLower(str), "week") {
		numStr := str[:len(str)-4]
		num, _ := strconv.ParseUint(numStr, 10, 64)
		return time.Hour * 24 * 7 * time.Duration(num)
	}

	if strings.HasSuffix(strings.ToUpper(str), "D") {
		numStr := str[:len(str)-1]
		num, _ := strconv.ParseUint(numStr, 10, 64)
		return time.Hour * 24 * time.Duration(num)
	} else if strings.HasSuffix(strings.ToLower(str), "day") {
		numStr := str[:len(str)-3]
		num, _ := strconv.ParseUint(numStr, 10, 64)
		return time.Hour * 24 * time.Duration(num)
	}

	if strings.HasSuffix(strings.ToUpper(str), "H") {
		numStr := str[:len(str)-1]
		num, _ := strconv.ParseUint(numStr, 10, 64)
		return time.Hour * time.Duration(num)
	} else if strings.HasSuffix(strings.ToLower(str), "hour") {
		numStr := str[:len(str)-4]
		num, _ := strconv.ParseUint(numStr, 10, 64)
		return time.Hour * time.Duration(num)
	}

	if strings.HasSuffix(strings.ToUpper(str), "Min") { // 不能用M，否则会和 Month 冲突
		numStr := str[:len(str)-3]
		num, _ := strconv.ParseUint(numStr, 10, 64)
		return time.Minute * time.Duration(num)
	} else if strings.HasSuffix(strings.ToLower(str), "minute") {
		numStr := str[:len(str)-6]
		num, _ := strconv.ParseUint(numStr, 10, 64)
		return time.Minute * time.Duration(num)
	}

	if strings.HasSuffix(strings.ToUpper(str), "S") {
		numStr := str[:len(str)-1]
		num, _ := strconv.ParseUint(numStr, 10, 64)
		return time.Second * time.Duration(num)
	} else if strings.HasSuffix(strings.ToLower(str), "second") {
		numStr := str[:len(str)-6]
		num, _ := strconv.ParseUint(numStr, 10, 64)
		return time.Second * time.Duration(num)
	}

	if strings.HasSuffix(strings.ToUpper(str), "MS") {
		numStr := str[:len(str)-2]
		num, _ := strconv.ParseUint(numStr, 10, 64)
		return time.Millisecond * time.Duration(num)
	} else if strings.HasSuffix(strings.ToLower(str), "millisecond") {
		numStr := str[:len(str)-11]
		num, _ := strconv.ParseUint(numStr, 10, 64)
		return time.Millisecond * time.Duration(num)
	}

	if strings.HasSuffix(strings.ToUpper(str), "MiS") { // 不能用 MS , 否则会和 millisecond 冲突
		numStr := str[:len(str)-3]
		num, _ := strconv.ParseUint(numStr, 10, 64)
		return time.Microsecond * time.Duration(num)
	} else if strings.HasSuffix(strings.ToUpper(str), "MicroS") {
		numStr := str[:len(str)-6]
		num, _ := strconv.ParseUint(numStr, 10, 64)
		return time.Microsecond * time.Duration(num)
	} else if strings.HasSuffix(strings.ToLower(str), "microsecond") {
		numStr := str[:len(str)-11]
		num, _ := strconv.ParseUint(numStr, 10, 64)
		return time.Microsecond * time.Duration(num)
	}

	if strings.HasSuffix(strings.ToUpper(str), "NS") {
		numStr := str[:len(str)-2]
		num, _ := strconv.ParseUint(numStr, 10, 64)
		return time.Nanosecond * time.Duration(num)
	} else if strings.HasSuffix(strings.ToLower(str), "nanosecond") {
		numStr := str[:len(str)-10]
		num, _ := strconv.ParseUint(numStr, 10, 64)
		return time.Nanosecond * time.Duration(num)
	}

	num, _ := strconv.ParseUint(str, 10, 64)
	return time.Duration(num) * time.Second
}

func TimeDurationToStringCN(t time.Duration) string {
	const day = 24 * time.Hour
	const year = 365 * day

	if t > year {
		return fmt.Sprintf("%d年", t/year)
	} else if t > day {
		return fmt.Sprintf("%d天", t/day)
	} else if t > time.Hour {
		return fmt.Sprintf("%d小时", t/time.Hour)
	} else if t > time.Minute {
		return fmt.Sprintf("%d分钟", t/time.Minute)
	} else if t > time.Second {
		return fmt.Sprintf("%d秒", t/time.Second)
	}

	return "0秒"
}

func TimeDurationToString(t time.Duration) string {
	const day = 24 * time.Hour
	const year = 365 * day

	if t > year {
		return fmt.Sprintf("%dY", t/year)
	} else if t > day {
		return fmt.Sprintf("%dD", t/day)
	} else if t > time.Hour {
		return fmt.Sprintf("%dh", t/time.Hour)
	} else if t > time.Minute {
		return fmt.Sprintf("%dmin", t/time.Minute)
	} else if t > time.Second {
		return fmt.Sprintf("%ds", t/time.Second)
	}

	return "0s"
}
