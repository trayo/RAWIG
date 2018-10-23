/*
This file is slightly changed Caleb Doxsey's implementation of stack.
It is available on github (access 20181023):
https://github.com/badgerodon/collections/blob/master/stack/stack.go
It follows MIT License:

Copyright (c) 2012 Caleb Doxsey

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be included
in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package main

type (
	Stack struct {
		Top    *Node
		Length int
	}
	Node struct {
		Value interface{}
		Prev  *Node
	}
)

func NewStack() *Stack {
	/*Function NewStack creates and returns new Stack.*/
	return &Stack{nil, 0}
}

func (s *Stack) Len() int {
	/*Method Len has pointer to the Stack as receiver,
	and returns its Length.*/
	return s.Length
}

func (s *Stack) Peek() interface{} {
	/*Method Peek has pointer to the Stack as receiver;
	if stack is empty, it returns nil; else it
	returns top item on the stack.*/
	if s.Length == 0 {
		return nil
	}
	return s.Top.Value
}

func (s *Stack) Pop() interface{} {
	/*Method Pop has pointer to the Stack as receiver;
	it checks if stack is empty; if not, it
	pops and returns top item from the stack.*/
	if s.Length == 0 {
		return nil
	}
	n := s.Top
	s.Top = n.Prev
	s.Length--
	return n.Value
}

func (s *Stack) Push(v interface{}) {
	/*Method Push has pointer to the Stack as receiver;
	it creates new Node, and pushes value on top of the stack.*/
	n := &Node{v, s.Top}
	s.Top = n
	s.Length++
}
