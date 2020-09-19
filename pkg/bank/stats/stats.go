package stats

import (
	"github.com/Shahlojon/bank/pkg/bank/types"
)

//Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) (avg types.Money) {
	n := types.Money(len(payments))
	for _, payment := range payments {
		avg += (payment.Amount)
	}
	avg = avg/n
	return avg
}