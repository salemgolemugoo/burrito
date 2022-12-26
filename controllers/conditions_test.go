/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"strconv"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/padok-team/burrito/annotations"
	configv1alpha1 "github.com/padok-team/burrito/api/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	//+kubebuilder:scaffold:imports
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

func TestConditions(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Conditions Suite")
}

var _ = BeforeSuite(func() {
	logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true)))
})

var _ = AfterSuite(func() {
})

var _ = Describe("TerraformLayer", func() {
	var t *configv1alpha1.TerraformLayer

	BeforeEach(func() {
		t = &configv1alpha1.TerraformLayer{
			Spec: configv1alpha1.TerraformLayerSpec{
				Path:   "/test",
				Branch: "main",
				Repository: configv1alpha1.TerraformLayerRepository{
					Name:      "test-repository",
					Namespace: "default",
				},
			},
		}
		t.SetAnnotations(map[string]string{})
	})

	Describe("TerraformRunningCondition", func() {
		var condition TerraformRunning
		BeforeEach(func() {
			condition = TerraformRunning{}
		})
		Context("with lock", func() {
			It("should return true", func() {
				t.Annotations[annotations.Lock] = "1"
				Expect(condition.Evaluate(t)).To(Equal(true))
			})
		})
		Context("without lock", func() {
			It("should return false", func() {
				Expect(condition.Evaluate(t)).To(Equal(false))
			})
		})
	})
	Describe("TerraformPlanArtifactCondition", func() {
		var condition TerraformPlanArtifactUpToDate
		BeforeEach(func() {
			condition = TerraformPlanArtifactUpToDate{}
		})
		Context("without last timestamp", func() {
			It("should return false", func() {
				Expect(condition.Evaluate(t)).To(Equal(false))
			})
		})
		Context("with last timestamp < 20min", func() {
			It("should return true", func() {
				t.Annotations[annotations.LastPlanDate] = strconv.Itoa(int(time.Now().Add(-time.Minute * 15).Unix()))
				Expect(condition.Evaluate(t)).To(Equal(true))
			})
		})
		Context("with last timestamp > 20min", func() {
			It("should return false", func() {
				t.Annotations[annotations.LastPlanDate] = strconv.Itoa(int(time.Now().Add(-time.Minute * 60).Unix()))
				Expect(condition.Evaluate(t)).To(Equal(false))
			})
		})
	})
	Describe("TerraformApplyUpToDateCondition", func() {
		var condition TerraformApplyUpToDate
		BeforeEach(func() {
			condition = TerraformApplyUpToDate{}
		})
		Context("without plan", func() {
			It("should return true", func() {
				Expect(condition.Evaluate(t)).To(Equal(true))
			})
		})
		Context("with plan but no apply", func() {
			It("should return false", func() {
				t.Annotations[annotations.LastPlanSum] = "ThisIsAPlanArtifact"
				Expect(condition.Evaluate(t)).To(Equal(false))
			})
		})
		Context("with same plan and apply", func() {
			It("should return true", func() {
				t.Annotations[annotations.LastPlanSum] = "ThisIsAPlanArtifact"
				t.Annotations[annotations.LastApplySum] = "ThisIsAPlanArtifact"
				Expect(condition.Evaluate(t)).To(Equal(true))
			})
		})
		Context("with different plan and apply", func() {
			It("should return false", func() {
				t.Annotations[annotations.LastPlanSum] = "ThisIsAPlanArtifact"
				t.Annotations[annotations.LastApplySum] = "ThisIsAnotherPlanArtifact"
				Expect(condition.Evaluate(t)).To(Equal(false))
			})
		})
	})
	Describe("TerraformFailureCondition", func() {
		var condition TerraformFailure
		BeforeEach(func() {
			condition = TerraformFailure{}
		})
		Context("without run result", func() {
			It("should return false", func() {
				Expect(condition.Evaluate(t)).To(Equal(false))
			})
		})
		Context("with terraform failure", func() {
			It("should return true", func() {
				t.Annotations[annotations.Failure] = "1"
				Expect(condition.Evaluate(t)).To(Equal(true))
			})
		})
	})
	Describe("TerraformLayerConditions", func() {
		var conditions TerraformLayerConditions
		BeforeEach(func() {
			conditions = TerraformLayerConditions{Resource: t}
		})
		Context("terraform is running", func() {
			It("", func() {
				t.Annotations[annotations.Lock] = "1"
				_, out := conditions.Evaluate(context.TODO())
				Expect(out[0].Status).To(Equal(metav1.ConditionTrue))
			})
		})
		Context("terraform not running and everything is up to date", func() {
			It("", func() {
				t.Annotations[annotations.LastPlanDate] = strconv.Itoa(int(time.Now().Add(-time.Minute * 15).Unix()))
				t.Annotations[annotations.LastPlanSum] = "ThisIsAPlanArtifact"
				t.Annotations[annotations.LastApplySum] = "ThisIsAPlanArtifact"
				_, out := conditions.Evaluate(context.TODO())
				Expect(out[0].Status).To(Equal(metav1.ConditionFalse))
				Expect(out[1].Status).To(Equal(metav1.ConditionTrue))
				Expect(out[2].Status).To(Equal(metav1.ConditionTrue))
			})
		})
		Context("terraform not running, plan up to date, apply not up to date, terraform has failed", func() {
			It("", func() {
				t.Annotations[annotations.LastPlanDate] = strconv.Itoa(int(time.Now().Add(-time.Minute * 15).Unix()))
				t.Annotations[annotations.LastPlanSum] = "ThisIsAPlanArtifact"
				t.Annotations[annotations.LastApplySum] = "ThisIsAnotherPlanArtifact"
				t.Annotations[annotations.Failure] = "1"
				_, out := conditions.Evaluate(context.TODO())
				Expect(out[0].Status).To(Equal(metav1.ConditionFalse))
				Expect(out[1].Status).To(Equal(metav1.ConditionTrue))
				Expect(out[2].Status).To(Equal(metav1.ConditionFalse))
				Expect(out[3].Status).To(Equal(metav1.ConditionTrue))
			})
		})
		Context("terraform not running, plan up to date, apply noy up to date, terraform has not failed", func() {
			It("", func() {
				t.Annotations[annotations.LastPlanDate] = strconv.Itoa(int(time.Now().Add(-time.Minute * 15).Unix()))
				t.Annotations[annotations.LastPlanSum] = "ThisIsAPlanArtifact"
				t.Annotations[annotations.LastApplySum] = "ThisIsAnotherPlanArtifact"
				t.Annotations[annotations.Failure] = "0"
				_, out := conditions.Evaluate(context.TODO())
				Expect(out[0].Status).To(Equal(metav1.ConditionFalse))
				Expect(out[1].Status).To(Equal(metav1.ConditionTrue))
				Expect(out[2].Status).To(Equal(metav1.ConditionFalse))
				Expect(out[3].Status).To(Equal(metav1.ConditionFalse))
			})
		})
		Context("terraform not running, plan not up to date, terraform has failed", func() {
			It("", func() {
				t.Annotations[annotations.LastPlanDate] = strconv.Itoa(int(time.Now().Add(-time.Minute * 60).Unix()))
				t.Annotations[annotations.Failure] = "1"
				_, out := conditions.Evaluate(context.TODO())
				Expect(out[0].Status).To(Equal(metav1.ConditionFalse))
				Expect(out[1].Status).To(Equal(metav1.ConditionFalse))
				Expect(out[3].Status).To(Equal(metav1.ConditionTrue))
			})
		})
		Context("terraform not running, plan not up to date, terraform hasn't failed", func() {
			It("", func() {
				t.Annotations[annotations.LastPlanDate] = strconv.Itoa(int(time.Now().Add(-time.Minute * 60).Unix()))
				t.Annotations[annotations.Failure] = "0"
				_, out := conditions.Evaluate(context.TODO())
				Expect(out[0].Status).To(Equal(metav1.ConditionFalse))
				Expect(out[1].Status).To(Equal(metav1.ConditionFalse))
				Expect(out[3].Status).To(Equal(metav1.ConditionFalse))
			})
		})
	})
})
