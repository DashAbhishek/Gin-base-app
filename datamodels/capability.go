package datamodels

type Capability struct {
	CapabilityName string        `bson:"capability_name"`
	Description    string        `bson:"description"`
	Version        string        `bson:"version"`
	TenantName     string        `bson:"tenant_name"`
	ContextPath    string        `bson:"context_path"`
	HealthCheckURL string        `bson:"health_check_URL"`
	OpenAPISpecs   interface{}   `bson:"open_api_specs"`
	Category       string        `bson:"category"`
	EnableCORS     bool          `bson:"enable_CORS"`
	Labels         []string      `bson:"labels"`
	Authorization  interface{}   `bson:"authorization"`
	Owner          interface{}   `bson:"owner"`
	Consumers      []interface{} `bson:"consumers"`
}
