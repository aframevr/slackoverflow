package stackexchange

import "fmt"

// Parameter Stack Exchange parameter
type Parameter struct {
	param       string
	value       interface{}
	description string
}

// String value of the parameter
func (p *Parameter) String() string {
	return fmt.Sprintf("%v", p.value)
}

// Decription of this parameter
func (p *Parameter) Decription() string {
	return p.description
}

// Parameters Stack Exchange query parameters
type Parameters struct {
	applied map[string]Parameter
	allowed map[string]Parameter
}

// Allow parameter to be set
func (p *Parameters) Allow(param string, value interface{}, desc string) {
	if p.allowed == nil {
		p.allowed = make(map[string]Parameter)
	}
	p.allowed[param] = Parameter{param, value, desc}
}

// IsSet return true is parameter is set
func (p *Parameters) IsSet(param string) bool {
	if _, ok := p.applied[param]; ok {
		return true
	}
	return false
}

// Set adds or moifies a parameter given by the key and value
func (p *Parameters) Set(param string, value interface{}) {
	if !p.IsAllowed(param) || value == "" {
		return
	}
	if p.applied == nil {
		p.applied = make(map[string]Parameter)
	}
	p.applied[param] = Parameter{param, value, ""}
}

// ValueOf returns value of given paramter if it is set
func (p *Parameters) ValueOf(param string) string {
	if val, ok := p.applied[param]; ok {
		return val.String()
	}
	return ""
}

// Delete the given parameter
func (p *Parameters) Delete(param string) {
	delete(p.applied, param)
}

// IsAllowed return true is parameter is allowed by this endpoint
func (p *Parameters) IsAllowed(param string) bool {
	if _, ok := p.allowed[param]; ok {
		return true
	}
	return false
}

// GetAllowed returns parameters accepted by this endpoint
func (p *Parameters) GetAllowed() map[string]Parameter {
	return p.allowed
}

// ApplyDefaults apply default parameters
func (p *Parameters) ApplyDefaults() {
	// Do nothing if no defaults are set
	if p.allowed == nil {
		return
	}
	for param, value := range p.GetAllowed() {
		// Set default value if parameter has not been set
		if !p.IsSet(param) {
			p.Set(param, value.String())
		}
	}
}

// GetApplied returns parameters which have been set
func (p *Parameters) GetApplied() map[string]Parameter {
	return p.applied
}
