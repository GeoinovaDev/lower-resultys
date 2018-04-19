package log

// WARNING : aviso de erro no sistema
// PANIC   : erro crítico no sistema
const (
	WARNING  = 1
	PANIC    = 2
	CRITICAL = 3
)

// ILogger interface para salvar dados de log
type ILogger interface {
	Save(string, int, string)
}

// DefaultLogger logger default
type DefaultLogger struct {
}

// Logger variavel logger default para salvar informações de log
var Logger ILogger = DefaultLogger{}

// Save salva mensagem, tipo e origem do erro
func (d DefaultLogger) Save(message string, tpe int, ip string) {

}
