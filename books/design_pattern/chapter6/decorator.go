package chapter6

// 装饰器模式（Decorator）
//	动态地给一个对象添加一些额外的职责，就增加功能而言，装饰模式比生成子类更为灵活
//	对于已经存在的类型，在不侵入其原有实现的情况下，实现对其对象的修饰功能，同时还能维持一定的顺序
//	简单的讲，装饰器抽象类和实体类都从被装饰对象派生，通过某个Set接口层层包装被装饰对象

type AbstractPerson interface {
	Show() string
}

type Person struct {
	Name string
}

func (p Person) Show() string {
	return p.Name + " wearing"
	// fmt.Print(p.Name + " wearing")
}

// 装饰器模式不太理想的一点，装饰器基类或实体类需要从被装饰的类型派生，比如Person，这里需要从代码组织上进行识别，不然会造成理解上的困难
type Decorator struct {
	P AbstractPerson
}

func (d *Decorator) Decorate(p *Person) {
	d.P = p
}

func (d Decorator) Show() string {
	if d.P != nil {
		return d.P.Show()
	}
	return ""
}

type TShirt struct {
	Decorator
}

func (t TShirt) Show() string {
	return t.Decorator.Show() + " t-shirt"
	//fmt.Print(" t-shirt")
}

type Shorts struct {
	Decorator
}

func (s Shorts) Show() string {
	return s.Decorator.Show() + " shorts"
}
