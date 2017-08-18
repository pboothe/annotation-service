//Preforms linear search on a list of nodes for a geolocation
package search

import (
	"bytes"
	"errors"
	"net"

	"github.com/m-lab/annotation-service/parser"
)

func SearchList(list []parser.Node, ipLookUp string) (*parser.Node, error) {
	userIP := net.ParseIP(ipLookUp)
	if userIP == nil {
		return nil, errors.New("Ivalid search IP")
	}
	for i := range list {
		if bytes.Compare(userIP, list[i].LowRangeBin) >= 0 && bytes.Compare(userIP, list[i].HighRangeBin) <= 0 {
			return &list[i], nil
		}
	}
	return nil, errors.New("not found\n")
}
