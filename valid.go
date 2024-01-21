package dragonvalid

import (
	"container/list" //see https://www.topgoer.cn/docs/golangstandard/golangstandard-1cmkstg63ir38
	"errors"
	"strconv"
)

type (
	// 验证引擎结构体
	Engine struct {
		err          error //错误信息
		defaultValue interface{}
		queue        *list.List
		name         string
		value        string
		sep          string
		valueInt     int
		valueFloat   float64
		setRawValue  bool
		silent       bool
		result       bool
	}
	queueT func(v *Engine) *Engine
)

var (
	ErrNoValidationValueSet = errors.New("未设置验证值")
)

// 初始化验证引擎
func New() Engine {
	return Engine{queue: list.New()}
}

// Int use int new valid
func Int(value int, name ...string) Engine {
	return Text(strconv.FormatInt((int64(value)), 10), name...)
}

// Text use int new valid
func Text(value string, name ...string) Engine {
	var obj Engine
	obj.value = value
	obj.setRawValue = true
	obj.queue = list.New()
	if len(name) > 0 {
		obj.name = name[0]
	}
	return obj
}

// 通过将一个验证函数推送到队列中，实现了对字段的非空验证。如果值为空，则设置相应的错误信息。
func (v Engine) Required(customError ...string) Engine {
	return pushQueue(&v, func(v *Engine) *Engine {
		if v.value == "" {
			v.err = setError(v, "不能为空", customError...)
		}
		return v
	})
}

// 将一个验证函数推送到队列中，并根据 DisableCheckErr 参数来决定是否执行错误检查。如果禁用了错误检查或者已经存在错误，则直接返回当前引擎。
func pushQueue(v *Engine, fn queueT, DisableCheckErr ...bool) Engine {
	pFn := fn
	if !(len(DisableCheckErr) > 0 && DisableCheckErr[0]) {
		pFn = func(v *Engine) *Engine {
			if v.err != nil {
				return v
			}
			return fn(v)
		}
	}
	queue := list.New()
	if v.queue != nil {
		queue.PushBackList(v.queue)
	}
	queue.PushBack(pFn)
	v.queue = queue
	return *v
}

// 生成一个错误信息，可以附加自定义错误信息。如果有自定义错误信息，则返回自定义错误，否则使用默认错误信息。
func setError(v *Engine, msg string, customError ...string) error {
	if len(customError) > 0 && customError[0] != "" {
		return errors.New(customError[0])
	}
	if v.name != "" {
		msg = v.name + " " + msg
	}
	return errors.New(msg)
}
