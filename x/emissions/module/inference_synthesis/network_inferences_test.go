package inference_synthesis_test

import (
	"log"

	inference_synthesis "github.com/allora-network/allora-chain/x/emissions/module/inference_synthesis"
	emissions "github.com/allora-network/allora-chain/x/emissions/types"
)

func (s *InferenceSynthesisTestSuite) TestCalcWeightedInference() {
	topicId := inference_synthesis.TopicId(1)

	tests := []struct {
		name                                  string
		inferenceByWorker                     map[string]*emissions.Inference
		forecastImpliedInferenceByWorker      map[string]*emissions.Inference
		maxRegret                             inference_synthesis.Regret
		epsilon                               float64
		pInferenceSynthesis                   float64
		infererNetworkRegrets                 map[string]inference_synthesis.Regret
		expectedNetworkCombinedInferenceValue float64
		expectedErr                           error
	}{
		{ // EPOCH 3
			name: "normal operation",
			inferenceByWorker: map[string]*emissions.Inference{
				"worker0": {Value: -0.0514234892489971},
				"worker1": {Value: -0.0316532211989242},
				"worker2": {Value: -0.1018014248041400},
			},
			forecastImpliedInferenceByWorker: map[string]*emissions.Inference{
				"worker3": {Value: -0.0707517711518230},
				"worker4": {Value: -0.0646463841210426},
				"worker5": {Value: -0.0634099113416666},
			},
			maxRegret:           0.9871536722074480,
			epsilon:             1e-4,
			pInferenceSynthesis: 2,
			infererNetworkRegrets: map[string]inference_synthesis.Regret{
				"worker0": 0.6975029322458370,
				"worker1": 0.910174442412618,
				"worker2": 0.9871536722074480,
				"worker3": 0.8308330665491310,
				"worker4": 0.8396961220162480,
				"worker5": 0.8017696138115460,
			},
			expectedNetworkCombinedInferenceValue: -0.06470631905627390,
			expectedErr:                           nil,
		},
		{ // EPOCH 4
			name: "normal operation2",
			inferenceByWorker: map[string]*emissions.Inference{
				"worker0": {Value: -0.14361768314408600},
				"worker1": {Value: -0.23422685055675900},
				"worker2": {Value: -0.18201270373970600},
			},
			forecastImpliedInferenceByWorker: map[string]*emissions.Inference{
				"worker3": {Value: -0.19840891048468800},
				"worker4": {Value: -0.19696044261177800},
				"worker5": {Value: -0.20289734770434400},
			},
			maxRegret:           0.9737035757621540,
			epsilon:             1e-4,
			pInferenceSynthesis: 2,
			infererNetworkRegrets: map[string]inference_synthesis.Regret{
				"worker0": 0.5576393860961080,
				"worker1": 0.8588215562008240,
				"worker2": 0.9737035757621540,
				"worker3": 0.7535724745797420,
				"worker4": 0.7658774622830770,
				"worker5": 0.7185104293863190,
			},
			expectedNetworkCombinedInferenceValue: -0.19466636004515200,
			expectedErr:                           nil,
		},
	}

	for _, tc := range tests {
		s.Run(tc.name, func() {
			for inferer, regret := range tc.infererNetworkRegrets {
				s.emissionsKeeper.SetInfererNetworkRegret(
					s.ctx,
					topicId,
					[]byte(inferer),
					emissions.TimestampedValue{BlockHeight: 0, Value: regret},
				)
			}

			networkCombinedInferenceValue, err := inference_synthesis.CalcWeightedInference(
				s.ctx,
				s.emissionsKeeper,
				topicId,
				tc.inferenceByWorker,
				tc.forecastImpliedInferenceByWorker,
				tc.maxRegret,
				tc.epsilon,
				tc.pInferenceSynthesis,
			)

			if tc.expectedErr != nil {
				s.Require().ErrorIs(err, tc.expectedErr)
			} else {
				s.Require().NoError(err)
				s.Require().InEpsilon(tc.expectedNetworkCombinedInferenceValue, networkCombinedInferenceValue, 1e-5, "Network combined inference value should match expected value within epsilon")
			}
		})
	}
}

func (s *InferenceSynthesisTestSuite) TestCalcOneOutInferences() {
	topicId := inference_synthesis.TopicId(1)

	test := struct {
		name                             string
		inferenceByWorker                map[string]*emissions.Inference
		forecastImpliedInferenceByWorker map[string]*emissions.Inference
		forecasts                        *emissions.Forecasts
		maxRegret                        inference_synthesis.Regret
		networkCombinedLoss              inference_synthesis.Loss
		epsilon                          float64
		pInferenceSynthesis              float64
		expectedOneOutInferences         []*emissions.WithheldWorkerAttributedValue
		expectedOneOutImpliedInferences  []*emissions.WithheldWorkerAttributedValue
	}{ // ROW 5
		name: "basic functionality, multiple workers",
		inferenceByWorker: map[string]*emissions.Inference{
			"worker0": &emissions.Inference{Value: 0.09688553736890290},
			"worker1": &emissions.Inference{Value: 0.15603487178220000},
			"worker2": &emissions.Inference{Value: 0.00987426948965807},
		},
		forecastImpliedInferenceByWorker: map[string]*emissions.Inference{
			"worker0": &emissions.Inference{Value: 0.09590746110637150},
			"worker1": &emissions.Inference{Value: 0.09199706634747750},
			"worker2": &emissions.Inference{Value: 0.07867746964190580},
		},
		forecasts: &emissions.Forecasts{
			Forecasts: []*emissions.Forecast{
				{
					Forecaster: "forecaster0",
					ForecastElements: []*emissions.ForecastElement{
						{Inferer: "worker0", Value: 9.65209481504552e-06},
						{Inferer: "worker1", Value: 0.0013204058258572500},
						{Inferer: "worker2", Value: 0.009498919738615450},
					},
				},
				{
					Forecaster: "forecaster1",
					ForecastElements: []*emissions.ForecastElement{
						{Inferer: "worker0", Value: 1.57700563929882e-05},
						{Inferer: "worker1", Value: 0.002446373314877150},
						{Inferer: "worker2", Value: 0.00426518781753509},
					},
				},
			},
		},
		maxRegret:           0.5,
		networkCombinedLoss: 10.0,
		epsilon:             1e-4,
		expectedOneOutInferences: []*emissions.WithheldWorkerAttributedValue{
			{Worker: "worker0", Value: 0.07868265511452390},
			{Worker: "worker1", Value: 0.05882929409106640},
			{Worker: "worker2", Value: 0.12094791926963100},
		},
		expectedOneOutImpliedInferences: []*emissions.WithheldWorkerAttributedValue{
			{Worker: "worker0", Value: 0.11562305562592500},
			{Worker: "worker1", Value: 0.07351778409912410},
			{Worker: "worker2", Value: 0.11683957010303600},
		},
		pInferenceSynthesis: 2.0,
	}

	s.Run(test.name, func() {
		oneOutInferences, oneOutImpliedInferences, err := inference_synthesis.CalcOneOutInferences(
			s.ctx,
			s.emissionsKeeper,
			topicId,
			test.inferenceByWorker,
			test.forecastImpliedInferenceByWorker,
			test.forecasts,
			test.maxRegret,
			test.networkCombinedLoss,
			test.epsilon,
			test.pInferenceSynthesis,
		)

		s.Require().NoError(err, "CalcOneOutInferences should not return an error")

		s.Require().Len(oneOutInferences, len(test.expectedOneOutInferences), "Unexpected number of one-out inferences")
		s.Require().Len(oneOutImpliedInferences, len(test.expectedOneOutImpliedInferences), "Unexpected number of one-out implied inferences")

		for i, expected := range test.expectedOneOutInferences {
			log.Printf("Expected value: %v, Actual value: %v, Epsilon: %v", expected.Value, oneOutInferences[i].Value, 1e-5)
			s.Require().InEpsilon(expected.Value, oneOutInferences[i].Value, 1e-5, "Mismatch in value for one-out inference of worker %s", expected.Worker)
		}

		for i, expected := range test.expectedOneOutImpliedInferences {
			log.Printf("Expected value: %v, Actual value: %v, Epsilon: %v", expected.Value, oneOutImpliedInferences[i].Value, 1e-5)
			s.Require().InEpsilon(expected.Value, oneOutImpliedInferences[i].Value, 1e-5, "Mismatch in value for one-out implied inference of worker %s", expected.Worker)
		}
	})
}

func (s *InferenceSynthesisTestSuite) TestCalcOneInInferences() {
	topicId := inference_synthesis.TopicId(1)

	tests := []struct {
		name                        string
		inferences                  map[string]*emissions.Inference
		forecastImpliedInferences   map[string]*emissions.Inference
		maxRegretsByOneInForecaster map[string]inference_synthesis.Regret
		epsilon                     float64
		pInferenceSynthesis         float64
		expectedOneInInferences     []*emissions.WorkerAttributedValue
		expectedErr                 error
	}{
		{ // ROW 6
			name: "basic functionality, single worker",
			inferences: map[string]*emissions.Inference{
				"worker0": {Value: 0.10711562728325500},
				"worker1": {Value: 0.03008145586124120},
				"worker2": {Value: 0.09269114998018040},
			},
			forecastImpliedInferences: map[string]*emissions.Inference{
				"worker0": {Value: 0.08584946856167300},
				"worker1": {Value: 0.08215179314806270},
				"worker2": {Value: 0.0891905081396791},
			},
			maxRegretsByOneInForecaster: map[string]inference_synthesis.Regret{
				"worker0": 0.1,
				"worker1": 0.2,
			},
			epsilon:             0.0001,
			pInferenceSynthesis: 2.0,
			expectedOneInInferences: []*emissions.WorkerAttributedValue{
				{Worker: "worker0", Value: 0.07646863529477600},
				{Worker: "worker1", Value: 0.0755370605649977},
				{Worker: "worker2", Value: 0.07705216278952520},
			},
			expectedErr: nil,
		},
	}

	for _, tc := range tests {
		s.Run(tc.name, func() {
			oneInInferences, err := inference_synthesis.CalcOneInInferences(
				s.ctx,
				s.emissionsKeeper,
				topicId,
				tc.inferences,
				tc.forecastImpliedInferences,
				tc.maxRegretsByOneInForecaster,
				tc.epsilon,
				tc.pInferenceSynthesis,
			)

			if tc.expectedErr != nil {
				s.Require().ErrorIs(err, tc.expectedErr)
			} else {
				s.Require().NoError(err)
				s.Require().Len(oneInInferences, len(tc.expectedOneInInferences), "Unexpected number of one-in inferences")

				for _, expected := range tc.expectedOneInInferences {
					found := false
					for _, actual := range oneInInferences {
						if expected.Worker == actual.Worker {
							s.Require().InEpsilon(expected.Value, actual.Value, 1e-5, "Mismatch in value for one-in inference of worker %s", expected.Worker)
							found = true
							break
						}
					}
					if !found {
						s.FailNow("Matching worker not found", "Worker %s not found in actual inferences", expected.Worker)
					}
				}
			}
		})
	}
}
