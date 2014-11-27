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

// Creates an tree with a universe size (u) with u=2^x, where x is chosen,
// so that 2^(x-1) < size <= 2^x. Returns the pointer on the root element
// or nil on failure
func CreateTree(size int) *VEB {
	if size < 0 {
		return nil
	}

	// choose x so that 2^(x-1) < u <= 2^x
	x := math.Ceil(math.Log2(float64(size)))
	u := int(math.Pow(2, x))

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

// Returns the maximum of the tree
func (V VEB) Max() int { return V.max }

// Returns the minimum of the tree
func (V VEB) Min() int { return V.min }

// Calculates and returns the cluster number where x is stored in the tree  
func (V VEB) High(x int) int {
	return int(math.Floor(float64(x) / float64(LowerSqrt(V.u))))
}

// Calculates and returns the position in which x appears in its cluster
func (V VEB) Low(x int) int {
	return x % LowerSqrt(V.u)
}

// Calculates and returns the index for x and y
func (V VEB) Index(x, y int) int {
	return x*LowerSqrt(V.u) + y
}

// Checks if x is stored in the tree. Returns true if so, false if not
func (V VEB) IsMember(x int) bool {
	if x == V.min || x == V.max {
		return true
	} else if V.u == 2 {
		return false
	} else {
		return V.cluster[V.High(x)].IsMember(V.Low(x))
	}
}

// Inserts x into the tree
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

// Removes x from the tree
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

// Finds the successor element of x in the tree
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

// Finds the predecessor element of x in the tree
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

// Counts and returns all objects (structs) of the tree
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

// Prints the tree to std out
func (V VEB) Print() {
	//function just acts as wrapper, since default parameters are not possible in go
	V.PrintFunc(0, 0)
}

// Prints the tree to std out, where level is used to keep track of idention and clusterNo to 
// track the number of the cluster
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

// Clears the tree by setting min and max to -1 for every node of the tree
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

// Fills tree with all keys (inserts all keys from 0 to u)
func (V *VEB) Fill() {
	for i := 0; i < V.u; i++ {
		V.Insert(i)
	}
}

// Get array of tree members
func (V VEB) Members() []int{
	members := []int{}
	for i := 0; i < V.u; i++ {
		if V.IsMember(i) {
			members = append(members, i)
		}
	}
	return members
}

// Calculate lower square root (helper function)
func LowerSqrt(u int) int {
	return int(math.Pow(2.0, math.Floor(math.Log2(float64(u))/2)))
}

// Calculate higher square root (helper function)
func HigherSqrt(u int) int {
	return int(math.Pow(2.0, math.Ceil(math.Log2(float64(u))/2)))
}
