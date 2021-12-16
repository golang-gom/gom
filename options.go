package gom

type Option struct {
	Name  string
	Value interface{}
}

var (
	IsFinite *Option = &Option{
		Name: "is finite",
	}
	NoClose *Option = &Option{
		Name: "no close",
	}
)
