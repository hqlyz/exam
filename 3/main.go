package main

import "log"

var (
	// 未偿还金额
	outstandingPrincipal float64 = 1000000
	// 每月偿还金额
	monthlyPay float64 = 1500
	// 年利率
	annualInterest float64 = 0.03
	// 每3个月多支付的比率
	morePayRate float64 = 1.05
	// 总还款金额
	totalPay float64 = 0
	// 总还款月数
	totalMonths int64 = 0
	// 每年还款总额
	annualPay float64 = 0
	// 最后一月还款差额修正
	lastPayOffset float64 = 0
)

func main() {
L:
	for outstandingPrincipal > monthlyPay {
		for i := 0; i < 12; i++ {
			if totalMonths != 0 && i%3 == 0 {
				monthlyPay = monthlyPay * morePayRate
			}
			totalPay += monthlyPay
			annualPay += monthlyPay
			totalMonths++
			if annualPay >= outstandingPrincipal {
				lastPayOffset = annualPay - outstandingPrincipal
				break L
			}
		}
		annualPay = 0
		outstandingPrincipal = outstandingPrincipal*annualInterest + (outstandingPrincipal - annualPay)
	}
	log.Printf("It will take %d months", totalMonths)
	log.Printf("It will pay %.2f in total", totalPay - lastPayOffset)
}
