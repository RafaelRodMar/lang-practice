/**
 * @param {string} a
 * @param {string} b
 * @return {string}
 */
var addBinary = function(a, b) {
    a = a.split("").reverse().join("");
    b = b.split("").reverse().join("");
    var indexa = 0;
    var indexb = 0;
    var mellevo = 0;
    var newNum = "";

    while(indexa < a.length || indexb < b.length){
        var digitA = indexa < a.length ? a[indexa] - '0' : 0;
        var digitB = indexb < b.length ? b[indexb] - '0' : 0;
        var newDigit = digitA + digitB + mellevo;

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
        newNum += newDigit.toString();
        indexa++; indexb++;
    }

    if(mellevo==1) newNum += "1";

    newNum = newNum.split("").reverse().join("");
    return newNum;
};