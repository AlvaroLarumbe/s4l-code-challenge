package helper

import "encoding/json"

type Optional[T any] struct {
	Value   *T
	Defined bool
}

func (o *Optional[T]) IsDefined() bool {
	return o.Defined
}

func (o *Optional[T]) Get() (T, bool) {
	return *o.Value, o.Defined
}

func (o *Optional[T]) MustGet() T {
	return *o.Value
}

// UnmarshalJSON will be called only if the value is defined in the JSON payload.
func (o *Optional[T]) UnmarshalJSON(data []byte) error {
	o.Defined = true

	return json.Unmarshal(data, &o.Value)
}
