package math

type Knapsack struct {
	Num       int
	WeightArr []int
	PriceArr  []int
	MaxWeight int
	MaxPrice  int
	I         int
	mem       []bool
}

func (k *Knapsack) Run() {
	k.Init()
	k.re(0, 0, 0)
}

//动态规划
func (k *Knapsack) DpRun() {
	k.Init()
	k.dp()
}

func (k *Knapsack) Init() {
	k.I = 0
	k.MaxPrice = 0
	k.mem = []bool{}
}

func (k *Knapsack) re(index int, curWeight int, curPrice int) {
	if curWeight == k.MaxWeight || index == k.Num {
		if curPrice > k.MaxPrice {
			k.MaxPrice = curPrice
		}
		return
	}
	k.I++
	// 不选择选择当前物品
	k.re(index+1, curWeight, curPrice)
	// 选择当前物品
	if curWeight+k.WeightArr[index] <= k.MaxWeight {
		k.re(index+1, curWeight+k.WeightArr[index], curPrice+k.PriceArr[index])
	}
}

func (k *Knapsack) dp() {
	var states []int
	for i := 0; i <= k.MaxWeight; i++ {
		states = append(states, 0)
	}
	// 循环状态转移
	for i := 1; i < k.Num; i++ {
		for j := k.MaxWeight; j >= k.WeightArr[i]; j-- {
			states[j] = max(states[j], states[j-k.WeightArr[i]]+k.PriceArr[i])
			k.MaxPrice = max(k.MaxPrice, states[j])
			k.I++
		}
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
