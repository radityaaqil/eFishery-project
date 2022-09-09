package controller

import (
	"efishery-project/entities"
	"efishery-project/service"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Controller_Transaction struct {
	servicetransaction service.ServiceTransaction
}

func NewControllerTransaction(servicetransaction service.ServiceTransaction) *Controller_Transaction {
	return &Controller_Transaction{servicetransaction}
}

func (co *Controller_Transaction) CreateTransactionController(c echo.Context) error {
	var createInput entities.Transaction

	user_id_tkn := c.Get("currentUser").(entities.User)
	createInput.User_id = int(user_id_tkn.ID)

	createInputTransaction, errCreateInputTransaction := co.servicetransaction.CreateTransactionService(createInput, createInput.User_id)

	if errCreateInputTransaction != nil {
		log.Println("Failed to create transaction")
		return c.JSON(http.StatusBadRequest, errCreateInputTransaction.Error())
	}

	return c.JSON(http.StatusOK, createInputTransaction)
}

func (co *Controller_Transaction) InsertPaymentSlipController(c echo.Context) error {
	var update entities.Transaction

	transaction_id := c.Param("transaction_id")
	transaction_idInt, _ := strconv.Atoi(transaction_id)

	// log.Println("Line 44")

	//Get User id
	user_id_tkn := c.Get("currentUser").(entities.User)
	update.User_id = int(user_id_tkn.ID)
	// log.Println("Line 48")
	file, err := c.FormFile("payment_slip")
	if err != nil {
		return err
	}
	// log.Println("Line 53")
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	// log.Println("Line 59")

	// Destination
	path := "paymentslip/"
	now := time.Now()
	sec := now.Unix()
	secStr := strconv.Itoa(int(sec))
	dst, err := os.Create(path + "payment-slip-" + secStr + "-" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()
	// log.Println("Line 66")

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	// log.Println("Line 71")

	//Input filename
	update.Payment_slip = file.Filename

	updateTransaction, errupdateTransaction := co.servicetransaction.InsertPaymentSlipService(update.User_id, transaction_idInt, update.Payment_slip)

	if errupdateTransaction != nil {
		return c.String(http.StatusBadRequest, "Failed to upload payment slip")
	}
	// log.Println("Line 80")
	return c.JSON(http.StatusOK, updateTransaction)
}

func (co *Controller_Transaction) GetUserTransactionsController(c echo.Context) error {
	//Get User id
	user_id_tkn := c.Get("currentUser").(entities.User)
	user_id := int(user_id_tkn.ID)

	getUserTransactions, errgetUserTransactions := co.servicetransaction.GetUserTransactionsService(user_id)

	if errgetUserTransactions != nil {
		return c.String(http.StatusBadRequest, "Failde to load transactions")
	}

	return c.JSON(http.StatusOK, getUserTransactions)
}
