package configs

import (
	"fmt"
	"os"
	"strings"

	"github.com/HikariKnight/quickpassthrough/pkg/fileio"
)

func Set_Dracut() {
	config := GetConfig()

	// Set the dracut config file
	dracutConf := fmt.Sprintf("%s/vfio.conf", config.Path.DRACUT)

	// If the file already exists then delete it
	if fileio.FileExist(dracutConf) {
		os.Remove(dracutConf)
	}

	// Write the dracut config file
	fileio.AppendContent(fmt.Sprintf("add_drivers+=\" %s \"\n", strings.Join(vfio_modules(), " ")), dracutConf)

	// Add to our kernel arguments file that vfio_pci should load early (dracut does this using kernel arguments)
	fileio.AppendContent(" rd.driver.pre=vfio_pci", config.Path.CMDLINE)
}
