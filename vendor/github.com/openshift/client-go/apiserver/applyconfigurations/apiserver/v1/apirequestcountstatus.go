// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// APIRequestCountStatusApplyConfiguration represents a declarative configuration of the APIRequestCountStatus type for use
// with apply.
type APIRequestCountStatusApplyConfiguration struct {
	Conditions       []metav1.ConditionApplyConfiguration         `json:"conditions,omitempty"`
	RemovedInRelease *string                                      `json:"removedInRelease,omitempty"`
	RequestCount     *int64                                       `json:"requestCount,omitempty"`
	CurrentHour      *PerResourceAPIRequestLogApplyConfiguration  `json:"currentHour,omitempty"`
	Last24h          []PerResourceAPIRequestLogApplyConfiguration `json:"last24h,omitempty"`
}

// APIRequestCountStatusApplyConfiguration constructs a declarative configuration of the APIRequestCountStatus type for use with
// apply.
func APIRequestCountStatus() *APIRequestCountStatusApplyConfiguration {
	return &APIRequestCountStatusApplyConfiguration{}
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *APIRequestCountStatusApplyConfiguration) WithConditions(values ...*metav1.ConditionApplyConfiguration) *APIRequestCountStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}

// WithRemovedInRelease sets the RemovedInRelease field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RemovedInRelease field is set to the value of the last call.
func (b *APIRequestCountStatusApplyConfiguration) WithRemovedInRelease(value string) *APIRequestCountStatusApplyConfiguration {
	b.RemovedInRelease = &value
	return b
}

// WithRequestCount sets the RequestCount field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RequestCount field is set to the value of the last call.
func (b *APIRequestCountStatusApplyConfiguration) WithRequestCount(value int64) *APIRequestCountStatusApplyConfiguration {
	b.RequestCount = &value
	return b
}

// WithCurrentHour sets the CurrentHour field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CurrentHour field is set to the value of the last call.
func (b *APIRequestCountStatusApplyConfiguration) WithCurrentHour(value *PerResourceAPIRequestLogApplyConfiguration) *APIRequestCountStatusApplyConfiguration {
	b.CurrentHour = value
	return b
}

// WithLast24h adds the given value to the Last24h field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Last24h field.
func (b *APIRequestCountStatusApplyConfiguration) WithLast24h(values ...*PerResourceAPIRequestLogApplyConfiguration) *APIRequestCountStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithLast24h")
		}
		b.Last24h = append(b.Last24h, *values[i])
	}
	return b
}
