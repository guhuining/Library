package data

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
