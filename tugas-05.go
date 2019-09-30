package main

import "fmt"

type pelajar struct {
 nama string
 umur int
}
func main(){
var n1,n2,n3 pelajar;
n1.nama = "Ahmad"
n1.umur = 81
n2.nama = "Abi"
n2.umur = 87
n3.nama = "Zaka"
n3.umur = 47

fmt.Println(n1.nama, n1.umur)
fmt.Println(n2.nama, n2.umur)
fmt.Println(n3.nama, n3.umur)
}
