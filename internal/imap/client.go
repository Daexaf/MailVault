package imap

import (
	"crypto/tls"
	"fmt"

	"github.com/emersion/go-imap/client"
)

// type Client struct {
// 	Host string
// 	Port int

// 	Username string
// 	Password string

// 	UseSSL bool
// }

// func (c *Client) Connect() error

// func TestConnection(account *entities.MailAccount) error {

// 	address := fmt.Sprintf("%s:%d", account.Host, account.Port)

// 	c, err := client.DialTLS(address, &tls.Config{})
// 	if err != nil {
// 		return err
// 	}
// 	defer c.Logout()

// 	if err := c.Login(account.Email, account.Password); err != nil {
// 		return err
// 	}

// 	fmt.Println("✅ Connected to Gmail IMAP")

// 	return nil
// }

func TestConnection(
	host string,
	port int,
	email string,
	password string,
) error {
	address := fmt.Sprintf("%s:%d", host, port)

	c, err := client.DialTLS(address, &tls.Config{})
	if err != nil {
		return err
	}
	defer c.Logout()

	if err := c.Login(email, password); err != nil {
		return err
	}

	fmt.Println("✅ Connected to Gmail IMAP")

	return nil
}
