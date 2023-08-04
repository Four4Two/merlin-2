package cli_test

import (
	gocontext "context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/four4two/merlin/v16/app/apptesting"
	"github.com/four4two/merlin/v16/x/pool-incentives/types"
)

type QueryTestSuite struct {
	apptesting.KeeperTestHelper
	queryClient types.QueryClient
}

func (s *QueryTestSuite) SetupSuite() {
	s.Setup()
	s.queryClient = types.NewQueryClient(s.QueryHelper)

	// set up pool
	s.PrepareBalancerPool()
	s.Commit()
}

func (s *QueryTestSuite) TestQueriesNeverAlterState() {
	testCases := []struct {
		name   string
		query  string
		input  interface{}
		output interface{}
	}{
		{
			"Query distribution info",
			"/merlin.poolincentives.v1beta1.Query/DistrInfo",
			&types.QueryDistrInfoRequest{},
			&types.QueryDistrInfoResponse{},
		},
		{
			"Query external incentive gauges",
			"/merlin.poolincentives.v1beta1.Query/ExternalIncentiveGauges",
			&types.QueryExternalIncentiveGaugesRequest{},
			&types.QueryExternalIncentiveGaugesResponse{},
		},
		{
			"Query all gauge ids",
			"/merlin.poolincentives.v1beta1.Query/GaugeIds",
			&types.QueryGaugeIdsRequest{PoolId: 1},
			&types.QueryGaugeIdsResponse{},
		},
		{
			"Query all incentivized pools",
			"/merlin.poolincentives.v1beta1.Query/IncentivizedPools",
			&types.QueryIncentivizedPoolsRequest{},
			&types.QueryIncentivizedPoolsResponse{},
		},
		{
			"Query lockable durations",
			"/merlin.poolincentives.v1beta1.Query/LockableDurations",
			&types.QueryLockableDurationsRequest{},
			&types.QueryLockableDurationsResponse{},
		},
		{
			"Query params",
			"/merlin.poolincentives.v1beta1.Query/Params",
			&types.QueryParamsRequest{},
			&types.QueryParamsResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			s.SetupSuite()
			err := s.QueryHelper.Invoke(gocontext.Background(), tc.query, tc.input, tc.output)
			s.Require().NoError(err)
			s.StateNotAltered()
		})
	}
}

func TestQueryTestSuite(t *testing.T) {
	suite.Run(t, new(QueryTestSuite))
}
