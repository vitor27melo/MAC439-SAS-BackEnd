package model

type Oferecimento struct {
	Id_disciplina string `json:"id_disciplina"`
	Sigla         string `json:"sigla"`
	Nome          string `json:"nome"`
	Inicio        string `json:"inicio"`
	Fim           string `json:"fim"`
}
