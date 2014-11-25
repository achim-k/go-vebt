package vebt

import (
	"math"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestLowerSqrt(t *testing.T) {
	const in, out = 32, 4
	if x := LowerSqrt(in); x != out {
		t.Errorf("LowerSqrt(%v) = %v, want %v", in, x, out)
	}
}

func TestHigherSqrt(t *testing.T) {
	const in, out = 32, 8
	if x := HigherSqrt(in); x != out {
		t.Errorf("HigherSqrt(%v) = %v, want %v", in, x, out)
	}
}

func TestCreateTree(t *testing.T) {
	maxUpower := 16
	// Create trees with different universe sizes
	for i := 1; i < maxUpower; i++ {
		u := int(math.Pow(2, float64(i)))
		V := CreateTree(u)

		if V == nil {
			t.Errorf("CreateTree(%v) returned %v", u, nil)
		}

		// TODO: compare count with calculated number
		/*
			if count := V.Count(); count != 1 {
				t.Errorf("CreateTree(%v) created %v VEB structures, want %v", u, count, count)
			}
		*/
	}
}

func TestMax(t *testing.T) {
	maxUpower := 10
	// Create trees with different universe sizes
	for i := 1; i < maxUpower; i++ {
		u := int(math.Pow(2, float64(i)))
		V := CreateTree(u)
		// Insert random, sorted keys and check if max corresponds to that key
		keys := createRandomSortedKeys(u)
		for j := 0; j < len(keys); j++ {
			V.Insert(keys[j]) // Insert key into tree
			if out := V.Max(); out != keys[j] {
				t.Errorf("Max() = %v, want %v", out, keys[j])
				break
			}
		}
	}
}

func TestMin(t *testing.T) {
	maxUpower := 10
	// Create trees with different universe sizes
	for i := 1; i < maxUpower; i++ {
		u := int(math.Pow(2, float64(i)))
		V := CreateTree(u)
		// Insert random, sorted keys and check if min corresponds to that key
		keys := createRandomSortedKeys(u)
		for j := len(keys) - 1; j >= 0; j-- {
			V.Insert(keys[j]) // Insert key into tree
			if out := V.Min(); out != keys[j] {
				t.Errorf("Min() = %v, want %v", out, keys[j])
				break
			}
		}
	}
}

// Tests insert, membership + delete
func TestIsMember(t *testing.T) {
	maxUpower := 10
	// Create trees with different universe sizes
	for i := 1; i < maxUpower; i++ {
		u := int(math.Pow(2, float64(i)))
		V := CreateTree(u)
		// check empty tree for membership first
		for j := 0; j < u; j++ {
			if out := V.IsMember(j); out != false {
				t.Errorf("IsMember(%v) = %v, want %v", j, out, false)
				break
			}
		}

		// insert random keys
		keys := createRandomSortedKeys(u)
		for j := 0; j < len(keys); j++ {
			V.Insert(keys[j]) // Insert key into tree
		}

		// check membership again
		for j := 0; j < u; j++ {
			expect := false
			if arrayContains(keys, j) {
				expect = true // key was inserted, so expect IsMember to return true
			}

			if out := V.IsMember(j); out != expect {
				t.Errorf("IsMember(%v) = %v, want %v", j, out, expect)
				break
			}
		}

	}
}

func TestSuccessor(t *testing.T) {
	maxUpower := 10
	// Create trees with different universe sizes
	for i := 1; i < maxUpower; i++ {
		u := int(math.Pow(2, float64(i)))
		V := CreateTree(u)

		// check emptry tree for successors
		for j := 0; j < u; j++ {
			if out := V.Successor(j); out != -1 {
				t.Errorf("Successor(%v) = %v, want %v", j, out, -1)
				break
			}
		}

		// insert random keys
		keys := createRandomSortedKeys(u)
		for j := 0; j < len(keys); j++ {
			V.Insert(keys[j]) // Insert key into tree
		}

		for j := 0; j < u; j++ {
			// check if successor matches next bigger inserted key
			nextBiggerKey := -1
			//find next bigger key
			for k := 0; k < len(keys); k++ {
				if keys[k] > j {
					nextBiggerKey = keys[k]
					break
				}
			}
			expect := nextBiggerKey
			if foundSuccessor := V.Successor(j); foundSuccessor != expect {
				t.Errorf("Successor(%v) = %v, want %v", j, foundSuccessor, expect)
				break
			}
		}
	}
}

func TestPredecessor(t *testing.T) {
	maxUpower := 10
	// Create trees with different universe sizes
	for i := 1; i < maxUpower; i++ {
		u := int(math.Pow(2, float64(i)))
		V := CreateTree(u)

		// check emptry tree for predecessors
		for j := 0; j < u; j++ {
			if out := V.Predecessor(j); out != -1 {
				t.Errorf("Predecessor(%v) = %v, want %v", j, out, -1)
				break
			}
		}

		// insert random keys
		keys := createRandomSortedKeys(u)
		for j := 0; j < len(keys); j++ {
			V.Insert(keys[j]) // Insert key into tree
		}

		for j := 0; j < u; j++ {
			//find next smaller key
			nextSmallerKey := -1
			for k := len(keys) - 1; k >= 0; k-- {
				if keys[k] < j {
					nextSmallerKey = keys[k]
					break
				}
			}
			expect := nextSmallerKey
			if foundPred := V.Predecessor(j); foundPred != expect {
				t.Errorf("Predecessor(%v) = %v, want %v \t %v", j, foundPred, expect, keys)
				break
			}
		}
	}
}

func TestDelete(t *testing.T) {
	maxUpower := 10
	// Create trees with different universe sizes
	for i := 1; i < maxUpower; i++ {
		u := int(math.Pow(2, float64(i)))
		V := CreateTree(u)

		/* fill tree */
		V.Fill()

		memberCount := 0
		for m := 0; m < V.u; m++ {
			if V.IsMember(m) {
				memberCount++
			}
		}

		// delete random keys
		keys := createRandomSortedKeys(u)
		for j := 0; j < len(keys); j++ {
			V.Delete(keys[j])
		}

		memberCount = 0
		for m := 0; m < V.u; m++ {
			if V.IsMember(m) {
				memberCount++
			}
		}

		/* Check if all elements excepted deleted keys are members */
		for k := 0; k < V.u; k++ {
			expect := true
			if arrayContains(keys, k) {
				expect = false
			}
			if member := V.IsMember(k); member != expect {
				t.Errorf("m=%v: \tIsMember(%v) = %v, want %v after keys %v were deleted before", i, k, member, expect, keys)
			}
		}

		/* remove everything */
		for k := 0; k < V.u; k++ {
			V.Delete(k)
		}

		for k := 0; k < V.u; k++ {
			expect := false
			if member := V.IsMember(k); member != expect {
				t.Errorf("m=%v: \tIsMember(%v) = %v, want %v after all keys were deleted before", i, k, member, expect)
			}
		}
	}
}

func TestClear(t *testing.T) {
	maxUpower := 10
	// Create trees with different universe sizes
	for i := 1; i < maxUpower; i++ {
		u := int(math.Pow(2, float64(i)))
		V := CreateTree(u)

		// insert random keys
		keys := createRandomSortedKeys(u)
		for j := 0; j < len(keys); j++ {
			V.Insert(keys[j]) // Insert key into tree
		}

		// clear tree
		V.Clear()

		// check if a member exists
		for j := 0; j < V.u; j++ {
			if check := V.IsMember(j); check == true {
				t.Errorf("IsMember(%v) = %v, want %v because tree was cleared before %v", j, check, false, keys)
				break
			}
		}

		/* Check if all data structs min and max is -1 */
		if V.IsTreeEmpty() == false {
			t.Errorf("IsTreeEmpty() = %v, want %v after all elements were deleted", true, false)
			break
		}
	}
}

/*
func TestPrint(t *testing.T) {
	keys := createRandomSortedKeys(16)

	fmt.Printf("Printing veb tree (u=16) with random keys %v inserted:\n", keys)

	V := CreateTree(128)
	for i := 0; i < len(keys); i++ {
		V.Insert(keys[i])
	}
	V.Print()
}
*/

func arrayContains(ar []int, value int) bool {
	for i := 0; i < len(ar); i++ {
		if ar[i] == value {
			return true
		}
	}
	return false
}

func createRandomSortedKeys(max int) []int {
	rnd := rand.New(rand.NewSource(int64(time.Now().Nanosecond()) * 1))
	keys := []int{}
	keyNo := rnd.Intn(max)
	//create random keys
	for i := 0; i < keyNo; i++ {
		rndKey := rnd.Intn(max - 1)
		if arrayContains(keys, rndKey) == false {
			keys = append(keys, rndKey)
		}
	}
	sort.Ints(keys)

	return keys
}

// Traverse tree down and check if every min/max is -1
func (V *VEB) IsTreeEmpty() bool {
	if V.u == 2 {
		if V.min != -1 || V.max != -1 {
			return false
		}
		return true
	}

	check := true
	for i := 0; i < len(V.cluster); i++ {
		check = check && V.cluster[i].IsTreeEmpty()
	}
	check = check && V.summary.IsTreeEmpty()
	return check
}
