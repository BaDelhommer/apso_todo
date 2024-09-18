package main

type config struct {
	items map[int]string
}

func main() {
	cfg := config{
		items: map[int]string{1: ""},
	}
	startRepl(&cfg)
}
