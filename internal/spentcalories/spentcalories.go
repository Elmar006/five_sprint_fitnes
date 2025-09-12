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
	// Разделяем строку на слайс строк
	parts := strings.Split(data, ",")

	// Проверяем, что длина слайса равна 3
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("неверный формат данных, ожидается три значения, получено %d", len(parts))
	}

	// Парсим шаги
	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка при парсинге шагов: %v", err)
	}

	// Проверяем, что количество шагов больше 0
	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("количество шагов должно быть положительным числом, получено %d", steps)
	}

	// Проверяем вид активности
	activity := parts[1]
	if activity == "" {
		return 0, "", 0, fmt.Errorf("неверный вид активности")
	}

	// Парсим продолжительность
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка при парсинге продолжительности: %v", err)
	}

	// Проверяем, что продолжительность положительная
	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("продолжительность должна быть положительной, получено %s", duration)
	}

	return steps, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	// Рассчитываем длину шага
	stepLength := height * stepLengthCoefficient
	// Вычисляем дистанцию в километрах
	distance := float64(steps) * stepLength / mInKm
	return distance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// Проверяем, что продолжительность больше 0
	if duration <= 0 {
		return 0
	}
	// Вычисляем дистанцию
	distance := distance(steps, height)
	// Переводим продолжительность в часы
	durationHours := duration.Hours()
	// Вычисляем среднюю скорость
	return distance / durationHours
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// Проверяем вес и рост
	if weight <= 0 {
		log.Println("вес должен быть положительным числом")
		return "", fmt.Errorf("вес должен быть положительным числом")
	}
	if height <= 0 {
		log.Println("рост должен быть положительным числом")
		return "", fmt.Errorf("рост должен быть положительным числом")
	}

	// Парсим входные данные
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// Вычисляем дистанцию и скорость
	dist := distance(steps, height)
	durationHours := duration.Hours()
	speed := meanSpeed(steps, height, duration)

	// Вычисляем калории в зависимости от типа тренировки
	var calories float64
	switch activity {
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
	default:
		log.Printf("неизвестный тип тренировки: %s", activity)
		return "", fmt.Errorf("неизвестный тип тренировки: %s", activity)
	}

	if err != nil {
		log.Println(err)
		return "", err
	}

	// Формируем результирующую строку
	info := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		activity, durationHours, dist, speed, calories,
	)

	return info, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверяем входные параметры
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

	// Рассчитываем среднюю скорость
	speed := meanSpeed(steps, height, duration)
	// Переводим продолжительность в минуты
	minutes := duration.Minutes()
	// Вычисляем калории
	calories := (weight * speed * minutes) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверяем входные параметры
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

	// Рассчитываем среднюю скорость
	speed := meanSpeed(steps, height, duration)
	// Переводим продолжительность в минуты
	minutes := duration.Minutes()
	// Вычисляем калории с учетом коэффициента
	calories := (weight * speed * minutes) / minInH * walkingCaloriesCoefficient
	return calories, nil
}
