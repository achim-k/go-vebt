package vebt

import (
	//"fmt"
	"math"
)

type VEB struct {
    u, min, max int //universe size, min-, max value
    summary *VEB 	//pointer to summary
    cluster []*VEB 	// array of pointers to each child cluster
}

func (V VEB) Max() int { return V.max }
func (V VEB) Min() int { return V.min }
func (V VEB) High(x int) int { 
	return int(math.Floor(float64(x)/float64(LowerSqrt(V.u))))
}
func (V VEB) Low(x int) int { 
	return x % LowerSqrt(V.u)
}
func (V VEB) Index(x, y int) int { 
	return x * LowerSqrt(V.u) + y
}
func (V VEB) IsMember(x int) bool {
	if x == V.min || x == V.max {
		return true
	} else if V.u == 2 {
        return false
    } else {
        return V.cluster[V.High(x)].IsMember(V.Low(x))
    } 
}
func (V *VEB) Insert(x int) {
	if V.min == -1 {
		V.min, V.max = x, x
	} else {
		if x < V.min {
			temp := V.min
			V.min = x
			x = temp
		}
		if V.u > 2 {
			if V.cluster[V.High(x)].Min() == -1 {
				V.summary.Insert(V.High(x))
				V.cluster[V.High(x)].min, V.cluster[V.High(x)].max = V.Low(x), V.Low(x)
			} else {
				V.cluster[V.High(x)].Insert(V.Low(x))
			}
		}
		if x > V.max {
			V.max = x
		}
	}
}

func CreateVEBTree(u int, testCounter *int) *VEB {
	*testCounter++

	V := new(VEB)
	V.min, V.max = -1, -1
	V.u = u

	if u == 2 {
		return V
	}

	// create clusters
	clusterCount := HigherSqrt(u)
	clusterUnivSize := LowerSqrt(u)
	for i := 0; i < clusterCount; i++ {
		V.cluster = append(V.cluster, CreateVEBTree(clusterUnivSize, testCounter))
	}

	// create summary
	summaryUnivSize := HigherSqrt(u)
	V.summary = CreateVEBTree(summaryUnivSize, testCounter)
	
	return V
}


// Calculate lower square root (helper function)
func LowerSqrt(u int) int {
	return int(math.Pow(2.0, math.Floor(math.Log2(float64(u)) / 2)))
}
// Calculate higher square root (helper function)
func HigherSqrt(u int) int {
	return int(math.Pow(2.0, math.Ceil(math.Log2(float64(u)) / 2)))
}
