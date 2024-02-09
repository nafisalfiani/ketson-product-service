package domain

import (
	"github.com/nafisalfiani/ketson-product-service/domain/category"
	"github.com/nafisalfiani/ketson-product-service/domain/region"
	"github.com/nafisalfiani/ketson-product-service/domain/ticket"
	"github.com/nafisalfiani/ketson-product-service/domain/wishlist"
	"github.com/nafisalfiani/ketson-product-service/lib/broker"
	"github.com/nafisalfiani/ketson-product-service/lib/cache"
	"github.com/nafisalfiani/ketson-product-service/lib/log"
	"github.com/nafisalfiani/ketson-product-service/lib/parser"
	"go.mongodb.org/mongo-driver/mongo"
)

type Domains struct {
	Ticket   ticket.Interface
	Category category.Interface
	Region   region.Interface
	Wishlist wishlist.Interface
}

func Init(logger log.Interface, json parser.JSONInterface, db *mongo.Client, cache cache.Interface, broker broker.Interface) *Domains {
	return &Domains{
		Ticket:   ticket.Init(logger, json, db.Database("ketson_product_db").Collection("ticket"), cache, broker),
		Category: category.Init(logger, json, db.Database("ketson_product_db").Collection("category"), cache),
		Region:   region.Init(logger, json, db.Database("ketson_product_db").Collection("region"), cache),
		Wishlist: wishlist.Init(logger, json, db.Database("ketson_product_db").Collection("wishlist")),
	}
}
