// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin,amd64,!go1.12

package unix

//sys  Getdirentries(fd int, buf []byte, basep *uintptr) (n int, err error) = SYS_GETDIRENTRIES64
