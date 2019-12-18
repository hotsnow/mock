package api

type Action interface {
	Double(i *int) error
}

type NUM struct{}

func (n NUM) Double(i *int) error {
	*i *= 2

	return nil
}
