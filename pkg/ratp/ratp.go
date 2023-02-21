package ratp

import ()


type Commander interface {
	Request()
	Stop()
}

type Getter interface {
	Request()
}

type Stringer interface {
	String()
}

