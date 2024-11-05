package usecase

import (
	"context"
	"fmt"

	"github.com/SatoKeiju/shiharai-kun/app/domain/model/invoice"
	"github.com/SatoKeiju/shiharai-kun/app/domain/repository"
)

// InvoiceUseCase : 請求書に関するユースケースが満たすべきユースケース
type InvoiceUseCase interface {
	Create(ctx context.Context, userID string, clientID string, issueDate string, paymentAmount int, paymentDueDate string) (invoice.Invoice, error)
	FetchList(ctx context.Context, userID string, from string, to string) ([]invoice.Invoice, error)
}

type invoiceUseCase struct {
	ri repository.Invoice
	ru repository.User
}

// NewInvoiceUseCase : 請求書に関するユースケースを作成
func NewInvoiceUseCase(ri repository.Invoice, ru repository.User) InvoiceUseCase {
	return invoiceUseCase{ri: ri, ru: ru}
}

// Create : 請求書データを作成
func (u invoiceUseCase) Create(ctx context.Context, userID string, clientID string, issueDate string, paymentAmount int, paymentDueDate string) (invoice.Invoice, error) {
	user, err := u.ru.FetchByID(ctx, userID)
	if err != nil {
		return invoice.Invoice{}, fmt.Errorf("UserRepository.FetchByID(userID: %s): %w", userID, err)
	}

	i, err := invoice.New(user.CompanyID, clientID, issueDate, paymentAmount, paymentDueDate)
	if err != nil {
		return invoice.Invoice{}, fmt.Errorf("invoiceModel.New(companyID: %s, clientID: %s, issueDate: %s, paymentAmount: %d, paymentDueDate: %s): %w", user.CompanyID, clientID, issueDate, paymentAmount, paymentDueDate, err)
	}

	created, err := u.ri.Create(ctx, i)
	if err != nil {
		return invoice.Invoice{}, fmt.Errorf("InvoiceRepository.Create(model: %+v): %w", i, err)
	}

	return created, nil
}

// FetchList : ユーザーが所属する企業に紐づく請求書を一覧で取得
func (u invoiceUseCase) FetchList(ctx context.Context, userID string, from string, to string) ([]invoice.Invoice, error) {
	user, err := u.ru.FetchByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("UserRepository.FetchByID(userID: %s): %w", userID, err)
	}

	list, err := u.ri.FetchListByCompanyID(ctx, user.CompanyID, from, to)
	if err != nil {
		return nil, fmt.Errorf("InvoiceRepository.FetchListByCompanyID(companyID: %s, from: %s, to: %s): %w", user.CompanyID, from, to, err)
	}

	return list, nil
}
