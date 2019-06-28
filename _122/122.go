package _122

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	var total int
	var index = -1 //买入的索引
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] < prices[i] {
			// 如果明天比今天少, 则第i天执行卖出
			if index >= 0 {
				total += prices[i] - prices[index]
				// 重新计数
				index = -1
			}
		} else {
			// 如果从来没有买入过, 则执行买入
			if index == -1 {
				index = i
			}
		}
	}

	// 如果买入, 但是一直是上升势头, 则最后需要再计算后面的收益
	if index != -1 && index < len(prices)-1 {
		total += prices[len(prices)-1] - prices[index]
	}
	return total
}
