package client

import (
	"errors"
	"lcs/util"
	_"sync"
)

var (
		ERR_MSG_SEND_FAIELD = errors.New("msg send failed")
		ERR_UID_INVALID = errors.New("uid invalid")
		ERR_MSG_PASSWORD_ERROR = errors.New("password error")
	)

/*
* 客户端
 */

type Client struct {
	UserInfo   util.UserInfo
	UserMsg    util.UserMsg
	Header util.Header
}

func (c *Client) getUserMsg() (msg util.UserMsg){
	return c.UserMsg
}

func (c *Client) getUser() (util.UserInfo){
	return c.UserInfo
}

func (c *Client) getHeader() (util.Header){
	return c.Header
}

func (c *Client) setUserInfo(userInfo util.UserInfo){
	c.UserInfo = userInfo
}

func (c *Client) setUserMsg(userMsg util.UserMsg){
	c.UserMsg = userMsg
}

func (c *Client) sendHeader(header util.Header){
	c.Header = header
}



