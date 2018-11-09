package main

import (
	"fmt"
	"time"
)

func main() {
	ti := time.Date(2018, time.November, 3, 16, 16, 16, 0, time.Local)
	// 03 Nov 18 16:16:16 CET
	fmt.Printf("Time: %s\n", ti.Format(time.RFC822))
	// 03 Nov 18 16:16:16 CET +0100
	fmt.Printf("Time: %s\n", ti.Format(time.RFC822Z))
	// Saturday, 03-Nov-18 16:16:16 CET
	fmt.Printf("Time: %s\n", ti.Format(time.RFC850))
	// Sat, o3 Nov 2018 16:16:16 CET
	fmt.Printf("Time: %s\n", ti.Format(time.RFC1123))
	// Sat, 03 Nov 2018 16:16:16 +0100
	fmt.Printf("Time: %s\n", ti.Format(time.RFC1123Z))
	// 2018-11-03T16:16:16+01:00
	fmt.Printf("Time: %s\n", ti.Format(time.RFC3339))
	// 2018-11-03T16:16:16+01:00
	fmt.Printf("Time: %s\n", ti.Format(time.RFC3339Nano))
	// Sat Nov 3 16:16:16 CET 2018
	fmt.Printf("Time: %s\n", ti.Format(time.UnixDate))
}
