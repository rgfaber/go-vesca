package contract

type IQry interface {
	TraceId()
}

func ImplementsIQry(qry IQry) bool {
	return true
}
