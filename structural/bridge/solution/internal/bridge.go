package internal

type Data struct{}

func ParseAndSaveData(parser DataParser, storage DataPersistent) error {
	data, err := parser.Parse()
	if err != nil {
		return err
	}

	if err := storage.Save(data); err != nil {
		return err
	}

	return nil
}
