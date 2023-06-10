package cashDispenser


func NewDispenser() Dispenser{
	return &dispenser{}
}

type Dispenser interface{
	AddCash(int) error
	DispenseCash() error
	GetAmount() (int, error)
}

type dispenser struct{
	
}

func (d *dispenser) AddCash(i int) error {
	//TODO implement me
	panic("implement me")
}

func (d *dispenser) DispenseCash() error {
	//TODO implement me
	panic("implement me")
}

func (d *dispenser) GetAmount() (int, error) {
	//TODO implement me
	panic("implement me")
}
