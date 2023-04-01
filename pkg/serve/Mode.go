package serve

type Mode uint8

const (
	FileOnly Mode = iota
	InMemoryOnly
	InMemoryFirst
)
