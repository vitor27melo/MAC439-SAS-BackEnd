package model

type Oferecimento struct {
	Id_disciplina string `json:"id_disciplina"`
	Sigla         string `json:"sigla"`
	DiaSemana     string `json:dia_semana`
	Nome          string `json:"nome"`
	Inicio        string `json:"inicio"`
	Fim           string `json:"fim"`
}
