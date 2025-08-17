package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps                 int
	TrainingType          string
	Duration              time.Duration
	personaldata.Personal // импортированная структура
}

func (t *Training) Parse(datastring string) (err error) {
	// формат: "3456,Ходьба,3h00m"
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("неверный формат")
	}
	stepsStr := parts[0]
	kind := parts[1]
	durStr := parts[2]

	// минимальные проверки: парсим число шагов и длительность
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return errors.New("invalid steps")
	}
	if steps <= 0 {
		return errors.New("invalid steps value")
	}
	d, err := time.ParseDuration(durStr)
	if err != nil || d <= 0 {
		return errors.New("invalid duration")
	}

	t.Steps = steps
	t.TrainingType = kind
	t.Duration = d
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// валидация исходных данных
	if t.Steps <= 0 || t.Duration <= 0 || t.Weight <= 0 || t.Height <= 0 {
		return "", errors.New("неверные входные данные")
	}

	distancia := spentenergy.Distance(t.Steps, t.Height)
	avgSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var cal float64
	var err error
	// через switch для разнообразия
	switch t.TrainingType {
	case "Бег":
		cal, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		cal, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}

	// вывод был правлиьный но на чтото ругались тесты ( на формат, но почему не понял)
	// 	// https://golang.cafe/blog/how-to-write-multiline-string-in-go-golang
	// 	text := `Тип тренировки: %s
	// Длительность: %f ч.
	// Дистанция: %f км.
	// Скорость: %f км/ч
	// Сожгли калорий: %f
	// `
	// 	fmt.Printf(text, t.TrainingType, t.Duration.Hours(), distancia, avgSpeed, cal)

	text := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distancia,
		avgSpeed,
		cal,
	)
	return text, nil
}
