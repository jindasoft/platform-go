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

func (l MongoLocale) LocalizeString(ctx context.Context) string {
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

func (l MongoLocale) UpdateLocale(culture string, s string) (MongoLocale, error) {
	c, err := xenums.FromCultureCode(culture)
	if err != nil {
		return l, err
	}

	switch c {
	case xenums.CultureEnUS:
		l.EnUS = s
	case xenums.CultureThTH:
		l.ThTH = s
	case xenums.CultureMsMY:
		l.MsMY = s
	case xenums.CultureLoLA:
		l.LoLA = s
	case xenums.CultureViVN:
		l.ViVN = s
	case xenums.CultureJaJP:
		l.JaJP = s
	case xenums.CultureKoKR:
		l.KoKR = s
	case xenums.CultureZhTW:
		l.ZhTW = s
	case xenums.CultureZhCN:
		l.ZhCN = s
	case xenums.CultureRuRU:
		l.RuRU = s
	default:
		l.EnUS = s
	}

	return l, nil
}

func NewMongoLocale(culture string, s string) (MongoLocale, error) {
	c, err := xenums.FromCultureCode(culture)
	if err != nil {
		return MongoLocale{}, err
	}

	switch c {
	case xenums.CultureEnUS:
		return MongoLocale{EnUS: s}, nil
	case xenums.CultureThTH:
		return MongoLocale{ThTH: s}, nil
	case xenums.CultureMsMY:
		return MongoLocale{MsMY: s}, nil
	case xenums.CultureLoLA:
		return MongoLocale{LoLA: s}, nil
	case xenums.CultureViVN:
		return MongoLocale{ViVN: s}, nil
	case xenums.CultureJaJP:
		return MongoLocale{JaJP: s}, nil
	case xenums.CultureKoKR:
		return MongoLocale{KoKR: s}, nil
	case xenums.CultureZhTW:
		return MongoLocale{ZhTW: s}, nil
	case xenums.CultureZhCN:
		return MongoLocale{ZhCN: s}, nil
	case xenums.CultureRuRU:
		return MongoLocale{RuRU: s}, nil
	default:
		return MongoLocale{EnUS: s}, nil
	}
}
