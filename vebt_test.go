package vebt

import(
	"testing"
	"fmt"
	"math/rand"
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

func TestCreateVEBTree(t *testing.T) {
	counter := 0
	const in, out = 16, 21

	// TODO: Count number of structs and compare
	if CreateVEBTree(in); counter != out {
		//t.Errorf("CreateVEBTree(%v) created %v VEB structures, want %v", in, counter, out)
	}
}

func TestRun(t *testing.T) {
	runs := 100

	for i := 0; i < runs; i++ {

		u := 16
		rd := rand.New(rand.NewSource(int64(time.Now().Nanosecond())*2))
		keys := []int{}
		keyNo := rd.Intn(u)
		//create random keys
		for i := 0; i < keyNo; i++ {
			keys = append(keys, rd.Intn(u - 1))
		}

		V := CreateVEBTree(u)
		if V != nil {
			insertMembershipTest(t, V, keyNo, keys)	
		} else {
			t.Errorf("CreateVEBTree(%v) failure", u )
		}
		
	}
}

func insertMembershipTest(t *testing.T, V *VEB, keyNo int, keys []int) {
	//we expect an empty tree, test each key for emptiness first
	for i := 0; i < V.u; i++ {
		if V.IsMember(i) == true {
			t.Errorf("IsMember(%v) = %v, want %v", i, true, false)
		}
	}

	fmt.Printf("Random keys (%v): %v\n", keyNo, keys)

	// Insert keys
	for i := 0; i < len(keys); i++ {
		V.Insert(keys[i])
	}

	// Member Check
	for i := 0; i < V.u; i++ {
		keyExists := false
		for k := 0; k < len(keys); k++ {
			if keys[k] == i {
				keyExists = true	
				break
			}
		}
		// Check existing keys
		if keyExists == true {
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
