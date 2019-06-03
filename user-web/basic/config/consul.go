package config

type ConsulCOnfig interface {
	GetEnable() bool
	GetPort() int
	GetHost() string
}

type defaultConsulConnfig struct {
	Enabled bool `json:"enable"`
	Host string `json:"host"`
	Port int `json:"port"`
}

func (c defaultConsulConnfig) GetPort() int  {
	return c.Port
}

func (c defaultConsulConnfig) GetEnabled() bool  {
	return c.Enabled
}

func (c defaultConsulConnfig) GetPort() int  {
	return c.Port
}
