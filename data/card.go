package data

// @title Card.Insert
// @description	插入新借阅证
// @param	CardNO			Card.CardNO					借阅证号码
// @return
func (card *Card) DeleteCard() (err error) {
	statement := `DELETE FROM Card WHERE cardNO = ?`
	_, err = Db.Query(statement, card.CardNO)
	return
}
