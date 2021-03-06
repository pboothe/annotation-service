package parser

// Notes:
//   GeoLite1 is IPv4 only.
//   GeoLite2 supports IPv4 and IPv6.

import (
	"encoding/binary"
	"encoding/csv"
	"errors"
	"io"
	"log"
	"math"
	"net"
	"reflect"
	"strconv"
)

const (
	ipNumColumnsGlite1       = 3
	locationNumColumnsGlite1 = 9
	gLite1Prefix             = "GeoLiteCity"
)

// Glite1HelpNode defines IPNode data defined inside
// GeoLite1 Location files
type gLite1HelpNode struct {
	Latitude   float64
	Longitude  float64
	PostalCode string
}

// TODO: Add equivalent of LoadGeoLite2

// Create Location list, map, and Glite1HelpNode for GLite1 databases
// GLiteHelpNode contains information to help populate fields in IPNode
func LoadLocListGLite1(reader io.Reader) ([]LocationNode, []gLite1HelpNode, map[int]int, error) {
	r := csv.NewReader(reader)
	idMap := make(map[int]int, mapMax)
	list := []LocationNode{}
	glite := []gLite1HelpNode{}
	// Skip the first 2 lines
	_, err := r.Read()
	if err == io.EOF {
		log.Println("Empty input data")
		return nil, nil, nil, errors.New("Empty input data")
	}
	_, err = r.Read()
	if err == io.EOF {
		log.Println("Empty input data")
		return nil, nil, nil, errors.New("Empty input data")
	}
	// Sets csv's internal column length check
	r.FieldsPerRecord = locationNumColumnsGlite1
	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else if len(record) != r.FieldsPerRecord {
				log.Println(err)
				log.Println("\twanted: ", locationNumColumnsGlite1, " got: ", len(record), record)
				return nil, nil, nil, errors.New("Corrupted Data: wrong number of columns")
			} else {
				log.Println(err, ": ", record)
				return nil, nil, nil, errors.New("Error reading file")
			}
		}
		var lNode LocationNode
		lNode.GeonameID, err = strconv.Atoi(record[0])
		if err != nil {
			return nil, nil, nil, errors.New("Corrupted Data: GeonameID should be a number")
		}
		lNode.CountryCode, err = checkCaps(record[1], "Country code")
		if err != nil {
			return nil, nil, nil, err
		}
		lNode.CityName = record[3]
		var gNode gLite1HelpNode
		gNode.PostalCode = record[4]
		gNode.Latitude, err = stringToFloat(record[5], "Latitude")
		if err != nil {
			return nil, nil, nil, err
		}
		gNode.Longitude, err = stringToFloat(record[6], "Longitude")
		if err != nil {
			return nil, nil, nil, err
		}

		list = append(list, lNode)
		glite = append(glite, gNode)
		idMap[lNode.GeonameID] = len(list) - 1
	}
	return list, glite, idMap, nil
}

// LoadIPListGLite1 creates a List of IPNodes from a GeoLite1 reader
// Note that GeoLite1 is IPv4 only.
// TODO(gfr) Update to use recursion instead of stack.
func LoadIPListGLite1(reader io.Reader, idMap map[int]int, glite1 []gLite1HelpNode) ([]IPNode, error) {
	g1IP := []string{"startIpNum", "endIpNum", "locId"}
	list := []IPNode{}
	stack := []IPNode{}
	r := csv.NewReader(reader)
	// Skip first line
	title, err := r.Read()
	if err == io.EOF {
		log.Println("Empty input data")
		return nil, errors.New("Empty input data")
	}
	// Skip 2nd line, which contains column labels
	title, err = r.Read()
	if err == io.EOF {
		log.Println("Empty input data")
		return nil, errors.New("Empty input data")
	}
	if !reflect.DeepEqual(g1IP, title) {
		log.Println("Improper data format got: ", title, " wanted: ", g1IP)
		return nil, errors.New("Improper data format")
	}

	for {
		var newNode IPNode
		// Example:
		// GLite1 : record = [16777216,16777471,17]
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		err = checkNumColumns(record, ipNumColumnsGlite1)
		if err != nil {
			return nil, err
		}
		newNode.IPAddressLow, err = intToIPv4(record[0])
		if err != nil {
			return nil, err
		}
		newNode.IPAddressHigh, err = intToIPv4(record[1])
		if err != nil {
			return nil, err
		}
		// Look for GeoId within idMap and return index
		index, err := lookupGeoId(record[2], idMap)
		if err != nil {
			return nil, err
		}
		newNode.LocationIndex = index
		newNode.Latitude = glite1[index].Latitude
		newNode.Longitude = glite1[index].Longitude
		newNode.PostalCode = glite1[index].PostalCode
		stack, list = handleStack(stack, list, newNode)
	}
	var pop IPNode
	pop, stack = stack[len(stack)-1], stack[:len(stack)-1]
	for ; len(stack) > 0; pop, stack = stack[len(stack)-1], stack[:len(stack)-1] {
		peek := stack[len(stack)-1]
		peek.IPAddressLow = PlusOne(pop.IPAddressHigh)
		list = append(list, peek)
	}
	return list, nil
}

// Converts integer to net.IPv4
func intToIPv4(str string) (net.IP, error) {
	num, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		log.Println("Provided IP should be a number")
		return nil, errors.New("Inputed string cannot be converted to a number")
	}
	// TODO: get rid of floating point
	ft := float64(num)
	if ft > math.Pow(2, 32) || num < 0 {
		log.Println("Provided IP should be in the range of 0.0.0.1 and 255.255.255.255 ", str)
	}
	ip := make(net.IP, 4)
	// Split number into array of bytes
	binary.BigEndian.PutUint32(ip, uint32(num))
	ip = net.IPv4(ip[0], ip[1], ip[2], ip[3])
	return ip, nil
}
