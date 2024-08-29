

TestCleanHref(t * testing.T...) {
	tests := []struct {
			hostname string
			partial string
			....
			
		} {
			hostName: "www.usfca.edu"
			partial: "/academics"
			want: "https://www.usfca.edu/academics"
		}
}

CleanHref(href string) string {
	
}
