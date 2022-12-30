package structure

type Matrixable interface {
	DiagonalMatrix | Grid | OrthogonalLinkedList | QuadTree | Trigram
}
