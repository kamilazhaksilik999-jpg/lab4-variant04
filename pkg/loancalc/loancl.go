// Package loancalc предоставляет функции для расчёта кредита:
// вычисление ежемесячной ставки, платежа,
// применение досрочного погашения и формирование отчёта.
package loancalc

import (
	"fmt"
	"math"
)

// MonthlyRate вычисляет ежемесячную процентную ставку
// из годовой процентной ставки.
func MonthlyRate(annualPercent float64) (float64, error) {
	if annualPercent < 0 {
		return 0, fmt.Errorf("годовая ставка не может быть отрицательной")
	}
	return annualPercent / 12 / 100, nil
}

// MonthlyPayment рассчитывает аннуитетный ежемесячный платёж.
// sum — сумма кредита
// annualPercent — годовая ставка
// months — срок в месяцах
func MonthlyPayment(sum, annualPercent float64, months int) (float64, error) {
	if sum <= 0 {
		return 0, fmt.Errorf("сумма кредита должна быть больше 0")
	}
	if annualPercent < 0 {
		return 0, fmt.Errorf("годовая ставка не может быть отрицательной")
	}
	if months <= 0 {
		return 0, fmt.Errorf("срок кредита должен быть больше 0")
	}

	rate, err := MonthlyRate(annualPercent)
	if err != nil {
		return 0, err
	}

	if rate == 0 {
		return sum / float64(months), nil
	}

	payment := sum * (rate * math.Pow(1+rate, float64(months))) /
		(math.Pow(1+rate, float64(months)) - 1)

	return payment, nil
}

// ApplyEarlyPayment уменьшает платёж на сумму досрочного погашения.
// payment — указатель на ежемесячный платёж
// extra — сумма досрочного платежа
func ApplyEarlyPayment(payment *float64, extra float64) error {
	if payment == nil {
		return fmt.Errorf("указатель на платеж равен nil")
	}
	if extra < 0 {
		return fmt.Errorf("досрочный платеж не может быть отрицательным")
	}
	if extra > *payment {
		return fmt.Errorf("досрочный платеж не может превышать ежемесячный платёж")
	}

	*payment -= extra
	return nil
}

// FormatLoanReport формирует строку отчёта по кредиту.
// client — имя клиента
// payment — ежемесячный платёж
// months — срок кредита
func FormatLoanReport(client string, payment float64, months int) (string, error) {
	if client == "" {
		return "", fmt.Errorf("имя клиента не может быть пустым")
	}
	if payment <= 0 {
		return "", fmt.Errorf("платёж должен быть положительным")
	}
	if months <= 0 {
		return "", fmt.Errorf("срок кредита должен быть больше 0")
	}

	report := fmt.Sprintf(
		"Клиент: %s\nЕжемесячный платёж: %.2f руб.\nСрок: %d месяцев\n",
		client,
		payment,
		months,
	)

	return report, nil
}
