package waitport

import (
	"fmt"
	"time"

	"github.com/miklosn/procspy"
)

// WaitPort waits for a port to be in listening state, up to timeout duration
func WaitPort(port uint16, timeout time.Duration) error {

	over := time.After(timeout)
	found := make(chan bool, 1)
	notfound := make(chan bool, 1)

	go func() {
		for {
			select {
			case <-over:
				notfound <- true
				break
			default:
				cs, err := procspy.Connections(false)
				if err != nil {
					continue
				}
				for c := cs.Next(); c != nil; c = cs.Next() {
					if c.LocalPort == port {
						found <- true
					}
				}
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()

	select {
	case <-found:
		return nil
	case <-notfound:
		return fmt.Errorf("Timed out waiting for port %d", port)
	}
}
