package fake

type Issuer struct {
	Token string
	Err   error
}

func (fi *Issuer) Issue(_ int64) (string, error) {
	return fi.Token, fi.Err
}
