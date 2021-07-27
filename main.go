package main

import (
	"net/http"
	"fmt"
	"github.com/blitzblade/generic_ussd_api/utils"

)

func ussdHandler(w http.ResponseWriter, r *http.Request){
	utils.Log("Request received successfully!", utils.LogInfo)
	w.Write([]byte(utils.FrontPageUnregistered))
}

func main() {
	defer func(){
		fmt.Println("Server shutting down...")
		utils.EndLog()
	}()

	go utils.InitLogger()

	http.HandleFunc("/ussd", ussdHandler)
	
	utils.Log("USSD API running...", utils.LogInfo)
	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		panic(err.Error())
	}
}
