package plex

import (
	"testing"
	"net"
)

func TestPipePlex(t *testing.T) {
	c1, c2 := net.Pipe()
 	_, _ = NewPlex(c1), NewPlex(c2)
}
