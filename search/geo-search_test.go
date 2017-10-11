package search_test

import (
	"encoding/binary"
	"log"
	"math/rand"
	"net"
	"testing"

	"google.golang.org/appengine/aetest"

	"github.com/m-lab/annotation-service/loader"
	"github.com/m-lab/annotation-service/parser"
	"github.com/m-lab/annotation-service/search"
)

var (
	// Preloaded by init()
	//gl2locationList []parser.LocationNode
	gl2ipv4 []parser.IPNode
	gl2ipv6 []parser.IPNode
)

func TestGeoLite1(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		log.Println(err)
		t.Errorf("Failed to create aecontext")
	}
	defer done()
	reader, err := loader.CreateZipReader(ctx, "test-annotator-sandbox", "MaxMind/2017/09/07/Maxmind%2F2017%2F09%2F01%2F20170901T085044Z-GeoLiteCity-latest.zip")
	if err != nil {
		log.Println(err)
		t.Errorf("Failed to create zipReader")
	}

	// Load Location list
	rc, err := loader.FindFile("GeoLiteCity-Location.csv", reader)
	if err != nil {
		t.Errorf("Failed to create io.ReaderCloser")
	}
	defer rc.Close()

	locationList, glite1help, idMap, err := parser.LoadLocListGLite1(rc)
	if err != nil {
		t.Errorf("Failed to LoadLocationList")
	}
	if locationList == nil || idMap == nil {
		t.Errorf("Failed to create LocationList and mapID")
	}

	// Test IPv4
	rcIPv4, err := loader.FindFile("GeoLiteCity-Blocks.csv", reader)
	if err != nil {
		log.Println(err)
		t.Errorf("Failed to create io.ReaderCloser")
	}
	defer rcIPv4.Close()
	// TODO: update tests to use high level data loader functions instead of low level funcs
	ipv4, err := parser.LoadIPListGLite1(rcIPv4, idMap, glite1help)
	if err != nil {
		log.Println(err)
		t.Errorf("Failed to create ipv4")
	}
	i := 0
	for i < len(ipv4) {
		ipMiddle := findMiddle(ipv4[i].IPAddressLow, ipv4[i].IPAddressHigh)
		ipBin, errBin := search.SearchBinary(ipv4, ipMiddle.String())
		// Linear search, starting at current node, since it can't be earlier.
		ipLin, errLin := search.SearchList(ipv4[i:], ipMiddle.String())
		if errBin != nil && errLin != nil && errBin.Error() != errLin.Error() {
			log.Println(errBin.Error(), "vs", errLin.Error())
			t.Errorf("Failed Error")
		}
		if parser.IsEqualIPNodes(ipBin, ipLin) != nil {
			log.Println("bad ", ipBin, ipLin)
			t.Errorf("Failed Binary vs Linear")
		}
		i += 100
	}
}

func TestGeoLite2(t *testing.T) {

	i := 0
	for i < len(gl2ipv6) {
		ipMiddle := findMiddle(gl2ipv6[i].IPAddressLow, gl2ipv6[i].IPAddressHigh)
		ipBin, errBin := search.SearchBinary(gl2ipv6, ipMiddle.String())
		// Linear search, starting at current node, since it can't be earlier.
		ipLin, errLin := search.SearchList(gl2ipv6[i:], ipMiddle.String())
		if errBin != nil && errLin != nil && errBin.Error() != errLin.Error() {
			log.Println(errBin.Error(), "vs", errLin.Error())
			t.Errorf("Failed Error")
		}
		if parser.IsEqualIPNodes(ipBin, ipLin) != nil {
			log.Println("bad ", ipBin, ipLin)
			t.Errorf("Failed Binary vs Linear")
		}
		i += 100
	}

	// Test IPv4
	i = 0
	for i < len(gl2ipv4) {
		ipMiddle := findMiddle(gl2ipv4[i].IPAddressLow, gl2ipv4[i].IPAddressHigh)
		ipBin, errBin := search.SearchBinary(gl2ipv4, ipMiddle.String())
		// Linear search, starting at current node, since it can't be earlier.
		ipLin, errLin := search.SearchList(gl2ipv4[i:], ipMiddle.String())
		if errBin != nil && errLin != nil && errBin.Error() != errLin.Error() {
			log.Println(errBin.Error(), "vs", errLin.Error())
			t.Errorf("Failed Error")
		}
		if parser.IsEqualIPNodes(ipBin, ipLin) != nil {
			log.Println("bad ", ipBin, ipLin)
			t.Errorf("Failed Binary vs Linear")
		}
		i += 100
	}

}

// TODO(gfr) This needs good comment and validation?
func findMiddle(low, high net.IP) net.IP {
	lowInt := binary.BigEndian.Uint32(low[12:16])
	highInt := binary.BigEndian.Uint32(high[12:16])
	middleInt := int((highInt - lowInt) / 2)
	mid := low
	i := 0
	if middleInt < 100000 {
		for i < middleInt/2 {
			mid = parser.PlusOne(mid)
			i++
		}
	}
	return mid
}

func BenchmarkGeoLite2ipv4(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		i := rand.Intn(len(gl2ipv4))
		ipMiddle := findMiddle(gl2ipv4[i].IPAddressLow, gl2ipv4[i].IPAddressHigh)
		_, _ = search.SearchBinary(gl2ipv4, ipMiddle.String())
	}
}

func init() {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		log.Println(err)
	}
	defer done()
	reader, err := loader.CreateZipReader(ctx, "test-annotator-sandbox", "MaxMind/2017/09/07/Maxmind%2F2017%2F09%2F07%2F20170907T023620Z-GeoLite2-City-CSV.zip")
	if err != nil {
		log.Println(err)
	}

	// Load Location list
	rc, err := loader.FindFile("GeoLite2-City-Locations-en.csv", reader)
	if err != nil {
	}
	defer rc.Close()

	gl2locationList, idMap, err := parser.LoadLocListGLite2(rc)
	if err != nil {
		log.Println("Failed to LoadLocationList")
	}
	if gl2locationList == nil || idMap == nil {
		log.Println("Failed to create LocationList and mapID")
	}

	// Benchmark IPv4
	rcIPv4, err := loader.FindFile("GeoLite2-City-Blocks-IPv4.csv", reader)
	if err != nil {
		log.Println(err)
	}
	defer rcIPv4.Close()

	gl2ipv4, err = parser.LoadIPListGLite2(rcIPv4, idMap)
	if err != nil {
		log.Println(err)
	}

	// Test IPv6
	rcIPv6, err := loader.FindFile("GeoLite2-City-Blocks-IPv6.csv", reader)
	if err != nil {
		log.Println(err)
		log.Println("Failed to create io.ReaderCloser")
	}
	defer rcIPv6.Close()
	// TODO: update tests to use high level data loader functions instead of low level funcs
	gl2ipv6, err = parser.LoadIPListGLite2(rcIPv6, idMap)
	if err != nil {
		log.Println(err)
		log.Println("Failed to create ipv6")
	}
}
