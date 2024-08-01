package testing3lamp

type Lamp struct{
	State bool
}

func (l *Lamp) TurnOn(){
	l.State = true
}

func (l *Lamp) TurnOff(){
	l.State = false
}