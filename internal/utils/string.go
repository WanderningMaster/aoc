package utils

func StringReverse(str string) string {
    newStr := ""
    for idx := len(str) - 1; idx >= 0; idx -= 1 {
        newStr += string(str[idx]) 
    }

    return newStr
}
