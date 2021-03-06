go-vebt
=======

Go implementation of Van Emde Boas tree data structure. The implementation was part of a project for "advanced data structures" class at Seoul National University. License: MIT

## Features
The implemented Van Emde Boas tree data structure supports the following core operations:

* Insert
* Delete
* Member
* Successor
* Predecessor
* Mininum
* Maximum

The following operations are also supported, which might be useful (e.g. for debugging):
* Print
* Count
* Clear
* Fill
* Members

## Usage

### Install
Usage:
```
go get github.com/achim-k/go-vebt
```

### Import
Add library to import statement:
```
import (	
	"github.com/achim-k/go-vebt"
	...
)
```

### Example usage

#### CreateTree() & Insert() & IsMember()
Code:
```
// Create tree
u := 16
V := CreateTree(u)
fmt.Printf("Created tree with %v structs\n", V.Count())

// Insert some elements
V.Insert(3)
V.Insert(6)
V.Insert(8)
V.Insert(13)

// Check which elements are members
members := []int{}
for i := 0; i < u; i++ {
	if V.IsMember(i) {
		members = append(members, i)
	}
}
// Print members to console
fmt.Printf("Members: %v\n", members)
```

Output:
```
Created tree with 21 structs
Members: [3 6 8 13]
```

#### Min() & Max()
Code:
``` 
fmt.Printf("Min: %v\nMax: %v\n", V.Min(), V.Max())
```

Output:
```
Min: 3
Max: 13
```

#### Predecessor() & Successor()
Code:
``` 
for i := 0; i < len(members); i++ {
	e := members[i]
	fmt.Printf("Predecessor(%v): %v\t Successor(%v): %v\n", 
				e, V.Predecessor(e), e, V.Successor(e))	
}
```

Output:
```
Predecessor(3): -1	 Successor(3): 6
Predecessor(6): 3	 Successor(6): 8
Predecessor(8): 6	 Successor(8): 13
Predecessor(13): 8	 Successor(13): -1
```

#### Print()
Prints tree to std out (useful for debug)

Code:
``` 
V.Print()
```

Output:
```
R: {u: 16, min: 3, max: 13, clusters: 4}
	S:    {u: 4, min: 1, max: 3, clusters: 2}
		C[0]: {u: 2, min: -1, max: -1, clusters: 0}
		C[1]: {u: 2, min: 0, max: 1, clusters: 0}
	C[0]: {u: 4, min: -1, max: -1, clusters: 2}
		S:    {u: 2, min: -1, max: -1, clusters: 0}
		C[0]: {u: 2, min: -1, max: -1, clusters: 0}
		C[1]: {u: 2, min: -1, max: -1, clusters: 0}
	C[1]: {u: 4, min: 2, max: 2, clusters: 2}
		S:    {u: 2, min: -1, max: -1, clusters: 0}
		C[0]: {u: 2, min: -1, max: -1, clusters: 0}
		C[1]: {u: 2, min: -1, max: -1, clusters: 0}
	C[2]: {u: 4, min: 0, max: 0, clusters: 2}
		S:    {u: 2, min: -1, max: -1, clusters: 0}
		C[0]: {u: 2, min: -1, max: -1, clusters: 0}
		C[1]: {u: 2, min: -1, max: -1, clusters: 0}
	C[3]: {u: 4, min: 1, max: 1, clusters: 2}
		S:    {u: 2, min: -1, max: -1, clusters: 0}
		C[0]: {u: 2, min: -1, max: -1, clusters: 0}
		C[1]: {u: 2, min: -1, max: -1, clusters: 0}
```
R: Root element, S: Summary, C[x]: cluster x of a node

#### Delete() & Clear() & Member()
Code:
``` 
V.Delete(3)
V.Delete(13)
fmt.Printf("Members: %v \n", V.Members())

V.Clear() // Deletes all keys in tree
fmt.Printf("Members: %v \n", V.Members())
```

Output:
``` 
Members: [6 8] 
Members: [] 
```



