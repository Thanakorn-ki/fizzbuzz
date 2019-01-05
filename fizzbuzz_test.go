package fizzbuzz_test

import (
	"testing"

	"github.com/salapao2136/fizzbuzz"
)

func TestFizzBuzzShouldSayOne(t *testing.T) {
	//  Arrange
	expected := "1"

	// Action
	result := fizzbuzz.Say(1)

	// Assertion
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}

	result = fizzbuzz.Say(2)
	expected = "2"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}

	result = fizzbuzz.Say(58)
	expected = "58"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestFizzBuzzShouldSayTwo(t *testing.T) {
	result := fizzbuzz.Say(2)
	expected := "2"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestFizzBuzzShouldSayFizz(t *testing.T) {
	result := fizzbuzz.Say(3)
	expected := "Fizz"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}

	result = fizzbuzz.Say(6)
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}

	result = fizzbuzz.Say(39)
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}

	result = fizzbuzz.Say(57)
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestFizzBuzzShouldSayBuzz(t *testing.T) {
	result := fizzbuzz.Say(5)
	expected := "Buzz"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}

	result = fizzbuzz.Say(10)
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}

	result = fizzbuzz.Say(20)
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}

}

func TestFizzBuzzShouldSayFizzBuzz(t *testing.T) {
	result := fizzbuzz.Say(15)
	expected := "FizzBuzz"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}

	result = fizzbuzz.Say(30)
	expected = "FizzBuzz"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}
