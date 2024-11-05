package usecase

import (
	"context"
	"fmt"

	"github.com/SatoKeiju/shiharai-kun/app/domain/model"
	"github.com/SatoKeiju/shiharai-kun/app/domain/repository"
)

// InvoiceUseCase : 請求書に関するユースケースが満たすべきユースケース
type InvoiceUseCase interface {
	FetchList(ctx context.Context, userID string, from string, to string) ([]model.Invoice, error)
}

type invoiceUseCase struct {
	repository.Invoice
	repository.User
}

// NewInvoiceUseCase : 請求書に関するユースケースを作成
func NewInvoiceUseCase() InvoiceUseCase {
	return invoiceUseCase{}
}

// FetchList : ユーザーが所属する企業に紐づく請求書を一覧で取得
func (u invoiceUseCase) FetchList(ctx context.Context, userID string, from string, to string) ([]model.Invoice, error) {
	user, err := u.User.FetchByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("UserRepository.FetchByID(userID: %s): %w", userID, err)
	}

	list, err := u.Invoice.FetchListByCompanyID(ctx, user.CompanyID, from, to)
	if err != nil {
		return nil, fmt.Errorf("InvoiceRepository.FetchListByCompanyID(companyID: %s, from: %s, to: %s): %w", user.CompanyID, from, to, err)
	}

	return list, nil
}
