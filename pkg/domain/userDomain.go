package domain

type User struct {
	IdUser   int    `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	UserName string `json:"user_name" gorm:"unique" validate:"required,min=2,max=50"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"notnull;unique" validate:"email,required"`
}
