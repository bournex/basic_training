package chapterone

// 工厂模型
//	面向对象的三个核心：封装、继承、多态
//	封装通过struct实现
//	继承与c++、java稍有区别，go不支持struct直接从其他struct继承，而是子类内嵌父类的方式，继承父类的方法和成员
//	多态通过interface和struct实现

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
