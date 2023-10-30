package main

import (
  "bytes"
  "context"
  "encoding/json"
  "fmt"
  "github.com/JetBrains/ij-perf-report-aggregator/pkg/server/meta"
  "net/http"
  "time"
)

type InsertionResults struct {
  degradation Degradation
  wasInserted bool
  error       error
}

func postDegradation(ctx context.Context, analysisSettings AnalysisSettings, degradations []Degradation) []InsertionResults {
  url := URL + "/api/meta/accidents"
  insertionResults := make([]InsertionResults, len(degradations))
  for i, degradation := range degradations {
    date := time.UnixMilli(degradation.timestamp).UTC().Format("2006-01-02")
    medianMessage := getMessageBasedOnMedianChange(degradation.medianValues)
    insertParams := meta.AccidentInsertParams{Date: date, Test: analysisSettings.test + "/" + analysisSettings.metric, Kind: "Inferred", Reason: medianMessage, BuildNumber: degradation.build}
    params, err := json.Marshal(insertParams)
    if err != nil {
      insertionResults[i] = InsertionResults{error: fmt.Errorf("failed to marshal query: %w", err)}
      continue
    }

    req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(params))
    if err != nil {
      insertionResults[i] = InsertionResults{error: fmt.Errorf("failed to create request: %w", err)}
      continue
    }
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
      insertionResults[i] = InsertionResults{error: fmt.Errorf("failed to send POST request: %w", err)}
      continue
    }

    if resp.StatusCode != http.StatusOK {
      insertionResults[i] = InsertionResults{error: fmt.Errorf("failed to post Degradation: %v", resp.Status)}
      continue
    }

    // the accident already exists
    if resp.StatusCode == http.StatusConflict {
      insertionResults[i] = InsertionResults{}
      continue
    }
    insertionResults[i] = InsertionResults{degradation, true, nil}
    resp.Body.Close()
  }
  return insertionResults
}