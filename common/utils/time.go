package utils

import "time"

func GetDaysAfter(startDay, today string) ([]string, error) {
	t1, err := time.ParseInLocation("20060102", startDay, time.Local)
	if err != nil {
		return nil, err
	}

	t2, err := time.ParseInLocation("20060102", today, time.Local)
	if err != nil {
		return nil, err
	}

	diffSeconds := t2.Unix() - t1.Unix()
	if diffSeconds == 0 {
		return nil, nil
	}

	diffDays := diffSeconds / 3600 / 24

	var ret []string
	for i := 1; i <= int(diffDays); i++ {
		ret = append(ret, t1.AddDate(0, 0, i).Format("20060102"))
	}
	return ret, nil
}

func DayTimeAddDays(dayTM string, days int) (string, error) {
	t1, err := time.ParseInLocation("20060102", dayTM, time.Local)
	if err != nil {
		return "", err
	}
	t2 := t1.AddDate(0, 0, days)
	return t2.Format("20060102"), nil
}
