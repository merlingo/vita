package types

import (
	math "math"
	"strconv"
	"strings"

	labourtypes "vita/x/labour/types"

	types "github.com/cosmos/cosmos-sdk/types"
)

func (mechanism *Mechanism) Validate() error {
	//coefficients parsing error, metrics parsing error, slope is not double error, converge limit should be between 0 and 1.
	coefficients := strings.Split(mechanism.Coefficients, ",")
	metrics := strings.Split(mechanism.Metrics, ",")
	if coefficients[0] == mechanism.Coefficients {
		return ErrCoefficientParsing
	}
	if metrics[0] == mechanism.Metrics {
		return ErrMetricsParsing
	}
	if len(metrics) != len(coefficients) {
		return ErrMetricAndCoefficientNotCompatible
	}
	cl, err := strconv.ParseFloat(mechanism.ConvergeLimit, 64)
	if err != nil {
		return ErrConvergeLimitParsing
	}
	if cl > 1 || cl < 0 {
		return ErrOutOfConvergeLimit
	}
	_, err = strconv.ParseFloat(mechanism.Slope, 64)
	if err != nil {
		return ErrSlopeParsing
	}
	return nil
}

func (mechanism Mechanism) CalculateRewards(wager types.Coin, activities []labourtypes.Activity, task labourtypes.Task) types.Coin {
	//parse metrics and check whether activity, task timespan are in the list. use coefficients to calculate the reward
	var total_ratio float64 = 0
	if task.Deadline > task.FinishTask {
		total_ratio = total_ratio + 1
	}
	coefficients := strings.Split(mechanism.Coefficients, ",")
	metrics := strings.Split(mechanism.Metrics, ",")
	index := indexOf("activity", metrics)
	converge_limit, _ := strconv.ParseFloat(mechanism.ConvergeLimit, 64)
	slope, _ := strconv.ParseFloat(mechanism.Slope, 64)

	if index > -1 {
		a_coefficient, _ := strconv.ParseFloat(coefficients[index], 64)
		for _, activity := range activities {
			total_hour := (activity.FinishWork - activity.BeginWork) / 3600000
			if float64(total_hour-uint64(activity.WorkingTime)) < float64(total_hour)*converge_limit {
				//yaklasim limitinin altinda bir çalisma gerçekleştirildiği icin bu activitenin puanı hesaplamadan cikarilacak.
				a_coefficient = a_coefficient - a_coefficient/float64(len(activities))
			}
		}
		total_ratio = total_ratio + a_coefficient
	}
	index = indexOf("tasktime", metrics)
	if index > -1 {
		t_coefficient, _ := strconv.ParseFloat(coefficients[index], 64)
		total_hour := math.Abs(float64(task.FinishTask-task.BeginTask)) / 3600000
		differ_hour := math.Abs(float64(task.Deadline-task.FinishTask)) / 3600000

		if differ_hour > total_hour*converge_limit && differ_hour <= total_hour {
			t_coefficient = slope * ((total_hour - differ_hour) / total_hour) * t_coefficient
		} else if differ_hour > total_hour {
			t_coefficient = 0
		}
		//t_coefficient =
		total_ratio = total_ratio + t_coefficient
	}

	return types.NewInt64Coin(wager.Denom, int64(float64(wager.Amount.Int64())*total_ratio))
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
