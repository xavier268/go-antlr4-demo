//Package generate is only used to hold the go:generate commands.
package generate

//The following lines will trigger the automatic generation
//go:generate rm -f -r internal/pkg/auto/
//go:generate antlr -Dlanguage=Go -visitor -listener -package auto -o internal/pkg/auto Expr.g4
