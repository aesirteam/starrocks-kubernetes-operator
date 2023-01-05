/*
Copyright 2021-present, StarRocks Inc.

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

package fe_controller

import (
	srapi "github.com/StarRocks/starrocks-kubernetes-operator/api/v1alpha1"
	rutils "github.com/StarRocks/starrocks-kubernetes-operator/pkg/common/resource_utils"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/klog/v2"
)

func feStatefulSetsLabels(src *srapi.StarRocksCluster) rutils.Labels {
	labels := rutils.Labels{}
	labels[srapi.OwnerReference] = src.Name
	labels[srapi.ComponentLabelKey] = srapi.DEFAULT_FE
	labels.AddLabel(src.Labels)
	return labels
}

func feStatefulSetName(src *srapi.StarRocksCluster) string {
	stname := src.Name + "-" + srapi.DEFAULT_FE
	if src.Spec.StarRocksFeSpec.Name != "" {
		stname = src.Spec.StarRocksFeSpec.Name
	}
	return stname
}

// buildStatefulSetParams generate the params of construct the statefulset.
func (fc *FeController) buildStatefulSetParams(src *srapi.StarRocksCluster, feconfig map[string]interface{}) rutils.StatefulSetParams {
	feSpec := src.Spec.StarRocksFeSpec
	var pvcs []corev1.PersistentVolumeClaim
	for _, vm := range feSpec.StorageVolumes {
		pvcs = append(pvcs, corev1.PersistentVolumeClaim{
			ObjectMeta: metav1.ObjectMeta{Name: vm.Name},
			Spec: corev1.PersistentVolumeClaimSpec{
				AccessModes: []corev1.PersistentVolumeAccessMode{
					corev1.ReadWriteOnce,
				},

				StorageClassName: vm.StorageClassName,
				Resources: corev1.ResourceRequirements{
					Requests: corev1.ResourceList{
						corev1.ResourceStorage: resource.MustParse(vm.StorageSize),
					},
				},
			},
		})
	}

	str, _ := json.Marshal(pvcs)
	klog.Info("FeController buildStatefulSetParams ", string(str))
	stname := feStatefulSetName(src)
	or := metav1.NewControllerRef(src, src.GroupVersionKind())
	podTemplateSpec := fc.buildPodTemplate(src, feconfig)

	return rutils.StatefulSetParams{
		Name:                 stname,
		Namespace:            src.Namespace,
		Replicas:             feSpec.Replicas,
		Annotations:          make(map[string]string),
		VolumeClaimTemplates: pvcs,
		ServiceName:          fc.getSearchService(),
		Labels:               feStatefulSetsLabels(src),
		PodTemplateSpec:      podTemplateSpec,
		Selector:             fePodLabels(src, stname),
		OwnerReferences:      []metav1.OwnerReference{*or},
	}
}
