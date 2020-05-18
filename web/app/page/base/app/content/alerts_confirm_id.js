import React, { Component } from 'react'
import { Button, Table, Pagination, Form, Tag } from 'antd'
import { getAlertsRules, confirmRules } from '@apis/alerts'
import moment from 'moment'
import { withRouter } from 'react-router-dom'
import { connect } from 'react-redux'
import { tableTimeWidth } from '@configs/const'
import AlertConfirmModal from './alerts_confirm/alerts_confirm'

@withRouter
@Form.create({})
@connect(state => ({
  loginData: state.loginData,
}))
export default class AlertConfirmById extends Component {
  state = {
    dataSource: [],
    keys: [],
    page: {
      current: 1,
      total: 0,
      size: 10,
    },
    labalWidth: 100,
  }
  paramsId = 0
  paramsStart = -1
  componentDidMount() {
    const { id } = this.props.match.params
    // console.log(this.props.history.location.search)
    const query = this.props.location.search
    // console.log(query)
    if (query != '') {
      const p = query.substr(1).split('=')
      if (p.length == 2 && p[0] === 'start') {
        this.paramsStart = p[1]
      }
    }
    // console.log(this.paramsStart)
    this.paramsId = id
    this.getList()
  }
  getList() {
    const { page } = this.state
    if (this.paramsId) {
      getAlertsRules({}, { id: this.paramsId }, (res) => {
        const labalWidth = Math.max(((this.calcLabelWidth(res.alerts) + 1) * 6.2) + 16, 80)
        this.setState({
          dataSource: res.alerts,
          page: { ...page, total: res.total },
          labalWidth,
        })
      }, { params: { start: this.paramsStart, page: page.current, pagesize: page.size } })
    }
  }
  calcLabelWidth(data) {
    let maxLength = 0
    data.forEach((item) => {
      const { labels } = item
      Object.keys(labels || {}).forEach((key) => {
        maxLength = Math.max(maxLength, key.length + labels[key].length)
      })
    })
    return maxLength
  }
  componentDidUpdate() {
    const { id } = this.props.match.params
    if (id !== this.paramsId) {
      this.paramsId = id
      this.resetPage()
      this.getList()
    }
  }
  handleConfirm = () => {
    const { keys } = this.state
    this.alertConfirmModal.updateValue(keys)
  }
  checkAll = () => {
    const { history } = this.props
    history.push('/alerts_confirm')
  }
  rowSelection = {
    onChange: (selectedRowKeys) => {
      this.setState({
        keys: selectedRowKeys,
      })
    },
    getCheckboxProps: record => ({
      disabled: record.status === 1,
    }),
  }
  onRef(component) {
    this.alertConfirmModal = component
  }
  updatePromethus = (values) => {
    const { duration } = values
    const { keys } = this.state
    const { loginData } = this.props
    return new Promise((resolve) => {
      confirmRules({
        duration,
        ids: keys,
        user: loginData.data,
      }, (res) => {
        resolve(true)
        this.getList()
      })
    })
  }
  /* page */
  pageChange = (page) => {
    this.resetPage({ current: page })
  }
  resetPage = (pageConfig = {}) => {
    this.setState(state => ({
      page: { ...state.page, current: 1, ...pageConfig },
    }), () => this.getList())
  }
  render() {
    const { dataSource, keys, page, labalWidth } = this.state
    const { id } = this.props.match.params
    const columns = [
      { title: 'ID', align: 'center', dataIndex: 'id' },
      { title: 'Rule ID', align: 'center', dataIndex: 'rule_id' },
      { title: '报警值', align: 'center', dataIndex: 'value' },
      {
        title: '当前状态',
        align: 'center',
        dataIndex: 'status',
        render: (status) => {
          let text
          if (status === 2) {
            text = '报警'
          } else if (status === 1) {
            text = '已确认'
          } else {
            text = '恢复'
          }
          return <span>{text}</span>
        },
      },
      {
        title: '异常分钟数',
        align: 'center',
        dataIndex: 'count',
        render: count => (
          <span>{count + 1}</span>
        ),
      },
      {
        title: '标题',
        align: 'center',
        dataIndex: 'summary',
      },
      {
        title: 'label',
        align: 'center',
        dataIndex: 'labels',
        width: labalWidth,
        render: (labels) => {
          if (labels && Object.prototype.toString.call(labels) === '[object Object]') {
            return Object.keys(labels).map(key => <Tag color="cyan" style={{ marginTop: '5px' }}>{key}: {labels[key]}</Tag>)
          }
          return '-'
        },
      },
      { title: '描述', align: 'center', dataIndex: 'description' },
      { title: '确认人', align: 'center', dataIndex: 'confirmed_by' },
      {
        title: '触发时间',
        align: 'center',
        dataIndex: 'fired_at',
        width: tableTimeWidth,
        render: firedAt => (
          <span>{firedAt === '0001-01-01T00:00:00Z' ? '--' : moment(firedAt).format('YYYY.MM.DD HH:mm:ss')}</span>
        ),
      },
      {
        title: '确认时间',
        align: 'center',
        dataIndex: 'confirmed_at',
        width: tableTimeWidth,
        render: confirmedAt => (
          <span>{confirmedAt === '0001-01-01T00:00:00Z' ? '--' : moment(confirmedAt).format('YYYY.MM.DD HH:mm:ss')}</span>
        ),
      },
      {
        title: '确认截止时间',
        align: 'center',
        dataIndex: 'confirmed_before',
        width: tableTimeWidth,
        render: confirmedBefore => (
          <span>{confirmedBefore === '0001-01-01T00:00:00Z' ? '--' : moment(confirmedBefore).format('YYYY.MM.DD HH:mm:ss')}</span>
        ),
      },
      {
        title: '恢复时间',
        align: 'center',
        dataIndex: 'resolved_at',
        width: tableTimeWidth,
        render: resolvedAt => (
          <span>{resolvedAt === '0001-01-01T00:00:00Z' ? '--' : moment(resolvedAt).format('YYYY.MM.DD HH:mm:ss')}</span>
        ),
      },
    ]
    return (
      <div>
        <div id="top-section">
          <Button type="primary" onClick={this.handleConfirm} disabled={!keys.length}>确认报警</Button>
          <Button type="primary" onClick={this.checkAll}>{id ? '查看全部' : '刷新列表'}</Button>
        </div>
        <Table scroll={{ x: 1300 }} pagination={false} rowSelection={this.rowSelection} dataSource={dataSource} columns={columns} rowKey="id" />
        <div style={{ display: 'flex', justifyContent: 'flex-end', paddingTop: '15px' }}>
          <Pagination onChange={this.pageChange} current={page.current} pageSize={page.size} total={page.total} />
        </div>
        <AlertConfirmModal OnRef={c => this.onRef(c)} onSubmit={this.updatePromethus} />
      </div>
    )
  }
}
