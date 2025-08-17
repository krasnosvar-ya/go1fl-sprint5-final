package personaldata

import "fmt"

type Personal struct {
	// TODO: добавить поля
	// Создайте экспортируемую структуру Personal, у которой три поля:
	// Name — имя пользователя;
	// Weight— вес пользователя;
	// Height — рост пользователя.
	// Эта структура будет встроена в другие структуры.
	// Для этой структуры нужно будет создать метод, который будет выводить данные структуры на экран.
	Name   string
	Weight float64
	Height float64
}

func (p Personal) Print() {
	// TODO: реализовать функцию
	fmt.Printf("Имя: %s\n", p.Name)
	fmt.Printf("Вес: %.2f кг.\n", p.Weight)
	fmt.Printf("Рост: %.2f м.\n", p.Height)
}
