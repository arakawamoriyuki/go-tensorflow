package inception

func SetUp() error {
	modelExists := modelExists()
	if modelExists == false {
		modelZipExists := modelZipExists()
		if modelZipExists == false {
			if err := downloadModelZip(); err != nil {
				return err
			}
		}
		if err := unzip(modelFile.ZipFilePath, modelFile.UnzipDestPath); err != nil {
			return err
		}
	}

	if err := loadLabelFile(); err != nil {
		return err
	}

	if err := startSession(); err != nil {
		return err
	}

	return nil
}


