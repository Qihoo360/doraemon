import React, { Component } from 'react'
import { BrowserRouter, HashRouter, Route, Switch } from 'react-router-dom'
import Base from '@/page/base/app' // 基础
import { Spin } from 'antd'
import { loginAction } from '@actions/common'
import { connect } from 'react-redux'

export const Router = BrowserRouter

@connect((state, props) => ({
  loginData: state.loginData,
}), dispatch => ({
  loginAction: () => dispatch(loginAction()),
}))
export default class Routes extends Component {
  componentDidMount() {
    this.props.loginAction()
  }
  render() {
    const { loaded } = this.props.loginData
    return (
      <Router>
        <Switch>
          <Route path="/"
            render={routeProps => (
              loaded ? <Base /> : <div id="loading-box"><Spin size="large" /></div>
            )}
          />
        </Switch>
      </Router>
    )
  }
}
