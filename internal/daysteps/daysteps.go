package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// количество шагов
	Steps int
	// длительность прогулки
	Duration time.Duration
	// персональные данные пользователя
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// формат: "678,0h50m"
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("invalid format: expected 2 parts")
	}
	stepsStr := parts[0]
	durStr := parts[1]

	var steps int
	steps, err = strconv.Atoi(stepsStr)
	if err != nil || steps <= 0 {
		return errors.New("invalid steps value")
	}
	// duration parse; must be >0
	d, err := time.ParseDuration(durStr)
	if err != nil || d <= 0 {
		return errors.New("invalid duration")
	}

	ds.Steps = steps
	ds.Duration = d
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 || ds.Duration <= 0 || ds.Weight <= 0 || ds.Height <= 0 {
		return "", errors.New("invalid input data")
	}

	dist := spentenergy.Distance(ds.Steps, ds.Height)
	cal, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	// многострочная строка одной строкой, а не как в trainings
	out := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, cal)
	return out, nil
}
