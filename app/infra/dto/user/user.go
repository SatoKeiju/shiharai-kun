package user

import "github.com/SatoKeiju/shiharai-kun/app/domain/model/user"

// DTO : ユーザー情報DTO
type DTO struct {
	ID          string `db:"id"`
	CompanyID   string `db:"company_id"`
	Name        string `db:"name"`
	MailAddress string `db:"mail_address"`
}

// ConvertFromDTO : ユーザー情報DTOからユーザードメインモデルを生成
func ConvertFromDTO(d DTO) user.User {
	return user.User{
		ID:          d.ID,
		CompanyID:   d.CompanyID,
		Name:        d.Name,
		MailAddress: d.MailAddress,
	}
}
