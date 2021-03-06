package parser_test

import (
	"archive/zip"
	"log"
	"net"
	"strings"
	"testing"

	"github.com/go-test/deep"

	"github.com/m-lab/annotation-service/loader"
	"github.com/m-lab/annotation-service/parser"
)

func init() {
	// Always prepend the filename and line number.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestIPLisGLite2(t *testing.T) {
	var ipv4, ipv6 []parser.IPNode
	var ipv6Expected = []parser.IPNode{
		parser.IPNode{
			net.ParseIP("600:8801:9400:5a1:948b:ab15:dde3:61a3"),
			net.ParseIP("600:8801:9400:5a1:948b:ab15:dde3:61a3"),
			4,
			"91941",
			32.7596,
			-116.994,
		},
		parser.IPNode{
			net.ParseIP("2001:5::"),
			net.ParseIP("2001:0005:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF"),
			4,
			"",
			47,
			8,
		},
		parser.IPNode{
			net.ParseIP("2001:200::"),
			net.ParseIP("2001:0200:00FF:FFFF:FFFF:FFFF:FFFF:FFFF"),
			4,
			"",
			36,
			138,
		},
	}
	var ipv4Expected = []parser.IPNode{
		parser.IPNode{
			net.ParseIP("1.0.0.0"),
			net.ParseIP("1.0.0.255"),
			0,
			"3095",
			-37.7,
			145.1833,
		},
		parser.IPNode{
			net.ParseIP("1.0.1.0"),
			net.ParseIP("1.0.1.255"),
			4,
			"",
			26.0614,
			119.3061,
		},
		parser.IPNode{
			net.ParseIP("1.0.2.0"),
			net.ParseIP("1.0.3.255"),
			4,
			"",
			26.0614,
			119.3061,
		},
	}

	locationIDMap := map[int]int{
		2151718: 0,
		1810821: 4,
		5363990: 4,
		6255148: 4,
		1861060: 4,
	}
	reader, err := zip.OpenReader("testdata/GeoLite2City.zip")
	if err != nil {
		t.Fatalf("Error opening zip file")
	}

	rcIPv4, err := loader.FindFile("GeoLite2-City-Blocks-IPv4.csv", &reader.Reader)
	if err != nil {
		t.Fatalf("Failed to create io.ReaderCloser")
	}
	defer rcIPv4.Close()
	ipv4, err = parser.LoadIPListGLite2(rcIPv4, locationIDMap)
	if err != nil {
		t.Errorf("Failed to create ipv4")
	}
	err = isEqualIPLists(ipv4Expected, ipv4)
	if err != nil {
		t.Errorf("Lists are not equal")
	}

	rcIPv6, err := loader.FindFile("GeoLite2-City-Blocks-IPv6.csv", &reader.Reader)
	if err != nil {
		t.Errorf("Failed to create io.ReaderCloser")
	}
	defer rcIPv6.Close()
	ipv6, err = parser.LoadIPListGLite2(rcIPv6, locationIDMap)
	if err != nil {
		log.Println(err)
		t.Errorf("Failed to create ipv6")
	}
	err = isEqualIPLists(ipv6Expected, ipv6)
	if err != nil {
		t.Errorf("Lists are not equal")
	}
}

func TestLocationListGLite2(t *testing.T) {
	var actualLocList []parser.LocationNode
	var actualIDMap map[int]int
	var expectedLocList = []parser.LocationNode{
		parser.LocationNode{
			32909,
			"AS", "IR", "Iran",
			"07", "Ostan-e Tehran",
			0, "Shahre Jadide Andisheh",
		},
		parser.LocationNode{
			49518,
			"AF", "RW", "Rwanda",
			"", "",
			0, "",
		},
		parser.LocationNode{
			51537,
			"AF", "SO", "Somalia",
			"", "",
			0, "",
		},
		parser.LocationNode{
			5127766,
			"NA", "US", "United States",
			"NY", "New York",
			538, "Mount Morris",
		},
	}
	expectedIDMap := map[int]int{
		5127766: 3,
		51537:   2,
		49518:   1,
		32909:   0,
	}

	reader, err := zip.OpenReader("testdata/GeoLite2City.zip")
	if err != nil {
		t.Fatalf("Error opening zip file")
	}

	rc, err := loader.FindFile("GeoLite2-City-Locations-en.csv", &reader.Reader)
	if err != nil {
		t.Fatalf("Failed to create io.ReaderCloser")
	}
	defer rc.Close()
	actualLocList, actualIDMap, err = parser.LoadLocListGLite2(rc)
	if err != nil {
		log.Println(err)
		t.Errorf("Failed to LoadLocationList")
	}
	if actualLocList == nil || actualIDMap == nil {
		t.Errorf("Failed to create LocationList and mapID")
	}

	if diff := deep.Equal(actualLocList, expectedLocList); diff != nil {
		log.Printf("%+v\n", actualLocList)
		log.Printf("%+v\n", expectedLocList)
		t.Error(diff)
	}
	err = isEqualLocLists(actualLocList, expectedLocList)
	if err != nil {
		t.Errorf("Location lists are not equal")
	}

	if diff := deep.Equal(expectedIDMap, actualIDMap); diff != nil {
		t.Error(diff)
	}
}

func TestCorruptData(t *testing.T) {
	reader, err := zip.OpenReader("testdata/GeoLite2CityCORRUPT.zip")
	if err != nil {
		t.Fatalf("Error opening zip file")
	}
	rc, err := loader.FindFile("GeoLite2-City-Locations-en.csv", &reader.Reader)
	if err != nil {
		t.Fatalf("Error finding file")
	}
	_, _, err = parser.LoadLocListGLite2(rc)
	if err.Error() != "Corrupted Data: wrong number of columns" {
		if err == nil {
			t.Errorf("Error inconsistent:\ngot: nil\nwanted: Corrupted Data: wrong number of columns")
		}
		if err != nil {
			output := strings.Join([]string{"Error inconsistent:\ngot:", err.Error(), "\nwanted: Corrupted Data: wrong number of columns"}, "")
			t.Errorf(output)
		}
	}
}
