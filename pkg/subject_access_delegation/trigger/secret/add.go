package secret

import (
	"fmt"

	"github.com/sirupsen/logrus"
	informer "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/tools/cache"

	authzv1alpha1 "github.com/joshvanl/k8s-subject-access-delegation/pkg/apis/authz/v1alpha1"
	"github.com/joshvanl/k8s-subject-access-delegation/pkg/interfaces"
	"github.com/joshvanl/k8s-subject-access-delegation/pkg/subject_access_delegation/utils"
)

const AddSecretKind = "AddSecret"

type AddSecret struct {
	log *logrus.Entry

	sad        interfaces.SubjectAccessDelegation
	secretName string
	replicas   int
	uid        int

	stopCh      chan struct{}
	completedCh chan struct{}

	count     int
	completed bool
	informer  informer.SecretInformer
}

var _ interfaces.Trigger = &AddSecret{}

func NewAddSecret(sad interfaces.SubjectAccessDelegation, trigger *authzv1alpha1.EventTrigger) (*AddSecret, error) {

	if !utils.ValidName(trigger.Value) {
		return nil, fmt.Errorf("not a valid name '%s', must contain only alphanumerics, '-', '.' and '*'", trigger.Value)
	}

	secretTrigger := &AddSecret{
		log:         sad.Log(),
		sad:         sad,
		secretName:  trigger.Value,
		replicas:    trigger.Replicas,
		stopCh:      make(chan struct{}),
		completedCh: make(chan struct{}),
		count:       0,
		completed:   trigger.Triggered,
		uid:         trigger.UID,
		informer:    sad.KubeInformerFactory().Core().V1().Secrets(),
	}

	secretTrigger.informer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: secretTrigger.addFunc,
	})

	return secretTrigger, nil
}

func (s *AddSecret) addFunc(obj interface{}) {

	secret, err := utils.GetSecretObject(obj)
	if err != nil {
		s.log.Errorf("failed to get added secret object: %v", err)
		return
	}
	if secret == nil {
		s.log.Error("failed to get secret, received nil object")
	}

	match, err := utils.MatchName(secret.Name, s.secretName)
	if err != nil {
		s.log.Error("failed to match secret name: %v", err)
		return
	}

	if !match || s.sad.SeenUid(secret.UID) || s.completed {
		return
	}

	s.sad.AddUid(secret.UID)

	s.log.Infof("A new secret '%s' has been added", secret.Name)
	s.count++
	if s.count >= s.replicas {
		s.log.Infof("Required replicas was met")
		s.completed = true
		close(s.completedCh)
	}
}

func (s *AddSecret) WaitOn() (forceClosed bool) {
	s.log.Debug("Trigger waiting")

	if s.watchChannels() {
		s.log.Debug("Add Secret Trigger was force closed")
		return true
	}

	s.log.Debug("Add Secret Trigger completed")

	if err := s.sad.UpdateTriggerFired(s.uid, true); err != nil {
		s.log.Errorf("error updating add secret trigger status: %v", err)
	}

	return false
}

func (s *AddSecret) watchChannels() (forceClose bool) {
	for {
		select {
		case <-s.stopCh:
			return true
		case <-s.completedCh:
			return false
		}
	}
}

func (s *AddSecret) Activate() {
	s.log.Debug("Add Secret Trigger Activated")
	s.completed = false

	go s.informer.Informer().Run(s.completedCh)

	return
}

func (s *AddSecret) Completed() bool {
	return s.completed
}

func (s *AddSecret) Delete() error {
	close(s.stopCh)
	return nil
}

func (s *AddSecret) Replicas() int {
	return s.replicas
}

func (s *AddSecret) Kind() string {
	return AddSecretKind
}
