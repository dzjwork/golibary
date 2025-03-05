package mg

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

// 表示一个可以用mg.Deps运行的函数
type Fn interface {
	// 函数的完全限定名
	Name() string

	// 额外的一个唯一限定符，防止Name不够唯一
	ID() string

	// 执行该函数
	Run(ctx context.Context) error
}

// F takes a function that is compatible as a mage target, and any args that need to be passed to
// it, and wraps it in an mg.Fn that mg.Deps can run. Args must be passed in the same order as they
// are declared by the function. Note that you do not need to and should not pass a context.Context
// to F, even if the target takes a context. Compatible args are int, bool, string, and
// time.Duration.
func F(target interface{}, args ...interface{}) Fn {
	hasContext, isNamespace, err := checkF(target, args)
	if err != nil {
		panic(err)
	}
	id, err := json.Marshal(args)
	if err != nil {
		panic(fmt.Errorf("can't convert args into a mage-compatible id for mg.Deps: %s", err))
	}
	return fn{
		name: funcName(target),
		id:   string(id),
		f: func(ctx context.Context) error {
			v := reflect.ValueOf(target)
			count := len(args)
			if hasContext {
				count++
			}
			if isNamespace {
				count++
			}
			vargs := make([]reflect.Value, count)
			x := 0
			if isNamespace {
				vargs[0] = reflect.ValueOf(struct{}{})
				x++
			}
			if hasContext {
				vargs[x] = reflect.ValueOf(ctx)
				x++
			}
			for y := range args {
				vargs[x+y] = reflect.ValueOf(args[y])
			}
			ret := v.Call(vargs)
			if len(ret) > 0 {
				// we only allow functions with a single error return, so this should be safe.
				if ret[0].IsNil() {
					return nil
				}
				return ret[0].Interface().(error)
			}
			return nil
		},
	}
}

type fn struct {
	name string
	id   string
	f    func(ctx context.Context) error
}

func (f fn) Name() string {
	return f.name
}

func (f fn) ID() string {
	return f.id
}

func (f fn) Run(ctx context.Context) error {
	return f.f(ctx)
}

func checkF(target interface{}, args []interface{}) (hasContext, isNamespace bool, _ error) {
	t := reflect.TypeOf(target)
	if t == nil || t.Kind() != reflect.Func {
		return false, false, fmt.Errorf("non-function passed to mg.F: %T. The mg.F function accepts function names, such as mg.F(TargetA, \"arg1\", \"arg2\")", target)
	}

	if t.NumOut() > 1 {
		return false, false, fmt.Errorf("target has too many return values, must be zero or just an error: %T", target)
	}
	if t.NumOut() == 1 && t.Out(0) != errType {
		return false, false, fmt.Errorf("target's return value is not an error")
	}

	// more inputs than slots is an error if not variadic
	if len(args) > t.NumIn() && !t.IsVariadic() {
		return false, false, fmt.Errorf("too many arguments for target, got %d for %T", len(args), target)
	}

	if t.NumIn() == 0 {
		return false, false, nil
	}

	x := 0
	inputs := t.NumIn()

	if t.In(0).AssignableTo(emptyType) {
		// nameSpace func
		isNamespace = true
		x++
		// callers must leave off the namespace value
		inputs--
	}
	if t.NumIn() > x && t.In(x) == ctxType {
		// callers must leave off the context
		inputs--

		// let the upper function know it should pass us a context.
		hasContext = true

		// skip checking the first argument in the below loop if it's a context, since first arg is
		// special.
		x++
	}

	if t.IsVariadic() {
		if len(args) < inputs-1 {
			return false, false, fmt.Errorf("too few arguments for target, got %d for %T", len(args), target)

		}
	} else if len(args) != inputs {
		return false, false, fmt.Errorf("wrong number of arguments for target, got %d for %T", len(args), target)
	}

	for _, arg := range args {
		argT := t.In(x)
		if t.IsVariadic() && x == t.NumIn()-1 {
			// For the variadic argument, use the slice element type.
			argT = argT.Elem()
		}
		if !argTypes[argT] {
			return false, false, fmt.Errorf("argument %d (%s), is not a supported argument type", x, argT)
		}
		passedT := reflect.TypeOf(arg)
		if argT != passedT {
			return false, false, fmt.Errorf("argument %d expected to be %s, but is %s", x, argT, passedT)
		}
		if x < t.NumIn()-1 {
			x++
		}
	}
	return hasContext, isNamespace, nil
}

// Here we define the types that are supported as arguments/returns
var (
	ctxType   = reflect.TypeOf(func(context.Context) {}).In(0)
	errType   = reflect.TypeOf(func() error { return nil }).Out(0)
	emptyType = reflect.TypeOf(struct{}{})

	intType    = reflect.TypeOf(int(0))
	stringType = reflect.TypeOf(string(""))
	boolType   = reflect.TypeOf(bool(false))
	durType    = reflect.TypeOf(time.Second)

	// don't put ctx in here, this is for non-context types
	argTypes = map[reflect.Type]bool{
		intType:    true,
		boolType:   true,
		stringType: true,
		durType:    true,
	}
)
