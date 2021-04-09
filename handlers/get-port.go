package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	clients "github.com/gospodinbodurov/ports-rest-client/clients"
)

type GetPortForm struct {
	PortKey string
}

func GetPort(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var form GetPortForm
	action := "GetPort"

	err := decoder.Decode(&form)
	if err != nil {
		log.Println(err.Error())

		Write(w, Response{
			Code:   http.StatusInternalServerError,
			Action: action,
			Data:   "500 - Can not decode form!",
		})

		return
	}

	port, err := clients.ServiceClient.GetPort(form.PortKey)

	if err != nil {
		log.Println(err.Error())

		Write(w, Response{
			Code:   http.StatusBadRequest,
			Action: action,
			Data:   "400 - Can not get port!",
		})

		return
	}

	Write(w, Response{
		Code:   http.StatusAccepted,
		Action: action,
		Data:   port,
	})
}
