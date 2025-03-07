package yaml

import (
	"encoding/base64"
	"math"
	"regexp"
	"strings"
	"time"
)

type resolveMapItem struct {
	value interface{}
	tag   string
}

var resolveTable = make([]byte, 256)
var resolveMap = make(map[string]resolveMapItem)

const (
	nullTag      = "!!null"
	boolTag      = "!!bool"
	strTag       = "!!str"
	intTag       = "!!int"
	floatTag     = "!!float"
	timestampTag = "!!timestamp"
	seqTag       = "!!seq"
	mapTag       = "!!map"
	binaryTag    = "!!binary"
	mergeTag     = "!!merge"
)

func init() {
	t := resolveTable
	t[int('+')] = 'S'
	t[int('-')] = 'S'

	for _, c := range "0123456789" {
		t[int(c)] = 'D'
	}

	for _, c := range "yYnNtTfFoO~" {
		t[int(c)] = 'M'
	}

	t[int('.')] = '.'

	var resolveMapList = []struct {
		v   interface{}
		tag string
		l   []string
	}{
		{true, boolTag, []string{"true", "True", "TRUE"}},
		{false, boolTag, []string{"false", "False", "FALSE"}},
		{nil, nullTag, []string{"", "~", "null", "Null", "NULL"}},
		{math.NaN(), floatTag, []string{".nan", ".NaN", ".NAN"}},
		{math.Inf(+1), floatTag, []string{".inf", ".Inf", ".INF"}},
		{math.Inf(+1), floatTag, []string{"+.inf", "+.Inf", "+.INF"}},
		{math.Inf(-1), floatTag, []string{"-.inf", "-.Inf", "-.INF"}},
		{"<<", mergeTag, []string{"<<"}},
	}

	m := resolveMap

	for _, item := range resolveMapList {
		for _, s := range item.l {
			m[s] = resolveMapItem{item.v, item.tag}
		}
	}
}

var longTagPrefix = "tag:yaml.org,2002:"
var longTags = make(map[string]string)
var shortTags = make(map[string]string)

// 生成标签的长名称
func longTag(tag string) string {
	if strings.HasPrefix(tag, "!!") {
		if ltag, ok := longTags[tag]; ok {
			return ltag
		}
		return longTagPrefix + tag[2:]
	}
	return tag
}

// 生成标签的短名称
func shortTag(tag string) string {
	if strings.HasPrefix(tag, longTagPrefix) {
		if stag, ok := shortTags[tag]; ok {
			return stag
		}
		return "!!" + tag[len(longTagPrefix):]
	}
	return tag
}

func init() {
	for _, stag := range []string{nullTag, boolTag, strTag, intTag, floatTag, timestampTag, seqTag, mapTag, binaryTag, mergeTag} {
		ltag := longTag(stag)
		longTags[stag] = ltag
		shortTags[ltag] = stag
	}
}

// 查看标签是表示""、null、bool、str、int、float、timestamp
func resolvableTag(tag string) bool {
	switch tag {
	case "", nullTag, boolTag, strTag, intTag, floatTag, timestampTag:
		return true
	}
	return false
}

var yamlStyleFloat = regexp.MustCompile(`^[-+]?(\.[0-9]+|[0-9]+(\.[0-9]*)?)([eE][-+]?[0-9]+)?$`)

var allowedTimestampFormats = []string{
	"2006-1-2T15:4:5.999999999Z07:00", // RCF3339Nano with short date fields.
	"2006-1-2t15:4:5.999999999Z07:00", // RFC3339Nano with short date fields and lower-case "t".
	"2006-1-2 15:4:5.999999999",       // space separated with no time zone
	"2006-1-2",                        // date only
}

// 将传入的时间解析为时间戳
func parseTimestamp(s string) (time.Time, bool) {
	// 所有日期都是以YYYY-开头
	i := 0
	for ; i < len(s); i++ {
		if c := s[i]; c < '0' || c > '9' {
			break
		}
	}

	if i != 4 || i == len(s) || s[i] != '-' {
		return time.Time{}, false
	}

	for _, format := range allowedTimestampFormats {
		if t, err := time.Parse(format, s); err == nil {
			return t, true
		}
	}
	return time.Time{}, false
}

// 将字符Base64处理
func encodeBase64(s string) string {
	const lineLen = 70
	encLen := base64.StdEncoding.EncodedLen(len(s))
	lines := encLen/lineLen + 1
	buf := make([]byte, encLen*2+lines)
	in := buf[0:encLen]
	out := buf[encLen:]
	base64.StdEncoding.Encode(in, []byte(s))
	k := 0
	for i := 0; i < len(in); i += lineLen {
		j := i + lineLen
		if j > len(in) {
			j = len(in)
		}
		k += copy(out[k:], in[i:j])
		if lines > 1 {
			out[k] = '\n'
			k++
		}
	}
	return string(out[:k])
}

func resolve(tag string, in string) (rtag string, out interface{}) {
	tag = shortTag(tag)

	if !resolvableTag(tag) {
		return tag, in
	}

	defer func() {
		switch tag {
		case "", rtag, strTag, binaryTag:
			return
		case floatTag:
			if rtag == intTag {
				switch v := out.(type) {
				case int64:
					rtag = floatTag
					out = float64(v)
					return
				case int:
					rtag = floatTag
					out = float64(v)
					return
				}
			}
		}
		failf("cannot decode %s `%s` as a %s", shortTag(rtag), in, shortTag(tag))
	}()

	return strTag, in
}
