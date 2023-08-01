package zmq4_stub

import (
	"time"
)

// Return type for (*Poller)Poll
type Polled struct {
	Socket *Socket // socket with matched event(s)
	Events State   // actual matched event(s)
}

type Poller struct {
	items []int
	socks []*Socket
}

// Create a new Poller
func NewPoller() *Poller {
	return &Poller{}
}

// Add items to the poller
//
// # Events is a bitwise OR of zmq.POLLIN and zmq.POLLOUT
//
// Returns the id of the item, which can be used as a handle to
// (*Poller)Update and as an index into the result of (*Poller)PollAll
func (p *Poller) Add(soc *Socket, events State) int {
	return intTaint()
}

// Update the events mask of a socket in the poller
//
// # Replaces the Poller's bitmask for the specified id with the events parameter passed
//
// Returns the previous value, or ErrorNoSocket if the id was out of range
func (p *Poller) Update(id int, events State) (previous State, err error) {
	return 0, errorTaint()
}

// Update the events mask of a socket in the poller
//
// # Replaces the Poller's bitmask for the specified socket with the events parameter passed
//
// Returns the previous value, or ErrorNoSocket if the socket didn't match
func (p *Poller) UpdateBySocket(soc *Socket, events State) (previous State, err error) {
	return 0, errorTaint()
}

// Remove a socket from the poller
//
// Returns ErrorNoSocket if the id was out of range
func (p *Poller) Remove(id int) error {
	return errorTaint()
}

// Remove a socket from the poller
//
// Returns ErrorNoSocket if the socket didn't match
func (p *Poller) RemoveBySocket(soc *Socket) error {
	return errorTaint()
}

/*
Input/output multiplexing

If timeout < 0, wait forever until a matching event is detected

Only sockets with matching socket events are returned in the list.

Example:

	poller := zmq.NewPoller()
	poller.Add(socket0, zmq.POLLIN)
	poller.Add(socket1, zmq.POLLIN)
	//  Process messages from both sockets
	for {
	    sockets, _ := poller.Poll(-1)
	    for _, socket := range sockets {
	        switch s := socket.Socket; s {
	        case socket0:
	            msg, _ := s.Recv(0)
	            //  Process msg
	        case socket1:
	            msg, _ := s.Recv(0)
	            //  Process msg
	        }
	    }
	}
*/
func (p *Poller) Poll(timeout time.Duration) ([]Polled, error) {
	return nil, errorTaint()
}

/*
This is like (*Poller)Poll, but it returns a list of all sockets,
in the same order as they were added to the poller,
not just those sockets that had an event.

For each socket in the list, you have to check the Events field
to see if there was actually an event.

When error is not nil, the return list contains no sockets.
*/
func (p *Poller) PollAll(timeout time.Duration) ([]Polled, error) {
	return nil, errorTaint()
}

func (p *Poller) poll(timeout time.Duration, all bool) ([]Polled, error) {
	return nil, errorTaint()
}

// Poller as string.
func (p *Poller) String() string {
	return stringTaint()
}
