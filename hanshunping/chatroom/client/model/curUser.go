<<<<<<< HEAD
package model

import (
	"net"
	"go_code/chatroom/common/message"
)
//因为在客户端，我们很多地方会使用到curUser,我们将其作为一个全局
type CurUser struct {
	Conn net.Conn
	message.User
} 
=======
package model

import (
	"go_code/chatroom/common/message"
	"net"
)

//因为在客户端，我们很多地方会使用到curUser,我们将其作为一个全局
type CurUser struct {
	Conn net.Conn
	message.User
}
>>>>>>> 34c4b0b7b2d411ca07c05d09bb616838a7be8dab
