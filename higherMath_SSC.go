//radius = 2.3, what is the perim?
package main
import (
	"fmt"
	"math"
)

type geometry interface{
	perim() float32
}

type wheel struct{

	radius float32

}

type wheel1 struct{

	radius float32

}

func (w wheel) perim() float32{

	return 2 * math.Pi * w.radius

}

func (w1 wheel1) perim() float32{

	return w1.radius * math.Pi * 6

}




func main() {

	r:= wheel{radius: 2.3}

	y:= wheel1{radius: 0.84}

	

	g:= func(g geometry){

	fmt.Println(g)
	fmt.Println(g.perim())
}

	g(r)
	g(y)
}