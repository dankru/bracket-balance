package main

import (
	"fmt"
	myStack "projects/bracket-balance/myStack"
)

// Глобальный словарь скобок для проверки на соответствие
var m = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

func main() {
	fmt.Println(isBalanced("(][)"))
	fmt.Println(isBalanced("qg"))
	fmt.Println(isBalanced(""))
	fmt.Println(isBalanced(")("))
	fmt.Println(isBalanced("[()]"))
}

// функция isBalanced проходит по каждому символу строки и с помощью стека проверяет, состоит ли строка из сбалансированных скобок
func isBalanced(inputStr string) string {
	var notBalanced = "Скобки не сбалансированы"

	// В случае если строка пустая, или длина нечётная, отбрасываем решение
	if inputStr == "" || len(inputStr)%2 != 0 {
		return notBalanced
	}

	// Структура stack определена в пакете myStack
	var stack myStack.Stack

	for _, bracket := range inputStr {

		_, ok := m[bracket]
		if ok {
			// При нахождении открывающей скобки пушим в стек соответствующую ей закрывающую скобку
			stack.Push(m[bracket])
		} else if !ok {
			// Если скобка закрывающая, вытаскиваем из стека последнюю скобку и проверяем их на соответствие
			if bracket == stack.Pop() {
				continue
			} else {
				return notBalanced
			}
		}
	}
	// Если стек оказался пуст после прохода по строке, значит вложенность и количество открывающих и закрывающих скобок валидные
	if len(stack.Items) == 0 {
		return "Скобки сбалансированы"
	}
	return notBalanced
}
