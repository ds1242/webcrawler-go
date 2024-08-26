package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	htmlString := `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`

	tests := []struct {
		name      string
		inputHTML string
		inputURL  string
		expected  []string
	}{
		{
			name:      "get absolute and relatives url from html body",
			inputHTML: htmlString,
			inputURL:  "https://blog.boot.dev",
			expected:  []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputHTML, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL expected URLs: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
