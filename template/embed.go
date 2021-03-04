package template

import "embed"

//go:embed field.gogo bean.gogo engine.gogo
var fsEmbed embed.FS

func GoGoModelsField() string {
	data, err := fsEmbed.ReadFile("field.gogo")
	if err != nil {
		return ""
	}

	return string(data)
}

func GoGoTypePoint() string {
	data, err := fsEmbed.ReadFile("type.gogo")
	if err != nil {
		return ""
	}

	return string(data)
}

func GoGoEngine() string {
	data, err := fsEmbed.ReadFile("engine.gogo")
	if err != nil {
		return ""
	}

	return string(data)
}

func GoGoBean() string {
	data, err := fsEmbed.ReadFile("bean.gogo")
	if err != nil {
		return ""
	}

	return string(data)
}
