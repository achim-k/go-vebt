package vebt

import (
	"fmt"
	"math"
)

type VEB struct {
	u, min, max int    //universe size, min-, max value
	summary     *VEB   //pointer to summary
	cluster     []*VEB // array of pointers to each child cluster
}

func CreateTree(u int) *VEB {
	if u < 2 {
		fmt.Println("Tree universe size u needs to be bigger than 2")
		return nil
	}

	pow := math.Log2(float64(u))
	if math.Trunc(pow) != pow {
		fmt.Println("Tree universe size u not power of 2 (u != 2^x)")
		return nil
	}

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
		V.cluster = append(V.cluster, CreateTree(clusterUnivSize))
	}

	// create summary
	summaryUnivSize := HigherSqrt(u)
	V.summary = CreateTree(summaryUnivSize)

	return V
}

func (V VEB) Max() int { return V.max }

func (V VEB) Min() int { return V.min }

func (V VEB) High(x int) int {
	return int(math.Floor(float64(x) / float64(LowerSqrt(V.u))))
}

func (V VEB) Low(x int) int {
	return x % LowerSqrt(V.u)
}

func (V VEB) Index(x, y int) int {
	return x*LowerSqrt(V.u) + y
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

func (V *VEB) Delete(x int) {
	if V.summary == nil || V.summary.Min() == -1 {
		// No nonempty cluster
		if x == V.min && x == V.max {
			// only element of V
			V.min, V.max = -1, -1
		} else if x == V.min {
			//2 elements in V: x is min element
			V.min = V.max
		} else {
			//2 elements in V: x is max element
			V.max = V.min
		}

	} else {
		// some nonempty cluster
		if x == V.min {
			// get smallest element in cluster
			y := V.Index(V.summary.min, V.cluster[V.summary.min].min)
			V.min = y
			// delete element from cluster
			V.cluster[V.High(y)].Delete(V.Low(y))
			if V.cluster[V.High(y)].min == -1 {
				// it was the only element in the cluster, update summary for this cluster
				V.summary.Delete(V.High(y))
			}
		} else if x == V.max {
			y := V.Index(V.summary.max, V.cluster[V.summary.max].max)
			// delete element from cluster
			V.cluster[V.High(y)].Delete(V.Low(y))
			if V.cluster[V.High(y)].min == -1 {
				// it was the only element in the cluster, update summary for this cluster
				V.summary.Delete(V.High(y))
			}

			if V.summary == nil || V.summary.min == -1 {
				// no summary anymore
				if V.min == y {
					V.min, V.max = -1, -1
				} else {
					V.max = V.min
				}
			} else {
				V.max = V.Index(V.summary.max, V.cluster[V.summary.max].max)
			}

		} else {
			// neither min nor max
			V.cluster[V.High(x)].Delete(V.Low(x))
			if V.cluster[V.High(x)].min == -1 {
				// cluster became empty, update summary
				V.summary.Delete(V.High(x))
			}
		}

	}
}

func (V VEB) Successor(x int) int {
	if V.u == 2 {
		if x == 0 && V.max == 1 {
			return 1
		} else {
			return -1
		}
	} else if V.min != -1 && x < V.min {
		return V.min
	} else {
		maxLow := V.cluster[V.High(x)].Max()
		if maxLow != -1 && V.Low(x) < maxLow {
			offset := V.cluster[V.High(x)].Successor(V.Low(x))
			return V.Index(V.High(x), offset)
		} else {
			succCluster := V.summary.Successor(V.High(x))
			if succCluster == -1 {
				return -1
			} else {
				offset := V.cluster[succCluster].Min()
				return V.Index(succCluster, offset)
			}
		}
	}
}

func (V VEB) Predecessor(x int) int {
	if V.u == 2 {
		if x == 1 && V.min == 0 {
			return 0
		} else {
			return -1
		}
	} else if V.max != -1 && x > V.max {
		return V.max
	} else {
		minLow := V.cluster[V.High(x)].Min()
		if minLow != -1 && V.Low(x) > minLow {
			offset := V.cluster[V.High(x)].Predecessor(V.Low(x))
			return V.Index(V.High(x), offset)
		} else {
			predCluster := V.summary.Predecessor(V.High(x))
			if predCluster == -1 {
				if V.min != -1 && x > V.min {
					return V.min
				} else {
					return -1
				}
			} else {
				offset := V.cluster[predCluster].Max()
				return V.Index(predCluster, offset)
			}
		}
	}
}

// Counts all struct objects of the tree (for testing purposes)
func (V VEB) Count() int {
	if V.u == 2 {
		return 1
	}

	sum := 1 // including structure itself
	for i := 0; i < len(V.cluster); i++ {
		sum += V.cluster[i].Count()
	}
	sum += V.summary.Count()
	return sum
}

// Prints the tree to std out (for debugging purposes only)
func (V VEB) Print() {
	//function just acts as wrapper, since default parameters are not possible in go
	V.PrintFunc(0, 0)
}

func (V VEB) PrintFunc(level, clusterNo int) {
	spacer := ""
	for i := 0; i < level; i++ {
		spacer += "\t"
	}

	if level == 0 {
		fmt.Printf("%vR: {u: %v, min: %v, max: %v, clusters: %v}\n", spacer, V.u, V.min, V.max, len(V.cluster))
	} else {
		fmt.Printf("%vC[%v]: {u: %v, min: %v, max: %v, clusters: %v}\n", spacer, clusterNo, V.u, V.min, V.max, len(V.cluster))
	}

	if len(V.cluster) > 0 {
		fmt.Printf("%v\tS:    {u: %v, min: %v, max: %v, clusters: %v}\n", spacer, V.summary.u, V.summary.min, V.summary.max, len(V.summary.cluster))
		for i := 0; i < len(V.summary.cluster); i++ {
			V.summary.cluster[i].PrintFunc(level+2, i)
		}
		for i := 0; i < len(V.cluster); i++ {
			V.cluster[i].PrintFunc(level+1, i)
		}
	}
}

// Counts all struct objects of the tree (for testing purposes)
func (V *VEB) Clear() {
	V.min, V.max = -1, -1

	if V.u == 2 {
		return
	}

	// clear clusters
	for i := 0; i < len(V.cluster); i++ {
		V.cluster[i].Clear()
	}
	// clear summary
	V.summary.Clear()
}

// Fills tree with all keys
func (V *VEB) Fill() {
	for i := 0; i < V.u; i++ {
		V.Insert(i)
	}
}

// Calculate lower square root (helper function)
func LowerSqrt(u int) int {
	return int(math.Pow(2.0, math.Floor(math.Log2(float64(u))/2)))
}

// Calculate higher square root (helper function)
func HigherSqrt(u int) int {
	return int(math.Pow(2.0, math.Ceil(math.Log2(float64(u))/2)))
}
