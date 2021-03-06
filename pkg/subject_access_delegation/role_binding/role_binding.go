package role_binding

import (
	"fmt"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/joshvanl/k8s-subject-access-delegation/pkg/interfaces"
)

const RoleBindingKind = "RoleBinding"

var _ interfaces.Binding = &RoleBinding{}

type RoleBinding struct {
	roleBinding *rbacv1.RoleBinding
	sad         interfaces.SubjectAccessDelegation
}

func NewRoleBinding(sad interfaces.SubjectAccessDelegation, roleRef *rbacv1.RoleRef) interfaces.Binding {
	name := fmt.Sprintf("%s-%s-%s-%s", sad.Name(), sad.OriginSubject().Name(), sad.Namespace(), roleRef.Name)

	return &RoleBinding{
		sad: sad,
		roleBinding: &rbacv1.RoleBinding{
			ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: sad.Namespace()},
			RoleRef:    *roleRef,
			Subjects:   sad.BindingSubjects(),
		},
	}
}

func NewFromRoleBinding(sad interfaces.SubjectAccessDelegation, roleBinding *rbacv1.RoleBinding) interfaces.Binding {
	name := fmt.Sprintf("%s-%s-%s-%s", sad.Name(), sad.OriginSubject().Name(), sad.Namespace(), roleBinding.RoleRef.Name)

	newBinding := new(rbacv1.RoleBinding)
	newBinding.Name = name
	newBinding.Namespace = roleBinding.Namespace
	newBinding.RoleRef = roleBinding.DeepCopy().RoleRef
	newBinding.Subjects = sad.BindingSubjects()

	return &RoleBinding{
		sad:         sad,
		roleBinding: newBinding,
	}
}

func (r *RoleBinding) CreateRoleBinding() (interfaces.Binding, bool, error) {
	options := metav1.GetOptions{}

	b, err := r.sad.Client().Rbac().RoleBindings(r.sad.Namespace()).Get(r.Name(), options)
	if err != nil && b != nil && b.UID != "" {
		return nil, true, nil
	}

	binding, err := r.sad.Client().Rbac().RoleBindings(r.sad.Namespace()).Create(r.roleBinding)
	if err != nil {
		return nil, false, err
	}

	r.roleBinding = binding

	return r, false, nil
}

func (r *RoleBinding) DeleteRoleBinding() error {
	options := &metav1.DeleteOptions{}
	if err := r.sad.Client().Rbac().RoleBindings(r.sad.Namespace()).Delete(r.roleBinding.Name, options); err != nil {
		return err
	}

	return nil
}

func (r *RoleBinding) Name() string {
	return r.roleBinding.Name
}

func (r *RoleBinding) RoleRef() *rbacv1.RoleRef {
	return &r.roleBinding.RoleRef
}

func (r *RoleBinding) Kind() string {
	return RoleBindingKind
}
