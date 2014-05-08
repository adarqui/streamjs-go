package streamjs

import (
	"fmt"
	"math"
	"testing"
)

func Test_NewStream(t *testing.T) {
	st := NewStream("1", nil)
	if l, _ := st.Length(); l != 1 {
		t.Error("Test_NewStream: st.Length() != 1")
	}

	if v, _ := st.Head(); v != "1" {
		t.Error("Test_NewStream: st.Head() != '1'")
	}

	st2 := NewStream("1", func(v interface{}, sfn STREAMFN) *Stream {
		return NewStream("2", nil)
	})
	if l, _ := st2.Length(); l != 2 {
		t.Error("Test_NewStream: st2.Length() != 2")
	}
	if v, _ := st2.Head(); v != "1" {
		t.Error("Test_NewStream: st2.Head() != '1'")
	}
	if v, _ := st2.Item(1); v != "2" {
		t.Error("Test_NewStream: st2.Item(1) != '2'")
	}
}

func Test_SJS_Make(t *testing.T) {
	st := Make("10", "20", "30", "40")
	if v, _ := st.Length(); v != 4 {
		t.Error("Test_SJS_1: st.Length() != 4.", v)
	}
	if v, _ := st.Head(); v != "10" {
		t.Error("Test_SJS_1: st.Head() != '10'")
	}
	if v, _ := st.Item(0); v != "10" {
		t.Error("Test_SJS_1: st.Item(0) != '10'")
	}
	if v, _ := st.Item(1); v != "20" {
		t.Error("Test_SJS_1: st.Item(1) != '20'")
	}
	if v, _ := st.Item(2); v != "30" {
		t.Error("Test_SJS_1: st.Item(2) != '30'")
	}
}

func Test_SJS_HeadTail(t *testing.T) {
	ini := ""
	fin := ""

	duo := Make(ini, fin)

	if duo.Head1() != ini {
		t.Error("Test_SJS_HeadTail: duo.Head1() != ini")
	}
	if duo.Tail1().Head1() != fin {
		t.Error("Test_SJS_HeadTail: duo.Tail1().Head1() != fin")
	}
	if duo.Tail1().Tail1().Empty() != true {
		t.Error("Test_SJS_HeadTail: duo.Tail1().Tail1().Empty() != true")
	}
}

func Test_SJS_Membership(t *testing.T) {
	stooges := Make("Curly", "Moe", "Larry")
	if v, _ := stooges.Member("Curly"); v != true {
		t.Error("Test_SJS_Membership: stooges.Member('Curly') != true.", v)
	}
	if v, _ := stooges.Member("Bobert"); v != false {
		t.Error("Test_SJS_Membership: stooges.Member('Bobert') != false.", v)
	}
}

func Test_SJS_Append(t *testing.T) {
	s0 := NewStream(nil, nil)
	s1 := Make(1)
	s2 := Make(2, 3)
	appended_s1 := s1.Append(s2)
	appended_s2 := s0.Append(s2)

	if v := appended_s1.Head1(); v != 1 {
		t.Error("Test_SJS_Append: appended_s1.Head1() != 1", v)
	}
	if v := appended_s1.Item1(1); v != 2 {
		t.Error("Test_SJS_Append: appended_s1.Item1(1) != 2", v)
	}
	if v := appended_s1.Item1(2); v != 3 {
		t.Error("Test_SJS_Append: appended_s1.Item1(2) != 3", v)
	}
	if v := appended_s1.Length1(); v != 3 {
		t.Error("Test_SJS_Append: appended_s1.Length1() != 3", v)
	}
	if s2 != appended_s2 {
		t.Error("Test_SJS_Append: s2 != appended_s2")
	}
}

func Test_SJS_Equality(t *testing.T) {
	s1 := Make(1)
	s2 := Make(1)
	s3 := Make(2, 3)

	if s1.Equals(s2) != true {
		t.Error("Test_SJS_Equality: s1.Equals(s2) != true")
	}
	if s1.Equals(s3) != false {
		t.Error("Test_SJS_Equality: s1.Equals(s3) != false")
	}
}

func Test_SJS_FromArray(t *testing.T) {
	arr := []interface{}{1, 2, 3}
	st := FromArray(arr)

	if v := st.Head1(); v != 1 {
		t.Error("Test_SJS_FromArray: st.Head1() != 1.", v)
	}
	if v := st.Item1(1); v != 2 {
		t.Error("Test_SJS_FromArray: st.Item(1) != 2.", v)
	}
	if v := st.Item1(2); v != 3 {
		t.Error("Test_SJS_FromArray: st.Item(2) != 3.", v)
	}
	if v := st.Length1(); v != 3 {
		t.Error("Test_SJS_FromArray: st.Length() != 3.", v)
	}
}

func Test_SJS_Range(t *testing.T) {
	st := RangeL(3, 7)
	if v := st.Length1(); v != 5 {
		t.Error("Test_SJS_Range: st.Length1() != 5.", v)
	}
	if v := st.Item1(0); v != 3 {
		t.Error("Test_SJS_Range: st.Item(0) != 3.", v)
	}
	if v := st.Item1(1); v != 4 {
		t.Error("Test_SJS_Range: st.Item(1) != 4.", v)
	}
	if v := st.Item1(2); v != 5 {
		t.Error("Test_SJS_Range: st.Item(2) != 5.", v)
	}
	if v := st.Item1(3); v != 6 {
		t.Error("Test_SJS_Range: st.Item(3) != 6.", v)
	}
	if v := st.Item1(4); v != 7 {
		t.Error("Test_SJS_Range: st.Item(4) != 7.", v)
	}

	st_infinite := Range(0, -1)
	if v := st_infinite.Item1(100); v != 100 {
		t.Error("Test_SJS_Range: st.Item1(100) != 100.", v)
	}

	stf64 := RangeL(3.1, 7.1)
	if v := stf64.Length1(); v != 5 {
		t.Error("Testf64_SJS_Range: stf64.Length1() != 5.0, ", v)
	}
	if v := stf64.Item1(0); v != 3.1 {
		t.Error("Testf64_SJS_Range: stf64.Item(0) != 3.1, ", v)
	}
	if v := stf64.Item1(1); v != 4.1 {
		t.Error("Testf64_SJS_Range: stf64.Item(1) != 4.1, ", v)
	}
	if v := stf64.Item1(2); v != 5.1 {
		t.Error("Testf64_SJS_Range: stf64.Item(2) != 5.1, ", v)
	}
	if v := stf64.Item1(3); v != 6.1 {
		t.Error("Testf64_SJS_Range: stf64.Item(3) != 6.1, ", v)
	}
	if v := stf64.Item1(4); v != 7.1 {
		t.Error("Testf64_SJS_Range: stf64.Item(4) != 7.1, ", v)
	}

	stf64_infinite := Range(0, -1)
	if v := stf64_infinite.Item1(100); v != 100 {
		t.Error("Testf64_SJS_Range: stf64.Item1(100) != 100.", v)
	}
}

func Test_SJS_RangeL(t *testing.T) {
	st := RangeL(3, 7)
	if v := st.Length1(); v != 5 {
		t.Error("Test_SJS_RangeL: st.Length1() != 5.", v)
	}
	if v := st.Item1(0); v != 3 {
		t.Error("Test_SJS_RangeL: st.Item(0) != 3.", v)
	}
	if v := st.Item1(1); v != 4 {
		t.Error("Test_SJS_RangeL: st.Item(1) != 4.", v)
	}
	if v := st.Item1(2); v != 5 {
		t.Error("Test_SJS_RangeL: st.Item(2) != 5.", v)
	}
	if v := st.Item1(3); v != 6 {
		t.Error("Test_SJS_RangeL: st.Item(3) != 6.", v)
	}
	if v := st.Item1(4); v != 7 {
		t.Error("Test_SJS_RangeL: st.Item(4) != 7.", v)
	}

	st_infinite := RangeL(0, -1)
	if v := st_infinite.Item1(100); v != 100 {
		t.Error("Test_SJS_RangeL: st.Item1(100) != 100.", v)
	}
}

func Test_SJS_RangeR(t *testing.T) {
	st := RangeR(7, 3)
	if v := st.Length1(); v != 5 {
		t.Error("Test_SJS_RangeR: st.Length1() != 5.", v)
	}
	if v := st.Item1(0); v != 7 {
		t.Error("Test_SJS_RangeR: st.Item(0) != 7.", v)
	}
	if v := st.Item1(1); v != 6 {
		t.Error("Test_SJS_RangeR: st.Item(1) != 6.", v)
	}
	if v := st.Item1(2); v != 5 {
		t.Error("Test_SJS_RangeR: st.Item(2) != 5.", v)
	}
	if v := st.Item1(3); v != 4 {
		t.Error("Test_SJS_RangeR: st.Item(3) != 4.", v)
	}
	if v := st.Item1(4); v != 3 {
		t.Error("Test_SJS_RangeR: st.Item(4) != 3.", v)
	}

	st_infinite := RangeR(0, 1)
	if v := st_infinite.Item1(100); v != -100 {
		t.Error("Test_SJS_RangeL: st.Item1(100) != -100.", v)
	}
}

func Test_SJS_Take(t *testing.T) {
	naturals := Range(1, -1)
	first_three_naturals := naturals.Take(3)

	if v := first_three_naturals.Length1(); v != 3 {
		t.Error("Test_SJS_Take: first_three_naturals.Length() != 3")
	}
	if v := first_three_naturals.Item1(0); v != 1 {
		t.Error("Test_SJS_Take: first_three_naturals.Item1(0) != 1")
	}
	if v := first_three_naturals.Item1(1); v != 2 {
		t.Error("Test_SJS_Take: first_three_naturals.Item1(1) != 2")
	}
	if v := first_three_naturals.Item1(2); v != 3 {
		t.Error("Test_SJS_Take: first_three_naturals.Item1(2) != 3")
	}
}

func Test_SJS_Drop(t *testing.T) {
	naturals := Range(1, -1)
	skip := naturals.Drop(3)

	if v := skip.Head1(); v != 4 {
		t.Error("Test_SJS_Drop: skip.Head1() != 4")
	}

	if v := skip.Item1(0); v != 4 {
		t.Error("Test_SJS_Take: skip.Item1(0) != 4")
	}
	if v := skip.Item1(1); v != 5 {
		t.Error("Test_SJS_Take: skip.Item1(1) != 5")
	}
	if v := skip.Item1(2); v != 6 {
		t.Error("Test_SJS_Take: skip.Item1(2) != 6")
	}
}

func Test_SJS_Map(t *testing.T) {
	alphabet_ascii := Range('A', 'Z')
	alphabet := alphabet_ascii.Map(func(code interface{}) interface{} {
		return 'A'
	})
	if v := alphabet.Head1(); v != 'A' {
		t.Error("Test_SJS_Map: alphabet.Head1() != 'A'.", v)
	}
	if v := alphabet.Tail1().Head1(); v != 'A' {
		t.Error("Test_SJS_Map: alphabet.Tail1().Head1() != 'A'.", v)
	}
	if v := alphabet.Item1(25); v != 'A' {
		t.Error("Test_SJS_Map: alphabet.Item(25) != 'A'.", v)
	}
}

func Test_SJS_Filter(t *testing.T) {
	first_ten_naturals := Range(1.0, 10.0)

	first_five_evens := first_ten_naturals.Filter(func(v interface{}) bool {
		switch d := v.(type) {
		case float64:
			return math.Mod(d, 2) == 0
		case int:
			return (d%2 == 0)
		default:
			return false
		}
	})

	if v := first_five_evens.Length1(); v != 5 {
		t.Error("Test_SJS_Filter: first_five_evens.Length1() != 5", v)
	}

	first_five_evens.Map(func(v interface{}) interface{} {
		switch d := v.(type) {
		case float64:
			if d/2 != math.Floor(d/2) {
				t.Error("Test_SJS_Filter: map test: d / 2 != math.Floor(d/2).", d)
			}
			return v
		default:
			return v
		}
	})
}

func Test_SJS_Reduce(t *testing.T) {
	first_twenty_naturals := Range(1, 20)
	twentieth_triangular_number_w_initial := first_twenty_naturals.Reduce(func(x, y interface{}) interface{} {
		switch vx := x.(type) {
		case int:
			{
				switch vy := y.(type) {
				case int:
					{
						return vx + vy
					}
				}
			}
		}
		return x
	}, 0)

	twentieth_triangular_number := first_twenty_naturals.Reduce(func(x, y interface{}) interface{} {
		switch vx := x.(type) {
		case int:
			{
				switch vy := y.(type) {
				case int:
					{
						return vx + vy
					}
				}
			}
		}
		return x
	}, nil)

	if twentieth_triangular_number_w_initial != 210 {
		t.Error("Test_SJS_Reduce: twentieth_triangular_number_w_initial != 210.", twentieth_triangular_number_w_initial)
	}

	if twentieth_triangular_number != nil {
		t.Error("Test_SJS_Reduce: twentieth_triangular_number != nil.", twentieth_triangular_number)
	}
}

func Test_SJS_Sum(t *testing.T) {

	first_twenty_naturals := Range(1, 20)
	twentieth_triangular_number := first_twenty_naturals.Sum()
	if twentieth_triangular_number != 210 {
		t.Error("Test_SJS_Sum: twentieth_triangular_number != 210.", twentieth_triangular_number)
	}

	first_twenty_naturalsf64 := Range(1.0, 20.0)
	twentieth_triangular_numberf64 := first_twenty_naturalsf64.Sum()
	if twentieth_triangular_numberf64 != 210.0 {
		t.Error("Test_SJS_Sum: twentieth_triangular_numberf64 != 210.0, ", twentieth_triangular_numberf64)
	}
}

func Test_SJS_Scale(t *testing.T) {
	first_ten_naturals := Range(1, 10)
	first_ten_evens := first_ten_naturals.Scale(2)
	if v := first_ten_evens.Length1(); v != 10 {
		t.Error("Test_SJS_Scale: first_ten_evens.Length1() != 10.", v)
	}
	if v := first_ten_evens.Head1(); v != 2 {
		t.Error("Test_SJS_Scale: first_ten_evens.Head1() != 2.", v)
	}
	if v := first_ten_evens.Item1(9); v != 20 {
		t.Error("Test_SJS_Scale: first_ten_evens.Item1(9) != 20.", v)
	}

	first_ten_naturals_f64 := Range(1.0, 10.0)
	first_ten_evens_f64 := first_ten_naturals_f64.Scale(2.0)
	if v := first_ten_evens_f64.Length1(); v != 10.0 {
		t.Error("Test_SJS_Scale: first_ten_evens_f64.Length1() != 10.0 ", v)
	}
	if v := first_ten_evens_f64.Head1(); v != 2.0 {
		t.Error("Test_SJS_Scale: first_ten_evens_f64.Head1() != 2.0 ", v)
	}
	if v := first_ten_evens_f64.Item1(9); v != 20.0 {
		t.Error("Test_SJS_Scale: first_ten_evens_f64.Item1(9) != 20.0 ", v)
	}

	first_ten_naturals_i64 := Range(int64(1), int64(10))
	first_ten_evens_i64 := first_ten_naturals_i64.Scale(int64(2))
	if v := first_ten_evens_i64.Length1(); v != 10 {
		t.Error("Test_SJS_Scale: first_ten_evens_i64.Length1() != 10.", v)
	}
	if v := first_ten_evens_i64.Head1(); v != int64(2) {
		t.Error("Test_SJS_Scale: first_ten_evens_i64.Head1() != 2.", v)
	}
	if v := first_ten_evens_i64.Item1(9); v != int64(20) {
		t.Error("Test_SJS_Scale: first_ten_evens_i64.Item1(9) != 20.", v)
	}
}

func Test_SJS_ConcatMap(t *testing.T) {

	// no test in stream.js lib. Used haskell example instead: concatMap (\x -> 1 `enumFromTo` x) [1,3,5] == [1,1,2,3,1,2,3,4,5]

	l := Make(1, 3, 5)
	cat := l.ConcatMap(func(x interface{}) interface{} {
		return Range(1, x)
	})

	if v := cat.Length1(); v != 9 {
		t.Error("Test_SJS_ConcatMap: cat.Length1() != 3.", v)
	}

	// ugly ehe!
	if cat.Item1(0) != 1 && cat.Item1(1) != 1 && cat.Item1(2) != 2 && cat.Item1(3) != 3 && cat.Item1(4) != 1 && cat.Item1(5) != 2 && cat.Item1(6) != 3 && cat.Item1(7) != 4 && cat.Item1(5) != 5 {
		t.Error("Test_SJS_ConcatMap: everything is wrecked.")
	}
}

func Test_Demo(demoT *testing.T) {

	fmt.Println("Example 1:")
	var s *Stream
	s = Make(10, 20, 30)
	fmt.Printf("\ts.Length() = %d\n", s.Length1())
	fmt.Printf("\ts.Head() = %d\n", s.Head1())
	fmt.Printf("\ts.Item1(0) = %d\n", s.Item1(0))
	fmt.Printf("\ts.Item1(1) = %d\n", s.Item1(1))
	fmt.Printf("\ts.Item1(2) = %d\n", s.Item1(2))

	fmt.Println("Example 2:")
	s = Make(10, 20, 30)
	t := s.Tail1()
	fmt.Printf("\ts.Tail() = t.Head() = %d\n", t.Head1())
	u := t.Tail1()
	fmt.Printf("\tt.Tail() = u.Head() = %d\n", u.Head1())
	v := u.Tail1()
	fmt.Printf("\tv.Empty() = %v\n", v.Empty())

	fmt.Println("Example 3:")
	s = Make(10, 20, 30)
	for ; !s.Empty(); s = s.Tail1() {
		fmt.Printf("\tValue = %d\n", s.Head1())
	}

	fmt.Println("Example 4:")
	s = Range(10, 20)
	s.Print(-1)

	fmt.Println("Example 5:")
	doubleNumber := func(x int) int {
		return 2 * x
	}
	doubleNumberInterface := func(x interface{}) interface{} {
		switch v := x.(type) {
		case int:
			return doubleNumber(v)
		default:
			return v
		}
	}

	fmt.Println("\tNumbers 10-15")
	numbers := Range(10, 15)
	numbers.Print(-1)

	fmt.Println("\tNumbers 10-15 doubled")
	doubles := numbers.Map(doubleNumberInterface)
	doubles.Print(-1)

	fmt.Println("Example 6:")
	checkIfOdd := func(x interface{}) bool {
		switch v := x.(type) {
		case int:
			if (v % 2) == 0 {
				return false
			} else {
				return true
			}
		default:
			return false
		}
	}

	fmt.Println("\tNumbers 10-15")
	numbers = Range(10, 15)
	numbers.Print(-1)

	fmt.Println("\tNumbers 10-15 Filtered for odd's")
	onlyOdds := numbers.Filter(checkIfOdd)
	onlyOdds.Print(-1)

	fmt.Println("Example 7:")
	printItem := func(x interface{}) interface{} {
		fmt.Printf("The element is: %v\n", x)
		return x
	}
	numbers = Range(10, 12)
	numbers.Walk(printItem)

	fmt.Println("Example 8:")
	numbers = Range(10, 100)
	fewerNumbers := numbers.Take(10)
	fewerNumbers.Print(10)

	fmt.Println("Example 9:")
	numbers = Range(1, 3)
	fmt.Println("\tNumbers 1-3 scaled by 10")
	multiplesOfTen := numbers.Scale(10)
	multiplesOfTen.Print(-1)
	fmt.Println("\tNumbers 1-3 scaled by 10, Add")
	numbers.Add(multiplesOfTen).Print(-1)

	fmt.Println("Example 10:")
	naturalNumbers := Range(1, -1)
	oneToTen := naturalNumbers.Take(10)
	oneToTen.Print(-1)

	fmt.Println("Example 11:")
	naturalNumbers = Range(1, -1)
	evenNumbers := naturalNumbers.Map(func(x interface{}) interface{} {
		switch v := x.(type) {
		case int:
			return 2 * v
		default:
			return x
		}
	})
	oddNumbers := naturalNumbers.Filter(func(x interface{}) bool {
		switch v := x.(type) {
		case int:
			return (v % 2) != 0
		default:
			return false
		}
	})
	fmt.Println("\tTake(3) of Even Numbers from 1 to Infiniti")
	evenNumbers.Take(3).Print(-1)
	fmt.Println("\tTake(3) of Odd Numbers from 1 to Infinity")
	oddNumbers.Take(3).Print(-1)

	fmt.Println("\tExample 12:")
	s = NewStream(10, func(v interface{}, fn STREAMFN) *Stream {
		return NewStream(nil, nil)
	})
	fmt.Printf("\ts.Head() = %d\n", s.Head1())
	s.Print(-1)

	t = NewStream(10, func(v interface{}, fn STREAMFN) *Stream {
		return NewStream(20, func(v interface{}, fn STREAMFN) *Stream {
			return NewStream(30, func(v interface{}, fn STREAMFN) *Stream {
				return NewStream(nil, nil)
			})
		})
	})

	fmt.Println("\tThree streams created manually via NewStream")
	t.Print(-1)

	fmt.Println("\tExample 13:")
	s = NewStream(1, Ones)
	s.Take(3).Print(-1)

	fmt.Println("\tExample 14:")
	NaturalNumbers(1, nil).Take(5).Print(-1)

	fmt.Println("\tExample 15: Sieve from 2 to 10!")
	Sieve(Range(2, -1)).Take(10).Print(-1)
}

func Ones(v interface{}, fn STREAMFN) *Stream {
	return NewStream(1, Ones)
}

func NaturalNumbers(v interface{}, fn STREAMFN) *Stream {
	return NewStream(1, func(v interface{}, fn STREAMFN) *Stream {
		return Ones(1, nil).Add(NaturalNumbers(v, fn))
	})
}

/*
 * haskell:
 *  sieve (p : xs) = p : sieve [x | x <- xs, x `mod` p > 0]
 *  take 10 $ sieve [2..]
 */

func Sieve(s *Stream) *Stream {
	h := s.Head1()
	return NewStream(h, func(v interface{}, fn STREAMFN) *Stream {
		return Sieve(s.Tail1().Filter(func(x interface{}) bool {
			switch d := x.(type) {
			case int:
				switch dh := h.(type) {
				case int:
					return d%dh != 0
				default:
					return false
				}
			default:
				return false
			}
		}))
	})
}
