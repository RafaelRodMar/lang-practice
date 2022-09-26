object Solution {
    def addBinary(a: String, b: String): String = {
        var as = a
        var bs = b
        as = a.reverse
        bs = b.reverse
        var indexa = 0
        var indexb = 0
        var mellevo = 0
        var newNum = ""
        
        while(indexa < as.length || indexb < bs.length){
            var digitA = if (indexa < as.length) as(indexa) - '0' else 0
            var digitB = if (indexb < bs.length) bs(indexb) - '0' else 0
            var newDigit = digitA + digitB + mellevo

            mellevo = 0
            
            if(newDigit == 2)
            {
                newDigit = 0
                mellevo = 1
            }
            if(newDigit == 3)
            {
                newDigit = 1
                mellevo = 1
            }
            newNum = newNum + newDigit.toString
            indexa = indexa + 1
            indexb = indexb + 1
        }
        
        if(mellevo==1)
        {
            newNum = newNum + "1"
        }
        
        newNum = newNum.reverse
        return newNum
    }
}