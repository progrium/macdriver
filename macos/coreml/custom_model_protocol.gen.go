// Code generated by DarwinKit. DO NOT EDIT.

package coreml

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// An interface that defines the behavior of a custom model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlcustommodel?language=objc
type PCustomModel interface {
	// optional
	PredictionsFromBatchOptionsError(inputBatch BatchProviderObject, options PredictionOptions, error foundation.Error) BatchProviderObject
	HasPredictionsFromBatchOptionsError() bool

	// optional
	PredictionFromFeaturesOptionsError(input FeatureProviderObject, options PredictionOptions, error foundation.Error) FeatureProviderObject
	HasPredictionFromFeaturesOptionsError() bool

	// optional
	InitWithModelDescriptionParameterDictionaryError(modelDescription ModelDescription, parameters map[string]objc.Object, error foundation.Error) objc.Object
	HasInitWithModelDescriptionParameterDictionaryError() bool
}

// ensure impl type implements protocol interface
var _ PCustomModel = (*CustomModelObject)(nil)

// A concrete type for the [PCustomModel] protocol.
type CustomModelObject struct {
	objc.Object
}

func (c_ CustomModelObject) HasPredictionsFromBatchOptionsError() bool {
	return c_.RespondsToSelector(objc.Sel("predictionsFromBatch:options:error:"))
}

// Predicts output values from the given batch of input features. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlcustommodel/2994298-predictionsfrombatch?language=objc
func (c_ CustomModelObject) PredictionsFromBatchOptionsError(inputBatch BatchProviderObject, options PredictionOptions, error foundation.Error) BatchProviderObject {
	po0 := objc.WrapAsProtocol("MLBatchProvider", inputBatch)
	rv := objc.Call[BatchProviderObject](c_, objc.Sel("predictionsFromBatch:options:error:"), po0, objc.Ptr(options), objc.Ptr(error))
	return rv
}

func (c_ CustomModelObject) HasPredictionFromFeaturesOptionsError() bool {
	return c_.RespondsToSelector(objc.Sel("predictionFromFeatures:options:error:"))
}

// Predicts output values from the given input features. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlcustommodel/2994297-predictionfromfeatures?language=objc
func (c_ CustomModelObject) PredictionFromFeaturesOptionsError(input FeatureProviderObject, options PredictionOptions, error foundation.Error) FeatureProviderObject {
	po0 := objc.WrapAsProtocol("MLFeatureProvider", input)
	rv := objc.Call[FeatureProviderObject](c_, objc.Sel("predictionFromFeatures:options:error:"), po0, objc.Ptr(options), objc.Ptr(error))
	return rv
}

func (c_ CustomModelObject) HasInitWithModelDescriptionParameterDictionaryError() bool {
	return c_.RespondsToSelector(objc.Sel("initWithModelDescription:parameterDictionary:error:"))
}

// Creates a custom model with the given description and parameters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlcustommodel/2994296-initwithmodeldescription?language=objc
func (c_ CustomModelObject) InitWithModelDescriptionParameterDictionaryError(modelDescription ModelDescription, parameters map[string]objc.Object, error foundation.Error) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("initWithModelDescription:parameterDictionary:error:"), objc.Ptr(modelDescription), parameters, objc.Ptr(error))
	return rv
}
