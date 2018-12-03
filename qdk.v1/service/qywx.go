package service

type (
	QywxMgr interface {
		ReceiveMessage()
		SendMessage()
	}

	QywxService struct {
	}
)

func (q *QywxService) ReceiveMessage() {

}

func (q *QywxService) SendMessage() {

}

func NewQywxMgr() (QywxMgr, error) {
	return &QywxService{}, nil
}
