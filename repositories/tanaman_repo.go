package repositories

import "kebunku-api/dto"

func (r *compRepository) RegisterTanaman(data dto.Tanaman) error {
	_, err := r.DB.Exec("INSERT INTO tanaman (name, description, class, plant_order, family) VALUES (?, ?, ?, ?, ?)", data.Name, data.Description, data.Class, data.Order, data.Family)
	if err != nil {
		return err
	}

	return nil
}

func (r *compRepository) GetTanaman() ([]dto.Tanaman, error) {
	rows, err := r.DB.Query("SELECT id, name, description, class, plant_order, family FROM tanaman")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []dto.Tanaman
	for rows.Next() {
		var t dto.Tanaman
		err := rows.Scan(&t.Id, &t.Name, &t.Description, &t.Class, &t.Order, &t.Family)
		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}
