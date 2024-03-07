package main

func OneFormat(path string, typesFile []ResponseData, toType string) (err error) {
	for _, val := range typesFile {
		if _, err := typer(val.Ct_Label); err != nil {
			continue
		}
		info, err := GetInfo(val.Path)
		if err != nil {
			return err
		}
		name, isDir := GetName(info)
		if isDir {
			continue
		}
		if err := ConverteFile(name, path, val, toType); err != nil {
			return err
		}

	}
	return nil
}

func MultipleFormat(path string, typesFile []ResponseData, toType []string) error {
	if err := CreateDirNeeded(path, toType); err != nil {
		return err
	}
	for _, val := range typesFile {
		if _, err := typer(val.Ct_Label); err != nil {
			continue
		}
		info, err := GetInfo(val.Path)
		if err != nil {
			return err
		}
		name, isDir := GetName(info)
		if isDir {
			continue
		}
		for _, tt := range toType {
			if err := ConverteFile(name, path+tt+"/", val, tt); err != nil {
				return err
			}
		}
	}
	return nil
}
