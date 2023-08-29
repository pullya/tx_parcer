package repository_test

import (
	"testing"

	"github.com/pullya/tx_parcer/internal/app/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RepositoryTestSuite struct {
	suite.Suite
	repo *repository.Repository
}

func (s *RepositoryTestSuite) SetupSuite() {
	s.repo = repository.NewRepository()
}

func (s *RepositoryTestSuite) TearDownSuite() {}

func (s *RepositoryTestSuite) TestAddSubscription() {
	res1, err1 := s.repo.AddSubscription("teststring1")
	assert.NoError(s.T(), err1)
	assert.EqualValues(s.T(), 1, res1)

	_, _ = s.repo.AddSubscription("teststring2")
	_, _ = s.repo.AddSubscription("teststring3")

	assert.EqualValues(s.T(), 3, s.repo.LastId)
	assert.EqualValues(s.T(), "teststring2", s.repo.Subscriptions["teststring2"].Address)

	delete(s.repo.Subscriptions, "teststring1")
	delete(s.repo.Subscriptions, "teststring2")
	delete(s.repo.Subscriptions, "teststring3")
	s.repo.LastId = 0
}

func (s *RepositoryTestSuite) TestGetSubsctiptionByAddress() {
	_, err := s.repo.AddSubscription("teststring4")
	assert.NoError(s.T(), err)
	_, _ = s.repo.AddSubscription("teststring5")
	_, _ = s.repo.AddSubscription("teststring6")

	_, err1 := s.repo.GetSubsctiptionByAddress("teststring100")
	assert.Error(s.T(), err1)

	id, err2 := s.repo.GetSubsctiptionByAddress("teststring5")
	assert.NoError(s.T(), err2)
	assert.EqualValues(s.T(), 2, id)

	delete(s.repo.Subscriptions, "teststring4")
	delete(s.repo.Subscriptions, "teststring5")
	delete(s.repo.Subscriptions, "teststring6")
	s.repo.LastId = 0
}

func (s *RepositoryTestSuite) TestDeleteSubscriptionByAddress() {
	_, err := s.repo.AddSubscription("teststring7")
	assert.NoError(s.T(), err)
	_, _ = s.repo.AddSubscription("teststring8")
	_, _ = s.repo.AddSubscription("teststring9")

	err1 := s.repo.DeleteSubscriptionByAddress("teststring900")
	assert.Error(s.T(), err1)

	err2 := s.repo.DeleteSubscriptionByAddress("teststring9")
	assert.NoError(s.T(), err2)

	_, err3 := s.repo.GetSubsctiptionByAddress("teststring9")
	assert.Error(s.T(), err3)

	delete(s.repo.Subscriptions, "teststring7")
	delete(s.repo.Subscriptions, "teststring8")
	delete(s.repo.Subscriptions, "teststring9")
	s.repo.LastId = 0
}

func TestRepo(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}
