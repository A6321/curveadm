/*
 *  Copyright (c) 2021 NetEase Inc.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

/*
 * Project: CurveAdm
 * Created Date: 2021-12-12
 * Author: Jingli Chen (Wine93)
 */

package module

import (
	"errors"

	ssh "github.com/melbahja/goph"
	"github.com/opencurve/curveadm/internal/log"
)

type FileManager struct {
	sshClient *ssh.Client
}

var (
	ERR_UNREACHED = errors.New("remote unreached")
)

func NewFileManager(sshClient *ssh.Client) *FileManager {
	return &FileManager{sshClient: sshClient}
}

func (f *FileManager) Upload(localPath, remotePath string) error {
	if f.sshClient == nil {
		return ERR_UNREACHED
	}

	err := f.sshClient.Upload(localPath, remotePath)
	log.SwitchLevel(err)("Upload",
		log.Field("remoteAddr", remoteAddr(f.sshClient)),
		log.Field("localPath", localPath),
		log.Field("remotePath", remotePath),
		log.Field("error", err))
	return err
}

func (f *FileManager) Download(remotePath, localPath string) error {
	if f.sshClient == nil {
		return ERR_UNREACHED
	}

	err := f.sshClient.Download(remotePath, localPath)
	log.SwitchLevel(err)("Download",
		log.Field("remoteAddr", remoteAddr(f.sshClient)),
		log.Field("remotePath", remotePath),
		log.Field("localPath", localPath),
		log.Field("error", err))
	return err
}
