package zmq4_stub

/*
Send multi-part message on socket.

Any `[]string' or `[][]byte' is split into separate `string's or `[]byte's

Any other part that isn't a `string' or `[]byte' is converted
to `string' with `fmt.Sprintf("%v", part)'.

Returns total bytes sent.
*/
func (soc *Socket) SendMessage(parts ...interface{}) (total int, err error) {
	return intTaint(), errorTaint()
}

/*
Like SendMessage(), but adding the DONTWAIT flag.
*/
func (soc *Socket) SendMessageDontwait(parts ...interface{}) (total int, err error) {
	return intTaint(), errorTaint()
}

func (soc *Socket) sendMessage(dontwait Flag, parts ...interface{}) (total int, err error) {
	return intTaint(), errorTaint()
}

/*
Receive parts as message from socket.

Returns last non-nil error code.
*/
func (soc *Socket) RecvMessage(flags Flag) (msg []string, err error) {
	return stringArrayTaint(), errorTaint()
}

/*
Receive parts as message from socket.

Returns last non-nil error code.
*/
func (soc *Socket) RecvMessageBytes(flags Flag) (msg [][]byte, err error) {
	return byteArrayArrayTaint(), errorTaint()

}

/*
Receive parts as message from socket, including metadata.

Metadata is picked from the first message part.

For details about metadata, see RecvWithMetadata().

Returns last non-nil error code.
*/
func (soc *Socket) RecvMessageWithMetadata(flags Flag, properties ...string) (msg []string, metadata map[string]string, err error) {
	return stringArrayTaint(), mapTaint(), errorTaint()
}

/*
Receive parts as message from socket, including metadata.

Metadata is picked from the first message part.

For details about metadata, see RecvBytesWithMetadata().

Returns last non-nil error code.
*/
func (soc *Socket) RecvMessageBytesWithMetadata(flags Flag, properties ...string) (msg [][]byte, metadata map[string]string, err error) {
	return byteArrayArrayTaint(), mapTaint(), errorTaint()
}
