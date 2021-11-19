package main 
import "fmt"
import "strings"

// func main ()  {
// 	var n [10]int /* n is an array of 10 integers */
//    var i,j int

//    /* initialize elements of array n to 0 */         
//    for i = 0; i < 10; i++ {
//       n[i] = i + 100 /* set element at location i to i + 100 */
//    }
   
//    /* output each array element's value */
//    for j = 0; j < 10; j++ {
//       fmt.Printf("Element[%d] = %d\n", j, n[j] )
//    }

// }

func max(num1, num2 int) int {
	var result int = num1
	if (num2 > num1){
		result = num2
	}
	return result
}

func swap(x,y string) (string, string){
	return y, x
}

func idk(){
	var greeting =  "Hello world!"
   
	fmt.Printf("normal string: ")
	fmt.Printf("%s", greeting)
	fmt.Printf("\n")
	fmt.Printf("hex bytes: ")
	
	for i := 0; i < len(greeting); i++ {
		fmt.Printf("%x ", greeting[i])
	}
	
	fmt.Printf("\n")
	const sampleText = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98" 
	
	/*q flag escapes unprintable characters, with + flag it escapses non-ascii 
	characters as well to make output unambigous */
	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", sampleText)
	fmt.Printf("\n")  

	greetings :=  []string{"Hello","world!"}
	fmt.Println(strings.Join(greetings, " "))


}