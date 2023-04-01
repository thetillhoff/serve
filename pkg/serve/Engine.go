package serve

type Engine struct {
	Mode          Mode
	Directory     string
	Verbose       bool
	Ipaddress     string
	Port          string
	InMemoryFiles map[string]string
}

func DefaultEngine() Engine {
	return Engine{
		Directory:     "",
		Verbose:       false,
		Ipaddress:     "0.0.0.0",
		Port:          "3000",
		Mode:          FileOnly,
		InMemoryFiles: map[string]string{},
	}
}
