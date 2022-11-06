package operations

import (
	"github.com/narutopig/neon-lang/value"
)

func Divide(left value.Value, right value.Value) value.Value {
	if left.Type == value.Number {
		lval := floatFromBytes(left.Data)

		if right.Type == value.Number {
			rval := floatFromBytes(right.Data)

			return value.NewNumber(lval / rval)
		}
	}

	invalidOp(left.Type, right.Type, "/")

	return value.Value{}
}
