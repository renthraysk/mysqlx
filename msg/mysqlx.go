package msg

const namespace = "mysqlx"

func Ping(buf []byte) Msg {
	s := NewStmtExecute(buf, "ping")
	s.SetNamespace(namespace)
	return s
}

func ListClients(buf []byte) Msg {
	s := NewStmtExecute(buf, "list_clients")
	s.SetNamespace(namespace)
	return s
}

func KillClient(buf []byte, id uint64) Msg {
	s := NewStmtExecute(buf, "kill_client")
	s.SetNamespace(namespace)
	s.AppendArgUint(id)
	return s
}

func ResetGlobalVariables(buf []byte) Msg {
	return NewStmtExecute(buf, "SELECT mysqlx_reset_global_status_variables()")
}

// Notices

type Notice string

const (
	NoticeWarnings          Notice = "warnings"
	NoticeAccountExpired    Notice = "account_expired"
	NoticeGeneratedInsertID Notice = "generated_insert_id"
	NoticeRowsAffected      Notice = "rows_affected"
	NoticeProcedMessage     Notice = "produced_messages"
)

func ListNotices(buf []byte) Msg {
	s := NewStmtExecute(buf, "list_notices")
	s.SetNamespace(namespace)
	return s
}

func DisableNotices(buf []byte, notice Notice) Msg {
	s := NewStmtExecute(buf, "disable_notices")
	s.SetNamespace(namespace)
	s.AppendArgString(string(notice), 0)
	return s
}

func EnableNotices(buf []byte, notice Notice) Msg {
	s := NewStmtExecute(buf, "enable_notices")
	s.SetNamespace(namespace)
	s.AppendArgString(string(notice), 0)
	return s
}

// Collection

func CreateCollection(buf []byte, database, name string) Msg {
	s := NewStmtExecute(buf, "create_collection")
	s.SetNamespace(namespace)
	s.AppendArgString(database, 0)
	s.AppendArgString(name, 0)
	return s
}

func CreateCollectionIndex(buf []byte) Msg {
	s := NewStmtExecute(buf, "create_collection_index")
	s.SetNamespace(namespace)
	// @TODO
	return s
}

func ListObjects(buf []byte, database, like string) Msg {
	s := NewStmtExecute(buf, "list_objects")
	s.SetNamespace(namespace)
	if database != "" {
		s.AppendArgString(database, 0)
		if like != "" {
			s.AppendArgString(like, 0)
		}
	}
	return s
}
