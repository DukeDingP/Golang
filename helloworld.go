package main
import("fmt")
fuc main(){
 fmt.Print("hello world") 
}

//python
//a=[1,2,3,4,5,6]
//for i in range(len(a)):
 
// print(i,a[i])

fuc main(){
 a:=[6]int{1,2,3,4,5,6}
 for i,x:=range a{
  fmt.Printf(i,x)
 }
 for j=0;j<len(a);j++{
 fmt.Printf(a[j])
 }
}
