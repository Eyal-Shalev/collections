package collections

type Container[I any] interface {
	Has(item I) bool
}

type Iterable[I any] interface {
	Iter() Iterator[I]
}

type Iterator[I any] interface {
	Iterable[I]
	Next() (I, bool)
}

type Sized interface {
	Len() uint
}

type Cloneable[T any] interface {
	Clone() T
}

type Collection[I any] interface {
	Container[I]
	Iterable[I]
	Sized
}

type Gettable[K comparable, V any] interface {
	Get(key K) (V, bool)
}

type Sliceable[K comparable, Self any] interface {
	Slice(start, end K) (Self, error)
}

type Sequence[I any] interface {
	Collection[I]
	Gettable[uint, I]
	Sliceable[uint, I]
}

type Settable[K comparable, I any] interface {
	Set(key K, value I)
}

type Deletable[K comparable, I any] interface {
	Delete(key K) (I, error)
}

type Appendable[I any] interface {
	Append(item I) error
}

type MutableSequence[I any] interface {
	Sequence[I]
	Settable[uint, I]
	Deletable[uint, I]
	Appendable[I]
}

type Set[K comparable] interface {
	Collection[K]
}

type Removable[K comparable] interface {
	Remove(item K) error
}

type MutableSet[K comparable] interface {
	Set[K]
	Appendable[K]
	Deletable[K, K]
}

type MapEntry[K comparable, V any] struct {
	Key   K
	Value V
}

type Map[K comparable, V any] interface {
	Collection[MapEntry[K, V]]
	Gettable[K, V]
	Keys() Sequence[K]
	Values() Sequence[V]
}

type MutableMap[K comparable, V any] interface {
	Map[K, V]
	Settable[K, V]
	Deletable[K, V]
}
