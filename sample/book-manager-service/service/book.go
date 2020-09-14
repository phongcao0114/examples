package service

func (b BookManagerServiceImpl) GetBookList() (interface{}, error) {
	return b.BookDBImpl.GetBookList()
}
