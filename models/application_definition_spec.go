// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApplicationDefinitionSpec ApplicationDefinitionSpec defines the desired state of ApplicationDefinition.
//
// swagger:model ApplicationDefinitionSpec
type ApplicationDefinitionSpec struct {

	// Default specifies if the application should be installed by default when a new user cluster is created. Default applications are
	// not enforced and users can update/delete them. KKP will only install them during cluster creation if the user didn't explicitly
	// opt out from installing default applications.
	// +optional
	Default bool `json:"default,omitempty"`

	// DefaultValuesBlock specifies default values for the UI which are passed to helm templating when creating an application. Comments are preserved.
	DefaultValuesBlock string `json:"defaultValuesBlock,omitempty"`

	// DefaultVersion of the application to use, if not specified the latest available version will be used.
	// +optional
	DefaultVersion string `json:"defaultVersion,omitempty"`

	// Description of the application. what is its purpose
	Description string `json:"description,omitempty"`

	// DisplayName is the name for the application that will be displayed in the UI.
	DisplayName string `json:"displayName,omitempty"`

	// DocumentationURL holds a link to official documentation of the Application
	// Alternatively this can be a link to the Readme of a chart in a git repository
	DocumentationURL string `json:"documentationURL,omitempty"`

	// Enforced specifies if the application is enforced to be installed on the user clusters. Enforced applications are
	// installed/updated by KKP for the user clusters. Users are not allowed to update/delete them. KKP will revert the changes
	// done by the application to the desired state specified in the ApplicationDefinition.
	// +optional
	Enforced bool `json:"enforced,omitempty"`

	// Logo of the Application as a base64 encoded svg
	Logo string `json:"logo,omitempty"`

	// LogoFormat contains logo format of the configured Application. Options are "svg+xml" and "png"
	// +kubebuilder:validation:Enum=svg+xml;png
	LogoFormat string `json:"logoFormat,omitempty"`

	// SourceURL holds a link to the official source code mirror or git repository of the application
	SourceURL string `json:"sourceURL,omitempty"`

	// Available version for this application
	Versions []*ApplicationVersion `json:"versions"`

	// default deploy options
	DefaultDeployOptions *DeployOptions `json:"defaultDeployOptions,omitempty"`

	// default values
	DefaultValues RawExtension `json:"defaultValues,omitempty"`

	// method
	Method TemplateMethod `json:"method,omitempty"`

	// selector
	Selector *DefaultingSelector `json:"selector,omitempty"`
}

// Validate validates this application definition spec
func (m *ApplicationDefinitionSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultDeployOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationDefinitionSpec) validateVersions(formats strfmt.Registry) error {
	if swag.IsZero(m.Versions) { // not required
		return nil
	}

	for i := 0; i < len(m.Versions); i++ {
		if swag.IsZero(m.Versions[i]) { // not required
			continue
		}

		if m.Versions[i] != nil {
			if err := m.Versions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("versions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationDefinitionSpec) validateDefaultDeployOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultDeployOptions) { // not required
		return nil
	}

	if m.DefaultDeployOptions != nil {
		if err := m.DefaultDeployOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultDeployOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defaultDeployOptions")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationDefinitionSpec) validateMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.Method) { // not required
		return nil
	}

	if err := m.Method.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("method")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("method")
		}
		return err
	}

	return nil
}

func (m *ApplicationDefinitionSpec) validateSelector(formats strfmt.Registry) error {
	if swag.IsZero(m.Selector) { // not required
		return nil
	}

	if m.Selector != nil {
		if err := m.Selector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selector")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("selector")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application definition spec based on the context it is used
func (m *ApplicationDefinitionSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVersions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefaultDeployOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMethod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationDefinitionSpec) contextValidateVersions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Versions); i++ {

		if m.Versions[i] != nil {
			if err := m.Versions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("versions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationDefinitionSpec) contextValidateDefaultDeployOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultDeployOptions != nil {
		if err := m.DefaultDeployOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultDeployOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defaultDeployOptions")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationDefinitionSpec) contextValidateMethod(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Method.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("method")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("method")
		}
		return err
	}

	return nil
}

func (m *ApplicationDefinitionSpec) contextValidateSelector(ctx context.Context, formats strfmt.Registry) error {

	if m.Selector != nil {
		if err := m.Selector.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selector")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("selector")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationDefinitionSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationDefinitionSpec) UnmarshalBinary(b []byte) error {
	var res ApplicationDefinitionSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
