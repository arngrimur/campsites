package main

import (
	"fmt"
	"strings"
	"strconv"
	"math"
	"bufio"
	"os"
	"log"
)

func main()  {
	//coordinate := "48° 43.135 N x 122° 30.947 W"
	//parts:=strings.Split(coordinate, " ")
	ReadLines()
	//getAsWgs84("48.0", "1.2")
	
	//lon:=getAsWgs84(parts[0], parts[1])
	//lat:= getAsWgs84(parts[4], parts[5])
	//fmt.Printf("%f -%f\n",lon,lat)
}

func ReadLines(){
	file,err := os.Open("./camp-coordinates-dwyer.txt")
	if err != nil {
		log.Fatal(err)
		defer file.Close()
	} else {
		scanner:= bufio.NewScanner(file)
			fmt.Printf("<gpx xmlns=\"http://www.topografix.com/GPX/1/1\" \n  xmlns:gpxx=\"http://www.garmin.com/xmlschemas/GpxExtensions/v3\" \n xmlns:gpxtpx=\"http://www.garmin.com/xmlschemas/TrackPointExtension/v1\"  \n creator=\"Oregon 400t\" \n  version=\"1.1\"  xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n  xsi:schemaLocation=\"http://www.topografix.com/GPX/1/1 http://www.topografix.com/GPX/1/1/gpx.xsd\n http://www.garmin.com/xmlschemas/GpxExtensions/v3 http://www.garmin.com/xmlschemas/GpxExtensionsv3.xsd\n http://www.garmin.com/xmlschemas/TrackPointExtension/v1\n http://www.garmin.com/xmlschemas/TrackPointExtensionv1.xsd\">\n <metadata>\n    <link href=\"http://www.garmin.com\">\n      <text>Garmin International</text>\n    </link>\n    <time>2009-10-17T22:58:43Z</time>\n  </metadata>\n")

		for scanner.Scan(){
			line := scanner.Text()
			name, coordinate := GetNameAndCoordinate(line)
		
			parts := strings.Split(coordinate, " ")
			lat :=  GetAsWgs84(parts[0], parts[1])
			lon := GetAsWgs84(parts[4], parts[5])
			fmt.Printf("<wpt lat=\"%f\" lon=\"-%f\"><name>%s</name></wpt>\n",lat,lon,name)
			

		}
		fmt.Printf("</gpx>");
		if scanner.Err() != nil {
			log.Fatal(scanner.Err())
		}
	}
	
}

func GetNameAndCoordinate(line string) (name , coordinate string) {
	firstNumber := strings.IndexAny(line, "0123456789")
	name = line[:firstNumber -1]
	coordinate = line[firstNumber:]
	return
}

func GetAsWgs84(degree string, minSec string )float64  {
	whole :=strings.Split(degree, "°")
	dd, ddErr:=strconv.ParseFloat(whole[0], 32)
	part, partErr := strconv.ParseFloat(minSec, 32)
	if(ddErr != nil || partErr != nil) {
		return math.NaN()
	}
	return dd + part / 60.0
}
