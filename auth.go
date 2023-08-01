/*

This file implements functionality very similar to that of the xauth module in czmq.

Notable differences in here:

 - domains are supported
 - domains are used in AuthAllow and AuthDeny too
 - usernames/passwords are read from memory, not from file
 - public keys are read from memory, not from file
 - an address can be a single IP address, or an IP address and mask in CIDR notation
 - additional functions for configuring server or client socket with a single command

*/

package zmq4_stub

const CURVE_ALLOW_ANY = "*"

// Start authentication.
//
// Note that until you add policies, all incoming NULL connections are allowed
// (classic ZeroMQ behaviour), and all PLAIN and CURVE connections are denied.
func AuthStart() (err error) {
	return errorTaint()
}

// Stop authentication.
func AuthStop() {
}

// Allow (whitelist) some addresses for a domain.
//
// An address can be a single IP address, or an IP address and mask in CIDR notation.
//
// For NULL, all clients from these addresses will be accepted.
//
// For PLAIN and CURVE, they will be allowed to continue with authentication.
//
// You can call this method multiple times to whitelist multiple IP addresses.
//
// If you whitelist a single address for a domain, any non-whitelisted addresses
// for that domain are treated as blacklisted.
//
// Use domain "*" for all domains.
//
// For backward compatibility: if domain can be parsed as an IP address, it will be
// interpreted as another address, and it and all remaining addresses will be added
// to all domains.
func AuthAllow(domain string, addresses ...string) {
}

// Deny (blacklist) some addresses for a domain.
//
// An address can be a single IP address, or an IP address and mask in CIDR notation.
//
// For all security mechanisms, this rejects the connection without any further authentication.
//
// Use either a whitelist for a domain, or a blacklist for a domain, not both.
// If you define both a whitelist and a blacklist for a domain, only the whitelist takes effect.
//
// Use domain "*" for all domains.
//
// For backward compatibility: if domain can be parsed as an IP address, it will be
// interpreted as another address, and it and all remaining addresses will be added
// to all domains.
func AuthDeny(domain string, addresses ...string) {
}

// Add a user for PLAIN authentication for a given domain.
//
// Set `domain` to "*" to apply to all domains.
func AuthPlainAdd(domain, username, password string) {
}

// Remove users from PLAIN authentication for a given domain.
func AuthPlainRemove(domain string, usernames ...string) {
}

// Remove all users from PLAIN authentication for a given domain.
func AuthPlainRemoveAll(domain string) {
}

// Add public user keys for CURVE authentication for a given domain.
//
// To cover all domains, use "*".
//
// Public keys are in Z85 printable text format.
//
// To allow all client keys without checking, specify CURVE_ALLOW_ANY for the key.
func AuthCurveAdd(domain string, pubkeys ...string) {
}

// Remove user keys from CURVE authentication for a given domain.
func AuthCurveRemove(domain string, pubkeys ...string) {
}

// Remove all user keys from CURVE authentication for a given domain.
func AuthCurveRemoveAll(domain string) {
}

// Enable verbose tracing of commands and activity.
func AuthSetVerbose(verbose bool) {
}

/*
This function sets the metadata handler that is called by the ZAP
handler to retrieve key/value properties that should be set on reply
messages in case of a status code "200" (succes).

Default properties are `Socket-Type`, which is already set, and
`Identity` and `User-Id` that are empty by default. The last two can be
set, and more properties can be added.

The `User-Id` property is used for the `user id` frame of the reply
message. All other properties are stored in the `metadata` frame of the
reply message.

The default handler returns an empty map.

For the meaning of the handler arguments, and other details, see:
http://rfc.zeromq.org/spec:27#toc10
*/
func AuthSetMetadataHandler(
	handler func(
		version, request_id, domain, address, identity, mechanism string, credentials ...string) (metadata map[string]string)) {
}

/*
This encodes a key/value pair into the format used by a ZAP handler.

Returns an error if key is more then 255 characters long.
*/
func AuthMetaBlob(key, value string) (blob []byte, err error) {
	return []byte{}, errorTaint()
}

//. Additional functions for configuring server or client socket with a single command

// Set NULL server role.
func (server *Socket) ServerAuthNull(domain string) error {
	return errorTaint()
}

// Set PLAIN server role.
func (server *Socket) ServerAuthPlain(domain string) error {
	return errorTaint()
}

// Set CURVE server role.
func (server *Socket) ServerAuthCurve(domain, secret_key string) error {
	return errorTaint()
}

// Set PLAIN client role.
func (client *Socket) ClientAuthPlain(username, password string) error {
	return errorTaint()
}

// Set CURVE client role.
func (client *Socket) ClientAuthCurve(server_public_key, client_public_key, client_secret_key string) error {
	return errorTaint()
}

// Helper function to derive z85 public key from secret key
//
// Returns ErrorNotImplemented42 with ZeroMQ version < 4.2
func AuthCurvePublic(z85SecretKey string) (z85PublicKey string, err error) {
	return stringTaint(), errorTaint()
}
