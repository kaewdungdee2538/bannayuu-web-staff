package utils

func IsNotStringAlphabet(str string) bool {
	for _, charVariable := range str {
		if (charVariable < 'a' || charVariable > 'z') && (charVariable < 'A' || charVariable > 'Z') && (charVariable<'0'||charVariable>'9'){
			return true
		}
	}
	return false
}
