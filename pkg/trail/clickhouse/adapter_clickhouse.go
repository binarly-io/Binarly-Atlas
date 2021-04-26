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
	databasesql "database/sql"
	"fmt"

	"github.com/MakeNowJust/heredoc"
	_ "github.com/mailru/go-clickhouse"
	log "github.com/sirupsen/logrus"

	"github.com/binarly-io/atlas/pkg/config/interfaces"
	"github.com/binarly-io/atlas/pkg/trail"
)

// JournalClickHouse
type AdapterClickHouse struct {
	connect *databasesql.DB
}

// Validate interface compatibility
var _ trail.Adapter = &AdapterClickHouse{}

// NewAdapterClickHouseConfig
func NewAdapterClickHouseConfig(cfg interfaces.ClickHouseConfigurator) (*AdapterClickHouse, error) {
	dsn := cfg.GetClickHouseEndpoint()
	return NewAdapterClickHouse(dsn)
}

// NewAdapterClickHouse
func NewAdapterClickHouse(dsn string) (*AdapterClickHouse, error) {
	if dsn == "" {
		str := "ClickHouse address in Config is empty"
		log.Errorf(str)
		return nil, fmt.Errorf(str)
	}

	log.Infof("connect to ClickHouse %s", dsn)

	connect, err := databasesql.Open("clickhouse", dsn)
	if err != nil {
		log.Errorf("unable to open ClickHouse err: %v", err)
		return nil, err
	}

	if err := connect.Ping(); err != nil {
		log.Errorf("unable to ping ClickHouse. err: %v", err)
		return nil, err
	}

	return &AdapterClickHouse{
		connect: connect,
	}, nil
}

// Insert
func (j *AdapterClickHouse) Insert(entry *trail.JournalEntry) error {
	e := NewAdapterEntryClickHouse().Import(entry)

	sql := heredoc.Doc(
		fmt.Sprintf(`
			INSERT INTO api_journal (
				%s
			) VALUES (
				%s
			)
			`,
			e.Fields(),
			e.StmtParamsPlaceholder(),
		),
	)

	tx, err := j.connect.Begin()
	if err != nil {
		log.Errorf("unable to begin tx. err: %v", err)
		return err
	}

	stmt, err := tx.Prepare(sql)
	if err != nil {
		log.Errorf("unable to prepare stmt. err: %v", err)
		return err
	}

	if _, err := stmt.Exec(e.AsUntypedSlice()...); err != nil {
		log.Errorf("exec failed. err: %v", err)
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("commit failed. err %v", err)
		return err
	}

	return nil
}

// FindAll
func (j *AdapterClickHouse) FindAll(entry *trail.JournalEntry) ([]*trail.JournalEntry, error) {
	e := NewAdapterEntryClickHouseSearch().Import(entry)
	placeholder, args := e.StmtSearchParamsPlaceholderAndArgs()
	sql := heredoc.Doc(
		fmt.Sprintf(`
			SELECT * FROM api_journal WHERE (1 == 1) %s 
			`,
			placeholder,
		),
	)

	stmt, err := j.connect.Prepare(sql)
	if err != nil {
		log.Errorf("unable to prepare stmt. err: %v", err)
		return nil, err
	}

	rows, err := stmt.Query(args...)
	if err != nil {
		log.Errorf("unable to query stmt. err: %v", err)
		return nil, err
	}
	defer rows.Close()

	var res []*trail.JournalEntry
	for rows.Next() {
		var ce AdapterEntryClickHouse
		if err := rows.Scan(
			&ce.d,
			&ce.endpointID,
			&ce.sourceID,
			&ce.contextID,
			&ce.taskID,
			&ce.typeID,
			&ce.duration,
			&ce._type,
			&ce.size,
			&ce.address,
			&ce.name,
			&ce.digest,
			&ce.data,
			&ce.error,
		); err == nil {
			res = append(res, ce.Export())
		}
	}

	return res, nil
}