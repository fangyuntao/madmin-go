//
// Copyright (c) 2015-2024 MinIO, Inc.
//
// This file is part of MinIO Object Storage stack
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.
//

package madmin

// Args - defines the arguments for the API.
type logArgs struct {
	Bucket   string            `json:"bucket,omitempty"`
	Object   string            `json:"object,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

// Trace - defines the trace.
type logTrace struct {
	Message   string                 `json:"message,omitempty"`
	Source    []string               `json:"source,omitempty"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

// API - defines the api type and its args.
type logAPI struct {
	Name string   `json:"name,omitempty"`
	Args *logArgs `json:"args,omitempty"`
}

// Entry - defines fields and values of each log entry.
type logEntry struct {
	DeploymentID string    `json:"deploymentid,omitempty"`
	Level        string    `json:"level"`
	LogKind      LogKind   `json:"errKind"`
	Time         string    `json:"time"`
	API          *logAPI   `json:"api,omitempty"`
	RemoteHost   string    `json:"remotehost,omitempty"`
	Host         string    `json:"host,omitempty"`
	RequestID    string    `json:"requestID,omitempty"`
	UserAgent    string    `json:"userAgent,omitempty"`
	Message      string    `json:"message,omitempty"`
	Trace        *logTrace `json:"error,omitempty"`
}
