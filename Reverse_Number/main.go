package main

import"fmt"

func Reverse(n int )int{

reverse:=0


for n !=0{
digit:=n%10
reverse=reverse*10+digit
n=n/10

}
return reverse


}

func main(){
number:=12345678

fmt.Println("Old_Number", number)

result:=Reverse(number)
fmt.Println("New_Number", result)

}