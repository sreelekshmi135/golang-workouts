package main
 import "fmt"
 
func find(ch interface{}){
	switch ch.(type) {
	case int:
		fmt.Println("Type:int,vaue:",ch.(int))
	case string:
		fmt.Println("Type:string,value:",ch.(string))
	case float64:
		fmt.Println("Type:float64,value:",ch.(float64))
	default:
		fmt.Println("Type not found")
		
	}
} 
 func main(){
	var ch interface{}
    fmt.Println("enter your charcater")
	fmt.Scanln(&ch)
	

 }
 