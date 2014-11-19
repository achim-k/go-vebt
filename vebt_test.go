package vebt

import(
	"testing"
	"fmt"
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
	counter := 0
	const in, out = 16, 21

	// TODO: Count number of structs and compare
	if CreateVEBTree(in); counter != out {
		//t.Errorf("CreateVEBTree(%v) created %v VEB structures, want %v", in, counter, out)
	}
}

func TestRun(t *testing.T) {
	runs := 20
	fmt.Printf("Performing %v tests...\n", runs)

	for i := 0; i < runs; i++ {
		fmt.Printf("#%v...\t", i+1)
		u := 16
		rd := rand.New(rand.NewSource(int64(time.Now().Nanosecond())*2))
		keys := []int{}
		keyNo := rd.Intn(u)
		//create random keys
		for i := 0; i < keyNo; i++ {
			rndKey := rd.Intn(u - 1)
			if arrayContains(keys, rndKey) == false {
				keys = append(keys, rndKey)	
			}
		}
		//sort keys
		sort.Ints(keys)

		V := CreateVEBTree(u)
		if V != nil {
			insertMembershipTest(t, V, keys)
			successorTest(t, V, keys)
			predecessorTest(t, V, keys)
		} else {
			t.Errorf("CreateVEBTree(%v) failure", u )
		}
		fmt.Printf("done\n")
		
	}
}

func insertMembershipTest(t *testing.T, V *VEB, keys []int) {
	//we expect an empty tree, test each key for emptiness first
	for i := 0; i < V.u; i++ {
		if V.IsMember(i) == true {
			t.Errorf("IsMember(%v) = %v, want %v", i, true, false)
		}
	}

	//fmt.Printf("Random keys (%v): %v\n", len(keys), keys)

	// Insert keys
	for i := 0; i < len(keys); i++ {
		V.Insert(keys[i])
	}

	// Member Check
	for i := 0; i < V.u; i++ {
		// Check existing keys
		if arrayContains(keys, i) {
			if V.IsMember(i) == false {
				t.Errorf("IsMember(%v) = %v, want %v", i, false, true)
			}	
		} else { // CHeck nonexisting keys
			if V.IsMember(i) == true {
				t.Errorf("IsMember(%v) = %v, want %v", i, true, false)
			}	
		}
	}
}

func successorTest(t *testing.T, V *VEB, keys []int) {
	for i := 0; i < V.u; i++ {
		//find next bigger key
		nextBiggerKey := -1
		for k := 0; k < len(keys); k++ {
			if keys[k] > i {
				nextBiggerKey = keys[k]
				break
			}
		}

		if nextBiggerKey == -1 {
			//No Successor should be there, hence we expect -1 as return
			expect := -1
			if foundSuccessor := V.Successor(i); foundSuccessor != expect {
				t.Errorf("Successor(%v) = %v, want %v", i, foundSuccessor, expect)
			}
		} else {
			//There is a key which is higher then the current tested one, we expect the successor to be that higher key
			expect := nextBiggerKey
			if foundSuccessor := V.Successor(i); foundSuccessor != expect {
				t.Errorf("Successor(%v) = %v, want %v", i, foundSuccessor, expect)
			}
		}
	}
}

func predecessorTest(t *testing.T, V *VEB, keys []int) {
	for i := 0; i < V.u; i++ {
		//find next smaller key
		nextSmallerKey := -1
		for k := len(keys)-1; k >= 0; k-- {
			if keys[k] < i {
				nextSmallerKey = keys[k]
				break
			}
		}

		if nextSmallerKey == -1 {
			//No Predecessor should be there, hence we expect -1 as return
			expect := -1
			if foundPredecessor := V.Predecessor(i); foundPredecessor != expect {
				t.Errorf("Predecessor(%v) = %v, want %v", i, foundPredecessor, expect)
			}
		} else {
			//There is a key which is smaller then the current tested one, we expect the Predecessor to be that smaller key
			expect := nextSmallerKey
			if foundPredecessor := V.Predecessor(i); foundPredecessor != expect {
				t.Errorf("Predecessor(%v) = %v, want %v", i, foundPredecessor, expect)
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
