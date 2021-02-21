package card

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)


type Card struct {
	Issuer   string `json:"issuer"`
	Number   string `json:"number"`
	Currency string `json:"currency"`
	Balance  int64 `json:"balance"`
	Virtual  bool `json:"virtual"`
	CardId   int64 `json:"card_id"`
	HolderId int64 `json:"holder_id"`
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
			CardId: 62826,
			HolderId: 1,
		}, &Card{
			Issuer:   "VISA",
			Number:   "0002",
			Currency: "RUB",
			Balance:  29_391_31,
			Virtual:  false,
			CardId: 85920,
			HolderId: 2,
		}, &Card{
			Issuer:   "MASTER",
			Number:   "0003",
			Currency: "RUB",
			Balance:  35_23,
			Virtual:  true,
			CardId: 14262,
			HolderId: 2,
		}},
	}
}

func (s *Service) CardAdding(yourId int64, issuer string, virtualCard bool) error {
	var n string
	if yourId > 10 && yourId < 100 {
		n = fmt.Sprintf("00%d", yourId + 1)
	}
	n = fmt.Sprintf("000%d", yourId + 1)

	rand.Seed(time.Now().UnixNano())
	randId := rand.Int63() % 100000

	for _, element := range s.cards{
		if element.HolderId == yourId {
			s.cards = append(s.cards, &Card{
				Issuer:   issuer,
				Number:   n,
				Currency: "RUB",
				Balance:  0,
				Virtual:  virtualCard,
				CardId: randId,
				HolderId: yourId,
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