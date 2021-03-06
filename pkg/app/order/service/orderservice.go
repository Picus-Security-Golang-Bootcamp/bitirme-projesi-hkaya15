package service

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/cart/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/cart/repository"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/order/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/order/repository"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/product/repository"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/errors"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/jwt"
)

type OrderService struct {
	OrderRepo   *OrderRepository
	CartRepo    *CartRepository
	ProductRepo *ProductRepository
}

func NewOrderService(o *OrderRepository, crt *CartRepository, pr *ProductRepository) *OrderService {
	return &OrderService{OrderRepo: o, CartRepo: crt, ProductRepo: pr}
}

func (o *OrderService) Migrate() {
	o.OrderRepo.Migrate()
}

func (o *OrderService) CompleteOrder(user *AccessTokenDetails) error {
	cart, err := o.CartRepo.CreateCart(ResponseToCart(user.UserID))
	if err != nil {
		return NewRestError(http.StatusBadRequest, os.Getenv("CREATE_CART_ISSUE"), err.Error())
	}

	userCart, err := o.CartRepo.GetCartList(cart)
	if err != nil {
		return NewRestError(http.StatusBadRequest, os.Getenv("GET_CART_ISSUE"), err.Error())
	}

	cartItems, err := o.CartRepo.GetCartItems(cart.UserID)
	if err != nil {
		return err
	}

	if len(cartItems) == 0 {
		return errors.New(os.Getenv("CART_EMPTY_FAIL"))
	}

	orderItems := make([]OrderItem, 0)
	for _, v := range cartItems {
		orderItems = append(orderItems, *NewOrderItem(user.UserID, v))
	}

	err = o.OrderRepo.CompleteOrder(NewOrder(user.UserID, orderItems))
	if err != nil {
		return err
	}

	for _, v := range userCart.Items {
		res, err := o.CartRepo.Delete(v)
		if err != nil {
			return errors.New(os.Getenv("DELETE_ITEM_ISSUE"))
		}
		if res == true {
			continue
		}
	}

	cart.CartLength=0
	cart.CartTotalPrice=0
	

	err = o.CartRepo.UpdateCart(cart)
	if err != nil {
		return NewRestError(http.StatusBadRequest, os.Getenv("ISSUE_ON_UPDATE_CART"), nil)
	}

	return nil
}

func (o *OrderService) GetAllOrders(userid string) ([]Order, error) {
	return o.OrderRepo.GetAllOrders(userid)
}

func (o *OrderService) CancelOrder(userid string, orderid string) error {
	order, err := o.OrderRepo.FindByOrderID(orderid)
	if err != nil {
		return err
	}
	cancelRule, _ := strconv.ParseFloat(os.Getenv("FOURTEEN_DAYS"), 64)

	if order.OrderDate.Sub(time.Now()).Hours() > cancelRule {
		return NewRestError(http.StatusNotAcceptable, os.Getenv("CANCEL_ORDER_DAY_ISSUE"), nil)
	}

	err = o.OrderRepo.CancelOrder(order)
	if err != nil {
		return err
	}
	return nil
}
