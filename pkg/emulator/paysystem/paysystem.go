package paysystem

import (
	"fmt"
	"github.com/go-openapi/errors"
	"github.com/mc_transaction/internal/logger"
	"math/rand"
	"strconv"
	"time"
)

const (
	StatusCreated     = "CREATED"
	StatusSuccess     = "SUCCESS"
	StatusNeedReverse = "NEED_TO_REVERSE"
)

var errTransactionCreate = errors.New(21, "transaction create failed")
var errTransactionReverse = errors.New(22, "transaction reverse failed")

type Transaction struct {
	PayId  int
	PayUrl string
}

func CreateTransaction(amount float64) (*Transaction, error) {
	time.Sleep(200 * time.Millisecond)

	rand.New(rand.NewSource(time.Now().UnixNano()))
	payId := rand.Intn(99999)
	payUrl := fmt.Sprintf("http://plati_padla/%d", payId)
	if payId > 95000 {
		logger.Warn("create transaction failed")
		return nil, errTransactionCreate
	}
	logger.Info("transaction created")
	return &Transaction{
		PayId:  payId,
		PayUrl: payUrl,
	}, nil
}

func CheckStatusTransaction(payId int) string {
	time.Sleep(50 * time.Millisecond)

	rand.New(rand.NewSource(time.Now().UnixNano()))
	randInt := rand.Intn(50)
	var status string
	if randInt%3 == 0 {
		status = StatusSuccess
	} else if randInt%4 == 0 {
		status = StatusNeedReverse
	} else {
		status = StatusCreated
	}
	logger.Info("check status for transaction with pay_id = " + strconv.Itoa(payId) + " status" + status)
	return status
}

func ReverseTransaction(payId int) error {
	time.Sleep(100 * time.Millisecond)

	rand.New(rand.NewSource(time.Now().UnixNano()))
	randInt := rand.Intn(99999)
	if randInt > 95000 {
		logger.Info("reverting transaction with pay_id = " + strconv.Itoa(payId) + " failed")
		return errTransactionReverse
	}
	logger.Info("reverting transaction with pay_id = " + strconv.Itoa(payId) + " success")
	return nil
}
