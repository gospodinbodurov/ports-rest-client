package readerjob

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	api "github.com/gospodinbodurov/ports-apis/port-domain-service/api"
	clients "github.com/gospodinbodurov/ports-rest-client/clients"
)

type ReaderJob struct {
	Filename string
}

func (rj *ReaderJob) Run() {
	filename := rj.Filename

	log.Println("Reading file: " + filename)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	reader := bufio.NewReader(file)

	dec := json.NewDecoder(reader)

	_, err = dec.Token()

	if err != nil {
		log.Fatal(err)
	}

	for {
		token, err := dec.Token()

		if err == io.EOF {
			break
		} else {
			if err != nil {
				log.Fatal(err)
			}
		}

		var port api.Port

		errDecode := dec.Decode(&port)

		if errDecode == io.EOF {
			break
		} else {
			if errDecode != nil {
				log.Fatal(errDecode)
			}
		}

		key := fmt.Sprintf("%v", token)

		port.PortKey = key

		rj.SendPort(&port)
	}
}

func (rj *ReaderJob) SendPort(port *api.Port) {
	err := clients.ServiceClient.PutPort(port)

	if err != nil {
		log.Fatal(err.Error())
	}
}
