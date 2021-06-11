package chapter7

import "fmt"

// 代理模式
//	为其他对象提供一种代理，以控制对该对象的访问
//	这个话容易产生歧义，代理代表被代理方，访问目标方。也就是说代理可以模拟被代理方的行为，而目标方无法反向对被代理方进行操作

type Chasing interface {
	GiveFlower() string
	GiveDolls() string
	GiveChocolate() string
}

type Target struct {
	Name string
}

func NewTarget(name string) Target {
	return Target{Name: name}
}

type Pursuit struct {
	Name   string
	Target Target
}

func NewPursuit(name string, t Target) *Pursuit {
	return &Pursuit{
		Name:   name,
		Target: t,
	}
}

func (p Pursuit) GiveFlower() string {
	return fmt.Sprintf("%s send flower to %s\n", p.Name, p.Target.Name)
}
func (p Pursuit) GiveDolls() string {
	return fmt.Sprintf("%s send dolls to %s\n", p.Name, p.Target.Name)
}
func (p Pursuit) GiveChocolate() string {
	return fmt.Sprintf("%s send chocolate to %s\n", p.Name, p.Target.Name)
}

type Proxy struct {
	c Chasing
}

func NewProxy(s string, t Target) Chasing {
	return &Proxy{c: NewPursuit(s, t)}
}
func (p Proxy) GiveFlower() string {
	return p.c.GiveFlower()
}
func (p Proxy) GiveDolls() string {
	return p.c.GiveDolls()
}
func (p Proxy) GiveChocolate() string {
	return p.c.GiveChocolate()
}
