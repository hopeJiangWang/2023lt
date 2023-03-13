package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Append(h *ListNode, i int) {
	for h.Next != nil {
		h = h.Next
	}
	h.Next = &ListNode{Val: i}
}

func Show(h *ListNode) {
	fmt.Println(h.Val)
	for h.Next != nil {
		h = h.Next
		fmt.Println(h.Val)
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	/*
		链表相加：
			2->4->3
			5->6->4
			==>
			7->0->8
	*/
	// 先生成下头尾节点（头节点用于标记，尾节点用于不断往后移动）
	head := &ListNode{Val: 0}
	tail := head

	// 记录下进位，默认为 0，下面循环切记需要加上carry作为循环退出条件 ！！！
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		// 默认节点值为 0，如果节点不为空，则更新节点值
		l1Val, l2Val := 0, 0
		if l1 != nil {
			// 直接在这里更新节点了，要不然结尾还得判断下是否为空节点
			l1Val, l1 = l1.Val, l1.Next
		}

		if l2 != nil {
			l2Val, l2 = l2.Val, l2.Next
		}

		// 计算当前节点的和，更新进位，获取值并生成节点，插入链表
		tmp := l1Val + l2Val + carry
		carry = tmp / 10

		tail.Next = &ListNode{Val: tmp % 10}
		tail = tail.Next
	}

	return head.Next
}

func genList(slice []int) *ListNode {
	if len(slice) == 0 {
		return nil
	}

	head := &ListNode{Val: slice[0]}
	tail := head

	for i := 1; i < len(slice); i++ {
		tail.Next = &ListNode{Val: slice[i]}
		tail = tail.Next
	}

	return head
}

func main() {
	//l1 := []int{2, 4, 3}
	//l2 := []int{5, 6, 4}
	l1 := []int{9, 9, 9, 9, 9, 9, 9}
	l2 := []int{9, 9, 9, 9}

	list1 := genList(l1)
	list2 := genList(l2)

	Show(list1)
	Show(list2)

	res := addTwoNumbers(list1, list2)

	Show(res)
}
