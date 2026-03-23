package logic

// Repository mendefinisikan kontrak data
type Repository interface {
	GetUserEmail(id int) (string, error)
}
