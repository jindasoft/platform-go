package xentities

import (
	"context"

	"github.com/jindasoft/jinda-platform/xenums"
	"github.com/jindasoft/jinda-platform/xutils"
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

func (l MongoLocale) String(ctx context.Context) string {
	cul := xutils.GetCulture(ctx)
	culture, err := xenums.FromCultureCode(cul.String())
	if err != nil {
		return l.EnUS // fallback
	}

	switch culture {
	case xenums.CultureEnUS:
		return l.EnUS
	case xenums.CultureThTH:
		return l.ThTH
	case xenums.CultureMsMY:
		return l.MsMY
	case xenums.CultureLoLA:
		return l.LoLA
	case xenums.CultureViVN:
		return l.ViVN
	case xenums.CultureJaJP:
		return l.JaJP
	case xenums.CultureKoKR:
		return l.KoKR
	case xenums.CultureZhTW:
		return l.ZhTW
	case xenums.CultureZhCN:
		return l.ZhCN
	case xenums.CultureRuRU:
		return l.RuRU

	default:
		return l.EnUS
	}
}

func SetLocale(locale MongoLocale) MongoLocale {
	return locale
}
