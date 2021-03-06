package serviceaccount

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/client-go/kubernetes"

	"github.com/previousnext/terraform-provider-k8s/utils/id"
)

func resourceCreate(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	serviceaccount, err := generateServiceAccount(d)
	if err != nil {
		return err
	}

	out, err := conn.CoreV1().ServiceAccounts(serviceaccount.ObjectMeta.Namespace).Create(&serviceaccount)
	if err != nil {
		return fmt.Errorf("failed to create core/v1/serviceaccount: %s", err)
	}

	d.SetId(id.Join(out.ObjectMeta))

	return nil
}
