package robot

import (
	"encoding/json"
	"time"

	"github.com/jonathanmoreiraa/2cents/internal/domain/model"
	"github.com/shopspring/decimal"
)

type RawData struct {
	Data  string `json:"data"`
	Valor string `json:"valor"`
}

func ParseAPIData(bytes []byte, investimentTypeID int) ([]model.Metric, error) {
	var rawList RawData
	if err := json.Unmarshal(bytes, &rawList); err != nil {

		return nil, err
	}

	metrics := make([]model.Metric, len([]RawData{rawList}))
	for i, r := range []RawData{rawList} {
		val, err := decimal.NewFromString(r.Valor)
		if err != nil {
			return nil, err
		}

		parsedDate, err := time.Parse("02/01/2006", r.Data)
		if err != nil {
			return nil, err
		}

		parsedDate = parsedDate.Add(time.Hour * 3)

		metrics[i] = model.Metric{
			InvestimentTypeID: investimentTypeID,
			Date:              parsedDate,
			Value:             val,
		}
	}
	return metrics, nil
}
