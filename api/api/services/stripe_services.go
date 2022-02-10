package services

import (
	"cramee/token"
	"cramee/util"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/client"
)

type StripeService interface {
	CreateCustomer(params *stripe.CustomerParams) (*stripe.Customer, error)
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
