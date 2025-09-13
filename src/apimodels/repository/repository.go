package repository

import (
	"api_model_cnn/src/apimodels/model"
	"database/sql"
	"fmt"
)

type Repository interface {
	GetEmailRepository(int) (string, error)
	GetGizi(label string) (model.RespGizi, error)
}

type DbRepository struct {
	db *sql.DB
}

// GetGizi implements Repository.
func (nr *DbRepository) GetGizi(label string) (model.RespGizi, error) {
	query := `SELECT "LABEL", "ENERGI", "PROTEIN", "LEMAK", "KARBO", "GIZI_UNGGULAN_1",
	"GIZI_UNGGULAN_2", "GIZI_UNGGULAN_3"
	FROM public."GIZI"
	WHERE "LABEL" = $1`

	var result model.RespGizi
	err := nr.db.QueryRow(query, label).Scan(
		&result.Label,
		&result.Energi,
		&result.Protein,
		&result.Lemak,
		&result.Karbo,
		&result.GiziUnggulan1,
		&result.GiziUnggulan2,
		&result.GiziUnggulan3,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.RespGizi{}, fmt.Errorf("tidak ditemukan data gizi dengan label '%s'", label)
		}
		return model.RespGizi{}, fmt.Errorf("error saat mengambil data gizi: %w", err)
	}

	return result, nil
}


// GetEmail implements Repository.
func (nr *DbRepository) GetEmailRepository(id int) (string, error) {
	query := `select "EMAIL" FROM public."USER_ACCESS" WHERE "USER_ID" = $1`

	var result string
	err := nr.db.QueryRow(query, id).Scan(&result)
	if err != nil {
		return "", err
	}

	return result, nil
}

func CreateRepository(db *sql.DB) Repository {
	return &DbRepository{db: db}
}
