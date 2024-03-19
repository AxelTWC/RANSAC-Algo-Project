// Student Name: Axel Tang
// Student Number: 300164095
package main

import (
	"bufio"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Point3D struct {
	X float64
	Y float64
	Z float64
}
type Plane3D struct {
	A float64
	B float64
	C float64
	D float64
}
type Plane3DwSupport struct {
	Plane3D
	SupportSize int
}

func ReadXYZ(filename string) []Point3D { //Reads the XYZ file
	var countSize int = 0
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// s := strings.Split(scanner.Text(), " ")
		// fmt.Print(s)
		countSize++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	countSize = countSize - 1

	point3DArray := make([]Point3D, 0)

	// test := len(point3DArray)
	// fmt.Print("test")
	// fmt.Print(test)

	g, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer g.Close()

	actual := bufio.NewScanner(g)

	for actual.Scan() {
		var point Point3D
		s := strings.Split(actual.Text(), "	")
		// fmt.Print(s)
		if s[0] != "X" && s[1] != "Y" && s[2] != "Z" {
			p1, err := strconv.ParseFloat(s[0], 64)
			p2, err := strconv.ParseFloat(s[1], 64)
			p3, err := strconv.ParseFloat(s[2], 64)
			point.X, point.Y, point.Z = p1, p2, p3
			point3DArray = append(point3DArray, point)
			// fmt.Print(point3DArray)
			if err != nil {
				// fmt.Println(err)
			} else {
				// fmt.Println("N/A")
			}
		}
	}

	point3DArray = point3DArray[1:]

	return point3DArray
}

func SaveXYZ(filename string, points []Point3D) { //Saves into the XYZ file
	e := os.Remove(filename)
	if e != nil {
		file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("failed creating file: %s", err)
		}

		datawriter := bufio.NewWriter(file)

		for _, data := range points {
			p1data := data.X
			p2data := data.Y
			p3data := data.Z

			p1dataString := strconv.FormatFloat(p1data, 'g', 9, 64)
			p2dataString := strconv.FormatFloat(p2data, 'g', 9, 64)
			p3dataString := strconv.FormatFloat(p3data, 'g', 9, 64)
			_, _ = datawriter.WriteString(p1dataString + " " + p2dataString + " " + p3dataString + "\n")
		}
		datawriter.Flush()
		file.Close()
	} else {

		file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("failed creating file: %s", err)
		}

		datawriter := bufio.NewWriter(file)
		datawriter.WriteString("x " + "y " + "z" + "\n")
		for _, data := range points {
			p1data := data.X
			p2data := data.Y
			p3data := data.Z

			p1dataString := strconv.FormatFloat(p1data, 'g', 9, 64)
			p2dataString := strconv.FormatFloat(p2data, 'g', 9, 64)
			p3dataString := strconv.FormatFloat(p3data, 'g', 9, 64)
			_, _ = datawriter.WriteString(p1dataString + " " + p2dataString + " " + p3dataString + "\n")
		}
		datawriter.Flush()
		file.Close()
	}

}
func (p1 *Point3D) GetDistance(p2 *Point3D) float64 { //Gets the distance
	distanceOfTwoPoints := math.Sqrt(math.Pow((p2.X-p1.X), 2) + math.Pow((p2.Y-p1.Y), 2) + math.Pow((p2.Z-p1.Z), 2))
	return distanceOfTwoPoints
}
func GetPlane(points []Point3D) Plane3D { //Gets the plane
	var plane3D Plane3D
	ax := points[1].X - points[0].X
	ay := points[1].Y - points[0].Y
	az := points[1].Z - points[0].Z

	bx := points[2].X - points[0].X
	by := points[2].Y - points[0].Y
	bz := points[2].Z - points[0].Z

	px := points[0].X
	py := points[0].Y
	pz := points[0].Z

	crossProduct := make([]float64, 4)
	crossProduct[0] = (ay)*(bz) - (az)*(by)
	crossProduct[1] = (az)*(bx) - (ax)*(bz)
	crossProduct[2] = (ax)*(by) - (ay)*(bx)
	crossProduct[3] = (-crossProduct[0] * px) - (crossProduct[1] * py) - (crossProduct[2] * pz)

	plane3D.A, plane3D.B, plane3D.C, plane3D.D = crossProduct[0], crossProduct[1], crossProduct[2], crossProduct[3]

	return plane3D
}

func GetNumberOfIterations(confidence float64, percentageOfPointsOnPlane float64) int { //Gets the Number Of Iteartions
	var iteration float64
	iteration = math.Ceil(math.Log(1-confidence) / math.Log(1-math.Pow(percentageOfPointsOnPlane, 3)))
	roundedIteration := int(math.Round(iteration*100) / 100)
	return roundedIteration
}

func GetSupport(plane Plane3D, points []Point3D, eps float64) Plane3DwSupport { //Gets the Supporting value
	var plane3DwSupport Plane3DwSupport
	plane3DwSupport.A, plane3DwSupport.B, plane3DwSupport.C, plane3DwSupport.D = plane.A, plane.B, plane.C, plane.D
	var counter int = 0
	for _, data := range points {
		p1data := data.X
		p2data := data.Y
		p3data := data.Z
		distance := math.Abs(plane3DwSupport.A*p1data+plane3DwSupport.B*p2data+plane3DwSupport.C*p3data+plane3DwSupport.D) / math.Sqrt(math.Pow(plane3DwSupport.A, 2)+math.Pow(plane3DwSupport.B, 2)+math.Pow(plane3DwSupport.C, 2))

		if distance < eps {
			counter++
		}
	}

	plane3DwSupport.SupportSize = counter

	return plane3DwSupport

}

func GetSupportingPoints(plane Plane3D, points []Point3D, eps float64) []Point3D { //Gets the Supporting Points
	var plane3DwSupport Plane3DwSupport
	var pointSupport Point3D
	plane3DwSupport.A, plane3DwSupport.B, plane3DwSupport.C, plane3DwSupport.D = plane.A, plane.B, plane.C, plane.D
	supportingPointsArray := make([]Point3D, 0)
	for _, data := range points {
		p1data := data.X
		p2data := data.Y
		p3data := data.Z

		pointSupport.X = data.X
		pointSupport.Y = data.Y
		pointSupport.Z = data.Z

		distance := math.Abs(plane3DwSupport.A*p1data+plane3DwSupport.B*p2data+plane3DwSupport.C*p3data+plane3DwSupport.D) / math.Sqrt(math.Pow(plane3DwSupport.A, 2)+math.Pow(plane3DwSupport.B, 2)+math.Pow(plane3DwSupport.C, 2))
		if distance < eps {
			supportingPointsArray = append(supportingPointsArray, pointSupport)
		}
	}
	supportingPointsArray = supportingPointsArray[1:]
	return supportingPointsArray

}

func RemovePlane(plane Plane3D, points []Point3D, eps float64) []Point3D { //Removes the current plane
	var plane3DwSupport Plane3DwSupport
	var pointNonSupport Point3D
	plane3DwSupport.A, plane3DwSupport.B, plane3DwSupport.C, plane3DwSupport.D = plane.A, plane.B, plane.C, plane.D
	supportingPointsArray := make([]Point3D, 0)
	for _, data := range points {
		p1data := data.X
		p2data := data.Y
		p3data := data.Z

		pointNonSupport.X = data.X
		pointNonSupport.Y = data.Y
		pointNonSupport.Z = data.Z

		distance := math.Abs(plane3DwSupport.A*p1data+plane3DwSupport.B*p2data+plane3DwSupport.C*p3data+plane3DwSupport.D) / math.Sqrt(math.Pow(plane3DwSupport.A, 2)+math.Pow(plane3DwSupport.B, 2)+math.Pow(plane3DwSupport.C, 2))
		if distance > eps {
			supportingPointsArray = append(supportingPointsArray, pointNonSupport)
		}
	}
	supportingPointsArray = supportingPointsArray[1:]

	return supportingPointsArray

}

// Start of PipeLine
func randomPointGenerator(pt []Point3D) <-chan Point3D {
	// defer wg.Done()
	pointRandom := make(chan Point3D)
	go func() {
		for {
			randomIndex := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(pt))
			pointRandom <- pt[randomIndex]
			time.Sleep(1 * time.Millisecond)
		}
	}()
	return pointRandom
}

func tripletOfPointsGenerator(in <-chan Point3D) chan [3]Point3D {
	triplePointsChannel := make(chan [3]Point3D)
	go func() {
		var triplePointsSlice []Point3D
		for point := range in {
			triplePointsSlice = append(triplePointsSlice, point)
			if len(triplePointsSlice) == 3 {
				triplePointsChannel <- [3]Point3D{triplePointsSlice[0], triplePointsSlice[1], triplePointsSlice[2]}
				triplePointsSlice = nil
			}
		}
		if len(triplePointsSlice) > 0 {
			for i := len(triplePointsSlice); i < 3; i++ {
				triplePointsSlice = append(triplePointsSlice, Point3D{0, 0, 0})
			}
			time.Sleep(1 * time.Millisecond)
			triplePointsChannel <- [3]Point3D{triplePointsSlice[0], triplePointsSlice[1], triplePointsSlice[2]}
		}
		close(triplePointsChannel)
	}()

	return triplePointsChannel
}

func takeN(n int, in <-chan [3]Point3D) chan [3]Point3D {
	out := make(chan [3]Point3D)
	go func() {
		for i := 0; i < n; i++ {
			arr := <-in
			out <- arr
		}
		close(out)
	}()
	return out
}

func planeEstimator(forPlane chan [3]Point3D) chan Plane3D {
	out := make(chan Plane3D)
	// GetPlane(<-forPlane[:])
	go func() {
		for points := range forPlane {
			var plane3D Plane3D
			ax := points[1].X - points[0].X
			ay := points[1].Y - points[0].Y
			az := points[1].Z - points[0].Z

			bx := points[2].X - points[0].X
			by := points[2].Y - points[0].Y
			bz := points[2].Z - points[0].Z

			px := points[0].X
			py := points[0].Y
			pz := points[0].Z

			crossProduct := make([]float64, 4)
			crossProduct[0] = (ay)*(bz) - (az)*(by)
			crossProduct[1] = (az)*(bx) - (ax)*(bz)
			crossProduct[2] = (ax)*(by) - (ay)*(bx)
			crossProduct[3] = (-crossProduct[0] * px) - (crossProduct[1] * py) - (crossProduct[2] * pz)

			plane3D.A, plane3D.B, plane3D.C, plane3D.D = crossProduct[0], crossProduct[1], crossProduct[2], crossProduct[3]

			out <- plane3D
		}
		close(out)
	}()
	return out
}

func supportingPointFinder(plane Plane3D, pt []Point3D) chan Plane3DwSupport {
	defer wg.Done()
	var plane3DwSupport Plane3DwSupport
	plane3DwSupport = GetSupport(plane, pt, 0.1)
	out := make(chan Plane3DwSupport)
	out <- plane3DwSupport
	return out
}

func fanIn(in chan Plane3DwSupport) chan Plane3DwSupport {
	out := make(chan Plane3DwSupport)
	var wg sync.WaitGroup
	fanInFunc := func(in <-chan Plane3DwSupport) {
		defer wg.Done()
		for p := range in {
			out <- p
		}
	}
	wg.Add(len(in))

	go fanInFunc(in)

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func dominantPlaneIdentifier(in chan Plane3DwSupport, bestPlane *Plane3DwSupport) Plane3D {
	defer wg.Done()
	for p := range in {
		if p.SupportSize > bestPlane.SupportSize {
			*bestPlane = p
		}
	}
	return bestPlane.Plane3D
}

// End of PipeLine

func main() { //Feel free to uncomment to test out functions here

	// fmt.Print(ReadXYZ("PointCloud1.xyz"))
	// ReadXYZ("PointCloud1.xyz")
	// SaveXYZ("test.xyz", ReadXYZ("PointCloud1.xyz"))
	// var points []Point3D
	// var a, b, c Point3D
	// a.X = 1
	// a.Y = 2
	// a.Z = -2
	// b.X = 3
	// b.Y = -2
	// b.Z = 1
	// c.X = 5
	// c.Y = 1
	// c.Z = -4
	// points = append(points, a, b, c)
	// fmt.Print(GetPlane(points))
	// wg.Add(1)
	// tripletOfPointsGenerator(pt)

	// parseN := make(chan [3]Point3D)
	// for i := 0; i < 3; i++ {
	// 	parseN <- <-tripletOfPointsGenerator(pt)
	// }
	// go takeN(parseN)
	// close(parseN)
	// fmt.Print(parseN)
	// wg.Wait()

	// fmt.Print(randomPointGenerator(pt))
	// fmt.Print(tripletOfPointsGenerator(pt))
	// var bestSupport Plane3DwSupport
	// bestSupport.A = 0
	// bestSupport.B = 0
	// bestSupport.C = 0
	// bestSupport.D = 0
	// bestSupport.Plane3D = Plane3D{0, 0, 0, 0}
	// bestSupport.SupportSize = 0

	// points2 := ReadXYZ("PointCloud2.xyz")

	// SaveXYZ("test1.xyz", points2)
	// fmt.Print(GetSupport(GetPlane(points2), points2, 0.02))
	// nIterations := GetNumberOfIterations(0.99, 0.05)
	// fmt.Print(nIterations)

	pt := ReadXYZ("PointCloud1.xyz")
	randomPointGen := randomPointGenerator(pt)
	triplePoints := tripletOfPointsGenerator(randomPointGen)
	out := takeN(GetNumberOfIterations(0.99, 0.05), triplePoints)
	planeEstimatorTemp := planeEstimator(out)
	var bestSupport Plane3DwSupport
	bestSupport.A = 0
	bestSupport.B = 0
	bestSupport.C = 0
	bestSupport.D = 0
	bestSupport.Plane3D = Plane3D{0, 0, 0, 0}
	bestSupport.SupportSize = 0
	supportPointChan := make(chan Plane3DwSupport)
	go func() {
		defer wg.Done()
		supportPointFinder := supportingPointFinder(<-planeEstimatorTemp, pt)
		supportPoints := <-supportPointFinder
		supportPointChan <- supportPoints
		fanInTemp := fanIn(supportPointChan)
		wg.Add(1)
		plane := dominantPlaneIdentifier(fanInTemp, &bestSupport)
		SaveXYZ("PointCloud1_p1.xyz", GetSupportingPoints(plane, pt, 0.1))

		time.Sleep(1 * time.Millisecond)
	}()
	wg.Wait()
	go func() {
		defer wg.Done()
		supportPointFinder := supportingPointFinder(<-planeEstimatorTemp, pt)
		supportPoints := <-supportPointFinder
		supportPointChan <- supportPoints
		fanInTemp := fanIn(supportPointChan)
		wg.Add(1)
		plane := dominantPlaneIdentifier(fanInTemp, &bestSupport)
		SaveXYZ("PointCloud1_p2.xyz", GetSupportingPoints(plane, pt, 0.1))
		RemovePlane(plane, pt, 0.1)
		time.Sleep(1 * time.Millisecond)
	}()
	wg.Wait()
	go func() {
		defer wg.Done()
		supportPointFinder := supportingPointFinder(<-planeEstimatorTemp, pt)
		supportPoints := <-supportPointFinder
		supportPointChan <- supportPoints
		fanInTemp := fanIn(supportPointChan)
		wg.Add(1)
		plane := dominantPlaneIdentifier(fanInTemp, &bestSupport)
		SaveXYZ("PointCloud1_p3.xyz", GetSupportingPoints(plane, pt, 0.1))
		p0plane := RemovePlane(plane, pt, 0.1)
		SaveXYZ("PointCloud1_p0.xyz", p0plane)
		time.Sleep(1 * time.Millisecond)
	}()
	wg.Wait()

	pt2 := ReadXYZ("PointCloud2.xyz")
	randomPointGen2 := randomPointGenerator(pt2)
	triplePoints2 := tripletOfPointsGenerator(randomPointGen2)
	out2 := takeN(GetNumberOfIterations(0.99, 0.05), triplePoints2)
	planeEstimator2 := planeEstimator(out2)
	var bestSupport2 Plane3DwSupport
	bestSupport2.A = 0
	bestSupport2.B = 0
	bestSupport2.C = 0
	bestSupport2.D = 0
	bestSupport2.Plane3D = Plane3D{0, 0, 0, 0}
	bestSupport2.SupportSize = 0
	supportPointChan2 := make(chan Plane3DwSupport)
	go func() {
		defer wg.Done()
		supportPointFinder := supportingPointFinder(<-planeEstimator2, pt)
		supportPoints := <-supportPointFinder
		supportPointChan2 <- supportPoints
		fanInTemp := fanIn(supportPointChan2)
		wg.Add(1)
		plane := dominantPlaneIdentifier(fanInTemp, &bestSupport)
		SaveXYZ("PointCloud2_p1.xyz", GetSupportingPoints(plane, pt, 0.1))
		RemovePlane(plane, pt, 0.1)
		time.Sleep(1 * time.Millisecond)
	}()
	wg.Wait()
	go func() {
		defer wg.Done()
		supportPointFinder := supportingPointFinder(<-planeEstimator2, pt)
		supportPoints := <-supportPointFinder
		supportPointChan2 <- supportPoints
		fanInTemp := fanIn(supportPointChan2)
		wg.Add(1)
		plane := dominantPlaneIdentifier(fanInTemp, &bestSupport)
		SaveXYZ("PointCloud2_p2.xyz", GetSupportingPoints(plane, pt, 0.1))
		RemovePlane(plane, pt, 0.1)
		time.Sleep(1 * time.Millisecond)
	}()
	wg.Wait()
	go func() {
		defer wg.Done()
		supportPointFinder := supportingPointFinder(<-planeEstimator2, pt)
		supportPoints := <-supportPointFinder
		supportPointChan2 <- supportPoints
		fanInTemp := fanIn(supportPointChan2)
		wg.Add(1)
		plane := dominantPlaneIdentifier(fanInTemp, &bestSupport)
		SaveXYZ("PointCloud2_p3.xyz", GetSupportingPoints(plane, pt, 0.1))
		p0plane := RemovePlane(plane, pt, 0.1)
		SaveXYZ("PointCloud1_p0.xyz", p0plane)
		time.Sleep(1 * time.Millisecond)
	}()
	wg.Wait()

	pt3 := ReadXYZ("PointCloud3.xyz")
	randomPointGen3 := randomPointGenerator(pt3)
	triplePoints3 := tripletOfPointsGenerator(randomPointGen3)
	out3 := takeN(GetNumberOfIterations(0.99, 0.05), triplePoints3)
	planeEstimator3 := planeEstimator(out3)
	var bestSupport3 Plane3DwSupport
	bestSupport3.A = 0
	bestSupport3.B = 0
	bestSupport3.C = 0
	bestSupport3.D = 0
	bestSupport3.Plane3D = Plane3D{0, 0, 0, 0}
	bestSupport3.SupportSize = 0
	supportPointChan3 := make(chan Plane3DwSupport)
	go func() {
		defer wg.Done()
		supportPointFinder := supportingPointFinder(<-planeEstimator3, pt)
		supportPoints := <-supportPointFinder
		supportPointChan3 <- supportPoints
		fanInTemp := fanIn(supportPointChan3)
		wg.Add(1)
		plane := dominantPlaneIdentifier(fanInTemp, &bestSupport)
		SaveXYZ("PointCloud3_p1.xyz", GetSupportingPoints(plane, pt, 0.1))
		RemovePlane(plane, pt, 0.1)
		time.Sleep(1 * time.Millisecond)
	}()
	wg.Wait()
	go func() {
		defer wg.Done()
		supportPointFinder := supportingPointFinder(<-planeEstimator3, pt)
		supportPoints := <-supportPointFinder
		supportPointChan2 <- supportPoints
		fanInTemp := fanIn(supportPointChan3)
		wg.Add(1)
		plane := dominantPlaneIdentifier(fanInTemp, &bestSupport)
		SaveXYZ("PointCloud3_p2.xyz", GetSupportingPoints(plane, pt, 0.1))
		RemovePlane(plane, pt, 0.1)
		time.Sleep(1 * time.Millisecond)
	}()
	wg.Wait()
	go func() {
		defer wg.Done()
		supportPointFinder := supportingPointFinder(<-planeEstimator3, pt)
		supportPoints := <-supportPointFinder
		supportPointChan2 <- supportPoints
		fanInTemp := fanIn(supportPointChan2)
		wg.Add(1)
		plane := dominantPlaneIdentifier(fanInTemp, &bestSupport)
		SaveXYZ("PointCloud3_p3.xyz", GetSupportingPoints(plane, pt, 0.1))
		p0plane := RemovePlane(plane, pt, 0.1)
		SaveXYZ("PointCloud1_p0.xyz", p0plane)
		time.Sleep(1 * time.Millisecond)
	}()
	wg.Wait()

	//Print out supporting points
	// for arr := range out {
	// 	fmt.Println(arr)
	// }
	// for arr := range out2 {
	// 	fmt.Println(arr)
	// }
	// for arr := range out3 {
	// 	fmt.Println(arr)
	// }
	//Print out plane equation
	// for plane := range planeEstimatorTemp {
	// 	fmt.Println(plane)
	// }
	// for plane2 := range planeEstimator2 {
	// 	fmt.Println(plane2)
	// }
	// for plane3 := range planeEstimator3 {
	// 	fmt.Println(plane3)
	// }
}
