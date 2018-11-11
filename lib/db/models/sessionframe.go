package models

import (
  "fmt"
  "github.com/a-rmz/private-api/lib/db"
)

type SessionFrame struct{
  SessionFrameId string `db:"session_frame_id"`
  SessionID string `db:"session_id"`
  SessionType string `db:"type"`
  StartTime string `db:"start_time"`
  EndTime string `db:"end_time"`
  Data float32 `db:"data"`
}

func InsertSessionFrame(frame SessionFrame) {
  db := db.Connect()
  defer db.Close()

  _, err := db.NamedExec(`
    INSERT INTO session_frames
    (session_frame_id, session_id, type, start_time, end_time, data)
    VALUES
    (:session_frame_id, :session_id, :type, to_timestamp(:start_time), to_timestamp(:end_time), :data);
  `, frame);
  if err != nil {
    fmt.Println(err)
  }
}

