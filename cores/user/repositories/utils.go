package repositories

import "strconv"

func gormIdConv(id string) (uint, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	idUint := uint(idInt)

	return idUint, nil
}
