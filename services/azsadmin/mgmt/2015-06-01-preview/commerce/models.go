package commerce

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
)

// Resource base Resource Object
type Resource struct {
	// ID - URI of the resource.
	ID *string `json:"id,omitempty"`
	// Name - Name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - Type of resource.
	Type *string `json:"type,omitempty"`
	// Location - Location where resource is location.
	Location *string `json:"location,omitempty"`
	// Tags - List of key value pairs.
	Tags *map[string]*string `json:"tags,omitempty"`
}

// UsageAggregate aggregate usage values for resource.
type UsageAggregate struct {
	// ID - URI of the resource.
	ID *string `json:"id,omitempty"`
	// Name - Name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - Type of resource.
	Type *string `json:"type,omitempty"`
	// Location - Location where resource is location.
	Location *string `json:"location,omitempty"`
	// Tags - List of key value pairs.
	Tags *map[string]*string `json:"tags,omitempty"`
	// UsageAggregateModel - Properties for aggregate usage.
	*UsageAggregateModel `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for UsageAggregate struct.
func (ua *UsageAggregate) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["properties"]
	if v != nil {
		var properties UsageAggregateModel
		err = json.Unmarshal(*m["properties"], &properties)
		if err != nil {
			return err
		}
		ua.UsageAggregateModel = &properties
	}

	v = m["id"]
	if v != nil {
		var ID string
		err = json.Unmarshal(*m["id"], &ID)
		if err != nil {
			return err
		}
		ua.ID = &ID
	}

	v = m["name"]
	if v != nil {
		var name string
		err = json.Unmarshal(*m["name"], &name)
		if err != nil {
			return err
		}
		ua.Name = &name
	}

	v = m["type"]
	if v != nil {
		var typeVar string
		err = json.Unmarshal(*m["type"], &typeVar)
		if err != nil {
			return err
		}
		ua.Type = &typeVar
	}

	v = m["location"]
	if v != nil {
		var location string
		err = json.Unmarshal(*m["location"], &location)
		if err != nil {
			return err
		}
		ua.Location = &location
	}

	v = m["tags"]
	if v != nil {
		var tags map[string]*string
		err = json.Unmarshal(*m["tags"], &tags)
		if err != nil {
			return err
		}
		ua.Tags = &tags
	}

	return nil
}

// UsageAggregateModel properties for aggregate usage.
type UsageAggregateModel struct {
	// SubscriptionID - Subscription id of tenant using plan.
	SubscriptionID *string `json:"subscriptionId,omitempty"`
	// UsageStartTime - UTC start time for the usage bucket to which this usage aggregate belongs.
	UsageStartTime *date.Time `json:"usageStartTime,omitempty"`
	// UsageEndTime - UTC end time for the usage bucket to which this usage aggregate belongs.
	UsageEndTime *date.Time `json:"usageEndTime,omitempty"`
	// InstanceData - Key-value pairs of instance details represented as a string.
	InstanceData *string `json:"instanceData,omitempty"`
	// Quantity - The amount of the resource consumption that occurred in this time frame.
	Quantity *string `json:"quantity,omitempty"`
	// MeterID - Unique ID for the resource that was consumed (aka ResourceID).
	MeterID *string `json:"meterId,omitempty"`
}

// UsageAggregatePage holds an array of usage aggregates and the continuation token.
type UsageAggregatePage struct {
	autorest.Response `json:"-"`
	// Value - Array of usage aggregates.
	Value *[]UsageAggregate `json:"value,omitempty"`
	// NextLink - Continuation token
	NextLink *string `json:"nextLink,omitempty"`
}