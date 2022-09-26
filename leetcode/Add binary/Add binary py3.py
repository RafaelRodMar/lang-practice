class Solution:
    def addBinary(self, a: str, b: str) -> str:
        a = a[::-1]
        b = b[::-1]
        indexa = 0
        indexb = 0
        mellevo = 0
        newNum = ""
        
        while indexa < len(a) or indexb < len(b) :
            digitA = (ord(a[indexa]) - ord('0')) if indexa < len(a) else 0
            digitB = (ord(b[indexb]) - ord('0')) if indexb < len(b) else 0
            newDigit = digitA + digitB + mellevo

            mellevo = 0
            
            if newDigit == 2:
                newDigit = 0
                mellevo = 1

            if newDigit == 3:
                newDigit = 1
                mellevo = 1

            newNum += str(newDigit)
            indexa += 1
            indexb += 1
        
        if mellevo==1:
            newNum += "1"
        
        newNum = newNum[::-1]
        return newNum