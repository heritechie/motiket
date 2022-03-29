// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateCustomer(ctx context.Context, arg CreateCustomerParams) (Customer, error)
	CreateCustomerOrder(ctx context.Context, arg CreateCustomerOrderParams) (CustomerOrder, error)
	CreateCustomerPayment(ctx context.Context, arg CreateCustomerPaymentParams) (CustomerPayment, error)
	CreateEvent(ctx context.Context, arg CreateEventParams) (Event, error)
	CreateOrderTicket(ctx context.Context, arg CreateOrderTicketParams) (OrderTicket, error)
	CreatePaymentOption(ctx context.Context, name string) (PaymentOption, error)
	CreateTicket(ctx context.Context, arg CreateTicketParams) (Ticket, error)
	CreateTicketCategory(ctx context.Context, arg CreateTicketCategoryParams) (TicketCategory, error)
	GetCustomer(ctx context.Context, id uuid.UUID) (Customer, error)
	GetCustomerOrder(ctx context.Context, id uuid.UUID) (CustomerOrder, error)
	GetCustomerPayment(ctx context.Context, id uuid.UUID) (CustomerPayment, error)
	GetEvent(ctx context.Context, id uuid.UUID) (Event, error)
	GetPaymentOption(ctx context.Context, id int32) (PaymentOption, error)
	GetTicket(ctx context.Context, id uuid.UUID) (Ticket, error)
	GetTicketCategory(ctx context.Context, id uuid.UUID) (TicketCategory, error)
	ListCustomer(ctx context.Context, arg ListCustomerParams) ([]Customer, error)
	ListCustomerOrder(ctx context.Context, arg ListCustomerOrderParams) ([]CustomerOrder, error)
	ListCustomerPayment(ctx context.Context, arg ListCustomerPaymentParams) ([]CustomerPayment, error)
	ListEvent(ctx context.Context, arg ListEventParams) ([]Event, error)
	ListOrderTicket(ctx context.Context, arg ListOrderTicketParams) ([]OrderTicket, error)
	ListPaymentOption(ctx context.Context, arg ListPaymentOptionParams) ([]PaymentOption, error)
	ListTicket(ctx context.Context, arg ListTicketParams) ([]Ticket, error)
	ListTicketCategory(ctx context.Context, arg ListTicketCategoryParams) ([]TicketCategory, error)
}

var _ Querier = (*Queries)(nil)