package schedule

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Clankyyy/scheduler-bot/internal/entity"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{
		Timeout: 10 * time.Second,
	}
}

func GetGroups() ([]entity.GroupDataReq, error) {
	type wrapper struct {
		result []entity.GroupDataReq
		err    error
	}
	ch := make(chan wrapper, 1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1*time.Second))
	defer cancel()

	go func() {
		result, err := requestGroupsWithContext(ctx)
		ch <- wrapper{result, err}
	}()

	select {
	case data := <-ch:
		if data.err != nil {
			return nil, data.err
		}
		return data.result, nil
	case <-ctx.Done():
		return nil, fmt.Errorf("timeout exceeded")
	}
}

func requestGroupsWithContext(ctx context.Context) ([]entity.GroupDataReq, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), time.Duration(time.Second))
		defer cancel()
	}

	res, err := requestWithContext(ctx, http.MethodGet, "http://localhost:8000/groups/", nil)

	if err != nil {
		return nil, err
	}
	var data []entity.GroupDataReq
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func requestWithContext(ctx context.Context, method string, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
