package services

import (
	"context"
	"net/http"

	"github.com/skrevolve/product-svc/pkg/db"
	"github.com/skrevolve/product-svc/pkg/models"
	"github.com/skrevolve/product-svc/pkg/pb"
)

type Server struct {
	H db.Handler
	pb.UnimplementedProductServiceServer
}

// 제품 생성
func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequset) (*pb.CreateProductResponse, error) {

	var product models.Product

	product.Name = req.Name
	product.Stock = req.Stock
	product.Price = req.Price

	if result := s.H.DB.Create(&product); result.Error != nil {
		return &pb.CreateProductResponse{
			Status: http.StatusConflict,
			Error: result.Error.Error(),
		}, nil
	}

	return &pb.CreateProductResponse{
		Status: http.StatusCreated,
		ProductInfoId: product.ProductInfoId,
	}, nil
}

// 제품 검색
func (s *Server) FindOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {

	var product models.Product

	if result := s.H.DB.First(&product, req.ProductInfoId); result.Error != nil {
		return &pb.FindOneResponse{
			Status: http.StatusNotFound,
			Error: result.Error.Error(),
		}, nil
	}

	data := &pb.FindOneData{
		ProductInfoId: product.ProductInfoId,
		Name: product.Name,
		Stock: product.Stock,
		Price: product.Price,
	}

	return &pb.FindOneResponse{
		Status: http.StatusOK,
		Data: data,
	}, nil
}

// 제품 재고 감소
func (s *Server) DecreaseStock(ctx context.Context, req *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error) {

	var product models.Product

	if result := s.H.DB.First(&product, req.ProductInfoId); result.Error != nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusNotFound,
			Error: result.Error.Error(),
		}, nil
	}

	if product.Stock <= 0 {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error: "Stock too low",
		}, nil
	}

	var log models.StockDecreaseLog

	if result := s.H.DB.Where(&models.StockDecreaseLog{OrderInfoId: req.OrderInfoId}).First(&log); result.Error != nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error: "Stock already decreased",
		}, nil
	}

	product.Stock -= 1

	s.H.DB.Save(&product)

	log.OrderInfoId = req.OrderInfoId
	log.ProductRefer = product.ProductInfoId

	s.H.DB.Create(&log)

	return &pb.DecreaseStockResponse{
		Status: http.StatusOK,
	}, nil
}