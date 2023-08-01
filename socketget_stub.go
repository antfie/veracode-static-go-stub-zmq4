package zmq4_stub

func (soc *Socket) GetFd() (int, error) {
	return intTaint(), errorTaint()
}
