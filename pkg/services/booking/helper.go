package booking

import (
	"strconv"
	"strings"
	"thegrace/pkg"
	"thegrace/pkg/helper"
)

func decryptBookingData(bookingData string) (bookingsData, error) {
	var dat bookingsData
	decryptedData, err := helper.DecryptData(bookingData, pkg.Conf.SigningKeyEncrypt)
	if err != nil {
		return bookingsData{}, err
	}
	split := strings.Split(decryptedData, ",")
	for i := range split {
		st := strings.Split(split[i], ":")
		if i == 0 {
			id, err := strconv.Atoi(st[1])
			if err != nil {
				return bookingsData{}, err
			}
			dat.Id = id
		} else if i == 1 {
			dat.Name = st[1]
		} else if i == 4 {
			ibadahId, err := strconv.Atoi(st[1])
			if err != nil {
				return bookingsData{}, err
			}
			dat.IbadahId = ibadahId
		}
	}
	return dat, nil
}
