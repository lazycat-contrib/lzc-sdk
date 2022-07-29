package appspec

type Permission string

type IngressConfig struct {
	Type    string `yaml:"type"`
	Port    string `yaml:"port"`
	Service string `yaml:"service"`
}

type Application struct {
	Subdomain  string   `yaml:"subdomain"`
	PublicPath string   `ymal:"public_path"`
	Routes     []string `ymal:"routes"`

	Ingress []IngressConfig `yaml:"ingress"`

	Image     string `yaml:"image"`
	DependsOn string `yaml:"depends_on"`
}

type Service struct {
	Image       string   `yaml:"image"`
	Environment []string `yaml:"environments"`
	Binds       []string `yaml:"binds"`
	Command     string   `yaml:"command"`
	DependsOn   string   `yaml:"depends_on"`
}

type LZCAppManifest struct {
	LZCSDKVersion string `yaml:"lzc-sdk-version"`
	Package       string `yaml:"package"`
	Version       string `yaml:"version"`

	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Icon        string `yaml:"icon"`

	License  string `yaml:"license"`
	Homepage string `ymal:"homepage"`
	Author   string `ymal:"author"`

	Permissions []Permission `yaml:"permissions"`

	Application Application        `yaml:"application"`
	Service     map[string]Service `yaml:"services"`
}
