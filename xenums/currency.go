package xenums

type Currency string

const (
	CurrencyTHB Currency = "THB" // Thai Baht
	CurrencyUSD Currency = "USD" // US Dollar
	CurrencySGD Currency = "SGD" // Singapore Dollar
	CurrencyMYR Currency = "MYR" // Malaysian Ringgit
	CurrencyLAK Currency = "LAK" // Lao Kip
	CurrencyVND Currency = "VND" // Vietnamese Dong
	CurrencyJPY Currency = "JPY" // Japanese Yen
	CurrencyKRW Currency = "KRW" // Korean Won
	CurrencyCNY Currency = "CNY" // Chinese Yuan
)

var CurrencyMap = map[string]Currency{
	"THB": CurrencyTHB,
	"USD": CurrencyUSD,
	"SGD": CurrencySGD,
	"MYR": CurrencyMYR,
	"LAK": CurrencyLAK,
	"VND": CurrencyVND,
	"JPY": CurrencyJPY,
	"KRW": CurrencyKRW,
	"CNY": CurrencyCNY,
}

func ParseCurrency(code string) (Currency, bool) {
	currency, ok := CurrencyMap[code]
	return currency, ok
}

func (c Currency) IsValid() bool {
	_, ok := CurrencyMap[string(c)]
	return ok
}

func (c Currency) String() string {
	return string(c)
}
