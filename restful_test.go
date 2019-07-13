package iotmaker_server_json

import "fmt"

func ExampleOut_SaveCache() {

	out := NewJSonOut()
	out.Id = "0000-0000-0000-0000"
	out.Meta.Success = true
	out.Meta.Limit = 30
	out.Meta.Offset = 0
	out.Meta.Next = "next.link"
	out.Meta.Previous = "previous.link"
	out.Objects = []string{"a", "b", "c"}

	err := out.SaveCache()
	if err != nil {
		panic(err)
	}

	id := out.Id

	out = NewJSonOut()
	err = out.LoadCache(id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", out)

	// output:
	// {{{ 30 next.link 0 previous.link 0 true []} [a b c]} 0000-0000-0000-0000}
}
