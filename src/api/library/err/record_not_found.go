
package err

type RecordNotFound struct {
	Op string
	Model string
	Err error
}

func (e *RecordNotFound) Error() string {
	return e.Op + " " + e.Model + ": " + e.Err.Error()
}