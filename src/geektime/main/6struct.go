package main

import "reflect"

// struct后面的tag --当json读到myTypename/nodeName，就会把知道Name属性，把go struct和json struct联系起来了
// 如果再将Yaml转成json  这就变成yaml结构到go对象的转变
type myType struct {
	Name string `json:"myTypename,nodeName" protobuf:"byte"`
}

func main() {
	mt := myType{Name: "Namevalue"}
	myType := reflect.TypeOf(mt) //反射机制
	tag := myType.Field(0).Tag.Get("protobuf")
	println(tag)
}
