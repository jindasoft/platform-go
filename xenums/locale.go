package xenums

import (
	"fmt"
	"strings"
)

type Locale string

const (
	LocaleEnUS    Locale = "en-US" // English (United States)
	LocaleThTH    Locale = "th-TH" // Thai
	LocaleMsMY    Locale = "ms-MY" // Malay
	LocaleLoLA    Locale = "lo-LA" // Lao
	LocaleViVN    Locale = "vi-VN" // Vietnamese
	LocaleJaJP    Locale = "ja-JP" // Japanese
	LocaleKoKR    Locale = "ko-KR" // Korean
	LocaleZhCN    Locale = "zh-CN" // Chinese (Simplified)
	LocaleZhTW    Locale = "zh-TW" // Chinese (Traditional)
	LocaleRuRU    Locale = "ru-RU" // Russian
	LocaleDefault Locale = LocaleEnUS
)

var LocaleMap = map[string]Locale{
	strings.ToLower(LocaleEnUS.String()): LocaleEnUS,
	strings.ToLower(LocaleThTH.String()): LocaleThTH,
	strings.ToLower(LocaleMsMY.String()): LocaleMsMY,
	strings.ToLower(LocaleLoLA.String()): LocaleLoLA,
	strings.ToLower(LocaleViVN.String()): LocaleViVN,
	strings.ToLower(LocaleJaJP.String()): LocaleJaJP,
	strings.ToLower(LocaleKoKR.String()): LocaleKoKR,
	strings.ToLower(LocaleZhCN.String()): LocaleZhCN,
	strings.ToLower(LocaleZhTW.String()): LocaleZhTW,
	strings.ToLower(LocaleRuRU.String()): LocaleRuRU,
}

func FromLocaleCode(locale string) (Locale, error) {
	locale = strings.TrimSpace(locale)
	locale = strings.ToLower(locale)

	c, ok := LocaleMap[locale]
	if !ok {
		return "", fmt.Errorf("invalid locale: %s", locale)
	}

	return c, nil
}

func (l Locale) IsValid() bool {
	_, ok := LocaleMap[strings.ToLower(string(l))]

	return ok
}

func (l Locale) IsValidOrDefault() Locale {
	if l.IsValid() {
		return l
	}

	return LocaleDefault
}

func (l Locale) String() string {
	return string(l)
}

func (l Locale) IsDefault() bool {
	return l == LocaleDefault
}
