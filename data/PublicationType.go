package data

// @title	PublicationType.Insert
// @description	新增出版物类型
// @param	publicationType	PublicationType.PublicationType	出版物类型
// @param	fine			PublicationType.Fine			超期罚金
// @return	err				error							错误信息
func (publicationType *PublicationType) Insert() (err error) {
	statement := `INSERT INTO PublicationType (PublicationType, fine) VALUES(?, ?)`
	_, err = Db.Query(statement, publicationType.PublicationType, publicationType.Fine)
	return
}
