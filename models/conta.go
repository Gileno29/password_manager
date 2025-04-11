package models

func NewConta(tipo string, descricao string) *conta {
	return &conta{Tipo: tipo, Descricao: descricao}
}

type conta struct {
	Tipo      string
	Descricao string
}

func (c *conta) GetTipo() string {
	return c.Tipo
}

func (c *conta) SetTipo(t string) {
	c.Tipo = t
}

func (c *conta) GetDescricao() string {
	return c.Descricao

}

func (c *conta) SetDescricao(d string) {
	c.Descricao = d

}
