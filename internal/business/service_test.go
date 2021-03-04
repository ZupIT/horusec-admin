package business

import (
	"os"
	"os/user"
	"testing"

	"github.com/ZupIT/horusec-admin/internal/kubernetes"
	"github.com/stretchr/testify/assert"
)

func TestNewConfigService_Expect_NotNil(t *testing.T) {
	defer tearDown()
	svc := setup()

	assert.NotNil(t, svc)
}

func TestConfigService_GetConfig_Expect_NotNil(t *testing.T) {
	defer tearDown()
	svc := setup()

	cfg, err := svc.GetConfig()
	if err != nil {
		panic(err)
	}

	assert.NotNil(t, cfg)
}

func TestConfigService_GetConfig_Expect_NoError(t *testing.T) {
	defer tearDown()
	svc := setup()

	_, err := svc.GetConfig()
	assert.NoError(t, err)
}

func setup() *ConfigService {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	_ = os.Setenv("KUBECONFIG", usr.HomeDir+"/.kube/config")
	_ = os.Setenv("NAMESPACE", "horusec-system")

	restcfg, err := kubernetes.NewRestConfig()
	if err != nil {
		panic(err)
	}
	client, err := kubernetes.NewHorusecManagerClient(restcfg)
	if err != nil {
		panic(err)
	}
	return NewConfigService(client)
}

func tearDown() {
	_ = os.Unsetenv("KUBECONFIG")
	_ = os.Unsetenv("NAMESPACE")
}
