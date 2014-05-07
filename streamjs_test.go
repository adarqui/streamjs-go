package streamjs

import (
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

	/*
	closure := func(v interface{}, sfn STREAMFN) (*Stream) {
		return NewStream("2", closure);
	}
	*/

	st2 := NewStream("1", func(v interface{}, sfn STREAMFN) (*Stream) {
		return NewStream("2", nil);
	});
	if l, _ := st2.Length(); l != 2 {
		t.Error("Test_NewStream: st2.Length() != 2")
	}
	if v, _ := st2.Head(); v != "1" {
		t.Error("Test_NewStream: st2.Head() != '1'")
	}
	if v, _ := st2.Item(1); v != "2" {
		t.Error("Test_NewStream: st2.Item(1) != '2'")
	}



	/*
	st3 := NewStream("1", closure)
	if l, _ := st3.Length(); l != 3 {
		t.Error("Test_NewStream: st2.Length() != 3")
	}
	*/
}


func Test_SJS_Make(t *testing.T) {
	st := Make("10", "20", "30", "40");
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
