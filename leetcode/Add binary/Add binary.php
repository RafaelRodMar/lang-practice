class Solution {

    /**
     * @param String $a
     * @param String $b
     * @return String
     */
    function addBinary($a, $b) {
        $a = strrev($a);
        $b = strrev($b);
        $indexa = 0;
        $indexb = 0;
        $mellevo = 0;
        $newNum = "";

        while($indexa < strlen($a) || $indexb < strlen($b)){
            $digitA = $indexa < strlen($a) ? ord($a[$indexa]) - ord('0') : 0;
            $digitB = $indexb < strlen($b) ? ord($b[$indexb]) - ord('0') : 0;
            $newDigit = $digitA + $digitB + $mellevo;

            $mellevo = 0;

            if($newDigit == 2)
            {
                $newDigit = 0;
                $mellevo = 1;
            }
            if($newDigit == 3)
            {
                $newDigit = 1;
                $mellevo = 1;
            }
            $newNum .= strval($newDigit);
            $indexa++; $indexb++;
        }

        if($mellevo==1) $newNum .= "1";

        $newNum = strrev($newNum);
        return $newNum;
    }
}