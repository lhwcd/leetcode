package circulardeque

import (
	"fmt"
	"testing"
)

func TestCircularDeque(t *testing.T) {
	obj := Constructor(3)
	fmt.Println(obj.InsertLast(1))  //true
	fmt.Println(obj.InsertLast(2))  //true
	fmt.Println(obj.InsertFront(3)) //true
	fmt.Println(obj.InsertFront(4)) //false
	fmt.Println(obj.GetRear())      //2
	fmt.Println(obj.IsFull())       //true
	fmt.Println(obj.DeleteLast())   //true
	fmt.Println(obj.InsertFront(4)) //true
	fmt.Println(obj.GetFront())     //4
	//
	//[null,true,true,true,false,2,true,true,true,4]
}
