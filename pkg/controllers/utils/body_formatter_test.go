package utils

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBodyFormatter_Success(t *testing.T) {
	// Données d'entrée
	body := bytes.NewReader([]byte(`{"champ1":"valeur1","champ2":123}`))
	data := struct {
		Champ1 string `json:"champ1"`
		Champ2 int    `json:"champ2"`
	}{}

	// Appel de la fonction
	err := BodyFormatter(io.NopCloser(body), &data)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, "valeur1", data.Champ1)
	assert.Equal(t, 123, data.Champ2)
}

func TestBodyFormatter_ErrorReadBody(t *testing.T) {
	// Données d'entrée
	body := bytes.NewReader([]byte(`boubobubou{"champ1":"=w=w=w=w=w=2""""""D"D"D"D"valeur1","champ2":`))
	data := struct{}{}

	// Appel de la fonction
	err := BodyFormatter(io.NopCloser(body), &data)

	// Assertions
	assert.Error(t, err)
}

func TestBodyFormatter_ErrorUnmarshal(t *testing.T) {
	// Données d'entrée
	body := bytes.NewReader([]byte(`{"champ1":"valeur1","champ2":`))
	data := struct{}{}

	// Appel de la fonction
	err := BodyFormatter(io.NopCloser(body), &data)

	// Assertions
	assert.Error(t, err)
}
