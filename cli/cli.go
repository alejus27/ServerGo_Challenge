package cli

import (
	"bufio"
	. "file-server/structs"
	. "file-server/utils"
	"fmt"
	"os"
	"strings"
)


type SubjectCli interface {
	Register(observer CliObserver)
	Unregister(observer CliObserver)
	NotifyAll(options ...string)
}

// Define la estructura de la linea de comando personalizada.
type Cli struct {
	Observers []CliObserver
	Source    string
	Active    chan bool
}

// 
func NewCli(source string) *Cli {
	return &Cli{
		Source: source,
		Active: make(chan bool),
	}
}

// Inicialización de linea de comando .
func (cli *Cli) Start() {
	for {
		cli.listenInput()
	}
}



// Escucha por nuevos datos de entrada a través de la linea de comandos.
func (cli *Cli) listenInput() {
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		cli.handleInput(input)
	}
}

/*
Comandos prefinidos:

subscribe channel:name
unsubscribe channel:name
send channel:name file:path

start server
stop server
*/

// Controla las cadenas de entrada.
func (cli *Cli) handleInput(input string) {
	input = SingleSpacePattern.ReplaceAllString(strings.TrimSpace(input), " ")

	if cli.Source == "server" {
		switch input {
		case START:
			cli.NotifyAll(START)
		case STOP:
			cli.NotifyAll(STOP)
			cli.Active <- false
		case HELP:
			cli.help()
		default:
			cli.invalid()
		}
		return
	}

	options := strings.Split(input, " ")
	if len(options) < 1 {
		return
	}
	action := strings.ToLower(options[0])

	// Coincidencia de cadenas de texto a través de patrones de expresiones regulares
	if RegexSubscribe.MatchString(input) || RegexUnsubscribe.MatchString(input) {
		cli.NotifyAll(action, cli.value(options[1]))
		return
	}

	if RegexSend.MatchString(input) {
		cli.NotifyAll(action, cli.value(options[1]), cli.value(options[2]))
		return
	}

	switch action {
	case EXIT:
		cli.Active <- false
	case HELP:
		cli.help()
	default:
		cli.invalid()
	}

}

// Devuelve el valor de un parámetro.
func (cli *Cli) value(param string) string {
	data := strings.SplitN(param, ":", 2)
	return data[1]
}

// Ayuda.
func (cli *Cli) help() {
	PrintHelp("----subscribe channel:name----",
		"----unsubscribe channel:name----",
		"----send channel:name file:path (max file size: 5 MB)----")
}

// Parametros invalidos.
func (cli *Cli) invalid() {
	PrintError("Invalid parameters - Run command 'help'")
}

// Agrega un observador a la linea de comandos.
func (cli *Cli) Register(observer CliObserver) {
	cli.Observers = append(cli.Observers, observer)
}

func (cli *Cli) Unregister(observer CliObserver) {
	for i, obs := range cli.Observers {
		if obs.Identifier() == observer.Identifier() {
			cli.Observers = append(cli.Observers[:i], cli.Observers[i+1:]...)
		}
		
	}
}

// Notifica a todos los observadores.
func (cli *Cli) NotifyAll(options ...string) {
	for _, obs := range cli.Observers {
		if obs.Identifier() == cli.Source {
			obs.OnEntry(options)
		}
	}
}
