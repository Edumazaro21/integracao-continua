package main

import "testing"

func TestSoma(t *testing.T) {

	result := Soma(15, 15)

	if result != 30 {
		t.Errorf("Soma(15, 15) = %d; want 30", result)
	}
}

func TestMain(m *testing.M) {
	main()

	m.Run()
}
