package service

import "testing"
import "fmt"

func TestProfile(t *testing.T) {
	icon, name, err := GetProfile("konojunya")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(icon)
	fmt.Println(name)
}
