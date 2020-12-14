package data

// @title	PublicationType.Insert
// @description	新增出版物类型
// @param	publicationType	PublicationType.PublicationType	出版物类型
// @param	fine			PublicationType.Fine			超期罚金
// @return	err				error							错误信息
func (publication *Publication) Insert() (err error) {
	statement := `INSERT INTO Publication (name, ISBN, price, total, inventory, publicationType, author) 
				  VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err = Db.Query(statement, publication.Name, publication.ISBN, publication.Price, publication.Total,
		publication.Inventory, publication.PublicationType.PublicationType, publication.Author)
	return
}

// @title	PublicationType.Delete
// @description	新增出版物类型
// @param	publicationID	Publication.PublicationID	出版物ID
// @return	err				error						错误信息
func (publication *Publication) Delete() (err error) {
	statement := `DELETE FROM Publication WHERE publicationID = ?`
	_, err = Db.Query(statement, publication.PublicationID)
	return
}
