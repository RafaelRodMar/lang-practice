class Solution {
    fun addBinary(a: String, b: String): String {
        var ast: String = ""
        var bst: String = ""
        ast = a.reversed()
        bst = b.reversed()
        var indexa: Int = 0
        var indexb: Int = 0
        var mellevo: Int = 0
        var newNum: String = ""
        
        while(indexa < ast.length || indexb < bst.length){
            var digitA: Int = if (indexa < ast.length) ast[indexa] - '0' else 0
            var digitB: Int = if (indexb < bst.length) bst[indexb] - '0' else 0
            var newDigit: Int = digitA + digitB + mellevo
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
            newNum += newDigit.toString()
            indexa++
            indexb++
        }
        
        if(mellevo==1)
        {
            newNum += "1"
        }
        
        newNum = newNum.reversed()
            
        return newNum
    }
}