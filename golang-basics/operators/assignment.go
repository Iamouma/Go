package main
  
import "fmt"
  
func main() {
   var p int = 45
    var q int = 50
     
   // “=”(Simple Assignment)
   p = q
   fmt.Println(p)
     
   // “+=”(Add Assignment)
    p += q
   fmt.Println(p)
     
   //“-=”(Subtract Assignment)
   p-=q
   fmt.Println(p)
     
   // “*=”(Multiply Assignment)
   p*= q
   fmt.Println(p)
     
   // “/=”(Division Assignment)
    p /= q
   fmt.Println(p)
    
    // “%=”(Modulus Assignment)
    p %= q
   fmt.Println(p)
    
}
