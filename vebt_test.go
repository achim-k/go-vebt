package vebt

import(
	"testing"
	"math"
	"math/rand"
	"time"
	"sort"
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

func TestCreateVEBTree(t *testing.T) {
	const in, out = 16, 21

	// TODO: Count number of structs and compare
	if CreateVEBTree(in) == nil {
		t.Errorf("CreateVEBTree(%v) created %v VEB structures, want", in, out)
	}
}

func TestMax(t *testing.T) {
	maxUpower := 10
	// Create trees with different universe sizes
	for i := 1; i < maxUpower; i++ {
		u := int(math.Pow(2, float64(i)))
		V := CreateVEBTree(u)
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
		V := CreateVEBTree(u)
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
		V := CreateVEBTree(u)
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

		// delete inserted keys again
		for j := 0; j < len(keys); j++ {
			V.Delete(keys[j]) // Insert key into tree
		}

		// check membership again
		for j := 0; j < u; j++ {
			expect := false
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
		V := CreateVEBTree(u)
		
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
				if keys[k] > i {
					nextBiggerKey = keys[k]
					break
				}
			}
			expect := nextBiggerKey
			if foundSuccessor := V.Successor(i); foundSuccessor != expect {
				t.Errorf("Successor(%v) = %v, want %v", i, foundSuccessor, expect)
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
		V := CreateVEBTree(u)
		
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
			for k := len(keys)-1; k >= 0; k-- {
				if keys[k] < i {
					nextSmallerKey = keys[k]
					break
				}
			}
			expect := nextSmallerKey
			if foundPred := V.Predecessor(i); foundPred != expect {
				t.Errorf("Predecessor(%v) = %v, want %v", i, foundPred, expect)
				break
			}
		}
	}
}

func arrayContains(ar []int, value int) bool {
	for i := 0; i < len(ar); i++ {
		if ar[i] == value {
			return true
		}
	}
	return false
}

func createRandomSortedKeys(max int) []int {
	rnd := rand.New(rand.NewSource(int64(time.Now().Nanosecond())*1))
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