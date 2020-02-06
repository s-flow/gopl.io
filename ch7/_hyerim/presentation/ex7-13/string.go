package eval

import "fmt"

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	//실수에서 지수가 크면 %e, %E로 표현하고, 지수가 작으면 %f, %F로 표현합니다.
	// fmt.Printf("%g %g\n", 1234567.1234567, 1.2) // 1.2345671234567e+06 1.2
	return fmt.Sprintf("%g", float64(l))
}

func (u unary) String() string {
	return fmt.Sprintf("(%c%s)", u.op, u.x.String())
}

func (b binary) String() string {
	return fmt.Sprintf("(%s %c %s)", b.x.String(), b.op, b.y.String())
}

func (c call) String() string {
	switch c.fn {
	case "pow":
		return fmt.Sprintf("%s(%s, %s)", c.fn, c.args[0].String(), c.args[1].String())
	case "sin":
	case "sqrt":
		return fmt.Sprintf("%s(%s)", c.fn, c.args[0].String())
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}
