package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	// Parse принимает строку данных и наполняет структуру
	Parse(data string) error
	// ActionInfo формирует строку с информацией об активности
	ActionInfo() (string, error)
}

// Функция принимает слайс строк с данными о тренировках или прогулках и
// экземпляр одной из ваших структур Training или DaySteps.
// Это возможно, потому что они обе реализуют интерфейс DataParser,
// то есть для каждой из этих структур вы реализовали методы, которые описаны в интерфейсе.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("parse error: %v", err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("action info error: %v", err)
			continue
		}
		if info != "" {
			fmt.Println(info)
		}
	}
}
