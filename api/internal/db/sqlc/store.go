package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Store interface {
	Querier
	CheckoutTx(ctx context.Context, arg CheckoutTxParams) (CheckoutTxResult, error)
	PaymentTx(ctx context.Context, arg PaymentTxParams) (PaymentTxResult, error)
	PaymentConfirmationTx(ctx context.Context, arg PaymentConfirmationTxParams) (PaymentConfirmationTxResult, error)
}

// Store - provides all functions to execute SQL queries and transactions
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore - creates a new Store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// execTx - executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type OrderTicketParams struct {
	Qty      int32     `json:"qty"`
	TicketID uuid.UUID `json:"ticket_id"`
}

type CheckoutTxParams struct {
	CustomerID      uuid.UUID           `json:"customer_id"`
	TotalPrice      int64               `json:"total_price"`
	Discount        int32               `json:"discount"`
	FinalPrice      int64               `json:"final_price"`
	ListOrderTicket []OrderTicketParams `json:"list_order_ticket"`
}

type CheckoutTxResult struct {
	CustomerOrder CustomerOrder                         `json:"customer_order"`
	ListTicket    []ListOrderTicketByCustomerOrderIdRow `json:"list_ticket"`
}

func (store *SQLStore) CheckoutTx(ctx context.Context, arg CheckoutTxParams) (CheckoutTxResult, error) {
	var result CheckoutTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.CustomerOrder, err = q.CreateCustomerOrder(ctx, CreateCustomerOrderParams{
			ID:         uuid.New(),
			CustomerID: arg.CustomerID,
			OrderTime:  time.Now(),
			TotalPrice: arg.TotalPrice,
			Discount: sql.NullInt32{
				Int32: arg.Discount,
				Valid: true,
			},
			FinalPrice: arg.FinalPrice,
		})

		if err != nil {
			return err
		}

		for _, ot := range arg.ListOrderTicket {
			_, err = q.CreateOrderTicket(ctx, CreateOrderTicketParams{
				Qty: sql.NullInt32{
					Int32: ot.Qty,
					Valid: true,
				},
				TicketID:        ot.TicketID,
				CustomerOrderID: result.CustomerOrder.ID,
			})

			if err != nil {
				return err
			}
		}

		result.ListTicket, err = q.ListOrderTicketByCustomerOrderId(ctx, ListOrderTicketByCustomerOrderIdParams{
			CustomerOrderID: result.CustomerOrder.ID,
			Limit:           10,
			Offset:          1,
		})

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

type PaymentTxParams struct {
	CustomerOrderID uuid.UUID `json:"customer_order_id"`
	PaymentOptionID int32     `json:"payment_option_id"`
}

type PaymentTxResult struct {
	CustomerPayment CustomerPayment `json:"customer_payment"`
}

func (store *SQLStore) PaymentTx(ctx context.Context, arg PaymentTxParams) (PaymentTxResult, error) {
	var result PaymentTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.CustomerPayment, err = q.CreateCustomerPayment(ctx, CreateCustomerPaymentParams{
			ID:              uuid.New(),
			Status:          "UNPAID",
			PaymentOptionID: arg.PaymentOptionID,
			CustomerOrderID: arg.CustomerOrderID,
		})

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

type PaymentConfirmationTxParams struct {
	PaymentID    uuid.UUID `json:"payment_id"`
	Status       string    `json:"status"`
	FailedReason string    `json:"failed_reason"`
}

type PaymentConfirmationTxResult struct {
	CustomerPayment CustomerPayment `json:"customer_payment"`
}

func (store *SQLStore) PaymentConfirmationTx(ctx context.Context, arg PaymentConfirmationTxParams) (PaymentConfirmationTxResult, error) {
	var result PaymentConfirmationTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		paymentStatus := "UNPAID"
		var successAt sql.NullTime
		if arg.Status == "SUCCESS" {
			successAt = sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			}
			paymentStatus = "PAID"
		}

		var failedReason sql.NullString
		if arg.Status == "FAILED" {
			failedReason = sql.NullString{
				String: arg.FailedReason,
				Valid:  true,
			}
		}

		result.CustomerPayment, err = q.UpdateCustomerPayment(ctx, UpdateCustomerPaymentParams{
			Status:       paymentStatus,
			SuccessAt:    successAt,
			FailedReason: failedReason,
		})

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
