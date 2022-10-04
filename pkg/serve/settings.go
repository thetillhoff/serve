package serve

type Setting struct {
	Type  SettingType
	Value interface{}
}

type SettingType uint

const (
	Directory SettingType = iota
	Verbose
	IPAddress
	Port
)
