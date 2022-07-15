package controllers

//
//import (
//	"context"
//	"fmt"
//	. "github.com/onsi/ginkgo"
//	. "github.com/onsi/gomega"
//
//	"github.com/slintes/demo-operator/api/v1alpha1"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"sigs.k8s.io/controller-runtime/pkg/client"
//)
//
//var _ = Describe("Node Health Check CR", func() {
//	Context("Finalizer", func() {
//
//		It("should work", func() {
//			hw := &v1alpha1.HelloWorld{
//				ObjectMeta: metav1.ObjectMeta{
//					Namespace: "default",
//					Name:      "test",
//				},
//				Spec: v1alpha1.HelloWorldSpec{
//					Message: "hello",
//				},
//			}
//			err := k8sClient.Create(context.Background(), hw)
//			Expect(err).ToNot(HaveOccurred())
//
//			key := client.ObjectKeyFromObject(hw)
//			//for {
//			err = k8sClient.Get(context.Background(), key, hw)
//			Expect(err).ToNot(HaveOccurred())
//			fmt.Fprintf(GinkgoWriter, "finalizers: %v\n", hw.Finalizers)
//			//time.Sleep(5 * time.Second)
//			//}
//		})
//
//	})
//
//})
