// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2022 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package configcore

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/snapcore/snapd/dirs"
	"github.com/snapcore/snapd/overlord/configstate/config"
	"github.com/snapcore/snapd/sysconfig"
)

func init() {
	supportedConfigurations["core.system.faillock"] = true
}

func validateFaillockSettings(tr config.ConfGetter) error {
	return validateBoolFlag(tr, "system.faillock")
}

func handleFaillockConfiguration(dev sysconfig.Device, tr config.ConfGetter, opts *fsOnlyContext) error {
	faillock, err := coreCfg(tr, "system.faillock")
	if err != nil {
		return err
	}

	marker := filepath.Join(dirs.GlobalRootDir, "/etc/writable/faillock.enabled")

	switch faillock {
	case "":
		// nothing to do if unset
	case "true":
		if err := ioutil.WriteFile(marker, nil, 0644); err != nil {
			return err
		}
	case "false":
		if err := os.Remove(marker); err != nil && !os.IsNotExist(err) {
			return err
		}
	default:
		return fmt.Errorf("unsupported system.faillock value: %q", faillock)
	}

	return nil
}
