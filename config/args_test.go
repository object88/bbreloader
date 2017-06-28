package config

import "testing"

func Test_Parse_Nil(t *testing.T) {
	result := parseArgs(nil)
	if result == nil {
		t.Fatal("Result is nil")
	}
	if len(*result) != 0 {
		t.Fatal("Result does not have 0 length")
	}
}

func Test_Parse_Empty(t *testing.T) {
	result := parseArgs(&[]string{})
	if result == nil {
		t.Fatal("Result is nil")
	}
	if len(*result) != 0 {
		t.Fatal("Result does not have 0 length")
	}
}

func Test_Parse_Single(t *testing.T) {
	result := parseArgs(&[]string{"a"})
	if result == nil {
		t.Fatal("Result is nil")
	}
	if len(*result) != 1 {
		t.Fatal("Result does not have 1 length")
	}
	if (*result)[0] != "a" {
		t.Fatal("Single element is not 'a'")
	}
}

func Test_Parse_Triple(t *testing.T) {
	result := parseArgs(&[]string{"a", "b", "c"})
	if result == nil {
		t.Fatal("Result is nil")
	}
	if len(*result) != 3 {
		t.Fatal("Result does not have 3 length")
	}
	if (*result)[0] != "a" {
		t.Fatal("First element is not 'a'")
	}
	if (*result)[1] != "b" {
		t.Fatal("Second element is not 'b'")
	}
	if (*result)[2] != "c" {
		t.Fatal("Third element is not 'c'")
	}
}

func Test_Prepend_To_Nil(t *testing.T) {
	var a Args
	result := a.prependArgs()
	if result == nil {
		t.Fatal("Result is nil")
	}
}

func Test_Prepend_Nil_To_Empty(t *testing.T) {
	original := parseArgs(nil)

	result := original.prependArgs()
	if result == nil {
		t.Fatal("Result is nil")
	}
	if len(*result) != 0 {
		t.Fatal("Result does not have 0 length")
	}
}

func Test_Prepend_Single_To_Empty(t *testing.T) {
	original := parseArgs(nil)

	result := original.prependArgs("a")
	if result == nil {
		t.Fatal("Result is nil")
	}
	if len(*result) != 1 {
		t.Fatal("Result does not have 1 length")
	}
	if (*result)[0] != "a" {
		t.Fatal("Single element is not 'a'")
	}
}

func Test_Prepend_Three_To_Empty(t *testing.T) {
	original := parseArgs(nil)

	result := original.prependArgs("a", "b", "c")
	if result == nil {
		t.Fatal("Result is nil")
	}
	if len(*result) != 3 {
		t.Fatal("Result does not have 3 length")
	}
	if (*result)[0] != "a" {
		t.Fatal("First element is not 'a'")
	}
	if (*result)[1] != "b" {
		t.Fatal("Second element is not 'b'")
	}
	if (*result)[2] != "c" {
		t.Fatal("Third element is not 'c'")
	}
}

func Test_Prepend_One_To_One(t *testing.T) {
	original := parseArgs(&[]string{"a"})

	result := original.prependArgs("a0")
	if result == nil {
		t.Fatal("Result is nil")
	}
	if len(*result) != 2 {
		t.Fatal("Result does not have 2 length")
	}
	if (*result)[0] != "a0" {
		t.Fatal("Single element is not 'a0'")
	}
	if (*result)[1] != "a" {
		t.Fatal("Single element is not 'a'")
	}
}

func Test_Prepend_Three_To_Three(t *testing.T) {
	original := parseArgs(&[]string{"a", "b", "c"})

	result := original.prependArgs("a0", "b0", "c0")
	if result == nil {
		t.Fatal("Result is nil")
	}
	if len(*result) != 6 {
		t.Fatal("Result does not have 6 length")
	}
	if (*result)[0] != "a0" {
		t.Fatal("First element is not 'a0'")
	}
	if (*result)[1] != "b0" {
		t.Fatal("Second element is not 'b0'")
	}
	if (*result)[2] != "c0" {
		t.Fatal("Third element is not 'c0'")
	}
	if (*result)[3] != "a" {
		t.Fatal("First element is not 'a'")
	}
	if (*result)[4] != "b" {
		t.Fatal("Second element is not 'b'")
	}
	if (*result)[5] != "c" {
		t.Fatal("Third element is not 'c'")
	}
}
