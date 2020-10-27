package common

// TemplateConfig this is parsed from config yaml
type TemplateConfig struct {
	HTTPSProxy    string `yaml:"https-proxy,omitempty"`
	HTTPProxy     string `yaml:"http-proxy,omitempty"`
	PubKeyDir     string `yaml:"pub-key-dir"`
	PrivKeyDir    string `yaml:"priv-key-dir"`
	HostPubKeyDir string `yaml:"host-pub-key-dir"`

	PubKey     string
	PrivKey    string
	HostPubKey string
}

// NodeConfig Config for each noed. "node" field  in config yaml
type NodeConfig struct {
	CPU  string `yaml:"cpu"`
	MEM  string `yaml:"mem"`
	DISK string `yaml:"disk"`
}

// Arg Argument
type Arg struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

// KubeInit Configuration for kubeadm init
type KubeInit struct {
	Args []Arg `yaml:"args"`
}

// OverallConfig Over-all configs
/*
This config is in same file with Template config, but when parsed,
the template part will be ignored.
*/
type OverallConfig struct {
	NodeConfig NodeConfig `yaml:"node"`
	NodeNum    int        `yaml:"node-num"`
	Image      string     `yaml:"image"`
	CloudInit  string     `yaml:"cloud-init"`
	KubeInit   KubeInit   `yaml:"kube-init"`
}