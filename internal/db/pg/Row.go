package pg

type Row interface {
	Scan(dest ...interface{}) error
}
