package schedule

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/Clankyyy/scheduler-bot/internal/entity"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{
		Timeout: 10 * time.Second,
	}
}

func GetDaily(slug string, day string, kind string) (entity.Daily, error) {
	type wrapper struct {
		result entity.Daily
		err    error
	}

	ch := make(chan wrapper, 1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1*time.Second))
	defer cancel()

	go func() {
		result, err := requestDailyWithContext(ctx, slug, day, kind)
		ch <- wrapper{result, err}
	}()

	select {
	case data := <-ch:
		if data.err != nil {
			return entity.Daily{}, data.err
		}
		return data.result, nil
	case <-ctx.Done():
		return entity.Daily{}, fmt.Errorf("timeout exceeded")
	}
}

func GetGroups() (entity.GroupsRes, error) {
	type wrapper struct {
		data entity.GroupsRes
		err  error
	}

	ch := make(chan wrapper, 1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2*time.Second))
	defer cancel()

	go func() {
		result, err := requestGroupsWithContext(ctx)
		ch <- wrapper{result, err}
	}()

	select {
	case result := <-ch:
		if result.err != nil {
			return result.data, result.err
		}
		return result.data, nil
	case <-ctx.Done():
		return entity.GroupsRes{}, errors.New("timeout exceeded")
	}
}

func requestDailyWithContext(ctx context.Context, slug string, day string, kind string) (entity.Daily, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), time.Duration(time.Second))
		defer cancel()
	}

	q := url.Values{}
	q.Add("day", day)
	q.Add("type", kind)
	url := "http://localhost:8000/schedule/daily/" + slug + "?"
	res, err := requestWithContext(ctx, http.MethodGet, url, nil, q)

	var data entity.Daily
	if err != nil {
		return data, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return data, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func requestGroupsWithContext(ctx context.Context) (entity.GroupsRes, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), time.Duration(time.Second)*2)
		defer cancel()
	}

	res, err := requestWithContext(ctx, http.MethodGet, "http://localhost:8000/groups/", nil, nil)

	var data entity.GroupsRes
	if err != nil {
		return data, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return data, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}

	err = json.NewDecoder(res.Body).Decode(&data.Data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func requestWithContext(ctx context.Context, method string, url string, body io.Reader, q url.Values) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = q.Encode()
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
