package interfaces

type Animal interface {
	Respirar()
	Comer()
	EsCarnivoro() bool
	Sexo() string
	EstaVivo() bool
}
