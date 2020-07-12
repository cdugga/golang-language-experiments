package data

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvalidProductMissingNameAndSKU(t *testing.T){
	p := &Product{
		Description: "",
		Price:       1.0,
	}
	err := p.Validate()
	assert.Len(t, err, 2)
}

func TestInvalidProductMissingName(t *testing.T){
	p := &Product{
		Description: "",
		Price:       1.0,
		SKU:         "abs-abc-def",
	}
	err := p.Validate()
	assert.Len(t, err, 1)
}

func TestValidProduct(t *testing.T){
	p := &Product{
		ID:          0,
		Name:        "Tea",
		Description: "",
		Price:       1.0,
		SKU:         "abs-abc-def",
		CreatedOn:   "",
		UpdatedOn:   "",
		DeletedOn:   "",
	}

	err:= p.Validate()
	assert.NoError(t,err)
}

func TestProducts_ToJSON(t *testing.T) {

	p := productList
	b := bytes.NewBufferString("")

	err := oJSON(p, b)
	assert.NoError(err)
}




