public class Solution {
    public static string Reverse( string s )
    {
        char[] charArray = s.ToCharArray();
        Array.Reverse(charArray);
        return new string(charArray);
    }
    
    public string AddBinary(string a, string b) {
        a = Reverse(a);
        b = Reverse(b);
        int indexa = 0;
        int indexb = 0;
        int mellevo = 0;
        string newNum = "";
        int asize = a.Length;
        int bsize = b.Length;
        
        while(indexa < asize || indexb < bsize){
            int digitA = indexa < asize ? a[indexa] - '0' : 0;
            int digitB = indexb < bsize ? b[indexb] - '0' : 0;
            int newDigit = digitA + digitB + mellevo;

            mellevo = 0;
            
            if(newDigit == 2)
            {
                newDigit = 0;
                mellevo = 1;
            }
            if(newDigit == 3)
            {
                newDigit = 1;
                mellevo = 1;
            }
            newNum += newDigit.ToString();
            indexa++; indexb++;
        }
        
        if(mellevo==1) newNum += "1";
        
        newNum = Reverse(newNum);

        return newNum;
    }
}