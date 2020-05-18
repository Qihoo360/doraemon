import React, { Component } from 'react'
import { connect } from 'react-redux'
import { withRouter } from 'react-router'
import { Row, Col, Form, message, Input, Button, Icon } from 'antd'
import '@/styles/login.less'
import { parseQueryString } from '@configs/common'
import { localLoginApi, ldapLoginApi, oauthLoginApi } from '@apis/login'
import { loginAction } from '@actions/common'

const layout = {
  labelCol: { span: 6 },
  wrapperCol: { span: 16 },
};

@withRouter
@connect(state => ({
  loginData: state.loginData,
}), dispatch => ({
  loginAction: c => dispatch(loginAction(c)),
}))
@Form.create({})
export default class Login extends Component {
  componentDidMount() {
    const { location } = this.props.history
    this.query = parseQueryString(location.search)
  }
  state = {
    chooseMethod: 'none',
    defaultName: '',
  }
  onFinish = (method) => {
    const { chooseMethod } = this.state
    const { history } = this.props
    const state = this.query.state ? this.query.state : encodeURIComponent(window.location.origin)
    if (!method) {
      // 直接登录
      history.push('/')
      return
    }
    if (method === 'oauth') {
      oauthLoginApi({}, (res) => {
        console.log(state, res)
        window.location.href = `${res}&state=${state}`
      })
      return
    }
    if (chooseMethod === 'none') {
      this.setState({
        chooseMethod: method,
      })
      return
    }
    if (chooseMethod !== 'oauth') {
      this.props.form.validateFields(async (err, values) => {
        if (!err) {
          const { username, password } = values
          const handleFn = chooseMethod === 'local' ? localLoginApi : ldapLoginApi
          handleFn({ username, password }, () => {
            this.props.loginAction()
            // 存在 hash，所以选择直接跳转
            window.location.href = decodeURIComponent(state)
          }, (error) => {
            message.error(error)
          })
        }
      })
    }
  }
  chooseOther() {
    this.setState({
      chooseMethod: 'none',
    })
  }
  render() {
    const { data } = this.props.loginData
    const { getFieldDecorator } = this.props.form
    const { chooseMethod, defaultName } = this.state
    return (
      <div className="login-container">
        <div className="login-bg" />
        <div className="login-content">
          <div className="login-title">登&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;录</div>
          <Form
            {...layout}
            name="basic"
          >
            {
              (chooseMethod === 'local' || chooseMethod === 'ldap') && (<div>
                <Form.Item label="用户名">
                  {getFieldDecorator('username', {
                    initialValue: defaultName,
                    rules: [{ required: true, message: 'Please input your name!' }],
                  })(<Input
                    prefix={<Icon type="user" style={{ color: 'rgba(0,0,0,.25)' }} />}
                    placeholder="Username"
                  />)}
                </Form.Item>
                <Form.Item label="密码">
                  {getFieldDecorator('password', {
                    rules: [{ required: true, message: 'Please input your Password!' }],
                  })(<Input
                    prefix={<Icon type="lock" style={{ color: 'rgba(0,0,0,.25)' }} />}
                    type="password"
                    placeholder="Password"
                  />)}
                </Form.Item>
                <div style={{ height: '20px' }} />
              </div>)
            }
            <Form.Item
              wrapperCol={{ offset: 2, span: 20 }}
            >
              {chooseMethod === 'none' && <p className="choose-text">请选择登录方式 ：</p>}
              {
                (data && chooseMethod === 'none') && (<Button type="primary" block onClick={() => this.onFinish()} style={{ marginBottom: '15px' }}>
                已有账号，直接登录<Icon type="right" />
                </Button>)
              }
              <Row type="flex" justify="space-between" style={{ color: 'white' }}>
                <Col span={chooseMethod === 'none' ? 6 : 24} className={(chooseMethod !== 'none' && chooseMethod !== 'local') ? 'hide' : ''}>
                  <Button type="primary" block onClick={() => this.onFinish('local')} className="login-form-button">
                    本地
                  </Button>
                </Col>
                <Col span={chooseMethod === 'none' ? 6 : 24} className={chooseMethod !== 'none' && chooseMethod !== 'ldap' ? 'hide' : ''}>
                  <Button type="primary" block onClick={() => this.onFinish('ldap')} className="login-form-button">
                    LDAP
                  </Button>
                </Col>
                <Col span={chooseMethod === 'none' ? 6 : 24} className={chooseMethod !== 'none' && chooseMethod !== 'oauth' ? 'hide' : ''}>
                  <Button type="primary" block onClick={() => this.onFinish('oauth')} className="login-form-button">
                    OAuth
                  </Button>
                </Col>
                <Col style={{ fontSize: '12px' }} className={chooseMethod === 'none' ? 'hide' : ''}>
                  <a onClick={() => this.chooseOther()}>选择其他登录方式<Icon type="right" /></a>
                </Col>
              </Row>
            </Form.Item>
          </Form>
        </div>
      </div>
    )
  }
}
