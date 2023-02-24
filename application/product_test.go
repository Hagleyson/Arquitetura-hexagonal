package application_test

import (
	"testing"

	"github.com/hagleyson/Arquitetura-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T){
	product := application.Product{}
	product.Name ="Hello"
	product.Status = application.DISABLED
	product.Price = 10
	err := product.Enable()

	require.Nil(t,err)

	product.Price = 0

	err = product.Enable()

	require.Equal(t,"The price must be greater than zero to enable product",err.Error())
	
}

func TestProduct_Disabled(t *testing.T){
	product := application.Product{}
	product.Name ="Hello"
	product.Status = application.ENABLE
	product.Price = 0

	err := product.Disable()

	require.Nil(t,err)

	product.Price = 10

	err = product.Disable()

	require.Equal(t,"The price must be zero in order to have the product disabled",err.Error())
	
}

func TestProduct_IsValid(t *testing.T){
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "hello"
	product.Status = application.DISABLED
	product.Price = 10

	_,err := product.IsValid()
	require.Nil(t,err)

	product.Status = "INVALID"
	_,err = product.IsValid()

	require.Equal(t,"the status must be enabled or disabled",err.Error())

	product.Status = application.ENABLE
	_,err = product.IsValid()
	
	require.Nil(t,err)

	product.Price = -10
	_,err = product.IsValid()
	require.Equal(t,"the price must be greater or equal zero",err.Error())
}