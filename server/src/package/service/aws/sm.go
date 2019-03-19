package aws

import (
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/aws"
	"sync"
)

var smOnce sync.Once
var smInstance *SMService

type SMService struct {
	client *secretsmanager.SecretsManager
}

func GetSMService() (*SMService) {
	smOnce.Do(func() {
		smInstance = &SMService{
			client: secretsmanager.New(Session),
		}
	})
	return smInstance
}

func (service *SMService) GetClient() (*secretsmanager.SecretsManager) {
	return service.client
}

func (service *SMService) GetRDSSecretArn(secretName string) (*string, error) {
	output, err := service.client.ListSecretVersionIds(&secretsmanager.ListSecretVersionIdsInput{
		SecretId: aws.String(secretName),
	})

	if err != nil {
		return nil, err
	}

	return output.ARN, nil
}
