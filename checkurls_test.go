package main

import (
	"fmt"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUrl(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockReposController(ctrl)
	setRepoController(m)
	a := "https://testURL.com"
	err, geturl := IsUrl(a)
	if err {
		fmt.Println(err)
	}
	fmt.Println(geturl)
	assert.Equal(t, geturl, a)
}
