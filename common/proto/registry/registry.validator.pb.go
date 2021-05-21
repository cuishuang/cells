// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: registry.proto

package go_micro_registry

import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Service) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	for _, item := range this.Endpoints {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Endpoints", err)
			}
		}
	}
	for _, item := range this.Nodes {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Nodes", err)
			}
		}
	}
	if this.Options != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Options); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Options", err)
		}
	}
	return nil
}
func (this *Node) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *Endpoint) Validate() error {
	if this.Request != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Request); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Request", err)
		}
	}
	if this.Response != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Response); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Response", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *Value) Validate() error {
	for _, item := range this.Values {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Values", err)
			}
		}
	}
	return nil
}
func (this *Options) Validate() error {
	return nil
}
func (this *Result) Validate() error {
	if this.Service != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Service); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Service", err)
		}
	}
	return nil
}
func (this *EmptyResponse) Validate() error {
	return nil
}
func (this *GetRequest) Validate() error {
	return nil
}
func (this *GetResponse) Validate() error {
	for _, item := range this.Services {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Services", err)
			}
		}
	}
	return nil
}
func (this *ListRequest) Validate() error {
	return nil
}
func (this *ListResponse) Validate() error {
	for _, item := range this.Services {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Services", err)
			}
		}
	}
	return nil
}
func (this *WatchRequest) Validate() error {
	return nil
}
func (this *Event) Validate() error {
	if this.Service != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Service); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Service", err)
		}
	}
	return nil
}
