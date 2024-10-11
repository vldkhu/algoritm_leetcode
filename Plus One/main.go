package main

import (
	"fmt"
)

// Функция plusOne принимает массив цифр и увеличивает число на 1.
func plusOne(digits []int) []int {
	// Получаем длину массива digits
	n := len(digits)

	// Проверяем, состоит ли массив только из цифр 9
	allNines := true
	for _, digit := range digits {
		if digit != 9 {
			allNines = false // Если хотя бы одна цифра не 9, устанавливаем флаг в false
			break            // Выходим из цикла, так как нам больше не нужно проверять
		}
	}

	// Если все цифры были 9, создаем новый массив длиной на 1 больше
	if allNines {
		// Создаем новый массив с первым элементом 1 и остальными нулями
		result := make([]int, n+1)
		result[0] = 1 // Устанавливаем первый элемент в 1
		return result // Возвращаем новый массив
	}

	// Начинаем с последнего элемента массива и двигаемся к началу
	for i := n - 1; i >= 0; i-- {
		// Если текущая цифра меньше 9, просто увеличиваем её на 1
		if digits[i] < 9 {
			digits[i]++   // Увеличиваем текущую цифру на 1
			return digits // Возвращаем измененный массив
		}
		// Если текущая цифра равна 9, устанавливаем её в 0
		digits[i] = 0
	}

	// Если мы дошли до этого места, значит все цифры были 9
	// Добавляем 1 в начало массива
	return append([]int{1}, digits...) // Возвращаем новый массив с 1 в начале и остальными цифрами
}

func main() {
	// Пример использования функции
	digits := []int{9}        // Исходный массив
	result := plusOne(digits) // Увеличиваем число на 1
	fmt.Println(result)       // Ожидаемый вывод: [1, 0]
}
