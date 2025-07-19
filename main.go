package main

func main(){
	err := godoenv.Load(".payEnv")
	if (err != nil){
		fmt.Println("Error loading .payEnv file");
	}

	

}