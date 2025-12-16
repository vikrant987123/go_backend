package service

import "time"

func CalcAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()

	if( now.YearDay() < dob.YearDay() ){
		age--
	}

	return age
}