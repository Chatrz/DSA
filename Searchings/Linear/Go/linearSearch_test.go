package linearsearch

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSearch(t *testing.T) {
	t0 := time.Now().UnixNano()
	arr := rand.Perm(1000)
	Search(arr[999], arr)
	t1 := time.Now().UnixNano()
	fmt.Printf("took %v ns", t1-t0)
}
