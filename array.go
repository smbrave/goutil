package goutil

import "reflect"

func Sort(data interface{}, fun interface{}) {
	fValue := reflect.ValueOf(fun)
	fType := fValue.Type()
	dValue := reflect.Indirect(reflect.ValueOf(data))
	dType := dValue.Type()

	if (dType.Kind() != reflect.Slice && dType.Kind() != reflect.Array) ||
		fType.Kind() != reflect.Func || fType.NumIn() != 2 || fType.NumOut() != 1 {
		panic("sort paramter type format error")
	}

	firstType := fType.In(0)
	secondType := fType.In(1)
	outType := fType.Out(0)
	elemType := dType.Elem()
	if elemType.Kind() != firstType.Kind() || elemType.Kind() != secondType.Kind() ||
		outType.Kind() != reflect.Bool {
		panic("sort elem type format error")
	}

	tmp := reflect.Indirect(reflect.New(elemType))
	for i := 0; i < dValue.Len(); i++ {
		for j := i + 1; j < dValue.Len(); j++ {
			out := fValue.Call([]reflect.Value{dValue.Index(i), dValue.Index(j)})
			if out[0].Bool() {
				continue
			}

			tmp.Set(dValue.Index(i))
			dValue.Index(i).Set(dValue.Index(j))
			dValue.Index(j).Set(tmp)
		}
	}

}

func Filter(data interface{}, fun interface{}) {
	fValue := reflect.ValueOf(fun)
	fType := fValue.Type()
	dValue := reflect.Indirect(reflect.ValueOf(data))
	dType := dValue.Type()

	if (dType.Kind() != reflect.Slice && dType.Kind() != reflect.Array) ||
		fType.Kind() != reflect.Func || fType.NumIn() != 1 || fType.NumOut() != 1 {
		panic("sort paramter type format error")
	}

	firstType := fType.In(0)
	outType := fType.Out(0)
	elemType := dType.Elem()
	if elemType.Kind() != firstType.Kind() || outType.Kind() != reflect.Bool {
		panic("sort elem type format error")
	}

	pos := 0
	for i := 0; i < dValue.Len(); i++ {
		out := fValue.Call([]reflect.Value{dValue.Index(i)})
		if !out[0].Bool() {
			continue
		}

		dValue.Index(pos).Set(dValue.Index(i))
		pos++
	}

	dValue.SetLen(pos)
}

func Traverse(data interface{}, fun interface{}) {
	fValue := reflect.ValueOf(fun)
	fType := fValue.Type()
	dValue := reflect.Indirect(reflect.ValueOf(data))
	dType := dValue.Type()

	if (dType.Kind() != reflect.Slice && dType.Kind() != reflect.Array) ||
		fType.Kind() != reflect.Func || fType.NumIn() != 1 || fType.NumOut() != 0 {
		panic("sort paramter type format error")
	}

	firstType := fType.In(0)
	elemType := dType.Elem()
	if elemType.Kind() != firstType.Kind() {
		panic("sort elem type format error")
	}

	for i := 0; i < dValue.Len(); i++ {
		fValue.Call([]reflect.Value{dValue.Index(i)})
	}
}
