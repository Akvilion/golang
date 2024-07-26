package main

import "errors"

func joinOld(s1, s2 string, max int) (string, error) { // багато вложеностей
	if s1 == "" {
		return "", errors.New("s1 is empty")
	} else {
		if s2 == "" {
			return "", errors.New("s2 is empty")
		} else {
			concat, err := concatenate(s1, s2)
			if err != nil {
				return "", err
			} else {
				if len(concat) > max {
					return concat[:max], nil
				} else {
					return concat, nil
				}
			}
		}
	}
}

func theWay(s string) error {
	// неправильно
	if s != "" { // тут s буде вести по нещасливому шляху
		// ...
	} else {
		return errors.New("empty string")
	}

	// правильно
	if s == "" { // тут s буде вести по щасливому шляху
		return errors.New("empty string")
	}
	// ...
}

func joinNew(s1, s2 string, max int) (string, error) { // причесали вложеності
	if s1 == "" {
		return "", errors.New("s1 is empty")
	}
	if s2 == "" {
		return "", errors.New("s2 is empty")
	}
	concat, err := concatenate(s1, s2)
	if err != nil {
		return "", err
	}
	if len(concat) > max {
		return concat[:max], nil
	}
	return concat, nil
}
