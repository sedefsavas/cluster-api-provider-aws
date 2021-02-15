package scope

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	infrav1 "sigs.k8s.io/cluster-api-provider-aws/api/v1alpha3"
	"time"
)

type AWSPrincipalTypeProvider interface {
	credentials.Provider
	// Hash returns a unique hash of the data forming the credentials
	// for this Principal
	Hash() (string, error)
	Name() string
}

func NewAWSControllerPrincipalTypeProvider() *AWSControllerPrincipalTypeProvider {
	def := defaults.Get()
	creds, _ := def.Config.Credentials.Get()
	fmt.Println(creds)

	return &AWSControllerPrincipalTypeProvider{
		credentials:     credentials.NewStaticCredentials(creds.AccessKeyID, creds.SecretAccessKey, creds.SessionToken),
		accessKeyID:     creds.AccessKeyID,
		secretAccessKey: creds.SecretAccessKey,
		sessionToken:     creds.SessionToken,
	}
}

type AWSControllerPrincipalTypeProvider struct {
	credentials *credentials.Credentials
	// these are for tests :/
	accessKeyID     string
	secretAccessKey string
	sessionToken    string
}
func (p *AWSControllerPrincipalTypeProvider) Retrieve() (credentials.Value, error) {
	return p.credentials.Get()
}
func (p *AWSControllerPrincipalTypeProvider) IsExpired() bool {
	return p.credentials.IsExpired()
}
func (p *AWSControllerPrincipalTypeProvider) Name() string {
	return "controller-principal"
}
func (p *AWSControllerPrincipalTypeProvider) Hash() (string, error) {
	var roleIdentityValue bytes.Buffer
	err := gob.NewEncoder(&roleIdentityValue).Encode(p)
	if err != nil {
		return "", err
	}
	hash := sha256.New()
	return string(hash.Sum(roleIdentityValue.Bytes())), nil
}

func NewAWSStaticPrincipalTypeProvider(principal *infrav1.AWSClusterStaticPrincipal, secret *corev1.Secret) *AWSStaticPrincipalTypeProvider {
	accessKeyID := string(secret.Data["AccessKeyID"])
	secretAccessKey := string(secret.Data["SecretAccessKey"])
	sessionToken := string(secret.Data["SessionToken"])

	return &AWSStaticPrincipalTypeProvider{
		Principal:       principal,
		credentials:     credentials.NewStaticCredentials(accessKeyID, secretAccessKey, sessionToken),
		accessKeyID:     accessKeyID,
		secretAccessKey: secretAccessKey,
		sessionToken:    sessionToken,
	}
}

func GetAssumeRoleCredentials(principal *infrav1.AWSClusterRolePrincipal, awsConfig *aws.Config) *credentials.Credentials {
	sess := session.Must(session.NewSession(awsConfig))

	creds := stscreds.NewCredentials(sess, principal.Spec.RoleArn, func(p *stscreds.AssumeRoleProvider) {
		if principal.Spec.ExternalID != "" {
			p.ExternalID = aws.String(principal.Spec.ExternalID)
		}
		p.RoleSessionName = principal.Spec.SessionName
		if principal.Spec.InlinePolicy != "" {
			p.Policy = aws.String(principal.Spec.InlinePolicy)
		}
		p.Duration = time.Duration(principal.Spec.DurationSeconds) * time.Second
	})
	return creds
}

func NewAWSRolePrincipalTypeProvider(principal *infrav1.AWSClusterRolePrincipal, sourceProvider *AWSPrincipalTypeProvider, log logr.Logger) *AWSRolePrincipalTypeProvider {
	if sourceProvider != nil{
		fmt.Printf("xx asil %v source %v\n", principal.Name, (*sourceProvider).Name())
	} else{
		fmt.Printf("xx asil %v source yok\n", principal.Name)
	}
	return &AWSRolePrincipalTypeProvider{
		credentials: nil,
		//credentials: GetAssumeRoleCredentials(principal, awsConfig),
		Principal:   principal,
		sourceProvider: sourceProvider,
		log:         log.WithName("AWSRolePrincipalTypeProvider"),
	}
}

type AWSStaticPrincipalTypeProvider struct {
	Principal   *infrav1.AWSClusterStaticPrincipal
	credentials *credentials.Credentials
	// these are for tests :/
	accessKeyID     string
	secretAccessKey string
	sessionToken    string
}

func (p *AWSStaticPrincipalTypeProvider) Hash() (string, error) {
	var roleIdentityValue bytes.Buffer
	err := gob.NewEncoder(&roleIdentityValue).Encode(p)
	if err != nil {
		return "", err
	}
	hash := sha256.New()
	return string(hash.Sum(roleIdentityValue.Bytes())), nil
}
func (p *AWSStaticPrincipalTypeProvider) Retrieve() (credentials.Value, error) {
	return p.credentials.Get()
}
func (p *AWSStaticPrincipalTypeProvider) Name() string {
	return p.Principal.Name
}
func (p *AWSStaticPrincipalTypeProvider) IsExpired() bool {
	return p.credentials.IsExpired()
}

type AWSRolePrincipalTypeProvider struct {
	Principal   *infrav1.AWSClusterRolePrincipal
	credentials *credentials.Credentials
	sourceProvider *AWSPrincipalTypeProvider
	log         logr.Logger
	name string
}

func (p *AWSRolePrincipalTypeProvider) Hash() (string, error) {
	var roleIdentityValue bytes.Buffer
	err := gob.NewEncoder(&roleIdentityValue).Encode(p)
	if err != nil {
		return "", err
	}
	hash := sha256.New()
	return string(hash.Sum(roleIdentityValue.Bytes())), nil
}

func (p *AWSRolePrincipalTypeProvider) Name() string {
	return p.Principal.Name
}
func (p *AWSRolePrincipalTypeProvider) Retrieve() (credentials.Value, error) {
	fmt.Println("xxooxxoo")
	a, err :=json.Marshal(p)
	if err != nil{
		fmt.Println("xx role principal object")
		fmt.Println(a)
	}
	if p.credentials == nil || p.IsExpired(){
		fmt.Printf("xx RETRIEVE: %v\n", p.Principal.Name)
		awsConfig := aws.NewConfig()
		if p.sourceProvider != nil {
			fmt.Printf("xx SOURCE RETRIEVE: %v\n", (*p.sourceProvider).Name())
			sourceCreds, err := (*p.sourceProvider).Retrieve()
			fmt.Printf("xx SOURCE RETRIEVE: %v\n", sourceCreds.ProviderName)
			if err != nil {
				return credentials.Value{}, err
			}
			awsConfig = awsConfig.WithCredentials(credentials.NewStaticCredentialsFromCreds(sourceCreds))
		}
		fmt.Printf("xx ASSUMING ROLE: %v\n", p.Principal.Name)

		creds := GetAssumeRoleCredentials(p.Principal, awsConfig)
		// Update credentials
		p.credentials = creds
	}
	return p.credentials.Get()
}

func (p *AWSRolePrincipalTypeProvider) IsExpired() bool {
	fmt.Printf("xx isexpired: %v\n",p.credentials.IsExpired())
	return p.credentials.IsExpired()
}
