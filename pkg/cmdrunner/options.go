// Copyright 2020 The Atlas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmdrunner

import (
	"io"
	"time"

	exe "github.com/go-cmd/cmd"
)

// Options
type Options struct {
	// Command has timeout to complete, otherwise it will be stopped
	Timeout time.Duration
	// Command logs it output each tick
	Tick time.Duration
	// Refers to github.com/go-cmd/cmd.Options.Buffered
	Buffered bool
	// Refers to github.com/go-cmd/cmd.Options.Streaming
	Streaming bool
	// Where to write stdout
	StdoutWriter io.Writer
	// Where to write stderr
	StderrWriter io.Writer
}

// GetTimeout gets timeout from options
func (opts *Options) GetTimeout() time.Duration {
	if opts == nil {
		return 0
	}

	return opts.Timeout
}

// HasTimeout checks whether timeout is specified
func (opts *Options) HasTimeout() bool {
	if opts == nil {
		return false
	}

	return opts.GetTimeout() > 0
}

// GetTick gets tick from options
func (opts *Options) GetTick() time.Duration {
	if opts == nil {
		return 0
	}

	return opts.Tick
}

// HasTick checks whether tick is specified
func (opts *Options) HasTick() bool {
	if opts == nil {
		return false
	}

	return opts.GetTick() > 0
}

// GetOptions cast options to github.com/go-cmd/cmd.Options
func (opts *Options) GetOptions() exe.Options {
	if opts == nil {
		// Use default options
		return exe.Options{
			Buffered: true,
		}
	}

	// Cast options to github.com/go-cmd/cmd.Options
	return exe.Options{
		Buffered:  opts.Buffered,
		Streaming: opts.Streaming,
	}
}

// GetStdoutWriter
func (opts *Options) GetStdoutWriter() io.Writer {
	if opts == nil {
		return nil
	}

	return opts.StdoutWriter
}

// GetStderrWriter
func (opts *Options) GetStderrWriter() io.Writer {
	if opts == nil {
		return nil
	}

	return opts.StderrWriter
}
