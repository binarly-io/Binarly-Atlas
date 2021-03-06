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

package clickhouse

import (
	"fmt"
	"github.com/binarly-io/atlas/pkg/api/atlas"

	_ "github.com/mailru/go-clickhouse"
	log "github.com/sirupsen/logrus"

	"github.com/binarly-io/atlas/pkg/config/sections"
	"github.com/binarly-io/atlas/pkg/journal"
	"github.com/binarly-io/atlas/pkg/journal/adapters/clickhouse"
)

// JournalClickHouse
type JournalClickHouse struct {
	journal.BaseJournal
}

// Validate interface compatibility
var _ journal.Journaller = &JournalClickHouse{}

// NewJournalClickHouseConfig
func NewJournalClickHouseConfig(cfg sections.ClickHouseConfigurator, endpointID int32, endpointInstanceID *atlas.UUID) (*JournalClickHouse, error) {
	dsn := cfg.GetClickHouseEndpoint()
	return NewJournalClickHouse(dsn, endpointID, endpointInstanceID)
}

// NewJournalClickHouse
func NewJournalClickHouse(dsn string, endpointID int32, endpointInstanceID *atlas.UUID) (*JournalClickHouse, error) {
	if dsn == "" {
		str := "ClickHouse address in Config is empty"
		log.Errorf(str)
		return nil, fmt.Errorf(str)
	}
	adapter, err := clickhouse.NewAdapterClickHouse(dsn)
	if err != nil {
		return nil, err
	}
	journal, err := journal.NewBaseJournal(endpointID, endpointInstanceID, adapter)
	if err != nil {
		return nil, err
	}
	return &JournalClickHouse{
		BaseJournal: *journal,
	}, nil
}
