package session

type SessionMgr interface {
	Init(addr string, options ...string) (err error)
	CreateSession() (session Session, err error)
	Get(sessionId string) (sesssion Session, err error)
}
