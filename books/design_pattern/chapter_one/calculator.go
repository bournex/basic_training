package chapterone

type Operation interface {
	SetA(float64)
	SetB(float64)
	GetResult() float64
}

type OperationBase struct {
	NumberA float64
	NumberB float64
}

func (ob *OperationBase) SetA(a float64) {
	ob.NumberA = a
}
func (ob *OperationBase) SetB(b float64) {
	ob.NumberB = b
}

type OperationAdd struct {
	OperationBase
}

func (oa OperationAdd) GetResult() float64 {
	return oa.NumberA + oa.NumberB
}

type OperationSub struct {
	OperationBase
}

func (os OperationSub) GetResult() float64 {
	return os.NumberA - os.NumberB
}

type OperationMul struct {
	OperationBase
}

func (os OperationMul) GetResult() float64 {
	return os.NumberA * os.NumberB
}

type OperationDiv struct {
	OperationBase
}

func (os OperationDiv) GetResult() float64 {
	if os.NumberB == 0 {
		panic("divisor can't be zero!")
	}
	return os.NumberA / os.NumberB
}

type OperationFactory struct {
}

func (of OperationFactory) createOperation(operate byte) Operation {
	switch operate {
	case '+':
		return new(OperationAdd)
	case '-':
		return new(OperationSub)
	case '*':
		return new(OperationMul)
	case '/':
		return new(OperationDiv)
	default:
		panic("only support +-*/ operations")
	}
}
