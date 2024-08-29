
TestExtract() {
	tests := []struct {
		doc []byte
		want []string
	}{
		{
			doc: `<html>hello</html`,
			want: {"hello"},
		},
		{
			doc: `<html>
			<body>
				<a href="/something">sth</a>
				<a href="/another">another</a>
				<a href="https://www.example.com">example</a>
			</body></html>
			`,
			want: { "sth", "another", "example"},
		}
	}
}

Extract(doc []byte)
