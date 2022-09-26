# @param {String} a
# @param {String} b
# @return {String}
def add_binary(a, b)
    a.reverse!
    b.reverse!
    @indexa = 0
    @indexb = 0
    @mellevo = 0
    @newNum = ""

    while (@indexa < a.length || @indexb < b.length)
        @digitA = @indexa < a.length ? a[@indexa].ord - '0'.ord : 0
        @digitB = @indexb < b.length ? b[@indexb].ord - '0'.ord : 0
        @newDigit = @digitA + @digitB + @mellevo
        
        @mellevo = 0
        
        if @newDigit == 2
            @newDigit = 0
            @mellevo = 1
        end
        if @newDigit == 3
            @newDigit = 1
            @mellevo = 1
        end
        @newNum += @newDigit.to_s
        @indexa += 1
        @indexb += 1
    end

    if @mellevo==1
        @newNum += "1"
    end

    @newNum.reverse!
    return @newNum;
end