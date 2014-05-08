package stream

import (
	"errors"
	"fmt"
)

type Stream struct {
	headValue   interface{}
	tailPromise STREAMFN
}

type STREAMFN func(interface{}, STREAMFN) *Stream
type ZIPFN func(interface{}, interface{}) interface{}
type MAPFN func(interface{}) interface{}
type FILTERFN func(interface{}) bool
type REDUCEFN func(interface{}, interface{}) interface{}
type WALKFN func(interface{}) interface{}

const (
	RANGE_OP_INC = 1
	RANGE_OP_DEC = -1
)

type RANGE_OP interface{}

func NewStream(head interface{}, tailPromise STREAMFN) *Stream {
	st := new(Stream)
	if head != nil {
		st.headValue = head
	}

	if tailPromise == nil {
		tailPromise = func(interface{}, STREAMFN) *Stream {
			return NewStream(nil, nil)
		}
	}

	st.tailPromise = tailPromise
	return st
}

func Make(v ...interface{}) *Stream {
	_len := len(v)
	if _len == 0 {
		return NewStream(nil, nil)
	}
	restArguments := v[1:_len]
	return NewStream(v[0], func(d interface{}, fn STREAMFN) *Stream {
		return Make(restArguments...)
	})
}

func FromArray(v []interface{}) *Stream {
	_len := len(v)
	if _len == 0 {
		return NewStream(nil, nil)
	}
	restArguments := v[1:_len]
	return NewStream(v[0], func(d interface{}, fn STREAMFN) *Stream {
		return FromArray(restArguments)
	})
}

func (this *Stream) Empty() bool {
	return this.headValue == nil
}

func (this *Stream) Head() (interface{}, error) {
	if this.Empty() {
		return nil, errors.New("Cannot get the head of the empty stream.")
	}
	return this.headValue, nil
}

func (this *Stream) Head1() interface{} {
	v, _ := this.Head()
	return v
}

func (this *Stream) Tail() (*Stream, error) {
	if this.Empty() == true {
		return nil, errors.New("Cannot get the tail of the empty stream.")
	}

	return this.tailPromise(this.headValue, this.tailPromise), nil
}

func (this *Stream) TailPromise() STREAMFN {
	return this.tailPromise
}

func (this *Stream) Tail1() *Stream {
	v, _ := this.Tail()
	return v
}

func (this *Stream) Item(n uint) (interface{}, error) {
	if this.Empty() == true {
		return nil, errors.New("Cannot use Item() on an empty stream.")
	}

	var st *Stream
	var err error

	st = this

	for n != 0 {
		n = n - 1
		st, err = st.Tail()
		if err != nil {
			return nil, errors.New("Item index does not exist in stream.")
		}
	}

	return st.Head()
}

func (this *Stream) Item1(n uint) interface{} {
	v, _ := this.Item(n)
	return v

}

func (this *Stream) Length() (int64, error) {
	var err error
	_len := int64(0)
	st := this
	for ; !st.Empty() ; st, err = st.Tail() {
		if err != nil {
			break
		}
		_len = _len + 1
	}
	return _len, nil
}

func (this *Stream) Length1() int64 {
	_len, _ := this.Length()
	return _len
}

/*
 * wtf?
 */
func (this *Stream) Add(s *Stream) *Stream {
	return this.Zip(func(x, y interface{}) interface{} {
		switch xv := x.(type) {
		case int:
			switch yv := y.(type) {
			case int:
				return xv + yv
			}
		case int64:
			switch yv := y.(type) {
			case int64:
				return xv + yv
			}
		case float64:
			switch yv := y.(type) {
			case float64:
				return xv + yv
			}
		case rune:
			switch yv := y.(type) {
			case rune:
				return xv + yv
			}
		}
		return 0
	}, s)
}

func (this *Stream) Append(s *Stream) *Stream {
	if this.Empty() {
		return s
	}
	self := this
	return NewStream(self.Head1(), func(v interface{}, sfn STREAMFN) *Stream {
		return self.Tail1().Append(s)
	})
}

func (this *Stream) Zip(f ZIPFN, s *Stream) *Stream {
	if this.Empty() {
		return s
	}
	if s.Empty() {
		return this
	}
	self := this
	return NewStream(f(s.Head1(), this.Head1()), func(v interface{}, sfn STREAMFN) *Stream {
		return self.Tail1().Zip(f, s.Tail1())
	})
}

func (this *Stream) Map(mfn MAPFN) *Stream {
	if this.Empty() {
		return this
	}
	self := this
	return NewStream(mfn(this.Head1()), func(v interface{}, fn STREAMFN) *Stream {
		return self.Tail1().Map(mfn)
	})
}

func (this *Stream) ConcatMap(fn MAPFN) *Stream {
	list := this.Reduce(func(a, x interface{}) interface{} {
		switch v := a.(type) {
		case *Stream:
			r := fn(x)
			switch vr := r.(type) {
			case *Stream:
				return v.Append(vr)
			}
		}
		return nil
	}, NewStream(nil, nil))
	l, ok := list.(*Stream)
	if ok {
		return l
	}
	return NewStream(nil, nil)
}

func (this *Stream) Reduce(aggregator REDUCEFN, initial interface{}) interface{} {
	if this.Empty() {
		return initial
	}

	return this.Tail1().Reduce(aggregator, aggregator(initial, this.Head1()))
}

func (this *Stream) Sum() interface{} {
	switch v := this.Head1().(type) {
	case int:
		return this.Reduce(func(a, b interface{}) interface{} {
			switch va := a.(type) {
			case int:
				switch vb := b.(type) {
				case int:
					return va + vb
				}
				break
			}
			return 0
		}, 0)
	case int64:
		return this.Reduce(func(a, b interface{}) interface{} {
			switch va := a.(type) {
			case int64:
				switch vb := b.(type) {
				case int64:
					return va + vb
				}
				break
			}
			return 0.0
		}, 0.0)
	case float64:
		return this.Reduce(func(a, b interface{}) interface{} {
			switch va := a.(type) {
			case float64:
				switch vb := b.(type) {
				case float64:
					return va + vb
				}
				break
			}
			return 0.0
		}, 0.0)
	default:
		return v
	}
}

func (this *Stream) Walk(fn WALKFN) {
	this.Map(func(x interface{}) interface{} {
		fn(x)
		return x
	}).Force()
}

func (this *Stream) Force() {
	var st *Stream
	var err error

	st = this
	for ; !st.Empty(); st, err = st.Tail() {
		if err != nil {
			break
		}
	}
}

func (this *Stream) Scale(factor interface{}) *Stream {
	return this.Map(func(a interface{}) interface{} {
		switch vfactor := factor.(type) {
		case int:
			switch va := a.(type) {
			case int:
				return va * vfactor
			}
		case int64:
			switch va := a.(type) {
			case int64:
				return va * vfactor
			}
		case float64:
			switch va := a.(type) {
			case float64:
				return va * vfactor
			}
		}
		return a
	})
}

func (this *Stream) Filter(ffn FILTERFN) *Stream {
	if this.Empty() {
		return this
	}
	h := this.Head1()
	t := this.Tail1()
	if ffn(h) {
		return NewStream(h, func(v interface{}, fn STREAMFN) *Stream {
			return t.Filter(ffn)
		})
	}
	return t.Filter(ffn)
}

func (this *Stream) Take(howmany int64) *Stream {
	if this.Empty() {
		return this
	}
	if howmany == 0 {
		return NewStream(nil, nil)
	}
	self := this
	return NewStream(this.Head1(), func(v interface{}, fn STREAMFN) *Stream {
		return self.Tail1().Take(howmany - 1)
	})
}

func (this *Stream) Drop(n int64) *Stream {
	self := this
	for ; n > 0; n-- {
		if self.Empty() {
			return NewStream(nil, nil)
		}
		self = self.Tail1()
	}
	return NewStream(self.Head1(), self.TailPromise())
}

func (this *Stream) Member(v interface{}) (bool, error) {
	var errt error
	self := this

	for ; !self.Empty() && errt == nil ; self, errt = self.Tail() {
		d, err := self.Head()
		if err != nil {
			break
		}
		if d == v {
			return true, nil
		}
	}

	return false, nil
}

func (this *Stream) Member1(v interface{}) bool {
	d, _ := this.Member(v)
	return d
}

func (this *Stream) ToString() {
}

func (this *Stream) Dump() {
	fmt.Printf("Dump: %v\n", this.Head1())
	v, err := this.Tail()
	if err == nil {
		v.Dump()
	}
}

func (this *Stream) Print(n int64) {
	var target *Stream
	if this.Empty() {
		target = NewStream(nil, nil)
	}
	target = this.Take(n)
	target.Walk(func(v interface{}) interface{} {
		fmt.Printf("%v\n", v)
		return v
	})
}

func (this *Stream) Equals(st *Stream) bool {
	if this.Empty() && st.Empty() {
		return true
	}
	if this.Empty() || st.Empty() {
		return false
	}
	a, a_err := this.Head()
	b, b_err := st.Head()

	if a_err == nil && b_err == nil && (a == b) {
		return this.Tail1().Equals(st.Tail1())
	}

	return false
}

// FIXME - need more than ints..
func Range(low, high interface{}) *Stream {
	return _range(RANGE_OP_INC, low, high)
}

func RangeL(low, high interface{}) *Stream {
	return _range(RANGE_OP_INC, low, high)
}

func RangeR(high, low interface{}) *Stream {
	return _range(RANGE_OP_DEC, high, low)
}

func _range(op RANGE_OP, low, high interface{}) *Stream {
	if low == high {
		return Make(low)
	}
	return NewStream(low, func(v interface{}, fn STREAMFN) *Stream {
		switch t := low.(type) {
		case float64:
			{
				switch op {
				case RANGE_OP_INC:
					return _range(op, t+1.0, high)
				case RANGE_OP_DEC:
					return _range(op, t-1.0, high)
				default:
					return NewStream(nil, nil)
				}
			}
		case int:
			{
				switch op {
				case RANGE_OP_INC:
					return _range(op, t+1, high)
				case RANGE_OP_DEC:
					return _range(op, t-1, high)
				default:
					return NewStream(nil, nil)
				}
			}
		case int64:
			{
				switch op {
				case RANGE_OP_INC:
					return _range(op, t+1, high)
				case RANGE_OP_DEC:
					return _range(op, t-1, high)
				default:
					return NewStream(nil, nil)
				}
			}
		case rune:
			{
				switch op {
				case RANGE_OP_INC:
					return _range(op, t+1, high)
				case RANGE_OP_DEC:
					return _range(op, t-1, high)
				default:
					return NewStream(nil, nil)
				}
			}
		default:
			return NewStream(nil, nil)
		}
	})
}
