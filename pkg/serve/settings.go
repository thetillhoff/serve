package serve

type Setting struct {
	Type  SettingType
	Value interface{}
}

type SettingType uint

const (
	// <string> The base directory from which the static files should be served
	Directory SettingType = iota

	// <bool> Whether to print additional information like which paths were accessed
	Verbose

	// <string> The ip address the webserver should listen on
	IPAddress

	// <string> The port the webserver should listen on
	Port

	// <map[string]string> A map[path]content of additional, in-memory files that are served beneath `/inmemory/`
	InMemoryFile
)
