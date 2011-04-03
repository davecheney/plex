package plex

import (
	"net"
)

type Plex struct {
	conn net.Conn
	// unexported
}

func runner(p *Plex) {

}


func NewPlex(conn net.Conn) *Plex {
	p := &Plex{conn}
	go runner(p)
	return p
}
