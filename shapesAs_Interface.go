package main

import ("fmt"
        "math"
)

type shape interface {
  area() float64
}

type square struct{
side float64
}

type circle struct{
radius float64
}

func (s square) area() float64{ //this function is in both
return s.side * s.side
}

func (c circle) area() float64{ //this function is in both
  return math.Pi * math.Sqrt(c.radius)
}

func interfaceTakes(s shape){ //takes the interface
  fmt.Println(s.area())
}

func main(){
fmt.Println("fdgd")
sq1 := square{
  44.0,
}
ci1 := circle{
  22.0,
}
 interfaceTakes(sq1)
 interfaceTakes(ci1)

}
