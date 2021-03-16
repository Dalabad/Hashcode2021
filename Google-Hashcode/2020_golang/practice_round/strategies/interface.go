package strategies

import (
	"bitbucket.org/crashtest-security/google-hash-code-2020-team-golang/practice_round/model"
)

type StrategyInterface interface {
	Run(dataset model.Dataset) model.Dataset
	GetName() string
}
