package main


import "fmt"


func main(){
	
	fmt.Println("Periodic Table is one of the most important and fundamental part of Chemistry. Improve your knowledge by knowing all the Information about Periodic Table")

	fmt.Println("There are total of 114 of Molcules in Periodic Table")

	fmt.Println("There are totally 7 Periods in Periodic Table")

	per1 := []string {"Hydrogen(H)","Helium(He)"}
	per2 := []string { "Lithium(Li)", "Beryliuam(Be)", "Boron(B)", "Carbon(C)", "Nitrogen(N)", "Oxygen(O)", "Flourine(F)", "Neon(Ne)"}
	per3 := []string { "Sodium(Na)", "Magnesium(Mg)", "Aluminium(Al)", "Silicon(Si)", "Phosphorus(P)", "Sulfur(S)", "Chlorine(Cl)", "Argon(Ar)"}
	
	fmt.Println(per1, per2, per3)



	fmt.Println("Enter the numbers of the molcules to know:")

	var mol int

	fmt.Scanln(&mol)

	fmt.Println("The number of the molecules you want to know is:", mol)

	 if mol<=10{

	fmt.Println("Hydrogen","Helium", "Lithium", "Berrylium", "Boron","Carbon", "Nitrogen","Oxygen", "Flourine", "Neon")

        }else{

	fmt.Println("We wil work for this later")


	
	}



}