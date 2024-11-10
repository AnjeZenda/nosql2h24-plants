// Copyright (c) 2023-2024, KNS Group LLC ("YADRO").
// All Rights Reserved.
// This software contains the intellectual property of YADRO
// or is licensed to YADRO from third parties. Use of this
// software and the intellectual property contained therein is expressly
// limited to the terms and conditions of the License Agreement under which
// it is provided by YADRO.
//

package storage

import "errors"

// User errors
var (
	ErrUserInvalidId = errors.New("invalid id")
	ErrUserNotFound  = errors.New("user not found")
)
