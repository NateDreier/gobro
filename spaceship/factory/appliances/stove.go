package appliances

type Stove struct {
	typeName string
}

func (sv *Stove) Start() {
	sv.typeName = " Stove "
}

func (sv *Stove) GetPurpose() string {
	return "imma stove"
}
