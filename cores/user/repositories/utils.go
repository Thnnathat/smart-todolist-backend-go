package repositories

import "strconv"

func gormIdConv(id string) (uint, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	idUint := uint(idInt)
	if err != nil {
		return 0, err
	}
	return idUint, nil
}
