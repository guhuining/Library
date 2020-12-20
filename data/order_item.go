package data

func (orderItem *OrderItem) Insert() (err error) {
	statement := `INSERT OrderItem (publicationID, cardNO, orderDate) VALUES (?, ?, NOW())`
	_, err = Db.Query(statement, orderItem.Publication.PublicationID, orderItem.Card.CardNO)
	return
}
