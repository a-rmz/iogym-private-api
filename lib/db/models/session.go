package models

import (
  "database/sql"
  "github.com/a-rmz/private-api/lib/db"
  "fmt"
  "errors"
)

type Session struct{
  SessionID string `db:"session_id"`
  Description sql.NullString `db:"description"`
  StartTime string `db:"start_time"`
  EndTime sql.NullString `db:"end_time"`
  UserRFID string `db:"user_rfid_code"`
  DeviceID string `db:"device_id"`
}

func GetSessionByRFIDAndDevice(rfid string, device string) (Session, error) {
  var session Session
  params := map[string]interface{}{"rfid": rfid, "device": device}

  rows := db.QueryOne(`
    SELECT *
    FROM sessions
    WHERE user_rfid_code = :rfid
    AND device_id = :device
    AND end_time IS NULL
    ORDER BY start_time DESC
    LIMIT 1;
  `, params);

  for rows.Next() {
    switch err := rows.StructScan(&session); err {
    case sql.ErrNoRows:
      return session, errors.New("Session was not found")
    case nil:
      return session, nil
    default:
      panic(err)
    }
  }
  return session, nil
}

func InsertSession(session Session) {
  db := db.Connect()
  defer db.Close()

  _, err := db.NamedExec(`
    INSERT INTO sessions
    (start_time, user_rfid_code, device_id)
    VALUES
    (to_timestamp(:start_time), :user_rfid_code, :device_id);
  `, session);
  if err != nil {
    fmt.Println(err)
  }
}

func UpdateSessionEndTime(userId string, deviceId string, endTime string) {
  db := db.Connect()
  defer db.Close()

params := map[string]interface{}{"user_rfid_code": userId, "device_id": deviceId, "end_time": endTime}
  _, err := db.NamedExec(`
    UPDATE sessions
    SET end_time = to_timestamp(:end_time)
    WHERE user_rfid_code = :user_rfid_code
    AND device_id = :device_id
    AND end_time IS NULL;
  `, params);
  if err != nil {
    fmt.Println(err)
  }
}

