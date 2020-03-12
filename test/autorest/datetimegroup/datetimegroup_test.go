// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package datetimegroup

import (
	"context"
	"generatortests/autorest/generated/datetimegroup"
	"generatortests/helpers"
	"net/http"
	"testing"
	"time"
)

func getDatetimeOperations(t *testing.T) datetimegroup.DatetimeOperations {
	client, err := datetimegroup.NewClient("http://localhost:3000", nil)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	return client.DatetimeOperations()
}

func TestGetInvalid(t *testing.T) {
	client := getDatetimeOperations(t)
	_, err := client.GetInvalid(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
}

func TestGetLocalNegativeOffsetLowercaseMaxDateTime(t *testing.T) {
	client := getDatetimeOperations(t)
	result, err := client.GetLocalNegativeOffsetLowercaseMaxDateTime(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	expected, err := time.Parse(time.RFC3339, "9999-12-31T23:59:59.999-14:00")
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Value, &expected)
}

func TestGetLocalNegativeOffsetMinDateTime(t *testing.T) {
	client := getDatetimeOperations(t)
	result, err := client.GetLocalNegativeOffsetMinDateTime(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	expected, err := time.Parse(time.RFC3339, "0001-01-01T00:00:00-14:00")
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Value, &expected)
}

func TestGetLocalNegativeOffsetUppercaseMaxDateTime(t *testing.T) {
	client := getDatetimeOperations(t)
	result, err := client.GetLocalNegativeOffsetUppercaseMaxDateTime(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	expected, err := time.Parse(time.RFC3339, "9999-12-31T23:59:59.999-14:00")
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Value, &expected)
}

func TestGetLocalPositiveOffsetLowercaseMaxDateTime(t *testing.T) {
	client := getDatetimeOperations(t)
	result, err := client.GetLocalPositiveOffsetLowercaseMaxDateTime(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	expected, err := time.Parse(time.RFC3339, "9999-12-31T23:59:59.999+14:00")
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Value, &expected)
}

func TestGetLocalPositiveOffsetMinDateTime(t *testing.T) {
	client := getDatetimeOperations(t)
	result, err := client.GetLocalPositiveOffsetMinDateTime(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	expected, err := time.Parse(time.RFC3339, "0001-01-01T00:00:00+14:00")
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Value, &expected)
}

func TestGetLocalPositiveOffsetUppercaseMaxDateTime(t *testing.T) {
	client := getDatetimeOperations(t)
	result, err := client.GetLocalPositiveOffsetUppercaseMaxDateTime(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	expected, err := time.Parse(time.RFC3339, "9999-12-31T23:59:59.999+14:00")
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Value, &expected)
}

func TestGetNull(t *testing.T) {
	client := getDatetimeOperations(t)
	result, err := client.GetNull(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	var expected *time.Time
	helpers.DeepEqualOrFatal(t, result.Value, expected)
}

func TestGetOverflow(t *testing.T) {
	t.Skip("API doesn't actually overflow")
	client := getDatetimeOperations(t)
	_, err := client.GetOverflow(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
}

func TestGetUnderflow(t *testing.T) {
	client := getDatetimeOperations(t)
	_, err := client.GetUnderflow(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
}

func TestGetUTCLowercaseMaxDateTime(t *testing.T) {
	client := getDatetimeOperations(t)
	result, err := client.GetUTCLowercaseMaxDateTime(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	expected, err := time.Parse(time.RFC3339, "9999-12-31T23:59:59.999Z")
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Value, &expected)
}

func TestGetUTCMinDateTime(t *testing.T) {
	client := getDatetimeOperations(t)
	result, err := client.GetUTCMinDateTime(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	expected, err := time.Parse(time.RFC3339, "0001-01-01T00:00:00Z")
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Value, &expected)
}

func TestGetUTCUppercaseMaxDateTime(t *testing.T) {
	client := getDatetimeOperations(t)
	result, err := client.GetUTCUppercaseMaxDateTime(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	expected, err := time.Parse(time.RFC3339, "9999-12-31T23:59:59.999Z")
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Value, &expected)
}

func TestGetUTCUppercaseMaxDateTime7Digits(t *testing.T) {
	client := getDatetimeOperations(t)
	result, err := client.GetUTCUppercaseMaxDateTime7Digits(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	expected, err := time.Parse(time.RFC3339, "9999-12-31T23:59:59.9999999Z")
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Value, &expected)
}

func TestPutLocalNegativeOffsetMaxDateTime(t *testing.T) {
	client := getDatetimeOperations(t)
	body, err := time.Parse(time.RFC3339, "9999-12-31T23:59:59.999-14:00")
	if err != nil {
		t.Fatal(err)
	}
	result, err := client.PutLocalNegativeOffsetMaxDateTime(context.Background(), body)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestPutLocalNegativeOffsetMinDateTime(t *testing.T) {
	client := getDatetimeOperations(t)
	body, err := time.Parse(time.RFC3339, "0001-01-01T00:00:00-14:00")
	if err != nil {
		t.Fatal(err)
	}
	result, err := client.PutLocalNegativeOffsetMinDateTime(context.Background(), body)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestPutLocalPositiveOffsetMaxDateTime(t *testing.T) {
	client := getDatetimeOperations(t)
	body, err := time.Parse(time.RFC3339, "9999-12-31T23:59:59.999+14:00")
	if err != nil {
		t.Fatal(err)
	}
	result, err := client.PutLocalPositiveOffsetMaxDateTime(context.Background(), body)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestPutLocalPositiveOffsetMinDateTime(t *testing.T) {
	client := getDatetimeOperations(t)
	body, err := time.Parse(time.RFC3339, "0001-01-01T00:00:00+14:00")
	if err != nil {
		t.Fatal(err)
	}
	result, err := client.PutLocalPositiveOffsetMinDateTime(context.Background(), body)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestPutUTCMaxDateTime(t *testing.T) {
	client := getDatetimeOperations(t)
	body, err := time.Parse(time.RFC3339, "9999-12-31T23:59:59.999Z")
	if err != nil {
		t.Fatal(err)
	}
	result, err := client.PutUTCMaxDateTime(context.Background(), body)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestPutUTCMaxDateTime7Digits(t *testing.T) {
	client := getDatetimeOperations(t)
	body, err := time.Parse(time.RFC3339, "9999-12-31T23:59:59.9999999Z")
	if err != nil {
		t.Fatal(err)
	}
	result, err := client.PutUTCMaxDateTime7Digits(context.Background(), body)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestPutUTCMinDateTime(t *testing.T) {
	client := getDatetimeOperations(t)
	body, err := time.Parse(time.RFC3339, "0001-01-01T00:00:00Z")
	if err != nil {
		t.Fatal(err)
	}
	result, err := client.PutUTCMinDateTime(context.Background(), body)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}