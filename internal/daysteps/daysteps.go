package daysteps

import (
	"fmt"
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
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")

	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("Ошибка! Длина слайса больше 2!")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("Ошибка при парсинге шагов! %v", err)
	}

	if steps <= 0 {
		return 0, 0, fmt.Errorf("Шагов меньше 0!")
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("Ошибка при парсинге продолжительности %v", err)
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println("Error", err)
		return ""
	}

	if steps <= 0 {
		return ""
	}
	distance := float64(steps) * stepLength / mInKm

	// Правильный вызов функции с четырьмя параметрами
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		fmt.Println("Error calculating calories:", err)
		return ""
	}
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", steps, distance, calories)
	return result
}
