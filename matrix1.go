package main

// import "fmt"

// func main(){
//    var a,b int
//   fmt.Printf("a= ")
//   fmt.Scanf("%d",&a)
//   fmt.Printf("b= ")
//   fmt.Scanf("%d",&b)

//   fmt.Printf("a=%d\t b=%d\n",a,b)

//   var S1 [][]float64
//   S1 = make([][]float64,a)
//   Read(S1,a,b)
//   P(S1,a,b)
//   fmt.Println("S1=",S1)

//   var S2 [][]float64
//   S2 = make([][]float64,a)
//   Read(S2,a,b)
//   P(S2,a,b)
//   fmt.Println("S2=",S2)

//   var Sum[][]float64
//   Sum = make([][]float64,a)
//   for i:=0; i<a; i++{
// 	Sum[i] = make([]float64, b)
//   }

//   for i:=0; i<a; i++{
// 	for j:= 0; j<b; j++{
// 		Sum[i][j] = S1[i][j] + S2[i][j]
// 	}
// }
// fmt.Println("Sum=",Sum)
// }

// func P(S [][]float64, a,b int){
// 	for i:=0; i < a; i++{
// 	  for e:=0; e < b; e++{
// 		  fmt.Printf("%.2f\t",S[i][e])
// 	  }
// 	  fmt.Println()
// 	}
//   }

// func Read(Slice [][]float64, c,d int){
// 	for i:=0; i<c; i++{
// 		Slice[i] = make([]float64, d)
// 		for j:= 0; j<d; j++{
// 			fmt.Printf("Slice[%d][%d]=",i,j)
// 			fmt.Scanf("%f",&Slice[i][j])
// 		}

// 	}
// }