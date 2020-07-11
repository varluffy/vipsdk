package vipsdk

type Param interface {
	ServiceName() string
	MethodName() string
	Version() string
	Params() interface{}
	Token() bool
}
