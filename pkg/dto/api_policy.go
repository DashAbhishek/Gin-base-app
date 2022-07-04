package dto

type Capability struct {
	CapabilityName string        `json:"capabilityName"`
	Description    string        `json:"description"`
	Version        string        `json:"version"`
	TenantName     string        `json:"tenantName"`
	ContextPath    string        `json:"contextPath"`
	HealthCheckURL string        `json:"healthCheckURL"`
	OpenAPISpecs   interface{}   `json:"openAPISpecs"`
	Category       string        `json:"category"`
	EnableCors     bool          `json:"enableCORS"`
	Labels         []string      `json:"labels"`
	Authorization  interface{}   `json:"authorization"`
	Owner          interface{}   `json:"owner"`
	Consumers      []interface{} `json:"consumers"`
}
