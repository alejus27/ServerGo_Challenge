package utils

import "regexp"

// Define las constantes generales del programa
const (
	maxSizeMb   = 5
	SUBSCRIBE   = "subscribe"
	UNSUBSCRIBE = "unsubscribe"
	RECEIVE = "receive"
	SEND        = "send"
	HELP        = "help"
	EXIT        = "exit"
	START       = "server start"
	STOP        = "server stop"
	MAX_SIZE    = 1024 * 1024 * maxSizeMb
)

// Patrones de expresiones regulares
// Analiza una expresión regular y devuelve, un objeto Regexp que se puede usar para compararlo con el texto.

var RegexSubscribe, _ = regexp.Compile("\\s*^" + SUBSCRIBE + "\\s*channel:\\w*\\s*")
var RegexUnsubscribe, _ = regexp.Compile("\\s*^" + UNSUBSCRIBE + "\\s*channel:\\w*\\s*")
var RegexSend, _ = regexp.Compile("\\s*^" + SEND + "\\s*channel:\\w*\\s*file:.*\\s*")
var SingleSpacePattern = regexp.MustCompile(`\s+`)
