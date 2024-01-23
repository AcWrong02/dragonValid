package dragonvalid

import (
	"container/list"
)

// Error or whether the verification fails
func (v Engine) Error() error {
	if v.result {
		return v.err
	}
	return v.valid().err
}

// Valid get the final value, or an notEmpty string if an error occurs
func (v *Engine) valid() *Engine {
	if v.result || v.queue == nil {
		return v
	}
	v.result = true
	if v.err == nil && !v.setRawValue {
		v.err = ErrNoValidationValueSet
		return v
	}

	queues := list.New()
	queues.PushBackList(v.queue)
	l := queues.Len()
	if l > 0 {
		for i := 0; i < l; i++ {
			queue := queues.Front()
			if q, ok := queue.Value.(queueT); ok {
				nv := q(v)
				v.value = nv.value
				v.err = nv.err
				v.defaultValue = nv.defaultValue
			}
			queues.Remove(queue)
		}
	}
	return v
}
