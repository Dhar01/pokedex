package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		got := cleanInput("    ")
		want := []string{}
		checkInput(t, got, want)
	})
	t.Run("spaces both side", func(t *testing.T) {
		got := cleanInput("  Hello  ")
		want := []string{"hello"}
		checkInput(t, got, want)
	})
	t.Run("multiple uppercase letter", func(t *testing.T) {
		got := cleanInput("HeLP")
		want := []string{"help"}
		checkInput(t, got, want)
	})
	t.Run("double string type", func(t *testing.T) {
		got := cleanInput(" HellO World ")
		want := []string{"hello", "world"}
		checkInput(t, got, want)
	})

}

func checkInput(t testing.TB, got, want []string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
