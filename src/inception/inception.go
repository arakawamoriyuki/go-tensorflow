package inception

func SetUp() error {
	modelExists := modelExists()
	if modelExists == false {
		modelZipExists := modelZipExists()
		if modelZipExists == false {
			err := downloadModelZip()
			if err != nil {
				return err
			}
		}
		err := unzip(modelFile.ZipPath, modelFile.UnzipPath)
		if err != nil {
			return err
		}
	}

	return nil
}
