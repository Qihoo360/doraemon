
import React, { Component } from 'react'
import { Dropdown, Menu, Icon, Avatar } from 'antd'
import { connect } from 'react-redux'
import { rejectLogin } from '@actions/common'
import { changeLangAction } from '@actions/lang'
import { withRouter } from 'react-router-dom'
import { CN, EN } from '@lang/lang'
import { localSet, localGet } from '@configs/store'
import { logoff } from '@apis/login'
import { ThemeContext } from '../context'

const localLang = localGet('lang')
@withRouter
@connect(state => ({
  loginData: state.loginData,
}), dispatch => ({
  changLangAction: c => dispatch(changeLangAction(c)),
  rejectLoginAction: c => dispatch(rejectLogin(c)),
}))
export default class Header extends Component {
  static contextType = ThemeContext
  state = {
    headerItem: [],
    lang: localLang || CN,
  }
  componentDidMount() {
    this.props.changLangAction(localLang || CN)
  }
  loginOut = () => {
    const { history } = this.props
    logoff({}, () => {
      this.props.rejectLoginAction()
      history.push('/login')
    })
  }
  changeLang(lang) {
    this.setState({
      lang,
    })
    localSet('lang', lang)
    this.props.changLangAction(lang)
  }
  render() {
    console.log(this.context)
    const { headerItem, lang } = this.state
    const { loginData } = this.props
    const menu = (<Menu>
      <Menu.Item>
        <a onClick={this.loginOut}>退出登陆</a>
      </Menu.Item></Menu>)
    return (
      <div className="header">
        <div className="left-section">
          {/* left */}
          <div className="icon-container">
            <span className="header-icon" />
            <p>&nbsp; Doraemon</p>
          </div>
          <div className="func-section">
            {
              headerItem && headerItem.map(item => (
                <Dropdown overlay={this.menu}>
                  <span className="ant-dropdown-link">
                    {item}<Icon type="down" />
                  </span>
                </Dropdown>
              ))
            }
          </div>
        </div>
        <div className="right-container">
          {/* right */}
          <div className="func-section">
            <span onClick={() => this.changeLang(lang === CN ? EN : CN)}>{lang === CN ? EN : CN}</span>
            <div className="divider" />
            <Dropdown overlay={menu}>
              <span className="ant-dropdown-link text">
                <Avatar style={{ justifyContent: 'center' }} icon="user" />{loginData.data}
              </span>
            </Dropdown>
          </div>
        </div>
      </div>
    );
  }
}
