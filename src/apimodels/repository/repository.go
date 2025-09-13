package repository

import (
	"api_model_cnn/src/apimodels/model"
	"database/sql"
	"fmt"
)

type Repository interface {
	GetEmailRepository(int) (string, error)
	GetGizi(label string) (model.RespGizi, error)
	GetGiziDetailRepository(label string) (model.RespGiziDetail, error)
	GetGiziAKgRepository(label string) (model.RespGiziDetail, error)
}

type DbRepository struct {
	db *sql.DB
}

// GetGiziAKgRepository implements Repository.
func (nr *DbRepository) GetGiziAKgRepository(label string) (model.RespGiziDetail, error) {
	query := `SELECT "LABEL", "ENERGI", "LEMAK", "VIT_A", "VIT_B1",
		"VIT_B2", "VIT_B3", "VIT_C", "KARBO", "PROTEIN", "SERAT_PANGAN", 
		"KALSIUM", "FOSFOR", "NATRIUM", "KALIUM", "TEMBAGA", "BESI", "SENG", 
		"B_KAROTEN", "KAROTEN_TOTAL", "AIR", "ABU"
	FROM public."GIZI_AKG"
	WHERE "LABEL" = $1`

	var result model.RespGiziDetail
	err := nr.db.QueryRow(query, label).Scan(
		&result.Label,
		&result.Energi,
		&result.Lemak,
		&result.VitA,
		&result.VitB1,
		&result.VitB2,
		&result.VitB3,
		&result.VitC,
		&result.Karbo,
		&result.Protein,
		&result.SeratPangan,
		&result.Kalsium,
		&result.Fosfor,
		&result.Natrium,
		&result.Kalium,
		&result.Tembaga,
		&result.Besi,
		&result.Seng,
		&result.BKarotene,
		&result.KarotenTotal,
		&result.Air,
		&result.Abu,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.RespGiziDetail{}, fmt.Errorf("tidak ditemukan data gizi AKG dengan label '%s'", label)
		}
		return model.RespGiziDetail{}, fmt.Errorf("error saat mengambil data gizi AKG: %w", err)
	}

	return result, nil
}


// GetGiziDetailReposistory implements Repository.
func (nr *DbRepository) GetGiziDetailRepository(label string) (model.RespGiziDetail, error) {
	query := `SELECT "LABEL", "ENERGI", "LEMAK", "VIT_A", "VIT_B1",
		"VIT_B2", "VIT_B3", "VIT_C", "KARBO", "PROTEIN", "SERAT_PANGAN", 
		"KALSIUM", "FOSFOR", "NATRIUM", "KALIUM", "TEMBAGA", "BESI", "SENG", 
		"B_KAROTEN", "KAROTEN_TOTAL", "AIR", "ABU"
	FROM public."GIZI_DETAIL"
	WHERE "LABEL" = $1`

	var result model.RespGiziDetail
	err := nr.db.QueryRow(query, label).Scan(
		&result.Label,
		&result.Energi,
		&result.Lemak,
		&result.VitA,
		&result.VitB1,
		&result.VitB2,
		&result.VitB3,
		&result.VitC,
		&result.Karbo,
		&result.Protein,
		&result.SeratPangan,
		&result.Kalsium,
		&result.Fosfor,
		&result.Natrium,
		&result.Kalium,
		&result.Tembaga,
		&result.Besi,
		&result.Seng,
		&result.BKarotene,
		&result.KarotenTotal,
		&result.Air,
		&result.Abu,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.RespGiziDetail{}, fmt.Errorf("tidak ditemukan data gizi detail dengan label '%s'", label)
		}
		return model.RespGiziDetail{}, fmt.Errorf("error saat mengambil data gizi detail: %w", err)
	}

	return result, nil
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
