package chapter1

import (
	"bufio"
	"errors"
	"strconv"
	"testing"
)

func TestCalculator(t *testing.T) {
	testCases := []struct {
		A      float64
		B      float64
		Opcode byte
		Expect float64
		desc   string
	}{
		{
			A:      3,
			B:      5,
			Opcode: '+',
			Expect: 8,
			desc:   "TESTING ADD",
		},
		{
			A:      3,
			B:      5,
			Opcode: '-',
			Expect: -2,
			desc:   "TESTING SUB",
		},
		{
			A:      3,
			B:      -5,
			Opcode: '*',
			Expect: -15,
			desc:   "TESTING MUL",
		},
		{
			A:      15,
			B:      -5,
			Opcode: '/',
			Expect: -3,
			desc:   "TESTING DIV",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// reader := bufio.NewReader(os.Stdin)
			// var (
			// 	a   float64
			// 	b   float64
			// 	o   byte
			// 	err error
			// )
			// if a, err = getFloat(reader); err != nil {
			// 	t.Errorf("read a error, %s", err.Error())
			// 	return
			// }
			// if o, err = getOperation(reader); err != nil {
			// 	t.Errorf("read a error, %s", err.Error())
			// 	return
			// }
			// if b, err = getFloat(reader); err != nil {
			// 	t.Errorf("read b error, %s", err.Error())
			// 	return
			// }

			operation := OperationFactory{}.createOperation(tC.Opcode)
			operation.SetA(tC.A)
			operation.SetB(tC.B)
			result := operation.GetResult()
			if result != tC.Expect {
				t.Errorf("input a %f, b %f, expect %f, got %f", tC.A, tC.B, tC.Expect, result)
				return
			}
		})
	}
}

func getFloat(reader *bufio.Reader) (float64, error) {
	s, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

func getOperation(reader *bufio.Reader) (byte, error) {
	s, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	if len(s) != 1 || (s[0] != '+' && s[0] != '-' && s[0] != '*' && s[0] != '/') {
		return 0, errors.New("unknown operation")
	}

	return s[0], nil
}
