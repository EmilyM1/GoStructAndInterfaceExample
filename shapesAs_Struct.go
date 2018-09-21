package main

import ("fmt"
"math"
)

type square struct{
sside float64
}

type circle struct{
cradius float64
}


func areaSquare(s square)float64{
return s.sside * s.sside
}

func circleArea(c circle)float64{
  return math.Pi * math.Sqrt(c.cradius)
}

func main(){
fmt.Println("fdgd")
sq1 := square{
  4.0,
}
ci1 := circle{
  2.0,
}
the := areaSquare(sq1)
fmt.Println(the)

tci := circleArea(ci1)
fmt.Println(tci)
}
