// Copyright 2021 Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// These tests are written in BDD-style using Ginkgo framework. Refer to
// http://onsi.github.io/ginkgo to learn more.

var _ = Describe("Schedule", func() {
	var (
		key              types.NamespacedName
		created, fetched *Schedule
	)

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
	})

	Context("Create API", func() {
		It("should create an object successfully", func() {
			key = types.NamespacedName{
				Name:      "foo",
				Namespace: "default",
			}

			created = &Schedule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: ScheduleSpec{
					Schedule: "* * * * 1",
					ScheduleItem: ScheduleItem{
						EmbedChaos: EmbedChaos{PodChaos: &PodChaosSpec{
							Action: PodKillAction,
							ContainerSelector: ContainerSelector{
								PodSelector: PodSelector{
									Mode: OnePodMode,
								},
							},
						},
						},
					},
					ConcurrencyPolicy: ForbidConcurrent,
					Type:              ScheduleTypePodChaos,
				},
				Status: ScheduleStatus{
					LastScheduleTime: metav1.Time{time.Now()},
				},
			}

			By("creating an API obj")
			Expect(k8sClient.Create(context.TODO(), created)).To(Succeed())

			fetched = &Schedule{}
			Expect(k8sClient.Get(context.TODO(), key, fetched)).To(Succeed())
			Expect(fetched).To(Equal(created))

			By("deleting the created object")
			Expect(k8sClient.Delete(context.TODO(), created)).To(Succeed())
			Expect(k8sClient.Get(context.TODO(), key, created)).ToNot(Succeed())
		})
	})
})