package commandlineargs

import (
	"flag"
	"fmt"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/formatutils"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/osutils"
	"github.com/SongZihuan/BackendServerTemplate/src/utils/reflectutils"
	"reflect"
	"strings"
)

const OptionIdent = "  "
const OptionPrefix = "--"
const OptionShortPrefix = "-"
const UsagePrefixWidth = 10

func (d *CommandLineArgsDataType) ready() {
	if d.isReady() {
		return
	}

	d.setFlag()
	d.writeUsage()
	d.parser()
	d.flagReady = true
}

func (d *CommandLineArgsDataType) writeUsage() {
	if len(d.Usage) != 0 {
		return
	}

	var result strings.Builder
	result.WriteString(formatutils.FormatTextToWidth(fmt.Sprintf("Usage of %s:", osutils.GetArgs0Name()), formatutils.NormalConsoleWidth))
	result.WriteString("\n")

	val := reflect.ValueOf(*d)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)

		if !strings.HasSuffix(field.Name, "Data") {
			continue
		}

		option := field.Name[:len(field.Name)-4]
		optionName := ""
		optionShortName := ""
		optionUsage := ""

		if reflectutils.HasFieldByReflect(typ, option+"Name") {
			var ok bool
			optionName, ok = val.FieldByName(option + "Name").Interface().(string)
			if !ok {
				panic("can not get option name")
			}
		}

		if reflectutils.HasFieldByReflect(typ, option+"ShortName") {
			var ok bool
			optionShortName, ok = val.FieldByName(option + "ShortName").Interface().(string)
			if !ok {
				panic("can not get option short name")
			}
		} else if len(optionName) > 1 {
			optionShortName = optionName[:1]
		}

		if reflectutils.HasFieldByReflect(typ, option+"Usage") {
			var ok bool
			optionUsage, ok = val.FieldByName(option + "Usage").Interface().(string)
			if !ok {
				panic("can not get option usage")
			}
		}

		var title string
		var title1 string
		var title2 string
		if field.Type.Kind() == reflect.Bool {
			var optionData bool
			if reflectutils.HasFieldByReflect(typ, option+"Data") {
				var ok bool
				optionData, ok = val.FieldByName(option + "Data").Interface().(bool)
				if !ok {
					panic("can not get option data")
				}
			}

			if optionData == true {
				panic("bool option can not be true")
			}

			if optionName != "" {
				title1 = fmt.Sprintf("%s%s%s", OptionIdent, OptionPrefix, formatutils.FormatTextToWidth(optionName, formatutils.NormalConsoleWidth-len(OptionIdent)-len(OptionPrefix)))
			}

			if optionShortName != "" {
				title2 = fmt.Sprintf("%s%s%s", OptionIdent, OptionShortPrefix, formatutils.FormatTextToWidth(optionShortName, formatutils.NormalConsoleWidth-len(OptionIdent)-len(OptionPrefix)))
			}
		} else if field.Type.Kind() == reflect.String {
			var optionData string
			if reflectutils.HasFieldByReflect(typ, option+"Data") {
				var ok bool
				optionData, ok = val.FieldByName(option + "Data").Interface().(string)
				if !ok {
					panic("can not get option data")
				}
			}

			if optionName != "" && optionData != "" {
				title1 = fmt.Sprintf("%s%s%s", OptionIdent, OptionPrefix, formatutils.FormatTextToWidth(fmt.Sprintf("%s string, default: '%s'", optionName, optionData), formatutils.NormalConsoleWidth-len(OptionIdent)-len(OptionPrefix)))
			} else if optionName != "" && optionData == "" {
				title1 = fmt.Sprintf("%s%s%s", OptionIdent, OptionPrefix, formatutils.FormatTextToWidth(fmt.Sprintf("%s string", optionName), formatutils.NormalConsoleWidth-len(OptionIdent)-len(OptionPrefix)))
			}

			if optionShortName != "" && optionData != "" {
				title2 = fmt.Sprintf("%s%s%s", OptionIdent, OptionShortPrefix, formatutils.FormatTextToWidth(fmt.Sprintf("%s string, default: '%s'", optionShortName, optionData), formatutils.NormalConsoleWidth-len(OptionIdent)-len(OptionPrefix)))
			} else if optionShortName != "" && optionData == "" {
				title2 = fmt.Sprintf("%s%s%s", OptionIdent, OptionShortPrefix, formatutils.FormatTextToWidth(fmt.Sprintf("%s string", optionShortName), formatutils.NormalConsoleWidth-len(OptionIdent)-len(OptionPrefix)))
			}
		} else if field.Type.Kind() == reflect.Uint || field.Type.Kind() == reflect.Int {
			var optionData uint
			if reflectutils.HasFieldByReflect(typ, option+"Data") {
				var ok bool
				optionData, ok = val.FieldByName(option + "Data").Interface().(uint)
				if !ok {
					panic("can not get option data")
				}
			}

			if optionName != "" && optionData != 0 {
				title1 = fmt.Sprintf("%s%s%s", OptionIdent, OptionPrefix, formatutils.FormatTextToWidth(fmt.Sprintf("%s number, default: %d", optionName, optionData), formatutils.NormalConsoleWidth-len(OptionIdent)-len(OptionPrefix)))
			} else if optionName != "" && optionData == 0 {
				title1 = fmt.Sprintf("%s%s%s", OptionIdent, OptionPrefix, formatutils.FormatTextToWidth(fmt.Sprintf("%s number", optionName), formatutils.NormalConsoleWidth-len(OptionIdent)-len(OptionPrefix)))
			}

			if optionShortName != "" && optionData != 0 {
				title2 = fmt.Sprintf("%s%s%s", OptionIdent, OptionShortPrefix, formatutils.FormatTextToWidth(fmt.Sprintf("%s number, default: %d", optionShortName, optionData), formatutils.NormalConsoleWidth-len(OptionIdent)-len(OptionPrefix)))
			} else if optionShortName != "" && optionData == 0 {
				title2 = fmt.Sprintf("%s%s%s", OptionIdent, OptionShortPrefix, formatutils.FormatTextToWidth(fmt.Sprintf("%s number", optionShortName), formatutils.NormalConsoleWidth-len(OptionIdent)-len(OptionPrefix)))
			}
		} else {
			panic(fmt.Sprintf("the flag type (%s) is not support", field.Type.Name()))
		}

		if title1 == "" && title2 == "" {
			continue
		} else if title1 != "" && title2 == "" {
			title = title1
		} else if title1 == "" {
			title = title2
		} else {
			title = fmt.Sprintf("%s\n%s", title1, title2)
		}

		result.WriteString(title)
		result.WriteString("\n")

		usage := formatutils.FormatTextToWidthAndPrefix(optionUsage, UsagePrefixWidth, formatutils.NormalConsoleWidth)
		result.WriteString(usage)
		result.WriteString("\n\n")
	}

	d.Usage = strings.TrimRight(result.String(), "\n")
}

func (d *CommandLineArgsDataType) parser() {
	if d.flagParser {
		return
	}

	if !d.isFlagSet() {
		panic("flag not set")
	}

	flag.Parse()
	d.flagParser = true
}

func (d *CommandLineArgsDataType) isReady() bool {
	return d.isFlagSet() && d.isFlagParser() && d.flagReady
}

func (d *CommandLineArgsDataType) isFlagSet() bool {
	return d.flagSet
}

func (d *CommandLineArgsDataType) isFlagParser() bool {
	return d.flagParser
}

func getData[T any](d *CommandLineArgsDataType, data T) T { // 泛型函数无法作为 “方法” 只能作为函数
	if !d.isReady() {
		panic("flag not ready")
	}

	return data
}
