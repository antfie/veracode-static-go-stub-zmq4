package zmq4_stub

// An Errno is an unsigned number describing an error condition as returned by a call to ZeroMQ.
// It implements the error interface.
// The number is either a standard system error, or an error defined by the C library of ZeroMQ.
type Errno uintptr

const (
	// Error conditions defined by the C library of ZeroMQ.

	// On Windows platform some of the standard POSIX errnos are not defined.
	EADDRINUSE      = Errno(0)
	EADDRNOTAVAIL   = Errno(0)
	EAFNOSUPPORT    = Errno(0)
	ECONNABORTED    = Errno(0)
	ECONNREFUSED    = Errno(0)
	ECONNRESET      = Errno(0)
	EHOSTUNREACH    = Errno(0)
	EINPROGRESS     = Errno(0)
	EMSGSIZE        = Errno(0)
	ENETDOWN        = Errno(0)
	ENETRESET       = Errno(0)
	ENETUNREACH     = Errno(0)
	ENOBUFS         = Errno(0)
	ENOTCONN        = Errno(0)
	ENOTSOCK        = Errno(0)
	ENOTSUP         = Errno(0)
	EPROTONOSUPPORT = Errno(0)
	ETIMEDOUT       = Errno(0)

	// Native 0MQ error codes.
	EFSM           = Errno(0)
	EMTHREAD       = Errno(0)
	ENOCOMPATPROTO = Errno(0)
	ETERM          = Errno(0)
)

// Return Errno as string.
func (errno Errno) Error() string {
	return stringTaint()
}

/*
Convert error to Errno.

Example usage:

	switch AsErrno(err) {

	case zmq.Errno(syscall.EINTR):
	    // standard system error

	    // call was interrupted

	case zmq.ETERM:
	    // error defined by ZeroMQ

	    // context was terminated

	}

See also: examples/interrupt.go
*/
func AsErrno(err error) Errno {
	return Errno(0)
}
