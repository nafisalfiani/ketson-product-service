package usecase

import (
	"github.com/nafisalfiani/ketson-product-service/domain"
	"github.com/nafisalfiani/ketson-product-service/lib/log"
	"github.com/nafisalfiani/ketson-product-service/usecase/category"
	"github.com/nafisalfiani/ketson-product-service/usecase/region"
	"github.com/nafisalfiani/ketson-product-service/usecase/ticket"
	"github.com/nafisalfiani/ketson-product-service/usecase/wishlist"
)

type Usecases struct {
	Ticket   ticket.Interface
	Category category.Interface
	Region   region.Interface
	Wishlist wishlist.Interface
}

func Init(logger log.Interface, dom *domain.Domains) *Usecases {
	return &Usecases{
		Ticket:   ticket.Init(logger, dom.Ticket, dom.Category, dom.Region),
		Category: category.Init(dom.Category),
		Region:   region.Init(dom.Region),
		Wishlist: wishlist.Init(dom.Wishlist),
	}
}
