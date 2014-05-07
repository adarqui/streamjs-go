package streamjs

import (
	"errors"
	"fmt"
)

type Stream struct {
	headValue interface{}
	tailPromise STREAMFN
}

type STREAMFN func(interface{}, STREAMFN) (*Stream)
type ZIPFN func(interface{}, interface{}) interface{}
type MAPFN func(interface{}) interface{}
type FILTERFN func(interface{}) bool

const (
	RANGE_OP_INC = 1
	RANGE_OP_DEC = -1
)

type RANGE_OP interface{}

func NewStream(head interface{}, tailPromise STREAMFN) (*Stream) {
	st := new (Stream)
	if head != nil {
		st.headValue = head
	}

	if tailPromise == nil {
		tailPromise = func(interface{}, STREAMFN) (*Stream) {
			return NewStream(nil, nil)
		}
	}

	st.tailPromise = tailPromise
	return st
}

func Make(v... interface{}) (*Stream) {
	_len := len(v)
	if _len == 0 {
		return NewStream(nil, nil)
	}
	restArguments := v[1:_len]
	return NewStream(v[0], func(d interface{}, fn STREAMFN) *Stream {
		return Make(restArguments...)
	})
}

func FromArray(v []interface{}) (*Stream) {
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
		return nil, errors.New("Cannot get the tail of the empty stream.");
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

func (this *Stream) Length () (uint, error) {
	var err error
	_len := uint(0)
	st := this
	for {
		v := st.Empty()
		if v == true {
			break
		}
		_len = _len + 1
		st, err = st.Tail()
		if err != nil {
			return _len, err
		}
	}
	return _len, nil
} 

func (this *Stream) Length1() uint {
	_len, _ := this.Length()
	return _len
}


/*
 * wtf?
 */
func (this *Stream) Add(s *Stream) interface{} {
	return this.Zip(func(x, y interface{}) interface{} {
		switch xv := x.(type) {
			case int: {
				switch yv := y.(type) {
					case int: {
						v := xv + yv
						return v
					}
				}
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

func (this *Stream) Zip(f ZIPFN, s *Stream) (*Stream) {
	if this.Empty() {
		return s
	}
	if s.Empty() {
		return this
	}
	self := this
	return NewStream(f(s.Head1(), this.Head1()), func(v interface{}, sfn STREAMFN) (*Stream) {
		return self.Tail1().Zip(f, s.Tail1() )
	})
}

func (this *Stream) Map(mfn MAPFN) (*Stream) {
	if this.Empty() {
		return this
	}
	self := this
	return NewStream(mfn(this.Head1()), func(v interface{}, fn STREAMFN) *Stream {
		return self.Tail1().Map(mfn)
	})
}

func (this *Stream) ConcatMap() {
}

func (this *Stream) Reduce() {
}

func (this *Stream) Sum() {
}

func (this *Stream) Walk() {
}

func (this *Stream) Force() {
}

func (this *Stream) Scale() {
}

func (this *Stream) Filter(ffn FILTERFN) *Stream {
	if this.Empty() {
		return this
	}
	h := this.Head1()
	t := this.Tail1()
	if ffn (h) {
		return NewStream(h, func(v interface{}, fn STREAMFN) *Stream {
			return t.Filter(ffn)
		})
	}	
	return t.Filter(ffn)
}

func (this *Stream) Take(howmany int) *Stream {
	if this.Empty() {
		return this
	}
	if howmany == 0 {
		return NewStream(nil, nil)
	}
	self := this
	return NewStream(this.Head1(), func(v interface{}, fn STREAMFN) *Stream {
		return self.Tail1().Take(howmany - 1);
	})
}

func (this *Stream) Drop(n int) *Stream {
	self := this
	for ; n > 0 ; n-- {
		if self.Empty() {
			return NewStream(nil, nil)
		}
		self = self.Tail1()
	}
	return NewStream(self.Head1(), self.TailPromise())
}

func (this *Stream) Member(v interface{}) (bool, error) {
	self := this

	for {
		if self.Empty() {
			break
		}
		d, err := self.Head()
		if err != nil {
			break
		}
		if d == v {
			return true, nil
		}
		self, err = self.Tail()
		if err != nil {
			break
		}
	}

	return false, nil
}

func (this *Stream) Member1(v interface{}) bool {
	d, _ := this.Member(v)
	return d
}

func (this *Stream) Print() {
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

func RangeL (low, high interface{}) *Stream {
	return _range(RANGE_OP_INC, low, high)
}

func RangeR (high, low interface{}) *Stream {
	return _range(RANGE_OP_DEC, high, low)
}

func _range (op RANGE_OP, low, high interface{}) *Stream {
	if low == high {
		return Make(low)
	}
	return NewStream(low, func(v interface{}, fn STREAMFN) *Stream {
		switch t := low.(type) {
			case float64: {
				switch op {
					case RANGE_OP_INC:
						return _range(op, t+1, high)
					case RANGE_OP_DEC:
						return _range(op, t-1, high)
					default:
						return NewStream(nil, nil)
				}
			}
			case int: {
				switch op {
					case RANGE_OP_INC:
						return _range(op, t+1, high)
					case RANGE_OP_DEC:
						return _range(op, t-1, high)
					default:
						return NewStream(nil, nil)
				}
			}
			case rune: {
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

