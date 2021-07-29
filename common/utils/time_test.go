package utils

import (
	"testing"
)

func TestGetDaysAfter(t *testing.T) {
	start, today := "20201017", "20201017"
	l, err := GetDaysAfter(start, today)
	if err != nil {
		t.Fatal(err)
	}
	if len(l) > 0 {
		t.Fatal(l)
	}

	start, today = "20201016", "20201017"
	l, err = GetDaysAfter(start, today)
	if err != nil {
		t.Fatal(err)
	}
	if len(l) != 1 && l[0] != "20201017" {
		t.Fatal(l)
	}

	start, today = "20201015", "20201017"
	l, err = GetDaysAfter(start, today)
	if err != nil {
		t.Fatal(err)
	}
	if len(l) != 2 && l[0] != "20201015" && l[1] != "20201016" {
		t.Fatal(l)
	}

	start, today = "20201016", "20201017"
	l, err = GetDaysAfter(start, today)
	if err != nil {
		t.Fatal(err)
	}
	if len(l) != 1 && l[0] != "20201017" {
		t.Fatal(l)
	}

	start, today = "20201001", "20201017"
	l, err = GetDaysAfter(start, today)
	if err != nil {
		t.Fatal(err)
	}
	if len(l) != 16 {
		t.Fatal(l)
	}
}

func TestDayTimeAddDays(t *testing.T) {
	dayTm, _ := DayTimeAddDays("20200101", 5)
	if dayTm != "20200106" {
		t.Fatal(dayTm)
	}
}
