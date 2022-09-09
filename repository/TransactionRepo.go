package repository

import (
	"efishery-project/entities"
	"errors"

	"gorm.io/gorm"
)

type RepositoryTransaction interface {
	CreateTransaction(transaction entities.Transaction, user_id int) (entities.Transaction, error)
	InsertPaymentSlip(user_id int, id int, path string) (entities.Transaction, error)
	GetUserTransactions(user_id int) ([]entities.Transaction, error)
}

type Repository_Transaction struct {
	db *gorm.DB
}

func NewRepositoryTransaction(db *gorm.DB) *Repository_Transaction {
	return &Repository_Transaction{db}
}

func (r *Repository_Transaction) CreateTransaction(transaction entities.Transaction, user_id int) (entities.Transaction, error) {
	var cart []entities.Cart
	var detailFromCart []entities.CartDetail
	var transactions entities.Transaction

	//Get Cart data
	resultCart := r.db.Raw("select * from carts where user_id = ? and deleted_at is null", user_id).Scan(&cart)

	resultCartRowsAffected := resultCart.RowsAffected

	resultCartRowsAffectedError := errors.New("There is no product found in cart")

	if resultCartRowsAffected == 0 {
		return transactions, resultCartRowsAffectedError
	}

	errResultCart := resultCart.Error

	if errResultCart != nil {
		return transaction, errResultCart
	}

	//Get product price and product name
	for i := 0; i < len(cart); i++ {
		var detail entities.CartProduct
		getDetail := r.db.Raw("select name, price from products where id = ?", cart[i].Product_id).Scan(&detail)

		errgetDetail := getDetail.Error

		if errgetDetail != nil {
			return transactions, errgetDetail
		}

		detail1 := entities.CartDetail{cart[i].User_id, cart[i].Product_id, cart[i].Quantity, detail.Name, detail.Price, cart[i].Quantity * detail.Price}

		detailFromCart = append(detailFromCart, detail1)
	}

	//Get total price
	totalPrice := 0

	for i := 0; i < len(detailFromCart); i++ {
		totalPrice += detailFromCart[i].Total_price
	}

	//Initialize status
	status := "MENUNGGU_PEMBAYARAN"

	//Insert into transactions
	transactions = entities.Transaction{transaction.Model, user_id, "", totalPrice, status}
	insertToTransactions := r.db.Create(&transactions)
	transactionsPK := transactions.ID

	errinsertToTransactions := insertToTransactions.Error

	if errinsertToTransactions != nil {
		return transactions, errinsertToTransactions
	}

	//Insert into transaction detail
	for i := 0; i < len(detailFromCart); i++ {
		var transaction_detail entities.Transaction_Detail
		transaction_detail = entities.Transaction_Detail{transaction.Model, int(transactionsPK), detailFromCart[i].Product_name, detailFromCart[i].Quantity, detailFromCart[i].Total_price}

		insertToTransactionDetail := r.db.Create(&transaction_detail)

		errinsertToTransactionDetail := insertToTransactionDetail.Error

		if errinsertToTransactionDetail != nil {
			return transactions, errinsertToTransactionDetail
		}
	}

	//Delete from cart
	// deleteCart := r.db.Raw("delete from carts where user_id = ? and deleted_at is null", user_id).Delete(&cart)
	deleteCart := r.db.Where("user_id = ? and deleted_at is null", user_id).Unscoped().Delete(&entities.Cart{})

	errDeleteCart := deleteCart.Error

	if errDeleteCart != nil {
		return transactions, errDeleteCart
	}

	return transactions, nil
}

func (r *Repository_Transaction) InsertPaymentSlip(user_id int, id int, path string) (entities.Transaction, error) {
	var transactions entities.Transaction

	result := r.db.Model(&entities.Transaction{}).Where("user_id = ? and id = ? and status like 'MENUNGGU_PEMBAYARAN'", user_id, id).Updates(entities.Transaction{Payment_slip: path, Status: "DIBAYAR"})

	err := result.Error

	if err != nil {
		return transactions, err
	}

	getResult := r.db.Where("user_id = ? and id = ?", user_id, id).First(&transactions)

	errgetResult := getResult.Error

	if errgetResult != nil {
		return transactions, errgetResult
	}

	return transactions, nil
}

func (r *Repository_Transaction) GetUserTransactions(user_id int) ([]entities.Transaction, error) {
	var transactions []entities.Transaction

	result := r.db.Raw("select * from transactions where user_id = ? and deleted_at is null", user_id).Scan(&transactions)

	resulRowAffected := result.RowsAffected

	errNoTransaction := errors.New("No transaction to be displayed")

	if resulRowAffected == 0 {
		return transactions, errNoTransaction
	}

	err := result.Error

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
