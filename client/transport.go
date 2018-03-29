package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	consulsd "github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	httptransport "github.com/go-kit/kit/transport/http"
	consulapi "github.com/hashicorp/consul/api"
)

func New() (DiffCalculator, error) {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	consulClient, err := consulapi.NewClient(consulapi.DefaultConfig())
	if err != nil {
		return nil, err
	}
	client := consulsd.NewClient(consulClient)

	retryMax := 3
	retryTimeout := time.Second * 5

	var instancer = consulsd.NewInstancer(client, logger, "diffcalculator", []string{}, true)

	d := &diffCalculator{}

	factory := makeFactory("/DiffCalculator/Calculate", encodeRequest, decodeCalculateResponse)
	endpointer := sd.NewEndpointer(instancer, factory, logger)
	balancer := lb.NewRoundRobin(endpointer)
	retry := lb.Retry(retryMax, retryTimeout, balancer)
	d.calculateEndpoint = retry

	return d, nil
}

type diffCalculator struct {
	calculateEndpoint endpoint.Endpoint
}

func (d *diffCalculator) Calculate(ctx context.Context, sourceID string, items []Item) error {
	request := calculateRequest{SourceID: sourceID, Items: items}
	response, err := d.calculateEndpoint(ctx, request)
	if err != nil {
		return err
	}
	_ = response.(calculateResponse)
	return nil
}

func makeFactory(path string, enc httptransport.EncodeRequestFunc, dec httptransport.DecodeResponseFunc) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		if !strings.HasPrefix(instance, "http") {
			instance = "https://" + instance
		}
		tgt, err := url.Parse("https://proxy.mmaks.me")
		if err != nil {
			return nil, nil, err
		}
		tgt.Path = path
		//log.Printf("%s", tgt.String())

		return httptransport.NewClient("POST", tgt, enc, dec).Endpoint(), nil, nil
	}
}

type calculateRequest struct {
	SourceID string
	Items    []Item
}

type calculateResponse struct {
}

func decodeCalculateResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response calculateResponse
	//err := json.NewDecoder(resp.Body).Decode(&response)
	if resp.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("bad http status: %d [%s]", resp.StatusCode, b)
	}
	return response, nil
}

func encodeRequest(_ context.Context, req *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(&buf)

	return nil
}
