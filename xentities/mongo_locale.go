package xentities

import (
	"context"

	"github.com/jindasoft/template-platform-go/xenums"
	"github.com/jindasoft/template-platform-go/xutils"
)

type MongoLocale struct {
	EnUS string `json:"en_us" bson:"en_us"`
	ThTH string `json:"th_th" bson:"th_th"`
	MsMY string `json:"ms_my" bson:"ms_my"`
	LoLA string `json:"lo_la" bson:"lo_la"`
	ViVN string `json:"vi_vn" bson:"vi_vn"`
	JaJP string `json:"ja_jp" bson:"ja_jp"`
	KoKR string `json:"ko_kr" bson:"ko_kr"`
	ZhCN string `json:"zh_cn" bson:"zh_cn"`
	ZhTW string `json:"zh_tw" bson:"zh_tw"`
	RuRU string `json:"ru_ru" bson:"ru_ru"`
}

func (l MongoLocale) LocalizeString(ctx context.Context) string {
	cul := xutils.GetLocale(ctx)
	locale, err := xenums.ParseLocale(cul.String())
	if err != nil {
		return l.EnUS // fallback
	}

	switch locale {
	case xenums.LocaleEnUS:
		return l.EnUS
	case xenums.LocaleThTH:
		return l.ThTH
	case xenums.LocaleMsMY:
		return l.MsMY
	case xenums.LocaleLoLA:
		return l.LoLA
	case xenums.LocaleViVN:
		return l.ViVN
	case xenums.LocaleJaJP:
		return l.JaJP
	case xenums.LocaleKoKR:
		return l.KoKR
	case xenums.LocaleZhTW:
		return l.ZhTW
	case xenums.LocaleZhCN:
		return l.ZhCN
	case xenums.LocaleRuRU:
		return l.RuRU

	default:
		return l.EnUS
	}
}

func (l MongoLocale) UpdateLocale(locale string, s string) (MongoLocale, error) {
	c, err := xenums.ParseLocale(locale)
	if err != nil {
		return l, err
	}

	switch c {
	case xenums.LocaleEnUS:
		l.EnUS = s
	case xenums.LocaleThTH:
		l.ThTH = s
	case xenums.LocaleMsMY:
		l.MsMY = s
	case xenums.LocaleLoLA:
		l.LoLA = s
	case xenums.LocaleViVN:
		l.ViVN = s
	case xenums.LocaleJaJP:
		l.JaJP = s
	case xenums.LocaleKoKR:
		l.KoKR = s
	case xenums.LocaleZhTW:
		l.ZhTW = s
	case xenums.LocaleZhCN:
		l.ZhCN = s
	case xenums.LocaleRuRU:
		l.RuRU = s
	default:
		l.EnUS = s
	}

	return l, nil
}

func NewMongoLocale(locale string, s string) (MongoLocale, error) {
	c, err := xenums.ParseLocale(locale)
	if err != nil {
		return MongoLocale{}, err
	}

	switch c {
	case xenums.LocaleEnUS:
		return MongoLocale{EnUS: s}, nil
	case xenums.LocaleThTH:
		return MongoLocale{ThTH: s}, nil
	case xenums.LocaleMsMY:
		return MongoLocale{MsMY: s}, nil
	case xenums.LocaleLoLA:
		return MongoLocale{LoLA: s}, nil
	case xenums.LocaleViVN:
		return MongoLocale{ViVN: s}, nil
	case xenums.LocaleJaJP:
		return MongoLocale{JaJP: s}, nil
	case xenums.LocaleKoKR:
		return MongoLocale{KoKR: s}, nil
	case xenums.LocaleZhTW:
		return MongoLocale{ZhTW: s}, nil
	case xenums.LocaleZhCN:
		return MongoLocale{ZhCN: s}, nil
	case xenums.LocaleRuRU:
		return MongoLocale{RuRU: s}, nil
	default:
		return MongoLocale{EnUS: s}, nil
	}
}
