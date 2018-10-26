func main() {
    m := gomail.NewMessage()

    m.SetAddressHeader("From", "xxx@foxmail.com" /*"发件人地址"*/, "发件人") // 发件人

    m.SetHeader("To",                                                            
        m.FormatAddress("xxxx@qq.com", "收件人")) // 收件人
     m.SetHeader("Cc",
        m.FormatAddress("xxxx@foxmail.com", "收件人")) //抄送
    m.SetHeader("Bcc",  
        m.FormatAddress("xxxx@gmail.com", "收件人")) /暗送

    m.SetHeader("Subject", "liic测试")     // 主题

    //m.SetBody("text/html",xxxxx ") // 可以放html..还有其他的
    m.SetBody("我是正文") // 正文

    m.Attach("我是附件")  //添加附件

    d := gomail.NewPlainDialer("smtp.qq.com", 465, "xxx@foxmail.com", "发件密码") // 发送邮件服务器、端口、发件人账号、发件人密码
    if err := d.DialAndSend(m); err != nil {
        log.Println("发送失败", err)
        return
    }

    log.Println("done.发送成功")
}

//很多邮件发送都需要先开启IMAP(QQ邮箱为例 通过：设置->账户)，需要注意的事qq邮箱来发送不是使用的邮箱密码，开启IMAP后会给一个授权码 
