package xenums

import (
	"fmt"
	"strings"
)

type Culture string

const (
	CultureEnUS    Culture = "en-US" // English (United States)
	CultureThTH    Culture = "th-TH" // Thai
	CultureMsMY    Culture = "ms-MY" // Malay
	CultureLoLA    Culture = "lo-LA" // Lao
	CultureViVN    Culture = "vi-VN" // Vietnamese
	CultureJaJP    Culture = "ja-JP" // Japanese
	CultureKoKR    Culture = "ko-KR" // Korean
	CultureZhCN    Culture = "zh-CN" // Chinese (Simplified)
	CultureZhTW    Culture = "zh-TW" // Chinese (Traditional)
	CultureRuRU    Culture = "ru-RU" // Russian
	CultureDefault Culture = CultureEnUS
)

var CultureMap = map[string]Culture{
	strings.ToLower(CultureEnUS.String()): CultureEnUS,
	strings.ToLower(CultureThTH.String()): CultureThTH,
	strings.ToLower(CultureMsMY.String()): CultureMsMY,
	strings.ToLower(CultureLoLA.String()): CultureLoLA,
	strings.ToLower(CultureViVN.String()): CultureViVN,
	strings.ToLower(CultureJaJP.String()): CultureJaJP,
	strings.ToLower(CultureKoKR.String()): CultureKoKR,
	strings.ToLower(CultureZhCN.String()): CultureZhCN,
	strings.ToLower(CultureZhTW.String()): CultureZhTW,
	strings.ToLower(CultureRuRU.String()): CultureRuRU,
}

func PairCulture(culture string) (Culture, error) {
	culture = strings.TrimSpace(culture)
	culture = strings.ToLower(culture)

	c, ok := CultureMap[culture]
	if !ok {
		return "", fmt.Errorf("invalid culture: %s", culture)
	}

	return c, nil
}

func (l Culture) IsValid() bool {
	_, ok := CultureMap[strings.ToLower(string(l))]

	return ok
}

func (l Culture) IsValidOrDefault() Culture {
	if l.IsValid() {
		return l
	}

	return CultureDefault
}

func (l Culture) String() string {
	return string(l)
}

func (l Culture) IsDefault() bool {
	return l == CultureDefault
}
