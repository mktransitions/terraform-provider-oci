// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service API
//
// **Generative AI Service**
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable LLMs that cover a wide range of use cases for text generation. Use the playground to try out the models out-of-the-box or create and host your own fine-tuned custom models based on your own data on dedicated AI clusters.
//

package generativeai

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VanillaTrainingConfig The Vanilla training config.
type VanillaTrainingConfig struct {

	// The maximum number of training epochs to run for.
	TotalTrainingEpochs *int `mandatory:"false" json:"totalTrainingEpochs"`

	// The initial learning rate to be used during training
	LearningRate *float64 `mandatory:"false" json:"learningRate"`

	// The batch size used during training.
	TrainingBatchSize *int `mandatory:"false" json:"trainingBatchSize"`

	// Stop training if the loss metric does not improve beyond 'early_stopping_threshold' for this many times of evaluation.
	EarlyStoppingPatience *int `mandatory:"false" json:"earlyStoppingPatience"`

	// How much the loss must improve to prevent early stopping.
	EarlyStoppingThreshold *float64 `mandatory:"false" json:"earlyStoppingThreshold"`

	// Determine how frequently to log model metrics.
	// First 20 steps will be logged every step and then will follow this parameter frequency. Set to 0 to disable it.
	LogModelMetricsIntervalInSteps *int `mandatory:"false" json:"logModelMetricsIntervalInSteps"`

	// The number of last layers to be finetuned.
	NumOfLastLayers *int `mandatory:"false" json:"numOfLastLayers"`
}

// GetTotalTrainingEpochs returns TotalTrainingEpochs
func (m VanillaTrainingConfig) GetTotalTrainingEpochs() *int {
	return m.TotalTrainingEpochs
}

// GetLearningRate returns LearningRate
func (m VanillaTrainingConfig) GetLearningRate() *float64 {
	return m.LearningRate
}

// GetTrainingBatchSize returns TrainingBatchSize
func (m VanillaTrainingConfig) GetTrainingBatchSize() *int {
	return m.TrainingBatchSize
}

// GetEarlyStoppingPatience returns EarlyStoppingPatience
func (m VanillaTrainingConfig) GetEarlyStoppingPatience() *int {
	return m.EarlyStoppingPatience
}

// GetEarlyStoppingThreshold returns EarlyStoppingThreshold
func (m VanillaTrainingConfig) GetEarlyStoppingThreshold() *float64 {
	return m.EarlyStoppingThreshold
}

// GetLogModelMetricsIntervalInSteps returns LogModelMetricsIntervalInSteps
func (m VanillaTrainingConfig) GetLogModelMetricsIntervalInSteps() *int {
	return m.LogModelMetricsIntervalInSteps
}

func (m VanillaTrainingConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VanillaTrainingConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m VanillaTrainingConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeVanillaTrainingConfig VanillaTrainingConfig
	s := struct {
		DiscriminatorParam string `json:"trainingConfigType"`
		MarshalTypeVanillaTrainingConfig
	}{
		"VANILLA_TRAINING_CONFIG",
		(MarshalTypeVanillaTrainingConfig)(m),
	}

	return json.Marshal(&s)
}
