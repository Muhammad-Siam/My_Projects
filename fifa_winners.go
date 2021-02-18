package main


import "fmt"


func main(){

	fmt.Println("Welcome to Statico! Improve your knowledge by taking all of the informations.")

	fmt.Println("You can know about all of the FIFA winners from the very begining")

	fmt.Println("Enter the name of the country:")

	var country string
	
	fmt.Scanln(&country)

	fmt.Println("The name of the country you want to know is:", country)


	switch country{

	case "Brazil":

	fmt.Println("Winner of 2002, 1994, 1970, 1962, 1958")

	case "Germany", "West Germany":

	fmt.Println("Winner of 2014, 1990, 1974, 1954")


	case "France":

	fmt.Println("Winner of 1998, 2018")

	case "Italy":

	fmt.Println("Winner of 2006, 1982, 1938, 1934")

	case "Argentina":

	fmt.Println("Winner of 1986, 1978")

	case "Uruguay":

	fmt.Println("Winner of 1930, 1950")

	case "England", "United Kingdom", "Britain":

	fmt.Println("Winner of 1966")

	case "Spain":

	fmt.Println("Winner of 2010")

	default:

	fmt.Println("Did not win any of the Fifa World Cup")




}


	





}