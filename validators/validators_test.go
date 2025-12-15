package validators

import "testing"

func TestValidAlias(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{"hello_world", true},
		{"hello-world123", true},
		{"hello world", false},
		{"hello@world", false},
		{"", false},    // empty string
		{"___", true},  // only underscores
		{"---", true},  // only dashes
		{"aB1_", true}, // mix of letters, numbers, underscore
		{"a b", false}, // contains space

	}

	for _, tt := range cases {
		// tt:= tt
		t.Run(tt.input, func(t *testing.T) {
			// t.Parallel()
			valid := ValidAlias(tt.input)
			if valid != tt.want {
				t.Errorf("ValidAlias(%q) = %v; want %v", tt.input, valid, tt.want)
			}
		})
	}

}

func TestValidUrl(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{"https://example.com", true},
		{"http://www.example.com/test", true},
		{"www.example.com", true},
		{"example", false},
		{"htp:/invalid.com", false},
		{"https://sub.domain.com/path?query=1", true}, // subdomain + query
		{"ftp://example.com", false},                  // unsupported scheme
		{"http://example.c", false},                   // invalid TLD
	}

	for _, tt := range cases {
		t.Run(tt.input, func(t *testing.T) {
			isValid := ValidLongURL(tt.input)
			if isValid != tt.want {
				t.Errorf("validLongURL(%q) = %v want %v", tt.input, isValid, tt.want)
			}
		})
	}

}
