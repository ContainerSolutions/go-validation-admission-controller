package server

import (
	"github.com/sirupsen/logrus"
	"k8s.io/api/admission/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NamespaceAdmission struct {
}

func (*NamespaceAdmission) HandleAdmission(review *v1beta1.AdmissionReview) error {
	logrus.Debugln(review.Request)
	review.Response = &v1beta1.AdmissionResponse{
		Allowed: true,
		UID: review.Request.UID,
		Result: &v1.Status{
			Message: "Welcome aboard!",
		},
	}
	return nil
}
