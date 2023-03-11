package main // principal package
import (
	"fmt" // print things of code in terminal
	"os"
)

// import "reflect" shows the type of the variable

func main() {

   displayIntro()

	 command := readCommand()

	 displayMenu()

	switch command {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}
}

func displayIntro() {
	name := "Fran"
	version := 1.1

	fmt.Println("Hello, miss.", name)
	fmt.Println("This program is in version", version)
}

func readCommand() int {
	var commandRead int
	fmt.Scan(&commandRead)

	return commandRead
}

func displayMenu()  {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}