//A struct is a collection of field or properties

package main

import "fmt"
type book struct{
title string
author string
isbn string
price float32
pages int
}
func main(){
var b1 book
b1.title="an introduction to programming in go"
b1.author="CALLEB DOXEY"
b1.isbn="978-1478355823"
b1.price=10.50
b1.pages=165
fmt.Println(b1)
fmt.Println(b1.title)
}