package card

import (
	"context"
	"errors"
	"fmt"
	"sync"
)


type Card struct {
	Issuer   string `json:"issuer"`
	Number   string `json:"number"`
	Currency string `json:"currency"`
	Balance  int64 `json:"balance"`
	Virtual  bool `json:"virtual"`
	Id       int64 `json:"id"`
}

type Service struct {
	ids int64
	mu sync.RWMutex
	cards []*Card
}

func NewService() *Service {
	return &Service{
		ids: 3,
		mu:    sync.RWMutex{},
		cards: []*Card{&Card{
			Issuer:   "VISA",
			Number:   "0001",
			Currency: "RUB",
			Balance:  325_325_33,
			Virtual:  false,
			Id:       1,
		}, &Card{
			Issuer:   "VISA",
			Number:   "0002",
			Currency: "RUB",
			Balance:  29_391_31,
			Virtual:  false,
			Id:       2,
		}, &Card{
			Issuer:   "MASTER",
			Number:   "0003",
			Currency: "RUB",
			Balance:  35_23,
			Virtual:  true,
			Id:       2,
		}},
	}
}

func (s *Service) CardAdding(yourId int64, virtualCard bool) error {
	var n string
	if yourId > 10 && yourId < 100 {
		n = fmt.Sprintf("00%d", yourId + 1)
	}
	n = fmt.Sprintf("000%d", yourId + 1)

	for _, element := range s.cards{
		if element.Id == yourId {
			s.cards = append(s.cards, &Card{
				Issuer:   "VISA",
				Number:   n,
				Currency: "RUB",
				Balance:  0,
				Virtual:  virtualCard,
				Id:       yourId,
			})
			return nil
		}
	}

	return errors.New("There is not cardHolder with this Id, before Adding the card, register in our Bank.")

}

func (s *Service) All(ctx context.Context) []*Card {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.cards
}