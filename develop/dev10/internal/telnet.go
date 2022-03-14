package utelnet

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	errTimeoutConnection = errors.New("go-telnet: Timeout connections")
	errEOF               = errors.New("go-telnet: EOF")
	errDisconnectPeer    = errors.New("go-telnet: Disconnect from peer")
)

// TelnetClient Структура Клиента telnet
type TelnetClient struct {
	*Config
	Err     chan error
	Conn    net.Conn
	inData  io.Reader
	outData io.Writer
}

// NewTelnetClient Создать клиент
func NewTelnetClient(c *Config, in io.Reader, out io.Writer) *TelnetClient {
	return &TelnetClient{
		Config:  c,
		Err:     make(chan error),
		inData:  in,
		outData: out,
	}
}

// GetAddress Возвращает полный адрес для подключения
func (t *TelnetClient) GetAddress() string {
	return net.JoinHostPort(t.Host, t.Port)
}

// Run Запуск telnet
func (t *TelnetClient) Run() {
	//fmt.Println(t.Timeout)
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	if err := t.Connect(); err != nil {
		log.Fatalln(errTimeoutConnection)
	}
	go func() {
		if err := t.Send(); err != nil {
			t.Err <- err
		}
	}()
	go func() {
		if err := t.Receive(); err != nil {
			t.Err <- err
		}
	}()

	select {
	case <-exit:
		t.Disconnect()
		fmt.Println("go-telnet: Exit..") // Ctrl C | Kill завершаем работу
	case err := <-t.Err:
		t.Disconnect()
		fmt.Println(err) // EOF
	}

}

// Connect Подключение к хосту
func (t *TelnetClient) Connect() error {
	fmt.Println("Try connect to " + t.GetAddress())
	conn, err := net.DialTimeout("tcp", t.GetAddress(), t.Timeout)
	if err != nil {
		time.Sleep(t.Timeout) // если адреса не существует, завершение все равно по таймауту
		return err
	}
	t.Conn = conn
	fmt.Println("Successfully connected!")
	return nil
}

// Send Отправка данных хосту
func (t *TelnetClient) Send() error {
	if _, err := io.Copy(t.Conn, t.inData); err != nil {
		return err
	}
	return errEOF
}

// Receive Получение данных от хоста
func (t *TelnetClient) Receive() error {
	if _, err := io.Copy(t.outData, t.Conn); err != nil {
		return err
	}
	return errDisconnectPeer
}

// Disconnect Закрыть сокет
func (t *TelnetClient) Disconnect() {
	if err := t.Conn.Close(); err != nil {
		log.Fatalln("Error with Disconnect!")
	}
}
