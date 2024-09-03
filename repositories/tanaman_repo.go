package repositories

import "kebunku-api/dto"

func (r *compRepository) RegisterTanaman(data dto.Tanaman) error {
	_, err := r.DB.Exec("INSERT INTO tanaman (name, description, class, plant_order, family) VALUES (?, ?, ?, ?, ?)", data.Name, data.Description, data.Class, data.Order, data.Family)
	if err != nil {
		return err
	}

	return nil
}
