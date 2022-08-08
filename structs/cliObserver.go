package structs

// Estructura del observador.
type CliObserver interface {
	OnEntry(options []string)
	Identifier() string
}
