package series

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestDoIdentityRequests(t *testing.T) {
	// expected return value is interface{}
	expectedResult := []interface{}{"task1", "task2"}

	ctx := context.TODO()
	res, err := Do(
		// TODO: test that functions were called only once
		func() (interface{}, error) { return identifyOK(ctx, "task1") },
		func() (interface{}, error) { return identifyOK(ctx, "task2") },
	)
	if err != nil {
		t.Errorf("series.Do failed. Expected %v, actual: %v", expectedResult, err)
	}

	if !reflect.DeepEqual(expectedResult, res) {
		t.Errorf("series.Do failed. Expected %v, actual: %v", expectedResult, res)
	}
}

func TestDoIdentityErrorRequests1(t *testing.T) {
	// expected return value is interface{}
	var expectedResult interface{}
	expectedError := errors.New("some error here")

	ctx := context.TODO()
	res, err := Do(
		// TODO: test that both function was called
		func() (interface{}, error) { return identifyOK(ctx, "task1") },
		func() (interface{}, error) { return identifyError(ctx, expectedError) },
	)
	if res != nil {
		t.Errorf("series.Do failed. Expected %v, actual: %v", expectedResult, res)
	}

	if err == nil {
		t.Errorf("series.Do failed. Expected %v, actual: %v", expectedError, res)
	}
}

func TestDoIdentityErrorRequests2(t *testing.T) {
	// expected return value is interface{}
	var expectedResult interface{}
	expectedError := errors.New("some error here")

	ctx := context.TODO()
	res, err := Do(
		// TODO: test that second function wasn't called
		func() (interface{}, error) { return identifyError(ctx, expectedError) },
		func() (interface{}, error) { return identifyOK(ctx, "task1") },
	)
	if res != nil {
		t.Errorf("series.Do failed. Expected %v, actual: %v", expectedResult, res)
	}

	if err == nil {
		t.Errorf("series.Do failed. Expected %v, actual: %v", expectedError, res)
	}
}

func TestDoIdentityErrorRequests3(t *testing.T) {
	// expected return value is interface{}
	var expectedResult interface{}
	expectedError := errors.New("some error here")

	ctx := context.TODO()
	res, err := Do(
		// TODO: test that second function wasn't called
		func() (interface{}, error) { return identifyError(ctx, expectedError) },
		func() (interface{}, error) { return identifyError(ctx, expectedError) },
	)
	if res != nil {
		t.Errorf("series.Do failed. Expected %v, actual: %v", expectedResult, res)
	}

	if err == nil {
		t.Errorf("series.Do failed. Expected %v, actual: %v", expectedError, res)
	}
}

func identifyOK(_ context.Context, str string) (string, error) {
	return str, nil
}

func identifyError(_ context.Context, err error) (string, error) {
	return "", err
}
