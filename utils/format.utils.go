package utils

import (
	"time"
	"strings"
)

func IsNotStringAlphabet(str string) bool {
	const alpha = `/!@#$%^&*()_+\-=\[\]{};':"|,.<>\/?~`
	// const alpha = `abcdefghijklmnopqrstuvwxyz0123456789กขฃคฅฆงจฉชซฌญฎฏฐฑฒณดตถทธนบปผฝพฟภมยรลวศษสหฬอฮ`
	for _, char := range str {  
		if strings.Contains(alpha, strings.ToLower(string(char))) {
		   return true
		}
	 }
	 return false
}

func IsNotStringEngOtNumber(str string) bool {
	for _, charVariable := range str {
		if (charVariable < 'a' || charVariable > 'z') && (charVariable < 'A' || charVariable > 'Z') && (charVariable < '0' || charVariable > '9') {
			return true
		}
	}
	return false
}

func IsNotFormatTime(str string) bool {
	if _, err := time.Parse("2006-01-02 15:04:00",str); err != nil{
		return true;
	}
	return false;
}
