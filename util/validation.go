package util

import (
	"regexp"
	"strconv"
	"strings"
)

var reNumber *regexp.Regexp = regexp.MustCompile(`^[\d]+$`)
var reEmail *regexp.Regexp = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func CheckIdCardFormat(idCard string) bool {

	var idCardReq = regexp.MustCompile(`^[\d]{17}[0-9xX]$`)
	if !idCardReq.MatchString(idCard) {
		return false
	}

	idCards := strings.Split(idCard, "")
	a := [...]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	sum := 0
	num := len(a)
	for i := 0; i < num; i++ {
		v := a[i]
		l := idCards[i]
		il, _ := strconv.Atoi(l)
		sum += il * v
	}

	sumMap := [...]string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}

	if sumMap[sum%11] == strings.ToUpper(idCard[17:18]) {
		return true
	}

	return false

}

func CheckNameFormat(name string) bool {
	var nameReg = regexp.MustCompile("^[\u4e00-\u9fa5]+(·?[\u4e00-\u9fa5]+)*$")
	return nameReg.MatchString(name)
}

func CheckBankCardFormat(bankCard string) bool {

	var invalidBankCardReg = regexp.MustCompile("[^0-9-\\s]+")
	if invalidBankCardReg.MatchString(bankCard) {
		return false
	}

	vs := strings.Split(bankCard, "")
	bCheck := 0
	bEven := false
	for i := len(vs) - 1; i >= 0; i-- {
		bi, _ := strconv.Atoi(vs[i])
		if bEven {
			bi = (bi * 2)
			if bi > 9 {
				bi -= 9
			}
		}

		bCheck += bi
		bEven = !bEven
	}

	if bCheck%10 == 0 {
		return true
	}

	return false
}

func CheckPhoneFormat(phone string) bool {

	var phoneReg = regexp.MustCompile(`^1[3|4|5|7|8]\d{9}$`)

	if !phoneReg.MatchString(phone) {
		return false
	}

	return true
}

func CheckEmailFormat(email string) bool {
	return reEmail.MatchString(email)
}

func CheckNumber(nubmer string) bool {
	return reNumber.MatchString(nubmer)

}
