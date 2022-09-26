char* strrev( char* s )
{
    char  c;
    char* s0 = s - 1;
    char* s1 = s;

    /* Find the end of the string */
    while (*s1) ++s1;

    /* Reverse it */
    while (s1-- > ++s0)
    {
        c   = *s0;
        *s0 = *s1;
        *s1 =  c;
    }

    return s;
}

char * addBinary(char * a, char * b){
    char *num1 = strrev(a);
    char *num2 = strrev(b);
    //printf("%s // %s", num1, num2);
    int indexa = 0;
    int indexb = 0;
    int mellevo = 0;
    /* with more space than needed */
    int alen = strlen(num1);
    int blen = strlen(num2);
    int siz = alen >= blen ? alen : blen;
    siz += 2;
    char *newNum = malloc(siz * sizeof(char));
    int newNumIndex = 0;
    
    while(indexa < strlen(a) || indexb < strlen(b)){
        int digitA = indexa < strlen(a) ? a[indexa] - '0' : 0;
        int digitB = indexb < strlen(b) ? b[indexb] - '0' : 0;
        //printf("digitA = %d, digitB = %d\n", digitA, digitB);
        int newDigit = digitA + digitB + mellevo;
        //printf("newDigit = %d\n", newDigit);
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
        //printf("newDigit = %d\n", newDigit);
        char c = newDigit + '0';
        //printf("addchar = %c\n", c);
        newNum[newNumIndex] = c;
        newNumIndex++;
        //printf("newNum después = %s \n", newNum);
        indexa++; indexb++;
    }
    
    if(mellevo==1)
    {
        newNum[newNumIndex] = '1';
        newNum[newNumIndex+1] = '\0';
    }
    else
    {
        newNum[newNumIndex] = '\0';
    }
    
    newNum = strrev(newNum);
    return newNum;
}