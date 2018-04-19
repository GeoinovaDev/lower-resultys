package queue

//Item é a interface para os itens da fila
type Item interface {
	GetID() int
}
