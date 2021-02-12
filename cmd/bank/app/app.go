package app

import (
	"encoding/json"
	"errors"
	"github.com/Geniuskaa/task10.1/cmd/bank/app/dto"
	"github.com/Geniuskaa/task10.1/pkg/card"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	cardSvc *card.Service
	mux *http.ServeMux
}

func NewServer(cardSvc *card.Service, mux *http.ServeMux) *Server {
	return &Server{
		cardSvc: cardSvc,
		mux:     mux,
	}
}

func (s *Server) Init() {
	s.mux.HandleFunc("/getCards", s.getCards)
	s.mux.HandleFunc("/addCard", s.addCard)
	//s.mux.HandleFunc("/editCard", s.editCard)
	//s.mux.HandleFunc("/removeCard", s.removeCard)
}

func (s *Server) getCards(w http.ResponseWriter, r *http.Request) {
	mapWithValues := r.URL.Query()
	value1 := mapWithValues["id"]
	id, err := strconv.Atoi(value1[0])
	if err != nil {
		log.Println(err)
		return
	}
	cards := s.cardSvc.All(r.Context())
	dtos := make([]*dto.CardDTO, len(cards))
	counter := 0
	for i, c := range cards {
		if c.Id == int64(id) {
			counter++
			dtos[i] = &dto.CardDTO{
				Card: card.Card{
					Issuer:   c.Issuer,
					Number:   c.Number,
					Currency: c.Currency,
					Balance:  c.Balance,
					Virtual:  c.Virtual,
					Id:       c.Id,
				},
			}
		}
	}

	if counter == 0 {
		w.WriteHeader(404)
		dtos := "There are not any cardHolders with this ID."
		respBody, _ := json.Marshal(dtos)
		w.Header().Add("Content-Type", "text/plain")
		_, err = w.Write(respBody)
		log.Println(errors.New("There are not any cardHolders with this ID."))
		return
	}

	respBody, err := json.Marshal(dtos)
	if err != nil {
		w.WriteHeader(404)
		log.Println(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(respBody)
	if err != nil {
		w.WriteHeader(404)
		log.Println(err)
		return
	}
}

func (s *Server) addCard(w http.ResponseWriter, r *http.Request) {
	mapWithValues := r.URL.Query()
	value1 := mapWithValues["id"]
	id, err := strconv.Atoi(value1[0])
	if err != nil {
		w.WriteHeader(404)
		log.Println(err)
		return
	}
	value2 := mapWithValues["virtualCard"]
	virtualCard, err := strconv.ParseBool(value2[0])
	if err != nil {
		w.WriteHeader(404)
		log.Println(err)
		return
	}

	err = s.cardSvc.CardAdding(int64(id), virtualCard)
	if err != nil {
		w.WriteHeader(404)
		log.Println(err)
		return
	}

	dtos := "Card was succesfully added!"
	respBody, err := json.Marshal(dtos)
	if err != nil {
		w.WriteHeader(404)
		log.Println(err)
		return
	}

	w.Header().Add("Content-Type", "text/plain")
	_, err = w.Write(respBody)
	if err != nil {
		w.WriteHeader(404)
		log.Println(err)
		return
	}
}
