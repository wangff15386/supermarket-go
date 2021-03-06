package bussiness

import "strconv"

// checkBarcode 校验标准条形码
func checkBarcode(str string) (int64, bool) {
	if len(str) <= 8 {
		return 0, false
	}
	v, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, false
	}
	return v, true
}

// checkPrice 校验价格
func checkPrice(str string) (float64, bool) {
	if len(str) > 8 {
		return 0, false
	}
	v, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0.00, false
	}
	return v, true
}
