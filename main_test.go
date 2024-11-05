package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		got := cleanInput("    ")
		want := []string{}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("spaces both side", func(t *testing.T) {
		got := cleanInput("  Hello  ")
		want := []string{"hello"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
