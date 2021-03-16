package strategies

import (
	"bitbucket.org/crashtest-security/google-hash-code-2020-team-golang/model"
)

type StrategyInterface interface {
	Run(dataset model.Dataset) model.Outputset
	GetName() string
}
