package observer

// Observer bla bla bla
type Observer interface {
	Update(value string)
}

// Observable bla bla bla
type Observable struct {
	Observers []Observer
}

// AddObserver bla bla bla
func (o *Observable) AddObserver(obs Observer) {
	o.Observers = append(o.Observers, obs)
}

// NotifyAll bla bla bla
func (o *Observable) NotifyAll(value string) {
	for _, ob := range o.Observers {
		ob.Update(value)
	}
}
