package main

import (
	"fmt"
	"github.com/zhangmingkai4315/GoWebStarter/InitTestWorkspace"
	"gopkg.in/mgo.v2"
)
var Session *(mgo.Session);
func init(){
	var err error
	Session,err=mgo.Dial("localhost")
	if err!=nil{
		panic(err)
	}
}
func dopanic(){
	defer func(){
	  if e:=recover();e!=nil{
		  fmt.Println("Recover from ",e)
	  }
	}()

	panic("I am panic")
	fmt.Println("Never be called")
}
func main() {
	fmt.Printf("Hello world :%v",InitTestWorkspace.Add(10,10))
	fmt.Printf("Mongodb is ready:%v",Session);
	defer Session.Close();

	x:=[]int{4:20};
	for _,i := range x{
		println(i);
	}
	y:=append(x,30);
	//fmt.Printf("%v",y)
	println(x,y); //[5/5]0xc8200f3e78 [6/10]0xc820010140
	println(cap(x),len(x));

	fmt.Println("Starting to panic")
	dopanic()
	fmt.Println("Recorve after panic");




}
