import React, { Component } from 'react'
import { Route, Switch, Redirect, withRouter } from 'react-router-dom'
import { ConfigProvider, DatePicker, message } from 'antd'
import enUS from 'antd/es/locale/en_US'
import zhCN from 'antd/es/locale/zh_CN'
import '@/styles/base.less'
import Rules from '@content/rules'
import Promethus from '@content/promethus'
import Strategy from '@content/strategy'
import Alerts from '@content/alerts'
import AlertsConfirm from '@content/alerts_confirm'
import AlertsConfirmByID from '@content/alerts_confirm_id'
import Group from '@content/group'
import User from '@content/user'
import Maintain from '@content/maintain'
import { CN, EN } from '@lang/lang'
import { GlobalClass } from '@configs/common'
import { connect } from 'react-redux'
import Login from './app/login'
import Header from './app/header'
import Sider from './app/sider'
import { ThemeContext } from './context'

@connect((state, props) => ({
  langInfo: state.langInfo,
}))
@withRouter
export default class App extends Component {
  componentDidMount() {
  }
  render() {
    const { path } = this.props.match
    const { history } = this.props
    const { lang } = this.props.langInfo
    // 挂在 histroy 到全局
    GlobalClass.history = history
    return (
      <ConfigProvider locale={lang === CN ? zhCN : enUS}>
        <Switch>
          <Route path={`${path}login`} component={Login} />
          <div id="container">
            <ThemeContext.Provider value="dark">
              <Header />
            </ThemeContext.Provider>
            <section id="content">
              <Sider />
              <div id="main">
                <div id="main-box">
                  <Route exact path={`${path}`}>
                    <Redirect to={`${path}rules`} />
                  </Route>
                  <Route exact path={`${path}rules`} component={Rules} />
                  <Route path={`${path}rules/:id`} component={Rules} />
                  <Route path={`${path}promethus`} component={Promethus} />
                  <Route path={`${path}strategy`} component={Strategy} />
                  <Route path={`${path}alerts`} component={Alerts} />
                  <Route exact path={`${path}alerts_confirm`} component={AlertsConfirm} />
                  <Route path={`${path}alerts_confirm/:id`} component={AlertsConfirmByID} />
                  <Route path={`${path}group`} component={Group} />
                  <Route path={`${path}maintain`} component={Maintain} />
                  <Route path={`${path}user`} component={User} />
                </div>
              </div>
            </section>
          </div>
        </Switch>
      </ConfigProvider>
    );
  }
}
