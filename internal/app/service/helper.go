package service

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"strconv"
	"strings"
)

func Int64FromPtr(value *int64) int64 {
	if value != nil {
		return *value
	}

	return 0
}

func StringFromPtr(value *string) string {
	if value != nil {
		return *value
	}

	return ""
}

func FormatRupiahFromFloat(amount float64) string {
	humanizeValue := humanize.CommafWithDigits(amount, 0)
	stringValue := strings.Replace(humanizeValue, ",", ".", -1)
	return "Rp " + stringValue
}

func FormatRupiahFromInt64(amount int64) string {
	return FormatRupiahFromFloat(float64(amount))
}

func FormattedDiscountPercentage(value int64) string {
	return fmt.Sprintf("%v%%", value)
}

func FormattedDiscountValue(value int64, discountTypeId int64) string {
	switch discountTypeId {
	case 1:
		return FormattedDiscountPercentage(value)
	case 2:
		return FormatRupiahFromInt64(value)
	}

	return ""
}

func FormatShortNumber(n float64, precision int) string {
	var formattedNumber string
	var symbol string

	switch {
	case n < 900:
		formattedNumber = strconv.FormatFloat(n, 'f', precision, 64)
	case n < 900_000:
		formattedNumber = strconv.FormatFloat(n/1_000, 'f', precision, 64)
		symbol = "rb"
	case n < 900_000_000:
		formattedNumber = strconv.FormatFloat(n/1_000_000, 'f', precision, 64)
		symbol = "jt"
	case n < 900_000_000_000:
		formattedNumber = strconv.FormatFloat(n/1_000_000_000, 'f', precision, 64)
		symbol = "M"
	default:
		formattedNumber = strconv.FormatFloat(n/1_000_000_000_000, 'f', precision, 64)
		symbol = "T"
	}

	// Remove unnecessary decimal places if precision > 0
	if precision > 0 {
		trimZeros := "." + strings.Repeat("0", precision)
		formattedNumber = strings.Replace(formattedNumber, trimZeros, "", -1)
	}

	return formattedNumber + symbol
}
