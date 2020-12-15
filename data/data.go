// @Title data.go
// @Description 数据库模型
package data

import "time"

type Administrator struct {
	// 系统管理员
	AdministratorID int64  `json:"administrator_id"`
	UserName        string `json:"user_name"`
	Password        string `json:"password"`
}

type Librarian struct {
	// 图书管理员
	LibrarianID int64  `json:"librarian_id"`
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
}

type BorrowerType struct {
	// 借阅者类型
	BorrowerType    string `json:"borrower_type"`
	Period          int64  `json:"period"`
	MaxBorrowNumber int64  `json:"max_borrow_number"`
}

type Card struct {
	// 借阅证
	CardNO              string       `json:"card_no"`
	Name                string       `json:"name"`
	Major               string       `json:"major"`
	BorrowerType        BorrowerType `json:"borrower_type"`
	CurrentBorrowNumber int64        `json:"current_borrow_number"`
}

type Borrower struct {
	// 借阅者
	UID      int64  `json:"uid"`
	Card     Card   `json:"card"`
	UserName string `json:"user_name"`
	Password string `json:"-"`
}

type PublicationType struct {
	// 出版物类型对应的罚金
	PublicationType string `json:"publication_type"`
	Fine            int64  `json:"fine"`
}

type Publication struct {
	// 出版物信息
	PublicationID   int64           `json:"publication_id"`
	Name            string          `json:"name"`
	ISBN            string          `json:"isbn"`
	Price           int64           `json:"price"`
	Total           int64           `json:"total"`
	Inventory       int64           `json:"inventory"`
	PublicationType PublicationType `json:"publication_Type"`
	Author          string          `json:"author"`
}

type BorrowItem struct {
	// 借书信息
	BorrowItemID int64       `json:"borrow_item_id"`
	Card         Card        `json:"card"`
	Publication  Publication `json:"publication"`
	BorrowDate   time.Time   `json:"borrow_date"`
	DueDate      time.Time   `json:"due_date"`
	Status       int64       `json:"status"`
}

type LostItem struct {
	// 丢失信息
	LostItemID int64      `json:"lost_item_id"`
	Card       Card       `json:"card"`
	BorrowItem BorrowItem `json:"borrow_item"`
	LostDate   time.Time  `json:"lost_date"`
}

type OrderItem struct {
	// 预定信息
	OrderItemID int64       `json:"order_item_id"`
	Publication Publication `json:"publication"`
	Card        Card        `json:"card"`
	OrderDate   time.Time   `json:"order_date"`
	Status      int64       `json:"status"`
}

type OverTimeItem struct {
	// 超时信息
	OvertimeID int64      `json:"overtime_id"`
	BorrowItem BorrowItem `json:"borrow_item"`
	Card       Card       `json:"card"`
	DueDate    time.Time  `json:"due_date"`
}
