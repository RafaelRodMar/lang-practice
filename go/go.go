package main
import "fmt" //i/o functions
//import "os"  //Package os provides a platform-independent interface to operating system functionality.
//import "bufio" //Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object, creating another object (Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O.
import "strings" //Package strings implements simple functions to manipulate UTF-8 encoded strings.
import "strconv"  //Package strconv implements conversions to and from string representations of basic data types.
import "unicode"

var PI float64 = 3.141592653589793
var mimetypelist [20000]string  //pares son extensiones, impares son el tipo

func checkarray(a string, n int) int {
    for x:=0; x<n*2 ; x++ {
        if mimetypelist[x] == a {
            return x
        }
    }
    return -1
}

func main(){
	var auxString string
	//fmt.Scan(&auxString) //get input
	auxString = "text string"
	fmt.Println(auxString)

	var aux int  //int definition
	fmt.Println(aux)

	//while using a for
	var aux2 = 0
	for{
		aux2++

		if aux2 == 3 {
			break
		}
	}

	//for loop
	//i:= is a new variable not previously declared
	for i:=0; i<8 ; i++ {
		fmt.Print(i, " ")
	}

	//multiple declarations
	var aux3, aux4, aux5, aux6 string
	//multiple assign
	aux5 = "Y"
	aux6 = "N"
	aux3, aux4 = aux5, aux6

	//formatted print
	fmt.Printf("%s%s\n", aux3, aux4)

	aux7 := 2
	if( aux2 == 3 ) { aux7 = 5 }
	aux2 = aux7

	//buffer creation
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Buffer(make([]byte, 1000000), 1000000)

    //var n int
    //scanner.Scan()
    //fmt.Sscan(scanner.Text(),&n)
	//scanner.Scan()
    //inputs := strings.Split(scanner.Text()," ") //split string with " " as separator
	//i64, err := strconv.ParseInt(s, 10, 32) //base 10, 32 bits

	var aux8 = "abc1def2ghi3jkl4mno5pqr6stu7vwx8yz"
	var index int
	for _, c := range aux8 {
		if unicode.IsLetter(c) {
			index = int(byte(unicode.ToUpper(rune(c))) - 'A')
		} else {
			index = 26
		}

		fmt.Print(index)
	}
	fmt.Println("")

	// Take substring of first word with runes.
	// ... This handles any kind of rune in the string.
	//runes := []rune(ROW)
	// ... Convert back into a string from rune slice.
	//safeSubstring := string(runes[a:b])
	
	MESSAGE := "e"
	var messageinbinary string

    for _, c := range MESSAGE {
        s := strconv.FormatInt(int64(c), 2) //convertir binario
        if len(s) < 7 {
            s = string('0') + s
        }
        messageinbinary += s
    }
   fmt.Println(messageinbinary)

   strings.ContainsAny( messageinbinary, "0")
   it := strings.LastIndexAny( messageinbinary, "0")
   newext := string(messageinbinary[it+1:])
   fmt.Println(newext)
}