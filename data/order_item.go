package data

import (
	"database/sql"
	"library/my_error"
)

func (orderItem *OrderItem) Insert() (err error) {
	statement := `INSERT OrderItem (publicationID, cardNO, orderDate) VALUES (?, ?, NOW())`
	_, err = Db.Query(statement, orderItem.Publication.PublicationID, orderItem.Card.CardNO)
	return
}

func (orderItem *OrderItem) Delete() (err error) {
	statement := "DELETE FROM OrderItem WHERE orderItemID=?"
	_, err = Db.Exec(statement, orderItem.OrderItemID)
	return
}

func (orderItem *OrderItem) RetrieveByCardNO() (results []OrderItem, err error) {
	statement := `SELECT o.orderItemID, o.publicationID, p.name, p.author, p.total, p.inventory FROM OrderItem o
    JOIN publication p on o.publicationID = p.publicationID
    WHERE cardNO=? and status=0`
	rows, err := Db.Query(statement, orderItem.Card.CardNO)
	if err != nil {
		return
	}
	for rows.Next() {
		result := &OrderItem{}
		err = rows.Scan(&result.OrderItemID, &result.Publication.PublicationID, &result.Publication.Name,
			&result.Publication.Author, &result.Publication.Total, &result.Publication.Inventory)
		if err != nil {
			return
		}
		results = append(results, *result)
	}
	return
}

func (orderItem *OrderItem) Borrow() (err error) {
	tx, err := Db.Begin()
	if err != nil {
		return
	}
	// 获取数据
	statement := `SELECT c.CardNO, p.publicationID FROM OrderItem o 
				  JOIN Publication p ON o.publicationID = p.publicationID
				  JOIN Card c ON o.cardNO = c.cardNO
				  WHERE o.orderItemID=?`
	err = tx.QueryRow(statement, orderItem.OrderItemID).Scan(&orderItem.Card.CardNO, &orderItem.Publication.PublicationID)
	if err != nil {
		tx.Rollback()
		return
	}
	// 检查是否有未归还图书
	statement = `SELECT COUNT(*) FROM BorrowItem JOIN Card ON BorrowItem.cardNO = Card.cardNO
				  								  JOIN BorrowerType ON Card.borrowerType = BorrowerType.borrowerType
                  WHERE BorrowItem.cardNO = ? AND DATE_ADD(borrowDate, INTERVAL BorrowerType.period DAY) < now()
                  AND status = 0`
	rows, err := tx.Query(statement, orderItem.Card.CardNO)
	if err != nil {
		tx.Rollback()
		return
	}
	if rows.Next() {
		var count = 0
		err = rows.Scan(&count)
		if err != nil {
			tx.Rollback()
			return
		} else if count != 0 {
			err = my_error.BorrowOutOfTimeError
			tx.Rollback()
			return
		}
	}
	rows.Close()
	// 库存-1
	statement = `UPDATE Publication SET inventory = inventory - 1 WHERE PublicationID = ?`
	_, err = tx.Query(statement, orderItem.Publication.PublicationID)
	if err != nil {
		tx.Rollback()
		err = my_error.InventoryNotEnoughError
		return
	}
	// 添加借阅信息
	statement = `INSERT INTO BorrowItem (cardNO, publicationID, borrowDate) VALUES(?, ?, NOW())`
	_, err = tx.Query(statement, orderItem.Card.CardNO, orderItem.Publication.PublicationID)
	if err != nil {
		tx.Rollback()
		return
	}
	// 查询借阅量是否达到上限
	statement = `SELECT currentBorrowNumber FROM Card INNER JOIN BorrowerType ON Card.borrowerType = BorrowerType.borrowerType 
				 WHERE cardNO = ? AND currentBorrowNumber < maxBorrowNumber`
	err = tx.QueryRow(statement, orderItem.Card.CardNO).Scan(&orderItem.Card.CurrentBorrowNumber)
	if err != nil && err.Error() == sql.ErrNoRows.Error() {
		err = my_error.MaxBorrowNumberError
		tx.Rollback()
		return
	} else if err != nil {
		return
	}
	statement = `UPDATE Card INNER JOIN BorrowerType ON Card.borrowerType = BorrowerType.borrowerType 
				 SET currentBorrowNumber = currentBorrowNumber + 1 WHERE cardNO = ? 
				 AND currentBorrowNumber < maxBorrowNumber`
	_, err = tx.Query(statement, orderItem.Card.CardNO)
	if err != nil {
		tx.Rollback()
		return
	}
	statement = `UPDATE OrderItem SET status=1 WHERE orderItemID=?`
	_, err = tx.Query(statement, orderItem.OrderItemID)
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}
