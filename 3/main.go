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
	morePayRate float64 = 0.05
	// 总还款金额
	totalPay float64 = 0
	// 总还款月数
	totalMonths int64 = 0
)

func main() {
	// 只要最终还没还完，就需要继续还
	for outstandingPrincipal > 0 {
		totalMonths++
		outstandingPrincipal -= monthlyPay
		totalPay += monthlyPay

		// 每过3个月增加还款金额
		if totalMonths%3 == 0 {
			monthlyPay += monthlyPay * morePayRate
		}

		// 每过12个月银行计算利息
		if totalMonths%12 == 0 {
			outstandingPrincipal += outstandingPrincipal * annualInterest
		}
	}
	log.Printf("It will take %d months", totalMonths)
	log.Printf("It will pay %.2f in total", totalPay+outstandingPrincipal)
}
