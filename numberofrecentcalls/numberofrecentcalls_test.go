package numberofrecentcalls

import (
	"fmt"
	"testing"
)

func TestPing(t *testing.T) {
	counter := Constructor()
	fmt.Println(counter.Ping(642))
	fmt.Println(counter.Ping(1849))
	fmt.Println(counter.Ping(4921))
	fmt.Println(counter.Ping(5936))
	fmt.Println(counter.Ping(5957))

	//
	// [1],[100],[3001],[3002]
	counter = Constructor()
	fmt.Println(counter.Ping(1))
	fmt.Println(counter.Ping(100))
	fmt.Println(counter.Ping(3001))
	fmt.Println(counter.Ping(3002))
	//
}

/*
["RecentCounter","ping","ping","ping","ping","ping"]
[[],[642],[1849],[4921],[5936],[5957]]
expect:[null,1,2,1,2,3]


*/
