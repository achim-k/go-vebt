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
	if CreateVEBTree(in, &counter); counter != out {
		t.Errorf("CreateVEBTree(%v) created %v VEB structures, want %v", in, counter, out)
	}
}

func TestRun(t *testing.T) {
	runs := 100

	for i := 0; i < runs; i++ {

		u := 128
		rd := rand.New(rand.NewSource(int64(time.Now().Nanosecond())*2))
		keys := []int{}
		keyNo := rd.Intn(u)
		//create random keys
		for i := 0; i < keyNo; i++ {
			keys = append(keys, rd.Intn(u - 1))
		}

		dummy := 0
		V := CreateVEBTree(u, &dummy)
		insertMembershipTest(t, V, keyNo, keys)
	}
}

func createTree(t *testing.T) (VEB) {
	//TODO: Would be nice to have u parameterized and tree created automatically

	const u = 16

	// VEB(2) objects (layer3)
	l2_0_0 := VEB{u: 2, min: -1, max: -1}
	l2_0_1 := VEB{u: 2, min: -1, max: -1}
	l2_1_0 := VEB{u: 2, min: -1, max: -1}
	l2_1_1 := VEB{u: 2, min: -1, max: -1}
	l2_2_0 := VEB{u: 2, min: -1, max: -1}
	l2_2_1 := VEB{u: 2, min: -1, max: -1}
	l2_3_0 := VEB{u: 2, min: -1, max: -1}
	l2_3_1 := VEB{u: 2, min: -1, max: -1}

	l2_0_s3 := VEB{u: 2, min: -1, max: -1}
	l2_1_s3 := VEB{u: 2, min: -1, max: -1}
	l2_2_s3 := VEB{u: 2, min: -1, max: -1}
	l2_3_s3 := VEB{u: 2, min: -1, max: -1}

	s2_0_s3_0 := VEB{u: 2, min: -1, max: -1}
	s2_0_s3_1 := VEB{u: 2, min: -1, max: -1}
	s2_0_s3_2 := VEB{u: 2, min: -1, max: -1}

	
	// VEB(4) objects
	l2_0 := VEB{u: 4, min: -1, max: -1, summary: &l2_0_s3, cluster: []*VEB{&l2_0_0, &l2_0_1}}
	l2_1 := VEB{u: 4, min: -1, max: -1, summary: &l2_1_s3, cluster: []*VEB{&l2_1_0, &l2_1_1}}
	l2_2 := VEB{u: 4, min: -1, max: -1, summary: &l2_2_s3, cluster: []*VEB{&l2_2_0, &l2_2_1}}
	l2_3 := VEB{u: 4, min: -1, max: -1, summary: &l2_3_s3, cluster: []*VEB{&l2_3_0, &l2_3_1}}
	s2_0 := VEB{u: 4, min: -1, max: -1, summary: &s2_0_s3_0, cluster: []*VEB{&s2_0_s3_1, &s2_0_s3_2}}

	// VEB(16)
	l1_0 := VEB{u: 16, min: -1, max: -1, summary: &s2_0, cluster: []*VEB{&l2_0, &l2_1, &l2_2, &l2_3}}

	return l1_0
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
