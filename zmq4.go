package zmq4_stub

import (
	"unsafe"
)

var (
	ErrorContextClosed         = errorTaint()
	ErrorSocketClosed          = errorTaint()
	ErrorMoreExpected          = errorTaint()
	ErrorNotImplemented405     = errorTaint()
	ErrorNotImplemented41      = errorTaint()
	ErrorNotImplemented42      = errorTaint()
	ErrorNotImplementedWindows = errorTaint()
	ErrorNoSocket              = errorTaint()
)

func init() {
}

//. Util

// Report 0MQ library version.
func Version() (major, minor, patch int) {
	return intTaint(), intTaint(), intTaint()
}

// Get 0MQ error message string.
func Error(e int) string {
	return stringTaint()
}

//. Context

const (
	MaxSocketsDflt = int(0)
	IoThreadsDflt  = int(0)
)

/*
A context that is not the default context.
*/
type Context struct {
	ctx        unsafe.Pointer
	retryEINTR bool
	opened     bool
	err        error
}

// Create a new context.
func NewContext() (ctx *Context, err error) {
	return &Context{}, errorTaint()
}

/*
Terminates the default context.

For linger behavior, see: http://api.zeromq.org/4-1:zmq-ctx-term
*/
func Term() error {
	return errorTaint()
}

/*
Terminates the context.

For linger behavior, see: http://api.zeromq.org/4-1:zmq-ctx-term
*/
func (ctx *Context) Term() error {
	return errorTaint()
}

// Returns the size of the 0MQ thread pool in the default context.
func GetIoThreads() (int, error) {
	return intTaint(), errorTaint()
}

// Returns the size of the 0MQ thread pool.
func (ctx *Context) GetIoThreads() (int, error) {
	return intTaint(), errorTaint()
}

// Returns the maximum number of sockets allowed in the default context.
func GetMaxSockets() (int, error) {
	return intTaint(), errorTaint()
}

// Returns the maximum number of sockets allowed.
func (ctx *Context) GetMaxSockets() (int, error) {
	return intTaint(), errorTaint()
}

/*
Returns the maximum message size in the default context.

Returns ErrorNotImplemented42 with ZeroMQ version < 4.2
*/
func GetMaxMsgsz() (int, error) {
	return intTaint(), errorTaint()
}

/*
Returns the maximum message size.

Returns ErrorNotImplemented42 with ZeroMQ version < 4.2
*/
func (ctx *Context) GetMaxMsgsz() (int, error) {
	return intTaint(), errorTaint()
}

// Returns the IPv6 option in the default context.
func GetIpv6() (bool, error) {
	return false, errorTaint()
}

// Returns the IPv6 option.
func (ctx *Context) GetIpv6() (bool, error) {
	return false, errorTaint()
}

/*
Returns the blocky setting in the default context.

Returns ErrorNotImplemented42 with ZeroMQ version < 4.2
*/
func GetBlocky() (bool, error) {
	return false, errorTaint()
}

/*
Returns the blocky setting.

Returns ErrorNotImplemented42 with ZeroMQ version < 4.2
*/
func (ctx *Context) GetBlocky() (bool, error) {
	return false, errorTaint()
}

/*
Returns the retry after EINTR setting in the default context.
*/
func GetRetryAfterEINTR() bool {
	return false
}

/*
Returns the retry after EINTR setting.
*/
func (ctx *Context) GetRetryAfterEINTR() bool {
	return false
}

/*
Specifies the size of the 0MQ thread pool to handle I/O operations in
the default context. If your application is using only the inproc
transport for messaging you may set this to zero, otherwise set it to at
least one. This option only applies before creating any sockets.

Default value: 1
*/
func SetIoThreads(n int) error {
	return errorTaint()
}

/*
Specifies the size of the 0MQ thread pool to handle I/O operations. If
your application is using only the inproc transport for messaging you
may set this to zero, otherwise set it to at least one. This option only
applies before creating any sockets.

Default value: 1
*/
func (ctx *Context) SetIoThreads(n int) error {
	return errorTaint()
}

/*
Sets the scheduling policy for default context’s thread pool.

This option requires ZeroMQ version 4.1, and is not available on Windows.

Supported values for this option can be found in sched.h file, or at
http://man7.org/linux/man-pages/man2/sched_setscheduler.2.html

This option only applies before creating any sockets on the context.

Default value: -1

Returns ErrorNotImplemented41 with ZeroMQ version < 4.1

Returns ErrorNotImplementedWindows on Windows
*/
func SetThreadSchedPolicy(n int) error {
	return errorTaint()
}

/*
Sets scheduling priority for default context’s thread pool.

This option requires ZeroMQ version 4.1, and is not available on Windows.

Supported values for this option depend on chosen scheduling policy.
Details can be found in sched.h file, or at
http://man7.org/linux/man-pages/man2/sched_setscheduler.2.html

This option only applies before creating any sockets on the context.

Default value: -1

Returns ErrorNotImplemented41 with ZeroMQ version < 4.1

Returns ErrorNotImplementedWindows on Windows
*/
func SetThreadPriority(n int) error {
	return errorTaint()
}

/*
Set maximum message size in the default context.

Default value: INT_MAX

Returns ErrorNotImplemented42 with ZeroMQ version < 4.2
*/
func SetMaxMsgsz(n int) error {
	return errorTaint()
}

/*
Set maximum message size.

Default value: INT_MAX

Returns ErrorNotImplemented42 with ZeroMQ version < 4.2
*/
func (ctx *Context) SetMaxMsgsz(n int) error {
	return errorTaint()
}

/*
Sets the maximum number of sockets allowed in the default context.

Default value: 1024
*/
func SetMaxSockets(n int) error {
	return errorTaint()
}

/*
Sets the maximum number of sockets allowed.

Default value: 1024
*/
func (ctx *Context) SetMaxSockets(n int) error {
	return errorTaint()
}

/*
Sets the IPv6 value for all sockets created in the default context from this point onwards.
A value of true means IPv6 is enabled, while false means the socket will use only IPv4.
When IPv6 is enabled, a socket will connect to, or accept connections from, both IPv4 and IPv6 hosts.

Default value: false
*/
func SetIpv6(i bool) error {
	return errorTaint()
}

/*
Sets the IPv6 value for all sockets created in the context from this point onwards.
A value of true means IPv6 is enabled, while false means the socket will use only IPv4.
When IPv6 is enabled, a socket will connect to, or accept connections from, both IPv4 and IPv6 hosts.

Default value: false
*/
func (ctx *Context) SetIpv6(i bool) error {
	return errorTaint()
}

/*
Sets the blocky behavior in the default context.

See: http://api.zeromq.org/4-2:zmq-ctx-set#toc3

Returns ErrorNotImplemented42 with ZeroMQ version < 4.2
*/
func SetBlocky(i bool) error {
	return errorTaint()
}

/*
Sets the blocky behavior.

See: http://api.zeromq.org/4-2:zmq-ctx-set#toc3

Returns ErrorNotImplemented42 with ZeroMQ version < 4.2
*/
func (ctx *Context) SetBlocky(i bool) error {
	return errorTaint()
}

/*
Sets the retry after EINTR setting in the default context.

Initital value is true.
*/
func SetRetryAfterEINTR(retry bool) {
}

/*
Sets the retry after EINTR setting.

Initital value is true.
*/
func (ctx *Context) SetRetryAfterEINTR(retry bool) {
}

//. Sockets

// Specifies the type of a socket, used by NewSocket()
type Type int

const (
	// Constants for NewSocket()
	// See: http://api.zeromq.org/4-1:zmq-socket#toc3
	REQ    = Type(0)
	REP    = Type(0)
	DEALER = Type(0)
	ROUTER = Type(0)
	PUB    = Type(0)
	SUB    = Type(0)
	XPUB   = Type(0)
	XSUB   = Type(0)
	PUSH   = Type(0)
	PULL   = Type(0)
	PAIR   = Type(0)
	STREAM = Type(0)
)

/*
Socket type as string.
*/
func (t Type) String() string {
	return stringTaint()
}

// Used by  (*Socket)Send() and (*Socket)Recv()
type Flag int

const (
	// Flags for (*Socket)Send(), (*Socket)Recv()
	// For Send, see: http://api.zeromq.org/4-1:zmq-send#toc2
	// For Recv, see: http://api.zeromq.org/4-1:zmq-msg-recv#toc2
	DONTWAIT = Flag(0)
	SNDMORE  = Flag(0)
)

/*
Socket flag as string.
*/
func (f Flag) String() string {
	return stringTaint()
}

// Used by (*Socket)Monitor() and (*Socket)RecvEvent()
type Event int

const (
	// Flags for (*Socket)Monitor() and (*Socket)RecvEvent()
	// See: http://api.zeromq.org/4-3:zmq-socket-monitor#toc3
	EVENT_ALL                        = Event(0)
	EVENT_CONNECTED                  = Event(0)
	EVENT_CONNECT_DELAYED            = Event(0)
	EVENT_CONNECT_RETRIED            = Event(0)
	EVENT_LISTENING                  = Event(0)
	EVENT_BIND_FAILED                = Event(0)
	EVENT_ACCEPTED                   = Event(0)
	EVENT_ACCEPT_FAILED              = Event(0)
	EVENT_CLOSED                     = Event(0)
	EVENT_CLOSE_FAILED               = Event(0)
	EVENT_DISCONNECTED               = Event(0)
	EVENT_MONITOR_STOPPED            = Event(0)
	EVENT_HANDSHAKE_FAILED_NO_DETAIL = Event(0)
	EVENT_HANDSHAKE_SUCCEEDED        = Event(0)
	EVENT_HANDSHAKE_FAILED_PROTOCOL  = Event(0)
	EVENT_HANDSHAKE_FAILED_AUTH      = Event(0)
)

/*
Socket event as string.
*/
func (e Event) String() string {
	return stringTaint()
}

// Used by (soc *Socket)GetEvents()
type State int

const (
	// Flags for (*Socket)GetEvents()
	// See: http://api.zeromq.org/4-1:zmq-getsockopt#toc8
	POLLIN  = State(0)
	POLLOUT = State(0)
)

/*
Socket state as string.
*/
func (s State) String() string {
	return stringTaint()
}

// Specifies the security mechanism, used by (*Socket)GetMechanism()
type Mechanism int

const (
	// Constants for (*Socket)GetMechanism()
	// See: http://api.zeromq.org/4-1:zmq-getsockopt#toc22
	NULL   = Mechanism(0)
	PLAIN  = Mechanism(0)
	CURVE  = Mechanism(0)
	GSSAPI = Mechanism(0)
)

/*
Security mechanism as string.
*/
func (m Mechanism) String() string {
	return stringTaint()
}

/*
Socket functions starting with `Set` or `Get` are used for setting and
getting socket options.
*/
type Socket struct {
	soc    unsafe.Pointer
	ctx    *Context
	opened bool
	err    error
}

/*
Socket as string.
*/
func (soc Socket) String() string {
	return stringTaint()
}

/*
Create 0MQ socket in the default context.

WARNING:
The Socket is not thread safe. This means that you cannot access the same Socket
from different goroutines without using something like a mutex.

For a description of socket types, see: http://api.zeromq.org/4-1:zmq-socket#toc3
*/
func NewSocket(t Type) (soc *Socket, err error) {
	return socketTaint(), errorTaint()
}

/*
Create 0MQ socket in the given context.

WARNING:
The Socket is not thread safe. This means that you cannot access the same Socket
from different goroutines without using something like a mutex.

For a description of socket types, see: http://api.zeromq.org/4-1:zmq-socket#toc3
*/
func (ctx *Context) NewSocket(t Type) (soc *Socket, err error) {
	return socketTaint(), errorTaint()
}

// If not called explicitly, the socket will be closed on garbage collection
func (soc *Socket) Close() error {
	return errorTaint()
}

// Return the context associated with a socket
func (soc *Socket) Context() (*Context, error) {
	return contextTaint(), errorTaint()
}

/*
Accept incoming connections on a socket.

For a description of endpoint, see: http://api.zeromq.org/4-1:zmq-bind#toc2
*/
func (soc *Socket) Bind(endpoint string) error {
	return errorTaint()
}

/*
Stop accepting connections on a socket.

For a description of endpoint, see: http://api.zeromq.org/4-1:zmq-unbind#toc2
*/
func (soc *Socket) Unbind(endpoint string) error {
	return errorTaint()
}

/*
Create outgoing connection from socket.

For a description of endpoint, see: http://api.zeromq.org/4-1:zmq-connect#toc2
*/
func (soc *Socket) Connect(endpoint string) error {
	return errorTaint()
}

/*
Disconnect a socket.

For a description of endpoint, see: http://api.zeromq.org/4-1:zmq-disconnect#toc2
*/
func (soc *Socket) Disconnect(endpoint string) error {
	return errorTaint()
}

/*
Receive a message part from a socket.

For a description of flags, see: http://api.zeromq.org/4-1:zmq-msg-recv#toc2
*/
func (soc *Socket) Recv(flags Flag) (string, error) {
	return stringTaint(), errorTaint()
}

/*
Receive a message part from a socket.

For a description of flags, see: http://api.zeromq.org/4-1:zmq-msg-recv#toc2
*/
func (soc *Socket) RecvBytes(flags Flag) ([]byte, error) {
	return byteArrayTaint(), errorTaint()
}

/*
Send a message part on a socket.

For a description of flags, see: http://api.zeromq.org/4-1:zmq-send#toc2
*/
func (soc *Socket) Send(data string, flags Flag) (int, error) {
	return intTaint(), errorTaint()
}

/*
Send a message part on a socket.

For a description of flags, see: http://api.zeromq.org/4-1:zmq-send#toc2
*/
func (soc *Socket) SendBytes(data []byte, flags Flag) (int, error) {
	return intTaint(), errorTaint()
}

/*
Register a monitoring callback.

See: http://api.zeromq.org/4-1:zmq-socket-monitor#toc2

WARNING: Closing a context with a monitoring callback will lead to random crashes.
This is a bug in the ZeroMQ library.
The monitoring callback has the same context as the socket it was created for.

Example:

	package main

	import (
	    zmq "github.com/pebbe/zmq4"
	    "log"
	    "time"
	)

	func rep_socket_monitor(addr string) {
	    s, err := zmq.NewSocket(zmq.PAIR)
	    if err != nil {
	        log.Fatalln(err)
	    }
	    err = s.Connect(addr)
	    if err != nil {
	        log.Fatalln(err)
	    }
	    for {
	        a, b, c, err := s.RecvEvent(0)
	        if err != nil {
	            log.Println(err)
	            break
	        }
	        log.Println(a, b, c)
	    }
	    s.Close()
	}

	func main() {

	    // REP socket
	    rep, err := zmq.NewSocket(zmq.REP)
	    if err != nil {
	        log.Fatalln(err)
	    }

	    // REP socket monitor, all events
	    err = rep.Monitor("inproc://monitor.rep", zmq.EVENT_ALL)
	    if err != nil {
	        log.Fatalln(err)
	    }
	    go rep_socket_monitor("inproc://monitor.rep")

	    // Generate an event
	    rep.Bind("tcp://*:5555")
	    if err != nil {
	        log.Fatalln(err)
	    }

	    // Allow some time for event detection
	    time.Sleep(time.Second)

	    rep.Close()
	    zmq.Term()
	}
*/
func (soc *Socket) Monitor(addr string, events Event) error {
	return errorTaint()
}

/*
Receive a message part from a socket interpreted as an event.

For a description of flags, see: http://api.zeromq.org/4-1:zmq-msg-recv#toc2

For a description of event_type, see: http://api.zeromq.org/4-1:zmq-socket-monitor#toc3

For an example, see: func (*Socket) Monitor
*/
func (soc *Socket) RecvEvent(flags Flag) (event_type Event, addr string, value int, err error) {
	return Event(0), stringTaint(), intTaint(), errorTaint()
}

/*
Start built-in ØMQ proxy

See: http://api.zeromq.org/4-1:zmq-proxy#toc2
*/
func Proxy(frontend, backend, capture *Socket) error {
	return errorTaint()
}

/*
Start built-in ØMQ proxy with PAUSE/RESUME/TERMINATE control flow

Returns ErrorNotImplemented405 with ZeroMQ version < 4.0.5

See: http://api.zeromq.org/4-1:zmq-proxy-steerable#toc2
*/
func ProxySteerable(frontend, backend, capture, control *Socket) error {
	return errorTaint()
}

//. CURVE

/*
Encode a binary key as Z85 printable text

See: http://api.zeromq.org/4-1:zmq-z85-encode
*/
func Z85encode(data string) string {
	return stringTaint()
}

/*
Decode a binary key from Z85 printable text

See: http://api.zeromq.org/4-1:zmq-z85-decode
*/
func Z85decode(s string) string {
	return stringTaint()
}

/*
Generate a new CURVE keypair

See: http://api.zeromq.org/4-1:zmq-curve-keypair#toc2
*/
func NewCurveKeypair() (z85_public_key, z85_secret_key string, err error) {
	return stringTaint(), stringTaint(), errorTaint()
}

/*
Receive a message part with metadata.

This requires ZeroMQ version 4.1.0. Lower versions will return the message part without metadata.

The returned metadata map contains only those properties that exist on the message.

For a description of flags, see: http://api.zeromq.org/4-1:zmq-msg-recv#toc2

For a description of metadata, see: http://api.zeromq.org/4-1:zmq-msg-gets#toc3
*/
func (soc *Socket) RecvWithMetadata(flags Flag, properties ...string) (msg string, metadata map[string]string, err error) {
	return stringTaint(), mapTaint(), errorTaint()
}

/*
Receive a message part with metadata.

This requires ZeroMQ version 4.1.0. Lower versions will return the message part without metadata.

The returned metadata map contains only those properties that exist on the message.

For a description of flags, see: http://api.zeromq.org/4-1:zmq-msg-recv#toc2

For a description of metadata, see: http://api.zeromq.org/4-1:zmq-msg-gets#toc3
*/
func (soc *Socket) RecvBytesWithMetadata(flags Flag, properties ...string) (msg []byte, metadata map[string]string, err error) {
	return byteArrayTaint(), mapTaint(), errorTaint()
}

func hasCap(s string) (value bool) {
	return false
}

// Returns false for ZeroMQ version < 4.1.0
//
// Else: returns true if the library supports the ipc:// protocol
func HasIpc() bool {
	return false
}

// Returns false for ZeroMQ version < 4.1.0
//
// Else: returns true if the library supports the pgm:// protocol
func HasPgm() bool {
	return false
}

// Returns false for ZeroMQ version < 4.1.0
//
// Else: returns true if the library supports the tipc:// protocol
func HasTipc() bool {
	return false
}

// Returns false for ZeroMQ version < 4.1.0
//
// Else: returns true if the library supports the norm:// protocol
func HasNorm() bool {
	return false
}

// Returns false for ZeroMQ version < 4.1.0
//
// Else: returns true if the library supports the CURVE security mechanism
func HasCurve() bool {
	return false
}

// Returns false for ZeroMQ version < 4.1.0
//
// Else: returns true if the library supports the GSSAPI security mechanism
func HasGssapi() bool {
	return false
}
