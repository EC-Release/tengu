//go:binary-only-package

/*
 * Copyright (c) 2016 General Electric Company. All rights reserved.
 *
 * The copyright to the computer software herein is the property of
 * General Electric Company. The software may be used and/or copied only
 * with the written permission of General Electric Company or in accordance
 * with the terms and conditions stipulated in the agreement/contract
 * under which the software has been supplied.
 *
 * author: apolo.yasuda@ge.com
 */

package wzapi
import (
	"archive/tar"
	"github.com/google/uuid"
	model "github.com/wzlib/wzschema"
	util "github.com/wzlib/wzutil"
	"net/http"
	"encoding/json"
	"fmt"
	"errors"
)
