package postgresql

import (
	utils "github.com/iboware/postgresql-operator/utils"

	databasev1alpha1 "github.com/iboware/postgresql-operator/apis/database/v1alpha1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//NewSecretForCR creates secrets for StatefulSet
func NewSecretForCR(cr *databasev1alpha1.PostgreSQL) *v1.Secret {
	labels := utils.LabelsForPostgreSQL(cr.ObjectMeta.Name)

	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{Name: cr.ObjectMeta.Name + "-postgresql", Labels: labels, Namespace: cr.ObjectMeta.Namespace},
		Type:       v1.SecretTypeOpaque,
		StringData: map[string]string{
			"postgresql-password": "WG9BbjB2Tk5yTA==",
			"repmgr-password":     "V0xWZGVOUnR6OQ==",
		},
	}
	return secret
}