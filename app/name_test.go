package main

import(
	"regexp"
	"testing"
)

func TestNameCasing(t *testing.T) {
	name := GetName()
	match, _ := regexp.MatchString("^[A-Z]{1}[a-z]+$", name)

	if(match) {
		t.Logf("GetName() PASSED, got %v", name)
	} else {
		t.Errorf("GetName() FAILED, got %v", name)
	}
}

func TestNameLength(t *testing.T) {
	for i := 0; i < 10; i++ {
		name := GetName()

		if(len(name) < 20) {
			t.Logf("GetName() PASSED, %s length is less than 20 characters", name)
		} else {
			t.Errorf("GetName() FAILED, %s length is more than 20 characters", name)
		}
	}
}

func TestRandomnessDistribution(t *testing.T) {
	gotName := make(map[string]int)

	for i := 0; i < 100; i++ {
		name := GetName()

		gotName[name] = gotName[name] + 1
	}

	// Expect between 1 and 5 names
	if(len(gotName) > 0 && len(gotName) < 6) {
		t.Logf("GetName() randomness distribution PASSED")
	} else {
		t.Errorf("GetName() randomness distribution FAILED")
	}
}