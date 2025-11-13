package lapsus

func (suite *ErrorsSuite) TestWithLabels() {
	var fac = NewBadRequestFactory("error")
	var newFactory = fac.WithLabels("TagFallbackable", "TagProcessing")
	var err = newFactory.New()
	suite.Require().NotNil(err)
	var typedErr, ok = err.(*implementation)
	suite.Require().True(ok)
	suite.Require().NotNil(typedErr)
	suite.Require().ElementsMatch(typedErr.labels, []Label{LabelUserFriendly, "TagFallbackable", "TagProcessing"})
}
