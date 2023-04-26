package cloud

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AliSMS struct {
	AccessKeyId     string
	AccessKeySecret string
	Endpoint        string
	SignName        string
	TemplateCode    string
	TemplateParam   string
}

func (m *AliSMS) Send(PhoneNumbers string, TemplateParam string) error {
	config := &openapi.Config{
		AccessKeyId:     &m.AccessKeyId,
		AccessKeySecret: &m.AccessKeySecret,
	}
	config.Endpoint = tea.String(m.Endpoint)
	client, err := dysmsapi.NewClient(config)
	if err != nil {
		return err
	}
	sendSmsRequest := &dysmsapi.SendSmsRequest{
		PhoneNumbers:  tea.String(PhoneNumbers),
		SignName:      tea.String(m.SignName),
		TemplateCode:  tea.String(m.TemplateCode),
		TemplateParam: tea.String(TemplateParam),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		_, _err := client.SendSmsWithOptions(sendSmsRequest, runtime)
		if _err != nil {
			return _err
		}
		return nil
	}()
	if tryErr != nil {
		var err = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			err = _t
		} else {
			err.Message = tea.String(tryErr.Error())
		}
		return err
	}
	return nil
}
