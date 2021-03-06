package service

import (
	"fmt"

	"github.com/sirupsen/logrus"
	informer "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/tools/cache"

	authzv1alpha1 "github.com/joshvanl/k8s-subject-access-delegation/pkg/apis/authz/v1alpha1"
	"github.com/joshvanl/k8s-subject-access-delegation/pkg/interfaces"
	"github.com/joshvanl/k8s-subject-access-delegation/pkg/subject_access_delegation/utils"
)

const DelServiceKind = "DelService"

type DelService struct {
	log *logrus.Entry

	sad         interfaces.SubjectAccessDelegation
	serviceName string
	replicas    int
	uid         int

	stopCh      chan struct{}
	completedCh chan struct{}

	count     int
	completed bool
	informer  informer.ServiceInformer
}

var _ interfaces.Trigger = &DelService{}

func NewDelService(sad interfaces.SubjectAccessDelegation, trigger *authzv1alpha1.EventTrigger) (*DelService, error) {

	if !utils.ValidName(trigger.Value) {
		return nil, fmt.Errorf("not a valid name '%s', must contain only alphanumerics, '-', '.' and '*'", trigger.Value)
	}

	serviceTrigger := &DelService{
		log:         sad.Log(),
		sad:         sad,
		serviceName: trigger.Value,
		replicas:    trigger.Replicas,
		stopCh:      make(chan struct{}),
		completedCh: make(chan struct{}),
		count:       0,
		completed:   trigger.Triggered,
		uid:         trigger.UID,
		informer:    sad.KubeInformerFactory().Core().V1().Services(),
	}

	serviceTrigger.informer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		DeleteFunc: serviceTrigger.delFunc,
	})

	return serviceTrigger, nil
}

func (p *DelService) delFunc(obj interface{}) {

	service, err := utils.GetServiceObject(obj)
	if err != nil {
		p.log.Errorf("failed to get deleted service object: %v", err)
		return
	}
	if service == nil {
		p.log.Error("failed to get service, received nil object")
	}

	match, err := utils.MatchName(service.Name, p.serviceName)
	if err != nil {
		p.log.Error("failed to match service name: %v", err)
		return
	}

	if !match || p.sad.DeletedUid(service.UID) || p.completed {
		return
	}

	p.sad.AddUid(service.UID)

	p.log.Infof("A service '%s' has been deleted", service.Name)
	p.count++
	if p.count >= p.replicas {
		p.log.Infof("Required replicas was met")
		p.completed = true
		close(p.completedCh)
	}
}

func (p *DelService) WaitOn() (forceClosed bool) {
	p.log.Debug("Trigger waiting")

	if p.watchChannels() {
		p.log.Debug("Del Service Trigger was force closed")
		return true
	}

	p.log.Debug("Del Service Trigger completed")

	if err := p.sad.UpdateTriggerFired(p.uid, true); err != nil {
		p.log.Errorf("error updating delete service trigger status: %v", err)
	}

	return false
}

func (p *DelService) watchChannels() (forceClose bool) {
	for {
		select {
		case <-p.stopCh:
			return true
		case <-p.completedCh:
			return false
		}
	}
}

func (p *DelService) Activate() {
	p.log.Debug("Del Service Trigger Activated")
	p.completed = false

	go p.informer.Informer().Run(p.completedCh)

	return
}

func (p *DelService) Completed() bool {
	return p.completed
}

func (p *DelService) Delete() error {
	close(p.stopCh)
	return nil
}

func (p *DelService) Replicas() int {
	return p.replicas
}

func (p *DelService) Kind() string {
	return DelServiceKind
}
