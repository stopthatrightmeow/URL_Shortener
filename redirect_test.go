package main

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestRedirect(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got %q Wanted %q", got, want)
		}
	}

	// Testing a valid URL
	t.Run("Testing redirect status code", func(t *testing.T) {
		// Makes a simple GET request to the URL specified
		httptest.NewRequest("GET", "http://localhost:9000/", nil)
		// Starts recording the results
		resp := httptest.NewRecorder().Result().StatusCode

		// Test results | The Sprint convers the int to str
		got := fmt.Sprint(resp)
		want := "301"
		// Perform the comparison.
		assertCorrectMessage(t, got, want)
	})

	t.Run("Testing non-existent URL", func(t *testing.T) {
		// Make a GET request to a random URL that shouldn't exist
		httptest.NewRequest("GET", "http://localhost:9000/potatoes", nil)
		// Get the results
		resp := httptest.NewRecorder().Result().StatusCode

		fmt.Println(a)
		// Setup Results for comparison
		got := fmt.Sprint(resp)
		want := "404"
		// Do the comparison
		assertCorrectMessage(t, got, want)
	})
}
