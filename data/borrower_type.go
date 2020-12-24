package data

// @title PublicationType.GetPublicationType
// @description 获取所有出版物类型
func (borrowerType *BorrowerType) GetBorrowerType() (results []BorrowerType, err error) {
	statement := `SELECT * FROM BorrowerType`
	rows, err := Db.Query(statement)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var temp BorrowerType
		err = rows.Scan(&temp.BorrowerType, &temp.Period, &temp.MaxBorrowNumber)
		if err != nil {
			return
		}
		results = append(results, temp)
	}
	return
}
