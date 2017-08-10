package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/stretchr/testify/suite"
)

type ResourceClientTestSuite struct {
	suite.Suite
}

func (s *ResourceClientTestSuite) TestClientInit() {
	var c *baremetal.Client
	c = &baremetal.Client{}
	s.Require().True(c != nil) // This won't compile if it's not
}
