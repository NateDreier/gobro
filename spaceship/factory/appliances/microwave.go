package appliances

type Microwave struct {
	typeName string
}

func (sv *Microwave) Start() {
	sv.typeName = " Stove "
}

func (sv *Microwave) GetPurpose() string {
	return "imma microwave"
}
