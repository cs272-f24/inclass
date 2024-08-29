func TestCleanHrefs(t *testing.T) {

	tests := []struct {
		hostName string // hostname to start at, and pre-pend to partial
		hrefs []string // input hrefs, could be absolute or partial
		want []string  // expected output, absolute URLs
	}{
		{
			"example.com",
			[]string{ "/", "https://example.com", "documents/"},
			[]string{"https://example.com/", "https://example.com/", "https://example.com/documents/"},
		},
		{},
		{},
	}

	range over tests
		call CleanHrefs for each one
		compare using reflect.DeepEqual() 
}

func CleanHrefs(hrefs []string) []string {
	range over hrefs
		u, err  := url.Parse(href) // takes a string and produces URL struct
		.....make it an absolute url including scheme and hostname

		u.String() // takes a struct and produces the string version
}
