package libnova

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/big"
)

// "Charlie" -> abcdef12345 -> A bucket -> Charlie
// [ Table ]
//    |    a          / - Record1
//    | - [ Bucket ] {  - Record2
//    |               \ - Record3

//    |    b          / - Record4
//    | - [ Bucket ] {  - Record5
//    |               \ - Record6

//    |    c          / - Record7
//    | - [ Bucket ] {  - Record8
//    |               \ - Record9
//    |

const DefaultBucketCount = 1024

// Table is used to effectively store
// key/value pairs
type Table struct {
	BucketCount int
	Name        string
	Buckets     []*Bucket
	Length      int // 0 indexed
}

// Bucket is used to improve efficency of
// lookups
type Bucket struct {
	ID      int
	Records *BinarySearchTree
	Length  int
}

// Record is a key/value pair where the value
// can be any type and is used to represent
// a single record in the table
type Record struct {
	Key   string
	Value interface{}
}

// NewHashTable is used to create a new data structure
// for key/value storage and lookup
func NewHashTable(name string, bucketCount int) *Table {
	if bucketCount < 1 {
		bucketCount = DefaultBucketCount
	}
	table := &Table{
		Name:        name,
		BucketCount: bucketCount,
	}
	// Initialize the Table with BucketCount buckets
	for i := bucketCount; i > 0; i-- {
		table.Buckets = append(table.Buckets, &Bucket{
			ID:      i - 1, // Pass in 512->511,1->0
			Records: NewBinarySearchTree(fmt.Sprintf("%d", i-1)),
		})
	}
	return table
}

// Get is used to search for a key and return it's
// corresponding Record
func (t *Table) Get(key string) *Record {
	n := t.Hash(key)
	bucket := t.Buckets[n]
	node := bucket.Records.Search(key)
	if node == nil {
		return nil
	}
	return node.Value.(*Record)
}

// Set is used to add a key/value pair to the hash table
//
// Currently uses a (slow) linear search to search for keys
//    Linear Search
//      Worst Case O(n) <-- Slow as fuck on large sets
//      Best  Case O(1)
//func (t *Table) Set(record *Record)
func (t *Table) Set(key string, value interface{}) {
	// Calculate n
	n := t.Hash(key)
	// Create a new record
	record := &Record{
		Key:   key,
		Value: value,
	}
	t.Length++
	t.Buckets[n].Length++
	t.Buckets[n].Records.Insert(&NodeBinary{
		Node: Node{
			Key:   key,
			Value: record,
		},
	})
}

// Hash is used to calculate an int value for a given string where the integer
// result is a factor of the table's dynamic bucket count
// TODO currently the Hash() method does not balance and there is a high standard deviation
// TODO md5 has risk of colliding - replace with sha256/512
func (t *Table) Hash(key string) int {
	// Calculate MD5 sum and form an int result
	h := md5.New()
	h.Write([]byte(key))
	md5 := hex.EncodeToString(h.Sum(nil))
	// Now that we have a idempotent hash, we can calculate an integer value
	bigInt := big.NewInt(0)
	bigInt.SetString(md5, 16) // Probably get away with base 2 or base 16
	i64 := bigInt.Int64()
	n := int(i64)
	// n is the int result of our arithmetic
	if n < 0 {
		n = n * -1 // Take the absolute value of n
	}
	// modulus for our bucket count
	//
	// We use modulus here so that our int value (n)
	// is always divisible by the bucketCount such that
	// bucketCount + 1 = bucketCount[0]
	//
	// n = 7
	// bucketCount = 3
	//     [0] 1,4,7 <--- n=7 gets bucket 0
	//     [1] 2,5
	//     [2] 3,6
	//
	// This allows for upwards maximum of values of n
	// while retaining a finite amount of buckets that
	// algorithmically can be calculated.
	x := n % (t.BucketCount) // < -- We need to consider probability
	return x                 // Result should alwayws be a factor of t.BucketCount
}

// String is used to stringify a Record
func (r *Record) String() string {
	if r == nil || r.Value == nil {
		return ""
	}
	return fmt.Sprintf("%v", r.Value)
}
