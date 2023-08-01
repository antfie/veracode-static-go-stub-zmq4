package zmq4_stub

import (
	"time"
)

type reactor_socket struct {
	e State
	f func(State) error
}

type reactor_channel struct {
	ch    <-chan interface{}
	f     func(interface{}) error
	limit int
}

type Reactor struct {
	sockets  map[*Socket]*reactor_socket
	channels map[uint64]*reactor_channel
	p        *Poller
	idx      uint64
	remove   []uint64
	verbose  bool
}

/*
Create a reactor to mix the handling of sockets and channels (timers or other channels).

Example:

	reactor := zmq.NewReactor()
	reactor.AddSocket(socket1, zmq.POLLIN, socket1_handler)
	reactor.AddSocket(socket2, zmq.POLLIN, socket2_handler)
	reactor.AddChannelTime(time.Tick(time.Second), 1, ticker_handler)
	reactor.Run(time.Second)

Warning:

There are problems with the reactor showing up with Go 1.14 (and later)
such as data race occurrences and code lock-up. Using SetRetryAfterEINTR
seems an effective fix, but at the moment there is no guaranty.
*/
func NewReactor() *Reactor {
	return &Reactor{}
}

// Add socket handler to the reactor.
//
// You can have only one handler per socket. Adding a second one will remove the first.
//
// The handler receives the socket state as an argument: POLLIN, POLLOUT, or both.
func (r *Reactor) AddSocket(soc *Socket, events State, handler func(State) error) {
}

// Remove a socket handler from the reactor.
func (r *Reactor) RemoveSocket(soc *Socket) {
}

// Add channel handler to the reactor.
//
// Returns id of added handler, that can be used later to remove it.
//
// If limit is positive, at most this many items will be handled in each run through the main loop,
// otherwise it will process as many items as possible.
//
// The handler function receives the value received from the channel.
func (r *Reactor) AddChannel(ch <-chan interface{}, limit int, handler func(interface{}) error) (id uint64) {
	return 0
}

// This function wraps AddChannel, using a channel of type time.Time instead of type interface{}.
func (r *Reactor) AddChannelTime(ch <-chan time.Time, limit int, handler func(interface{}) error) (id uint64) {
	return 0
}

// Remove a channel from the reactor.
//
// Closed channels are removed automatically.
func (r *Reactor) RemoveChannel(id uint64) {
}

func (r *Reactor) SetVerbose(verbose bool) {
}

// Run the reactor.
//
// The interval determines the time-out on the polling of sockets.
// Interval must be positive if there are channels.
// If there are no channels, you can set interval to -1.
//
// The run alternates between polling/handling sockets (using the interval as timeout),
// and reading/handling channels. The reading of channels is without time-out: if there
// is no activity on any channel, the run continues to poll sockets immediately.
//
// The run exits when any handler returns an error, returning that same error.
func (r *Reactor) Run(interval time.Duration) (err error) {
	return errorTaint()
}
