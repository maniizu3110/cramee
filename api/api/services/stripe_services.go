package services

import (
	"cramee/token"
	"cramee/util"
)

type StripeService interface {
}

type StripeRepository interface {
}

type stripeServiceImpl struct {
	config     util.Config
	tokenMaker token.Maker
}

func NewStripeService(config util.Config, tokenMaker token.Maker) StripeService {
	res := &stripeServiceImpl{}
	res.config = config
	res.tokenMaker = tokenMaker
	return res
}
