package repository

import (
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/order/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (o *OrderRepository) Migrate() {
	o.db.AutoMigrate(&Order{})
	o.db.AutoMigrate(&OrderItem{})
}

func (o *OrderRepository) CompleteOrder(order *Order) error{
	zap.L().Debug("order.repo.complete", zap.Reflect("order", order))
	if err := o.db.Where(Order{UserID: order.UserID}).Attrs(order).Create(order).Error; err != nil {
		zap.L().Error("order.repo.Create failed to create order", zap.Error(err))
		return err
	}
	return nil
}

func (o *OrderRepository)GetAllOrders(userid string) ([]Order, error){
	var orders []Order
	err:=o.db.Where("user_id",userid).
	Preload("OrderItems").
	Preload("OrderItems.User").
	Preload("OrderItems.Product").
	Preload("OrderItems.Product.Category").
		Find(&orders).Error
		if err != nil {
			zap.L().Error("order.repo.getAllOrders failed to get orderlist", zap.Error(err))
			return nil, err
		}
	return orders,nil
}

func (o *OrderRepository) CancelOrder(order *Order)error{
	zap.L().Debug("order.repo.cancelorder", zap.Reflect("orderid", order))
	if err:= o.db.Model(Order{ID:order.ID}).Update("is_canceled", true).Error; err!=nil{
		zap.L().Error("order.repo.cancelorder failed to cancel order", zap.Error(err))
		return err
	}
	return nil
}

func (o *OrderRepository) FindByOrderID(orderid string) (*Order, error) {
	zap.L().Debug("order.repo.findorderbyid", zap.String("orderid", orderid))
	var order *Order
	if err := o.db.Where("is_canceled = ?", false).Where("id", orderid).First(&order).Error; err != nil {
		zap.L().Error("order.repo.findorderbyid failed to get order", zap.Error(err))
		return nil, err
	}
	return order, nil

}