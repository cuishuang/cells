// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cellsapi-templates.proto

package rest

import (
	fmt "fmt"
	math "math"
	proto "google.golang.org/protobuf/proto"
	_ "github.com/pydio/cells/v4/common/proto/tree"
	_ "github.com/pydio/cells/v4/common/proto/service"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *TemplateNode) Validate() error {
	if this.Node != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Node); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Node", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *Template) Validate() error {
	if this.Node != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Node); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Node", err)
		}
	}
	for _, item := range this.Policies {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Policies", err)
			}
		}
	}
	return nil
}
func (this *ListTemplatesRequest) Validate() error {
	return nil
}
func (this *ListTemplatesResponse) Validate() error {
	for _, item := range this.Templates {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Templates", err)
			}
		}
	}
	return nil
}
