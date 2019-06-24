package main
import (
"fmt"
"math"
)
func main() {
fmt.Println("CHOOSE FROM THE FOLLOWING OPTIONS:\n"+
"                           \n"+
"1.Future Amount Calculator\n"+
"2.Principal Amount Calculator\n"+
"3.Interest Rate Calculator\n"+
"4.Number of Periods Calculator\n")
  
const FA = 1
const PA = 2
const IR = 3
const NP = 4
var CH float64
fmt.Print("Enter a choice(1-4): ")
fmt.Scanf("%f\n", &CH)
if CH == FA {
        fmt.Println("You have chosen Future Amount Calculator")
		calcFA()
} else if CH == PA {
	fmt.Println("You have chosen Principal Amount Calculator")
		calcPA()
} else if CH == IR {
	fmt.Println("You have chosen Interest Rate Calculator")
		calcIR()
} else if CH == NP {
	fmt.Println("You have chosen Number of Periods Calculator")
		calcNP()
} else {
	error := "Enter a number between 1 & 4"
        fmt.Println(error)
}
}
func calcFA() {
var p float64
var s string
var n float64
var r float64
fmt.Print("Enter a Principal Amount(£): ")
fmt.Scanf("%f\n", &p)
fmt.Print("Enter a Type of Period(days-weeks-months-years): ")
fmt.Scanf("%s\n", &s)
fmt.Print("Enter a Number of Periods("+s+"): ") 
fmt.Scanf("%f\n", &n) 
fmt.Print("Enter a Interest Rate(%): ")
fmt.Scanf("%f\n", &r)
f := math.Pow(float64(r)/100+1, float64(n))*p  
F := fmt.Sprintf("%.2f", f) 
fmt.Println("Future Amount:£",F)
}
func calcPA() {
var f float64
var s string
var n float64
var r float64
fmt.Print("Enter a Future Amount(£): ")
fmt.Scanf("%f\n", &f)
fmt.Print("Enter a Type of Period(days-weeks-months-years): ")
fmt.Scanf("%s\n", &s)
fmt.Print("Enter a Number of Periods("+s+"): ") 
fmt.Scanf("%f\n", &n) 
fmt.Print("Enter a Interest Rate(%): ")
fmt.Scanf("%f\n", &r)
p := math.Pow(float64(r)/100+1, float64(n)) 
fp := f/p
P := fmt.Sprintf("%.2f", fp) 
fmt.Println("Principal Amount:£",P)
}
func calcIR() {
var f float64
var s string
var n float64
var p float64
fmt.Print("Enter a Principal Amount(£): ")
fmt.Scanf("%f\n", &p)
fmt.Print("Enter a Future Amount(£): ")
fmt.Scanf("%f\n", &f)
fmt.Print("Enter a Type of Period(days-weeks-months-years): ")
fmt.Scanf("%s\n", &s)
fmt.Print("Enter a Number of Periods("+s+"): ") 
fmt.Scanf("%f\n", &n) 
fp := f/p
nd := 1/n
r := (math.Pow(float64(fp), float64(nd))-1)*100
R := fmt.Sprintf("%.2f", r) 
fmt.Println("Interest Rate:",R,"%")
}
func calcNP() {

}
