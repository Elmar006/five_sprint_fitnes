package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// Разделяем строку на слайс строк
	parts := strings.Split(data, ",")

	// Проверяем, что длина слайса равна 2
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("ошибка: ожидается два значения, получено %d", len(parts))
	}

	// Парсим количество шагов, учитывая возможные пробелы
	trimmedSteps := parts[0]
	steps, err := strconv.Atoi(trimmedSteps)
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка при парсинге шагов: %v", err)
	}

	// Проверяем, что количество шагов больше 0
	if steps <= 0 {
		return 0, 0, fmt.Errorf("ошибка: количество шагов должно быть положительным, получено %d", steps)
	}

	// Парсим продолжительность
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка при парсинге продолжительности: %v", err)
	}

	// Проверяем, что продолжительность положительная
	if duration <= 0 {
		return 0, 0, fmt.Errorf("ошибка: продолжительность должна быть положительной, получено %s", duration)
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// Получаем шаги и продолжительность с помощью parsePackage
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Printf("Error: %v", err)
		return ""
	}

	// Вычисляем дистанцию в километрах
	distance := float64(steps) * stepLength / mInKm

	// Вычисляем потраченные калории
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Printf("Error calculating calories: %v", err)
		return ""
	}

	// Формируем результирующую строку
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, calories)
}
