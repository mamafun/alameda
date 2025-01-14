package score

import (
	DaoScore "github.com/containers-ai/alameda/datahub/pkg/dao/score"
	EntityInfluxScore "github.com/containers-ai/alameda/datahub/pkg/entity/influxdb/score"
	RepoInflux "github.com/containers-ai/alameda/datahub/pkg/repository/influxdb"
	DBCommon "github.com/containers-ai/alameda/internal/pkg/database/common"
	InternalInflux "github.com/containers-ai/alameda/internal/pkg/database/influxdb"
	InfluxClient "github.com/influxdata/influxdb/client/v2"
	"github.com/pkg/errors"
)

// SimulatedSchedulingScoreRepository Repository of simulated_scheduling_score data
type SimulatedSchedulingScoreRepository struct {
	influxDB *InternalInflux.InfluxClient
}

// NewRepositoryWithConfig New SimulatedSchedulingScoreRepository with influxdb configuration
func NewRepositoryWithConfig(cfg InternalInflux.Config) SimulatedSchedulingScoreRepository {
	return SimulatedSchedulingScoreRepository{
		influxDB: InternalInflux.NewClient(&cfg),
	}
}

// ListScoresByRequest List scores from influxDB
func (r SimulatedSchedulingScoreRepository) ListScoresByRequest(request DaoScore.ListRequest) ([]*EntityInfluxScore.SimulatedSchedulingScoreEntity, error) {

	var (
		err error

		results      []InfluxClient.Result
		influxdbRows []*InternalInflux.InfluxRow
		scores       = make([]*EntityInfluxScore.SimulatedSchedulingScoreEntity, 0)
	)

	queryCondition := DBCommon.QueryCondition{
		StartTime:      request.QueryCondition.StartTime,
		EndTime:        request.QueryCondition.EndTime,
		StepTime:       request.QueryCondition.StepTime,
		TimestampOrder: request.QueryCondition.TimestampOrder,
		Limit:          request.QueryCondition.Limit,
	}

	influxdbStatement := InternalInflux.Statement{
		QueryCondition: &queryCondition,
		Measurement:    SimulatedSchedulingScore,
	}

	influxdbStatement.AppendWhereClauseFromTimeCondition()
	influxdbStatement.SetOrderClauseFromQueryCondition()
	influxdbStatement.SetLimitClauseFromQueryCondition()
	cmd := influxdbStatement.BuildQueryCmd()

	results, err = r.influxDB.QueryDB(cmd, string(RepoInflux.Score))
	if err != nil {
		return scores, errors.Wrap(err, "list scores failed")
	}

	influxdbRows = InternalInflux.PackMap(results)
	for _, influxdbRow := range influxdbRows {
		for _, data := range influxdbRow.Data {
			scoreEntity := EntityInfluxScore.NewSimulatedSchedulingScoreEntityFromMap(data)
			scores = append(scores, &scoreEntity)
		}
	}

	return scores, nil

}

// CreateScores Create simulated_scheduling_score data points into influxdb
func (r SimulatedSchedulingScoreRepository) CreateScores(scores []*DaoScore.SimulatedSchedulingScore) error {

	var (
		err error

		points = make([]*InfluxClient.Point, 0)
	)

	for _, score := range scores {

		time := score.Timestamp
		scoreBefore := score.ScoreBefore
		scoreAfter := score.ScoreAfter
		entity := EntityInfluxScore.SimulatedSchedulingScoreEntity{
			Time:        time,
			ScoreBefore: &scoreBefore,
			ScoreAfter:  &scoreAfter,
		}

		point, err := entity.InfluxDBPoint(string(SimulatedSchedulingScore))
		if err != nil {
			return errors.Wrap(err, "create scores failed")
		}
		points = append(points, point)
	}

	err = r.influxDB.WritePoints(points, InfluxClient.BatchPointsConfig{
		Database: string(RepoInflux.Score),
	})
	if err != nil {
		return errors.Wrap(err, "create scores failed")
	}

	return nil
}
