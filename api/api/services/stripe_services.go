package services

import (
	"cramee/token"
	"cramee/util"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/client"
)

type StripeService interface {
	CreateCustomer(params *stripe.CustomerParams) (*stripe.Customer, error)
	CreateCard(params *stripe.CardParams) (*stripe.Card, error)
	Charge(params *stripe.ChargeParams) (*stripe.Charge, error)
}

type StripeRepository interface {
}

type stripeServiceImpl struct {
	config       util.Config
	tokenMaker   token.Maker
	stripeClient *client.API
}

func NewStripeService(config util.Config, tokenMaker token.Maker, stripeClient *client.API) StripeService {
	res := &stripeServiceImpl{}
	res.config = config
	res.tokenMaker = tokenMaker
	res.stripeClient = stripeClient
	return res
}

func (s *stripeServiceImpl) CreateCustomer(params *stripe.CustomerParams) (*stripe.Customer, error) {
	customer, err := s.stripeClient.Customers.New(params)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
func (s *stripeServiceImpl) CreateCard(params *stripe.CardParams) (*stripe.Card, error) {
	card, err := s.stripeClient.Cards.New(params)
	if err != nil {
		return nil, err
	}
	return card, nil
}
func (s *stripeServiceImpl) Charge(params *stripe.ChargeParams) (*stripe.Charge, error) {
	charge, err := s.stripeClient.Charges.New(params)
	if err != nil {
		return nil, err
	}
	return charge, nil
}
