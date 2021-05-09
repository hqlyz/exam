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
		// 一年的还款周期
		for i := 0; i < 12; i++ {
			// 第一个还款月不算，每3个月还款数增加
			if totalMonths != 0 && i%3 == 0 {
				monthlyPay = monthlyPay * morePayRate
			}
			// 增加总还款数
			totalPay += monthlyPay
			// 增加每年累计还款数
			annualPay += monthlyPay
			// 增加总还款月
			totalMonths++
			// 如果当年累计还款数已经超过未还金额，说明已经还完
			if annualPay >= outstandingPrincipal {
				// 计算最后一个月还款数的偏差值
				lastPayOffset = annualPay - outstandingPrincipal
				break L
			}
		}
		// 重新计算未还金额
		outstandingPrincipal = outstandingPrincipal*annualInterest + (outstandingPrincipal - annualPay)
		// 一年还完后将每年累计还款金额清零
		annualPay = 0
	}
	log.Printf("It will take %d months", totalMonths)
	log.Printf("It will pay %.2f in total", totalPay-lastPayOffset)
}
