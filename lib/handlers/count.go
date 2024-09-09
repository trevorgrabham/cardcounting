package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/trevorgrabham/cardcounting/cardcounting/html"
	"github.com/trevorgrabham/cardcounting/cardcounting/lib"
)

func HandleCheckCount(w http.ResponseWriter, r *http.Request) {
	// userID, ok := r.Context().Value("user-id").(int64)
	// if !ok { panic(fmt.Errorf("error parsing 'user-id' from context")) }
	checkCount := template.Must(template.New("check-count").ParseFiles(html.IncludeFiles["check-count"]...))
	if err := checkCount.Execute(w, nil); err != nil { panic(err) }
}

func HandleValidateCount(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user-id").(int64)
	if !ok { panic(fmt.Errorf("error parsing 'user-id' from context")) }
	if err := r.ParseForm(); err != nil { panic(err) }
	countString := r.Form.Get("count")
	res, err := strconv.ParseInt(countString, 0, 16)
	if err != nil { panic(err) }
	count := int16(res)
	correct, correctCount := lib.UserData.CheckCount(userID, count)
	countValidation := template.Must(template.New("count-validation").ParseFiles(html.IncludeFiles["count-validation"]...))
	if err := countValidation.Execute(w, html.NewCountValidationData(correct, correctCount)); err != nil { panic(err) }
}