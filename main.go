package main

import (
	"fmt"
	tries2 "github.com/bvinayb/data-strucutres-in-go/tries"
)

func main() {

	//--------- Tries -----------------

	tries := tries2.InitTrie()
	tries.Insert("awesome")

	fmt.Println(tries)

	//fmt.Println(tries.Search("this"))
	fmt.Println(tries.Search("awesome"))

	/*
		//--------- Binary Tree -----------------

		bt := &trees.BinaryTreeNode{}
		bt.Insert(87)
		bt.Insert(28)
		bt.Insert(13)
		bt.Insert(754)
		bt.Insert(12)
		bt.Insert(6)
		bt.Insert(9)
		bt.Insert(1234)

		fmt.Println(bt)

		fmt.Println(bt.Search(9))
	*/
	/*
		//--------- Queue -----------------
		newQ := queue.Queue{}
		newQ.Enqueue(44)
		newQ.Enqueue(32)
		newQ.Enqueue(1)
		newQ.Enqueue(4545)
		newQ.Enqueue(67856)
		fmt.Println(newQ)
		newQ.Dqueue()
		fmt.Println(newQ)
	*/

	/*
		//--------- Stack -----------------
		newStack := stack.Stack{}
		newStack.Push(12)
		newStack.Push(152)
		newStack.Push(612)
		newStack.Push(7)
		newStack.Push(9)
		newStack.Push(1234)

		fmt.Println(newStack)
		newStack.Pop()
		fmt.Println(newStack)
	*/

	/*
		//--------- LinkedList -----------------
		myList := linked_list.LinkedList{}

		node1 := linked_list.Node{Key: 23}
		node2 := linked_list.Node{Key: 332}
		node3 := linked_list.Node{Key: 76}
		node4 := linked_list.Node{Key: 5}
		node5 := linked_list.Node{Key: 7}
		node6 := linked_list.Node{Key: 56}
		node7 := linked_list.Node{Key: 345}
		node8 := linked_list.Node{Key: 208}
		node9 := linked_list.Node{Key: 168}

		myList.Prepend(&node1)
		myList.Prepend(&node2)
		myList.Prepend(&node3)
		myList.Prepend(&node4)
		myList.Prepend(&node5)
		myList.Prepend(&node6)
		myList.Prepend(&node7)
		myList.Prepend(&node8)
		myList.Prepend(&node9)

		myList.PrintListData()

		myList.Delete(4234234)

		myList.PrintListData()
	*/

	/*
		//--------- Hash Table -----------------
		hashTables := hashtable.Init()

		list := []string{
			"2",
			"5",
			"1",
			"2",
			"3",
			"5",
			"1",
			"127",
		}

		for _, v := range list {
			hashTables.Insert(v)
		}

		//hashTables.Search("Joe") //false

	*/

	//var first = []string{"a", "b", "c", "x"}
	//var second = []string{"e", "g", "a", "y"}
	//
	//var third = map[string]bool{}
	//
	//for _, val := range second {
	//	third[val] = true
	//}
	//var test = checkMe(first, third)
	//fmt.Println("Result:%v", test)

	//var a = []int{1, 2, 3, 9}

	//HAS PAIR with SUM
	// c=0 and n=1 --> 1,2
	// c=1 and n=2 --> 2,3
	// c= 2 and n=3 --> 3,9
	// var b = []int{1, 2, 4, 4}
	//result2 := hasPairWithSum(a, 8)
	//fmt.Println("Result:%v", result2)

	//	??ROBO Position
	//	var roboMaps = [][]int{
	//		//{1, 0, 0, 1}, {0, 1, 1, 1},
	//		{1, 0, 0, 0, 1}, {1, 0, 1, 0, 0},
	//	}
	//
	//	var result = robotPositions(roboMaps)
	//	fmt.Println("Result:%v", result)

	/**
	Reverse String
	*/
	//fmt.Println(reverseString("Vinay"))

	/**
	Merge Sorted Array
	*/
	//var a = []int{0, 3, 4, 31}
	//var b = []int{4, 6, 30}
	//mergeSortedArray(a, b)

	/**
	Max Subarray
	*/
	//moveZeroes([]int{0, 1, 0, 3, 12})
}
