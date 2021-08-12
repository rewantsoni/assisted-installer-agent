package inventory

import (
	"encoding/json"
	"fmt"
	"github.com/openshift/assisted-installer-agent/src/util"
	"github.com/openshift/assisted-service/models"
)

func ReadInventory() *models.Inventory {
	d := util.NewDependencies()
	fmt.Println("Getting the disks")
	ret := models.Inventory{
		Disks:        GetDisks(d),
	}
	return &ret
}

func CreateInventoryInfo() []byte {
	in := ReadInventory()
	b, _ := json.Marshal(&in)
	return b
}
