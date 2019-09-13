package session
import(
	"fmt"
)


var (
	sessionMgr SessionMgr
)
//provider !:memory 2.redis
func Init(provider string, addr string,options ...string)(err error){

	switch (provider){
	case "memory":
		sessionMgr = NewMemorySessionMgr()
	case "redis":
		sessionMgr = NewRedisSessionMgr()
	default:
		err = fmt.Errorf("not support")
		return 
	}

	err = sessionMgr.Init(addr, options...)
	return 
}