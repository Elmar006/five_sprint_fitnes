package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")

	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("Неверный формат данных, переданных в слайс")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("%v", err)
	}

	aktive := parts[1]
	if aktive == "" {
		return 0, "", 0, fmt.Errorf("Неверный вид активности")
	}

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("%v", err)
	}

	return steps, aktive, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	distance := float64(steps) * stepLength / mInKm

	return distance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	distance := distance(steps, height)
	durationHours := duration.Hours()

	avarageSpeed := distance / durationHours

	return avarageSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, active, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", nil
	}

	if weight <= 0 {
		return "", fmt.Errorf("Вес должен быть положительным числом")
	}
	if height <= 0 {
		return "", fmt.Errorf("Рост должен быть положительным числом")
	}

	dist := distance(steps, height)
	durationHours := duration.Hours()
	speed := meanSpeed(steps, height, duration)

	var calories float64
	switch active {
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", active)
	}

	if err != nil {
		log.Println(err)
		return "", nil
	}

	info := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		active, durationHours, dist, speed, calories,
	)

	return info, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть положительным числом")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть положительным числом")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть положительным числом")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть положительной")
	}

	speed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	calories := (weight * speed * minutes) / minInH

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть положительным числом")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть положительным числом")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть положительным числом")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть положительной")
	}
	avSpeed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	userWeight := weight * avSpeed * minutes
	resultCal := userWeight / minInH
	coeffCalories := resultCal * walkingCaloriesCoefficient

	return coeffCalories, nil
}
