package chaptertwo

import "math"

const (
	StrategyTypeOne   = "strategy_one"
	StrategyTypeTwo   = "strategy_two"
	StrategyTypeThree = "strategy_three"
)

type CashContext struct {
	sa StrategyAbstract
}

func NewCashContext(t string) *CashContext {
	cc := new(CashContext)
	switch t {
	case StrategyTypeOne:
		cc.sa = newStrategyOne() // 原价
	case StrategyTypeTwo:
		cc.sa = newStrategyTwo(0.85) // 85折
	case StrategyTypeThree:
		cc.sa = newStrategyThree(200, 50) // 满200减50
	}
	return cc
}

func (cc CashContext) AcceptCash(money float64) float64 {
	return cc.sa.AcceptCash(money)
}

type StrategyAbstract interface {
	AcceptCash(money float64) float64
}

type StrategyOne struct {
}

func (so StrategyOne) AcceptCash(money float64) float64 {
	return money
}

func newStrategyOne() *StrategyOne {
	return new(StrategyOne)
}

type StrategyTwo struct {
	factor float64 // 折扣
}

func (st StrategyTwo) AcceptCash(money float64) float64 {
	return money * st.factor
}

func newStrategyTwo(factor float64) *StrategyTwo {
	return &StrategyTwo{factor: factor}
}

type StrategyThree struct {
	candidate float64 // 满减额度
	moneyRet  float64 // 满减金额
}

func (st StrategyThree) AcceptCash(money float64) float64 {
	if money > st.candidate {
		return money - math.Floor(money/st.candidate)*st.moneyRet
	}
	return money
}

func newStrategyThree(candidate, moneyRet float64) *StrategyThree {
	return &StrategyThree{candidate: candidate, moneyRet: moneyRet}
}
