package my_error

import "errors"

var InventoryNotEnoughError = errors.New("out of inventory")
var MaxBorrowNumberError = errors.New("borrow number has reached the max")
var BorrowOutOfTimeError = errors.New("there are publications having not been returned yet")
