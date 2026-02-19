package main

import (
	"fmt"

	"lab4_variant04/pkg/loancalc"

	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
	"github.com/google/uuid"
)

func main() {
	// Создаём уникальный ID кредита
	loanID := uuid.New()

	// Параметры кредита
	sum := 500000.0       // сумма кредита
	annualPercent := 12.0 // годовая ставка %
	months := 24          // срок в месяцах

	// Рассчитываем ежемесячный платёж
	payment, err := loancalc.MonthlyPayment(sum, annualPercent, months)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Printf("ID кредита: %s\n", loanID)
	fmt.Printf("Сумма кредита: %s руб.\n", humanize.Commaf(sum))
	fmt.Printf("Платёж до досрочного погашения: %.2f руб.\n", payment)

	// Применяем досрочный платёж
	err = loancalc.ApplyEarlyPayment(&payment, 1000)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Printf("Платёж после досрочного погашения: %.2f руб.\n", payment)

	// Формируем отчёт
	report, err := loancalc.FormatLoanReport("Иван Иванов", payment, months)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Цветной вывод отчёта
	color.Green("\n--- ОТЧЁТ ПО КРЕДИТУ ---")
	fmt.Println(report)
}
