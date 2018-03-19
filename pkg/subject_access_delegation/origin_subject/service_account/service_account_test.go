package service_account

import (
	"errors"
	//"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"

	"github.com/joshvanl/k8s-subject-access-delegation/pkg/subject_access_delegation/mocks"
)

var (
	noBindings        = &rbacv1.RoleBindingList{Items: nil}
	noClusterBindings = &rbacv1.ClusterRoleBindingList{Items: nil}
)

type fakeServiceAccount struct {
	*ServiceAccount
	ctrl *gomock.Controller

	fakeClient *mocks.MockInterface
	fakeRbac   *mocks.MockRbacV1Interface
	fakeCore   *mocks.MockCoreV1Interface

	fakeRoleBindingsInterface        *mocks.MockRoleBindingInterface
	fakeServiceAccountInterface      *mocks.MockServiceAccountInterface
	fakeClusterRoleBindingsInterface *mocks.MockClusterRoleBindingInterface
}

func newFakeServiceAccount(t *testing.T) *fakeServiceAccount {
	s := &fakeServiceAccount{
		ctrl: gomock.NewController(t),
		ServiceAccount: &ServiceAccount{
			namespace: "fakeNamespace",
			name:      "me",
		},
	}

	s.fakeClient = mocks.NewMockInterface(s.ctrl)
	s.fakeRbac = mocks.NewMockRbacV1Interface(s.ctrl)
	s.fakeCore = mocks.NewMockCoreV1Interface(s.ctrl)
	s.fakeRoleBindingsInterface = mocks.NewMockRoleBindingInterface(s.ctrl)
	s.fakeClusterRoleBindingsInterface = mocks.NewMockClusterRoleBindingInterface(s.ctrl)

	s.fakeServiceAccountInterface = mocks.NewMockServiceAccountInterface(s.ctrl)
	s.ServiceAccount.client = s.fakeClient

	s.fakeClient.EXPECT().Core().AnyTimes().Return(s.fakeCore)
	s.fakeClient.EXPECT().Rbac().AnyTimes().Return(s.fakeRbac)
	s.fakeCore.EXPECT().ServiceAccounts(s.ServiceAccount.namespace).AnyTimes().Return(s.fakeServiceAccountInterface)

	return s
}

func TestServiceAccount_ResolveDestination_Nil(t *testing.T) {
	s := newFakeServiceAccount(t)
	defer s.ctrl.Finish()

	s.fakeServiceAccountInterface.EXPECT().Get(s.ServiceAccount.name, gomock.Any()).Times(1).Return(nil, nil)

	if err := s.ResolveOrigin(); err == nil {
		t.Error("expected error but got none - rolebindings is nil")
	}
}

func TestServiceAccount_ResolveOrigin_Error(t *testing.T) {
	s := newFakeServiceAccount(t)
	defer s.ctrl.Finish()

	s.fakeServiceAccountInterface.EXPECT().Get(s.ServiceAccount.name, gomock.Any()).Times(1).Return(&corev1.ServiceAccount{}, errors.New("this is an error"))

	if err := s.ResolveOrigin(); err == nil {
		t.Error("expected error but got none - returned error")
	}
}

func TestServiceAccount_ResolveOrigin_Successful(t *testing.T) {
	s := newFakeServiceAccount(t)
	defer s.ctrl.Finish()

	s.fakeRbac.EXPECT().RoleBindings(gomock.Any()).Return(s.fakeRoleBindingsInterface)
	s.fakeRoleBindingsInterface.EXPECT().List(gomock.Any()).Return(noBindings, nil)

	s.fakeRbac.EXPECT().ClusterRoleBindings().Return(s.fakeClusterRoleBindingsInterface)
	s.fakeClusterRoleBindingsInterface.EXPECT().List(gomock.Any()).Return(noClusterBindings, nil)

	if err := s.roleBindings(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestServiceAccount_RoleRefs_Nil(t *testing.T) {
	s := newFakeServiceAccount(t)
	defer s.ctrl.Finish()

	refs, err := s.RoleRefs()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if refs != nil {
		t.Errorf("expected nil, got=%v", refs)
	}
}

func TestServiceAccount_ServiceAccountObject_Success(t *testing.T) {
	s := newFakeServiceAccount(t)
	defer s.ctrl.Finish()

	//s.fakeClient.EXPECT().Rbac().Times(1).Return(s.fakeRbac)
	//s.fakeRbac.EXPECT().RoleBindings(s.namespace).Times(1).Return(s.fakeRoleBindingsInterface)
	s.fakeRoleBindingsInterface.EXPECT().List(gomock.Any()).Return(nil, errors.New("this is an error"))

	_, err := s.RoleRefs()
	if err == nil {
		t.Error("expected error but got none - returned error")
	}
}

//func TestServiceAccount_RoleRefs_Successful_None(t *testing.T) {
//	s := newFakeServiceAccount(t)
//	defer s.ctrl.Finish()
//
//	roleRefsReturn := &rbacv1.RoleBindingList{}
//
//	s.fakeClient.EXPECT().Rbac().Times(1).Return(s.fakeRbac)
//	s.fakeRbac.EXPECT().RoleBindings(s.namespace).Times(1).Return(s.fakeRoleBindingsInterface)
//	s.fakeRoleBindingsInterface.EXPECT().List(gomock.Any()).Return(roleRefsReturn, nil)
//
//	refs, err := s.RoleRefs()
//	if err != nil {
//		t.Errorf("unexpected error: %v", err)
//	}
//
//	if len(refs) != 0 {
//		t.Errorf("unexpected role refs exp=nothing got=%+v", refs)
//	}
//}
//
//func TestServiceAccount_RoleRefs_Successful_All(t *testing.T) {
//	s := newFakeServiceAccount(t)
//	defer s.ctrl.Finish()
//
//	roleBindingsReturn := &rbacv1.RoleBindingList{}
//
//	var roleRef1, roleRef2, roleRef3 rbacv1.RoleRef
//	roleRef1.Name = "roleRef1"
//	roleRef1.Kind = "Role"
//	roleRef2.Name = "roleRef2"
//	roleRef2.Kind = "Role"
//	roleRef3.Name = "roleRef3"
//	roleRef3.Kind = "Role"
//
//	subjects := []rbacv1.Subject{
//		rbacv1.Subject{
//			Kind: "Pod",
//			Name: "foo",
//		},
//		rbacv1.Subject{
//			Kind: "ServiceAccount",
//			Name: "me",
//		},
//		rbacv1.Subject{
//			Kind: "Pod",
//			Name: "me",
//		},
//		rbacv1.Subject{
//			Kind: "ServiceAccount",
//			Name: "bar",
//		},
//	}
//
//	items := []rbacv1.RoleBinding{
//		rbacv1.RoleBinding{
//			RoleRef:  roleRef1,
//			Subjects: subjects,
//		},
//		rbacv1.RoleBinding{
//			RoleRef:  roleRef2,
//			Subjects: subjects,
//		},
//		rbacv1.RoleBinding{
//			RoleRef:  roleRef3,
//			Subjects: subjects,
//		},
//	}
//
//	roleBindingsReturn.Items = items
//
//	roleRefsExp := []*rbacv1.RoleRef{
//		&rbacv1.RoleRef{
//			Name: "roleRef1",
//			Kind: "Role",
//		},
//		&rbacv1.RoleRef{
//			Name: "roleRef2",
//			Kind: "Role",
//		},
//		&rbacv1.RoleRef{
//			Name: "roleRef3",
//			Kind: "Role",
//		},
//	}
//
//	s.fakeClient.EXPECT().Rbac().Times(1).Return(s.fakeRbac)
//	s.fakeRbac.EXPECT().RoleBindings(s.namespace).Times(1).Return(s.fakeRoleBindingsInterface)
//	s.fakeRoleBindingsInterface.EXPECT().List(gomock.Any()).Return(roleBindingsReturn, nil)
//
//	refs, err := s.RoleRefs()
//	if err != nil {
//		t.Errorf("unexpected error: %v", err)
//	}
//
//	if !reflect.DeepEqual(refs, roleRefsExp) {
//		t.Errorf("unexpected role refs\nexp=%+v\ngot=%+v", roleRefsExp, refs)
//	}
//}
//
//func TestServiceAccount_RoleRefs_Successful_Some(t *testing.T) {
//	s := newFakeServiceAccount(t)
//	defer s.ctrl.Finish()
//
//	roleBindingsReturn := &rbacv1.RoleBindingList{}
//
//	var roleRef1, roleRef2, roleRef3 rbacv1.RoleRef
//	roleRef1.Name = "roleRef1"
//	roleRef1.Kind = "Role"
//	roleRef2.Name = "roleRef2"
//	roleRef2.Kind = "Role"
//	roleRef3.Name = "roleRef3"
//	roleRef3.Kind = "Role"
//
//	subjects := []rbacv1.Subject{
//		rbacv1.Subject{
//			Kind: "Pod",
//			Name: "foo",
//		},
//		rbacv1.Subject{
//			Kind: "Pod",
//			Name: "me",
//		},
//		rbacv1.Subject{
//			Kind: "ServiceAccount",
//			Name: "bar",
//		},
//	}
//
//	items := []rbacv1.RoleBinding{
//		rbacv1.RoleBinding{
//			RoleRef:  roleRef1,
//			Subjects: subjects,
//		},
//		rbacv1.RoleBinding{
//			RoleRef:  roleRef2,
//			Subjects: subjects,
//		},
//	}
//
//		Name: "me",
//	})
//
//	items = append(items, rbacv1.RoleBinding{
//		RoleRef:  roleRef3,
//		Subjects: subjects,
//	})
//
//	roleBindingsReturn.Items = items
//
//	roleRefsExp := []*rbacv1.RoleRef{
//		&rbacv1.RoleRef{
//			Name: "roleRef3",
//			Kind: "Role",
//		},
//	}
//
//	s.fakeClient.EXPECT().Rbac().Times(1).Return(s.fakeRbac)
//	s.fakeRbac.EXPECT().RoleBindings(s.namespace).Times(1).Return(s.fakeRoleBindingsInterface)
//	s.fakeRoleBindingsInterface.EXPECT().List(gomock.Any()).Return(roleBindingsReturn, nil)
//
//	refs, err := s.RoleRefs()
//	if err != nil {
//		t.Errorf("unexpected error: %v", err)
//	}
//
//	if !reflect.DeepEqual(refs, roleRefsExp) {
//		t.Errorf("unexpected role refs\nexp=%+v\ngot=%+v", roleRefsExp, refs)
//	}
//}
//
//func TestServiceAccount_RoleRefs_Successful_NoRef(t *testing.T) {
//	s := newFakeServiceAccount(t)
//	defer s.ctrl.Finish()
//
//	roleBindingsReturn := &rbacv1.RoleBindingList{}
//
//	var roleRef1, roleRef2, roleRef3 rbacv1.RoleRef
//	roleRef1.Name = "roleRef1"
//	roleRef1.Kind = "Role"
//	roleRef2.Name = "roleRef2"
//	roleRef2.Kind = "Role"
//	roleRef3.Name = "roleRef3"
//	roleRef3.Kind = "Role"
//
//	subjects := []rbacv1.Subject{
//		rbacv1.Subject{
//			Kind: "Pod",
//			Name: "foo",
//		},
//		rbacv1.Subject{
//			Kind: "Pod",
//			Name: "me",
//		},
//		rbacv1.Subject{
//			Kind: "ServiceAccount",
//			Name: "bee",
//		},
//		rbacv1.Subject{
//			Kind: "ServiceAccount",
//			Name: "bar",
//		},
//	}
//
//	items := []rbacv1.RoleBinding{
//		rbacv1.RoleBinding{
//			RoleRef:  roleRef1,
//			Subjects: subjects,
//		},
//		rbacv1.RoleBinding{
//			RoleRef:  roleRef2,
//			Subjects: subjects,
//		},
//		rbacv1.RoleBinding{
//			RoleRef:  roleRef3,
//			Subjects: subjects,
//		},
//	}
//
//	roleBindingsReturn.Items = items
//
//	s.fakeClient.EXPECT().Rbac().Times(1).Return(s.fakeRbac)
//	s.fakeRbac.EXPECT().RoleBindings(s.namespace).Times(1).Return(s.fakeRoleBindingsInterface)
//	s.fakeRoleBindingsInterface.EXPECT().List(gomock.Any()).Return(roleBindingsReturn, nil)
//
//	refs, err := s.RoleRefs()
//	if err != nil {
//		t.Errorf("unexpected error: %v", err)
//	}
//
//	if len(refs) != 0 {
//		t.Errorf("unexpected role refs exp=nothing got=%+v", refs)
//	}
//}
