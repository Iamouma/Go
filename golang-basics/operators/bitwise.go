package main
 
import "fmt"
 
func main() {
   p:= 34
   q:= 20
    
   // & (bitwise AND)
   result1:= p & q
   fmt.Printf("Result of p & q = %d", result1)
    
   // | (bitwise OR)
   result2:= p | q
   fmt.Printf("\nResult of p | q = %d", result2)
    
   // ^ (bitwise XOR)
   result3:= p ^ q
   fmt.Printf("\nResult of p ^ q = %d", result3)
    
   // << (left shift)
   result4:= p << 1
   fmt.Printf("\nResult of p << 1 = %d", result4)
    
   // >> (right shift)
   result5:= p >> 1
   fmt.Printf("\nResult of p >> 1 = %d", result5)
    
   // &^ (AND NOT)
   result6:= p &^ q
   fmt.Printf("\nResult of p &^ q = %d", result6)    
}
