func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func addBinary(a string, b string) string {
    a = Reverse(a)
    b = Reverse(b)
    var indexa = 0
    var indexb = 0
    var mellevo = 0
    var newNum string = ""

    for indexa < len(a) || indexb < len(b) {
        var digitA = 0
        if indexa < len(a) {
            digitA = int(byte(rune(a[indexa]) - '0'))
        }
        var digitB = 0
        if indexb < len(b) {
            digitB = int(byte(rune(b[indexb]) - '0'))
        }
        var newDigit = digitA + digitB + mellevo

        mellevo = 0

        if newDigit == 2 {
            newDigit = 0
            mellevo = 1
        }
        if newDigit == 3 {
            newDigit = 1
            mellevo = 1
        }
        newNum += strconv.Itoa(newDigit)
        indexa++
        indexb++
    }

    if mellevo==1 {
        newNum += "1"
    }

    newNum = Reverse(newNum)

    return newNum
}