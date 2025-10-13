package robot

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/jonathanmoreiraa/2cents/internal/domain/repository"
	"github.com/jonathanmoreiraa/2cents/pkg/log"
)

func FetchData(urlTemplate string) ([]byte, error) {
	today := time.Now()
	startDate := today.AddDate(0, -6, 0)

	url := strings.ReplaceAll(urlTemplate, "${startDate}", startDate.Format("02/01/2006"))
	url = strings.ReplaceAll(url, "${today}", today.Format("02/01/2006"))

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var dataList []RawData
	err = json.Unmarshal(responseData, &dataList)
	if err != nil {
		log.NewLogger().Error(err)
		return nil, err
	}

	if len(dataList) == 0 {
		return nil, nil
	}

	last, err := json.Marshal(&dataList[len(dataList)-1])
	if err != nil {
		log.NewLogger().Error(err)
		return nil, err
	}

	return last, nil
}

func RunRobot(repo repository.MetricRepository, investimentTypeID int, url string) error {
	dataBytes, err := FetchData(url)
	if err != nil {
		return fmt.Errorf("erro ao buscar dados: %w", err)
	}

	metrics, err := ParseAPIData(dataBytes, investimentTypeID)
	if err != nil {
		return fmt.Errorf("erro ao parsear dados: %w", err)
	}

	lastMetric, err := repo.GetLastMetric(context.Background(), investimentTypeID)
	if err != nil {
		return fmt.Errorf("erro ao obter última métrica: %w", err)
	}

	for _, m := range metrics {
		if !m.Date.After(lastMetric.Date) {
			continue
		}

		_, err := repo.Create(context.Background(), m)
		if err != nil {
			log.NewLogger().Error(err)
			return err
		}
	}

	return nil
}
