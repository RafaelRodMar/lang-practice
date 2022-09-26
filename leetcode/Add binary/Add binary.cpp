// https://leetcode.com/problems/add-binary

class Solution {
public:
    string addBinary(string a, string b) {
        reverse(a.begin(),a.end());
        reverse(b.begin(),b.end());
        int indexa = 0;
        int indexb = 0;
        int mellevo = 0;
        string newNum = "";
        
        while(indexa < a.size() || indexb < b.size()){
            int digitA = indexa < a.size() ? a[indexa] - '0' : 0;
            int digitB = indexb < b.size() ? b[indexb] - '0' : 0;
            int newDigit = digitA + digitB + mellevo;
            cout << "digitA: " << digitA << "  digitB: " << digitB << " suma: " << newDigit << endl;
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
            newNum += to_string(newDigit);
            indexa++; indexb++;
        }
        
        if(mellevo==1) newNum += "1";
        
        reverse(newNum.begin(), newNum.end());
        cout << "new num : " << newNum << endl;
        return newNum;
    }
};