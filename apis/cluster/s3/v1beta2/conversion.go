// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package v1beta2

import (
	"encoding/json"
	"strconv"
)

// UnmarshalJSON for RuleFilterObservation handles empty strings in float64 fields
func (r *RuleFilterObservation) UnmarshalJSON(data []byte) error {
	type Alias RuleFilterObservation
	aux := &struct {
		ObjectSizeGreaterThan interface{} `json:"objectSizeGreaterThan,omitempty"`
		ObjectSizeLessThan    interface{} `json:"objectSizeLessThan,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	
	// Handle ObjectSizeGreaterThan
	if aux.ObjectSizeGreaterThan != nil {
		switch v := aux.ObjectSizeGreaterThan.(type) {
		case string:
			if v == "" {
				r.ObjectSizeGreaterThan = nil
			} else {
				if f, err := strconv.ParseFloat(v, 64); err == nil {
					r.ObjectSizeGreaterThan = &f
				}
			}
		case float64:
			if v != 0 {
				r.ObjectSizeGreaterThan = &v
			}
		}
	}
	
	// Handle ObjectSizeLessThan
	if aux.ObjectSizeLessThan != nil {
		switch v := aux.ObjectSizeLessThan.(type) {
		case string:
			if v == "" {
				r.ObjectSizeLessThan = nil
			} else {
				if f, err := strconv.ParseFloat(v, 64); err == nil {
					r.ObjectSizeLessThan = &f
				}
			}
		case float64:
			if v != 0 {
				r.ObjectSizeLessThan = &v
			}
		}
	}
	
	return nil
}

// UnmarshalJSON for RuleFilterParameters handles empty strings in float64 fields
func (r *RuleFilterParameters) UnmarshalJSON(data []byte) error {
	type Alias RuleFilterParameters
	aux := &struct {
		ObjectSizeGreaterThan interface{} `json:"objectSizeGreaterThan,omitempty"`
		ObjectSizeLessThan    interface{} `json:"objectSizeLessThan,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	
	// Handle ObjectSizeGreaterThan
	if aux.ObjectSizeGreaterThan != nil {
		switch v := aux.ObjectSizeGreaterThan.(type) {
		case string:
			if v == "" {
				r.ObjectSizeGreaterThan = nil
			} else {
				if f, err := strconv.ParseFloat(v, 64); err == nil {
					r.ObjectSizeGreaterThan = &f
				}
			}
		case float64:
			if v != 0 {
				r.ObjectSizeGreaterThan = &v
			}
		}
	}
	
	// Handle ObjectSizeLessThan
	if aux.ObjectSizeLessThan != nil {
		switch v := aux.ObjectSizeLessThan.(type) {
		case string:
			if v == "" {
				r.ObjectSizeLessThan = nil
			} else {
				if f, err := strconv.ParseFloat(v, 64); err == nil {
					r.ObjectSizeLessThan = &f
				}
			}
		case float64:
			if v != 0 {
				r.ObjectSizeLessThan = &v
			}
		}
	}
	
	return nil
}

// UnmarshalJSON for RuleFilterInitParameters handles empty strings in float64 fields
func (r *RuleFilterInitParameters) UnmarshalJSON(data []byte) error {
	type Alias RuleFilterInitParameters
	aux := &struct {
		ObjectSizeGreaterThan interface{} `json:"objectSizeGreaterThan,omitempty"`
		ObjectSizeLessThan    interface{} `json:"objectSizeLessThan,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	
	// Handle ObjectSizeGreaterThan
	if aux.ObjectSizeGreaterThan != nil {
		switch v := aux.ObjectSizeGreaterThan.(type) {
		case string:
			if v == "" {
				r.ObjectSizeGreaterThan = nil
			} else {
				if f, err := strconv.ParseFloat(v, 64); err == nil {
					r.ObjectSizeGreaterThan = &f
				}
			}
		case float64:
			if v != 0 {
				r.ObjectSizeGreaterThan = &v
			}
		}
	}
	
	// Handle ObjectSizeLessThan
	if aux.ObjectSizeLessThan != nil {
		switch v := aux.ObjectSizeLessThan.(type) {
		case string:
			if v == "" {
				r.ObjectSizeLessThan = nil
			} else {
				if f, err := strconv.ParseFloat(v, 64); err == nil {
					r.ObjectSizeLessThan = &f
				}
			}
		case float64:
			if v != 0 {
				r.ObjectSizeLessThan = &v
			}
		}
	}
	
	return nil
}

// UnmarshalJSON for AndObservation handles empty strings in float64 fields
func (a *AndObservation) UnmarshalJSON(data []byte) error {
	type Alias AndObservation
	aux := &struct {
		ObjectSizeGreaterThan interface{} `json:"objectSizeGreaterThan,omitempty"`
		ObjectSizeLessThan    interface{} `json:"objectSizeLessThan,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	
	// Handle ObjectSizeGreaterThan
	if aux.ObjectSizeGreaterThan != nil {
		switch v := aux.ObjectSizeGreaterThan.(type) {
		case string:
			if v == "" {
				a.ObjectSizeGreaterThan = nil
			} else {
				if f, err := strconv.ParseFloat(v, 64); err == nil {
					a.ObjectSizeGreaterThan = &f
				}
			}
		case float64:
			if v != 0 {
				a.ObjectSizeGreaterThan = &v
			}
		}
	}
	
	// Handle ObjectSizeLessThan
	if aux.ObjectSizeLessThan != nil {
		switch v := aux.ObjectSizeLessThan.(type) {
		case string:
			if v == "" {
				a.ObjectSizeLessThan = nil
			} else {
				if f, err := strconv.ParseFloat(v, 64); err == nil {
					a.ObjectSizeLessThan = &f
				}
			}
		case float64:
			if v != 0 {
				a.ObjectSizeLessThan = &v
			}
		}
	}
	
	return nil
}

// UnmarshalJSON for AndParameters handles empty strings in float64 fields
func (a *AndParameters) UnmarshalJSON(data []byte) error {
	type Alias AndParameters
	aux := &struct {
		ObjectSizeGreaterThan interface{} `json:"objectSizeGreaterThan,omitempty"`
		ObjectSizeLessThan    interface{} `json:"objectSizeLessThan,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	
	// Handle ObjectSizeGreaterThan
	if aux.ObjectSizeGreaterThan != nil {
		switch v := aux.ObjectSizeGreaterThan.(type) {
		case string:
			if v == "" {
				a.ObjectSizeGreaterThan = nil
			} else {
				if f, err := strconv.ParseFloat(v, 64); err == nil {
					a.ObjectSizeGreaterThan = &f
				}
			}
		case float64:
			if v != 0 {
				a.ObjectSizeGreaterThan = &v
			}
		}
	}
	
	// Handle ObjectSizeLessThan
	if aux.ObjectSizeLessThan != nil {
		switch v := aux.ObjectSizeLessThan.(type) {
		case string:
			if v == "" {
				a.ObjectSizeLessThan = nil
			} else {
				if f, err := strconv.ParseFloat(v, 64); err == nil {
					a.ObjectSizeLessThan = &f
				}
			}
		case float64:
			if v != 0 {
				a.ObjectSizeLessThan = &v
			}
		}
	}
	
	return nil
}

// UnmarshalJSON for AndInitParameters handles empty strings in float64 fields
func (a *AndInitParameters) UnmarshalJSON(data []byte) error {
	type Alias AndInitParameters
	aux := &struct {
		ObjectSizeGreaterThan interface{} `json:"objectSizeGreaterThan,omitempty"`
		ObjectSizeLessThan    interface{} `json:"objectSizeLessThan,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	
	// Handle ObjectSizeGreaterThan
	if aux.ObjectSizeGreaterThan != nil {
		switch v := aux.ObjectSizeGreaterThan.(type) {
		case string:
			if v == "" {
				a.ObjectSizeGreaterThan = nil
			} else {
				if f, err := strconv.ParseFloat(v, 64); err == nil {
					a.ObjectSizeGreaterThan = &f
				}
			}
		case float64:
			if v != 0 {
				a.ObjectSizeGreaterThan = &v
			}
		}
	}
	
	// Handle ObjectSizeLessThan
	if aux.ObjectSizeLessThan != nil {
		switch v := aux.ObjectSizeLessThan.(type) {
		case string:
			if v == "" {
				a.ObjectSizeLessThan = nil
			} else {
				if f, err := strconv.ParseFloat(v, 64); err == nil {
					a.ObjectSizeLessThan = &f
				}
			}
		case float64:
			if v != 0 {
				a.ObjectSizeLessThan = &v
			}
		}
	}
	
	return nil
}