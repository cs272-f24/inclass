func TestExtract {
	tests := []struct {
		doc string
		want []string
	}{
		doc: `...`,
		wantWords: []string{"foo", "bar", "baz"},
		wantHrefs: []string{"/", "index.html", "https://example.com"},
	}
}

func Extract() {
	// run html.Parse()
	// in recursion, when you encounter TextNode, split 
	// its Data into a slice of string

	strings.FieldsFunc(n.Data, f) // provide anon func like pkg example
}
