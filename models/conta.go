package models

func NewConta(tipo string, descricao string) *Conta {
	return &Conta{Tipo: tipo, Descricao: descricao}
}

type Conta struct {
	Tipo      string
	Descricao string
}

func (c *Conta) GetTipo() string {
	return c.Tipo
}

func (c *Conta) SetTipo(t string) {
	c.Tipo = t
}

func (c *Conta) GetDescricao() string {
	return c.Descricao

}

func (c *Conta) SetDescricao(d string) {
	c.Descricao = d

}
