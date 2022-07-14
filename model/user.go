package model

type User struct {
	Id_usuario    string `json:"id_usuario"`
	Nome          string `json:"nome"`
	Email         string `json:"email"`
	Cpf           string `json:"cpf"`
	Telefone      string `json:"telefone"`
	Token         string `json:"token"`
	Nivel_acesso  string `json:"nivel_acesso"`
	Senha         string `json:"senha"`
	Ultima_Sessao string `json:"ultima_sessao"`
}
