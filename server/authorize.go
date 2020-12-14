package server

import "net/http"

// @title	authorizeAdministrator
// @description	系统管理员鉴权
// @param	r	*http.Request
// @return	ok	bool			是否有此权限
// @return	err	error			错误信息
func authorizeAdministrator(r *http.Request) (ok bool, err error) {
	session, err := store.Get(r, "library")
	if err != nil {
		return
	}
	ok = session.Values["Roll"] == "Administrator"
	return
}

// @title	authorizeAdministrator
// @description	系统管理员鉴权
// @param	r	*http.Request
// @return	ok	bool			是否有此权限
// @return	err	error			错误信息
func authorizeLibrarian(r *http.Request) (ok bool, err error) {
	session, err := store.Get(r, "library")
	if err != nil {
		return
	}
	ok = session.Values["Roll"] == "Librarian"
	return
}
