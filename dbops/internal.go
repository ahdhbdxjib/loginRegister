package dbops

import (
	"database/sql"
	"log"
	"strconv"
	"sync"
	"video_server/api/defs"
)

func InsertSession(sid string, ttl int64, un string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("INSERT INTO sessions(session_id, TTL, login_name) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(sid, ttlstr, un)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func DeleteSession(sid string) error {
	stmtOut, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id=?")
	if err != nil {
		return err
	}
	if _, err = stmtOut.Query(sid); err != nil {
		return err
	}
	defer stmtOut.Close()
	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("SELECT TTL, login_name FROM sessions WHERE session_id=?")
	if err != nil {
		return nil, err
	}
	var ttl string
	var uname string
	err = stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if res, err := strconv.ParseInt(ttl, 10, 64); err != nil {
		ss.TTL = res
		ss.Username = uname
	} else {
		return nil, err
	}
	defer stmtOut.Close()
	return ss, nil
}

func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil {
		return nil, err
	}
	rows, err := stmtOut.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id string
		var ttlstr string
		var login_name string
		if er := rows.Scan(&id, &ttlstr, &login_name); er != nil {
			log.Printf("retrive sessions error: %s", er)
			break
		}
		if ttl, err1 := strconv.ParseInt(ttlstr, 10, 64); err1 == nil {
			ss := &defs.SimpleSession{Username: login_name, TTL: ttl}
			m.Store(id, ss)
			log.Printf("session id: %s, ttl: %d", id, ss.TTL)
		}
	}

	return m, nil
}