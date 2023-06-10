package oyun

type ISavaci interface{
	AtesEt()
	HasarAl()
}
type Savasci struct {
	Name string
	Can int
	MermiSayisi int
}

type Dunyali struct{
	Savasci
	Birlik string
}

func(d *Dunyali) AtesEt(){
	if d.MermiSayisi > 1{
		d.MermiSayisi --
	}
}

func (d *Dunyali) HasarAl(){
	if d.Can > 0 {
		d.Can -= 10
	}
}

type Marsli struct{
	Savasci
	UzayGemisi string
}

func (m *Marsli) AtesEt(){
	if m.MermiSayisi > 1 {
		m.MermiSayisi--
	}
}

func (m *Marsli) HasarAl(){
	if m.Can > 0 {
		m.Can -= 10
	}
}

