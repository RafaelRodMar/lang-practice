class Solution {
    public String addBinary(String a, String b) {
        a = new StringBuilder(a).reverse().toString();
        b = new StringBuilder(b).reverse().toString();
        System.out.print(a + "\n");
        System.out.print(b + "\n");
        int indexa = 0;
        int indexb = 0;
        int mellevo = 0;
        String newNum = "";
        
        while(indexa < a.length() || indexb < b.length()){
            int digitA = indexa < a.length() ? a.charAt(indexa) - '0' : 0;
            int digitB = indexb < b.length() ? b.charAt(indexb) - '0' : 0;
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
            newNum += Integer.toString(newDigit);
            indexa++; indexb++;
        }
        
        if(mellevo==1) newNum += "1";
        
        newNum = new StringBuilder(newNum).reverse().toString();
            
        return newNum;
    }
}