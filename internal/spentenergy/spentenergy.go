package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// скопировал с 4 спринта
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	switch {
	case steps <= 0:
		return 0, errors.New("steps must be greater than 0")
	case weight <= 0 || weight > 500:
		return 0, errors.New("weight must be between 0 and 500 kg")
	case height <= 0.5 || height > 2.5:
		return 0, errors.New("height must be between 0.5 and 2.5 meters")
	case duration <= 0:
		return 0, errors.New("duration must be greater than 0")
	}
	averageSpeed := MeanSpeed(steps, height, duration)
	caloriesCount := ((weight * averageSpeed * duration.Minutes()) / 60.0) * walkingCaloriesCoefficient
	return caloriesCount, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	switch {
	case steps <= 0:
		return 0, errors.New("steps must be greater than 0")
	case weight <= 0 || weight > 500:
		return 0, errors.New("weight must be between 0 and 500 kg")
	case height <= 0.5 || height > 2.5:
		return 0, errors.New("height must be between 0.5 and 2.5 meters")
	case duration <= 0:
		return 0, errors.New("duration must be greater than 0")
	}
	averageSpeed := MeanSpeed(steps, height, duration)
	caloriesCount := (weight * averageSpeed * duration.Minutes()) / minInH
	return caloriesCount, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	if steps <= 0 {
		return 0
	}
	distanceVariable := Distance(steps, height)
	hoursDuration := duration.Hours() // Get the duration in hours as a float64
	averageSpeed := distanceVariable / hoursDuration
	return averageSpeed
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	returnVal := (float64(steps) * stepLength) / float64(mInKm)
	return returnVal
}
