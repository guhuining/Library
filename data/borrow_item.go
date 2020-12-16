package data

import "library/my_error"

// @title BorrowItem.Borrow
// @description	借书
// @param	CardNO			BorrowItem.CardNO			借阅证号码
// @param	PublicationID	Publication.PublicationID	出版物ID
func (borrowItem *BorrowItem) Borrow() (err error) {
	tx, err := Db.Begin()
	if err != nil {
		return
	}
	statement := `UPDATE Publication SET inventory = inventory - 1 WHERE PublicationID = ?`
	_, err = tx.Query(statement, borrowItem.Publication.PublicationID)
	if err != nil {
		tx.Rollback()
		err = my_error.InventoryNotEnoughError
		return
	}
	statement = `INSERT INTO BorrowItem (cardNO, publicationID, borrowDate) VALUES(?, ?, NOW())`
	_, err = tx.Query(statement, borrowItem.Card.CardNO, borrowItem.Publication.PublicationID)
	if err != nil {
		tx.Rollback()
		return
	}
	statement = `UPDATE Card JOIN BorrowerType SET currentBorrowNumber = currentBorrowNumber + 1 WHERE cardNO = ? AND 
				 currentBorrowNumber < maxBorrowNumber`
	_, err = tx.Query(statement)
	if err != nil {
		tx.Rollback()
		err = my_error.MaxBorrowNumberError
		return
	}
	tx.Commit()
	return
}
