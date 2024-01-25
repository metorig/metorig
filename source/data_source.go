package source

import "github.com/metorig/metorig/model"

type DataSource interface {
	Store(metrics *model.Metrics) error
}
