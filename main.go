package main

import (
	"net/http"
	"fmt"
	"strings"
	"html/template"
	"github.com/blitzblade/generic_ussd_api/utils"

)

type USSDParams struct {
	SessionID string
    MobileNumber string
	MessageType string
	NetworkCode string

    Body []string
}

func unescape(s string) template.HTML {
	return template.HTML(s)
}

func ussdHandler(w http.ResponseWriter, r *http.Request){
	utils.Log("Request received successfully!", utils.LogInfo)
	urlQuery := r.URL.Query()
	fmt.Println(urlQuery)
	// body := utils.HTMLString(utils.FrontPageUnregistered)
	// body.MakeHTMLReady()
	body := strings.Split(utils.FrontPageUnregistered, "\n")
	params := USSDParams {MobileNumber: "+233247915505", Body: body}

	renderTemplate(w, "templates/index", &params)
}

func renderTemplate(w http.ResponseWriter, tmpl string, params *USSDParams) {
    t, _ := template.ParseFiles(tmpl + ".html") 
  
    t.Execute(w, params)
}

// func saveHandler(w http.ResponseWriter, r *http.Request){
// 	body := r.FormValue("body")
//     params := USSDParams{MobileNumber: r.FormValue("msisdn"), Body: []byte(body)}
//     // t, _ := template.ParseFiles("edit.html")
// 	utils.Log(params.MobileNumber, utils.LogInfo)
// 	// if err != nil {
//     //     http.Error(w, err.Error(), http.StatusInternalServerError)
//     //     return
//     // }
//     // t.Execute(w, p)
	
// 	http.Redirect(w, r, "/ussd", http.StatusFound)
// }

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
