package search_test

import (
	"archive/zip"
	"log"
	"net"
	"testing"

	"github.com/m-lab/annotation-service/loader"
	"github.com/m-lab/annotation-service/parser"
	"github.com/m-lab/annotation-service/search"
)
func TestSearchSmallRange(t *testing.T) {
	var ipv4 = []parser.IPNode{
		parser.IPNode{
			net.ParseIP("1.0.0.0"),
			net.ParseIP("1.0.0.255"),
			0,
			"",
			0,
			0,
		},
		parser.IPNode{
			net.ParseIP("1.0.0.2"),
			net.ParseIP("1.0.0.200"),
			4,
			"",
			0,
			0,
		},
		parser.IPNode{
			net.ParseIP("1.0.0.5"),
			net.ParseIP("1.0.0.100"),
			0,
			"",
			0,
			0,
		},
		parser.IPNode{
			net.ParseIP("1.0.0.120"),
			net.ParseIP("1.0.0.140"),
			0,
			"",
			0,
			0,
		},
		parser.IPNode{
			net.ParseIP("1.0.0.121"),
			net.ParseIP("1.0.0.125"),
			0,
			"",
			0,
			0,
		},
		parser.IPNode{
			net.ParseIP("1.0.0.129"),
			net.ParseIP("1.0.0.130"),
			0,
			"",
			0,
			0,
		},
		parser.IPNode{
			net.ParseIP("1.0.4.0"),
			net.ParseIP("1.0.7.255"),
			4,
			"",
			0,
			0,
		},
	}

	// Test IP node within several subsets 
	ip, err := search.SearchList(ipv4, "1.0.0.122")
	if err != nil {
		log.Println(err)
		t.Errorf("Search failed")
	}
	err = parser.IsEqualIPNodes(ipv4[4], ip)
	if err != nil {
		log.Println(err)
		t.Errorf("Found ", ip, " wanted", ipv4[4])
	}

	// Test IP node not in a subset 
	ip, err = search.SearchList(ipv4, "1.0.0.254")
	if err != nil {
		log.Println(err)
		t.Errorf("Search failed")
	}
	err = parser.IsEqualIPNodes(ipv4[0], ip)
	if err != nil {
		log.Println(err)
		t.Errorf("Found ", ip, " wanted", ipv4[0])
	}

	// Test first IP node
	ip, err = search.SearchList(ipv4, "1.0.0.254")
	if err != nil {
		log.Println(err)
		t.Errorf("Search failed")
	}
	err = parser.IsEqualIPNodes(ipv4[0], ip)
	if err != nil {
		log.Println(err)
		t.Errorf("Found ", ip, " wanted", ipv4[0])
	}

	// Test last IP node in the list
	ip, err = search.SearchList(ipv4, "1.0.6.0")
	if err != nil {
		log.Println(err)
		t.Errorf("Search failed")
	}
	err = parser.IsEqualIPNodes(ipv4[6], ip)
	if err != nil {
		log.Println(err)
		t.Errorf("Found ", ip, " wanted", ipv4[6])
	}

	// Test IP NOT in list 
	ip, err = search.SearchList(ipv4, "255.0.6.0")
	if err == nil {
		log.Println("Got ",ip," wanted: Node not found")
		t.Errorf("Search failed")
	}
}

func TestGeoLite2(t *testing.T) {
	reader, err := zip.OpenReader("testdata/GeoLite2.zip")
	if err != nil {
		t.Errorf("Error opening zip file")
	}

	// Create Location list 
	rc, err := loader.FindFile("GeoLite2-City-Locations-en.csv", &reader.Reader)
	if err != nil {
		t.Errorf("Failed to create io.ReaderCloser")
	}
	defer rc.Close()

	locationList, idMap, err := parser.CreateLocationList(rc)
	if err != nil {
		t.Errorf("Failed to CreateLocationList")
	}
	if locationList == nil || idMap == nil {
		t.Errorf("Failed to create LocationList and mapID")
	}

	/*rcIPv6, err := loader.FindFile("GeoLite2-City-Blocks-IPv6.csv", &reader.Reader)
	if err != nil {
		log.Println(err)
		t.Errorf("Failed to create io.ReaderCloser")
	}
	defer rcIPv6.Close() */
	rcIPv4, err := loader.FindFile("GeoLite2-City-Blocks-IPv4.csv", &reader.Reader)
	if err != nil {
		log.Println(err)
		t.Errorf("Failed to create io.ReaderCloser")
	}
	defer rcIPv4.Close()
	ipv4, err := parser.CreateIPList(rcIPv4, idMap, "GeoLite2-City-Blocks-IPv4.csv")
	if err != nil {
		log.Println(err)
		t.Errorf("Failed to create ipv4")
	}

	ip, err := search.SearchList(ipv4, "1.0.120.0")
	if err != nil {
		log.Println(err)
		t.Errorf("Search failed")
	}
	var n = parser.IPNode{
		net.ParseIP("1.0.120.0"),
		net.ParseIP("1.0.123.255"),
		11622,
		"690-0887",
		35.4722,
		133.0506,
	}
	err = parser.IsEqualIPNodes(n, ip)
	if err != nil {
		log.Println(err)
		t.Errorf("Found ", ip, " wanted", n)
	}

	ip, err = search.SearchList(ipv4,"80.231.5.200")
	if err != nil {
		log.Println(err) 
		t.Errorf("Search failed") 
	}
	var x = parser.IPNode{
		net.ParseIP("80.231.5.0"),
		net.ParseIP("80.231.5.255"),
		0,
		"",
		0,
		0,
	}
	err = parser.IsEqualIPNodes(x, ip)
	if err != nil {
		log.Println(err)
		t.Errorf("Found ", ip, " wanted", x)
	}


}
