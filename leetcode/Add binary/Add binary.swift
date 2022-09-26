class Solution {
    func addBinary(_ a: String, _ b: String) -> String {
        var sa = String(a.reversed())
        var sb = String(b.reversed())
        var indexa:Int = 0
        var indexb:Int = 0
        var mellevo:UInt8 = 0
        var newNum = ""
        var inda = sa.startIndex
        var indb = sb.startIndex
        
        while (indexa < sa.count || indexb < sb.count)
        {
            if indexa < sa.count
            {
                inda = sa.index(sa.startIndex, offsetBy: indexa)
            }
            //print("sa[inda].asciiValue!", sa[inda].asciiValue!)
            var digitA = indexa < sa.count ? sa[inda].asciiValue! - 48 : 0
            //print("digitA", digitA)
            if indexb < sb.count
            {
                indb = sb.index(sb.startIndex, offsetBy: indexb)
            }
            //print("sb[indb].asciiValue!", sb[indb].asciiValue!)
            var digitB = indexb < sb.count ? sb[indb].asciiValue! - 48 : 0
            //print("digitB", digitB)
            var newDigit = digitA + digitB + mellevo

            mellevo = 0
            
            if newDigit == 2
            {
                newDigit = 0
                mellevo = 1
            }

            if newDigit == 3
            {
                newDigit = 1
                mellevo = 1
            }

            newNum += String(newDigit)
            indexa += 1
            indexb += 1
        }
        
        if mellevo==1
        {
            newNum += "1"
        }
        
        newNum = String(newNum.reversed())
        return newNum
    }
}