import React, { Component } from 'react'
import { Button, Table, Input, Icon, Form, Select } from 'antd'
import { confirmRules, getAllAlerts } from '@apis/alerts'
import moment from 'moment'
import { withRouter, Link } from 'react-router-dom'
import { tableTimeWidth } from '@configs/const'
import { connect } from 'react-redux'
import AlertConfirmModal from './alerts_confirm/alerts_confirm'

const { Option } = Select

@withRouter
@Form.create({})
@connect(state => ({
  loginData: state.loginData,
}))
export default class AlertsConfirm extends Component {
  state = {
    dataSource: [],
    keys: [],
    filterItem: {
      summary: false,
    },
    keyList: [],
    valueList: [],
  }
  valueObj = {}
  componentDidMount() {
    const { id } = this.props.match.params
    this.paramsId = id
    this.getList()
  }
  /* list */
  getList() {
    getAllAlerts({}, (res) => {
      this.data = res
      this.setList()
      this.setState({
        keyList: Object.keys(res || {}),
      })
    })
  }
  setList() {
    const list = []
    const keyObj = {}
    const { form } = this.props
    const key = form.getFieldValue('key')
    if (key) {
      // 确认完报警之后key值依然存在
      if (this.data.hasOwnProperty(key)) {
        this.valueObj = this.data[key]
        this.setState({
          valueList: Object.keys(this.valueObj),
        })
      } else {
        window.location.reload()
      }
      this.handleQuery()
      // 清空选项
      this.setState({
        keys: [],
      })
    } else {
      Object.keys(this.data).forEach((k) => {
        Object.keys(this.data[k]).forEach((value) => {
          const items = this.data[k][value]
          items.forEach((item) => {
            if (keyObj[item.id] === undefined) {
              list.push(item)
              keyObj[item.id] = item.id
            }
          })
        })
      })
      this.setState({
        dataSource: list.sort((a, b) => b.id - a.id),
      })
    }
  }
  handleConfirm = () => {
    const { keys } = this.state
    this.alertConfirmModal.updateValue(keys)
  }
  rowSelection = {
    onChange: (selectedRowKeys) => {
      this.setState({
        keys: selectedRowKeys,
      })
    },
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
  getColumnSearchProps = dataIndex => ({
    filterDropdown: ({ setSelectedKeys, selectedKeys, confirm }) => (
      <div style={{ padding: 8 }}>
        <Input
          ref={node => this.searchInput = node}
          placeholder={`Search ${dataIndex}`}
          value={selectedKeys[0]}
          onInput={(e) => {
            setSelectedKeys(e.target.value ? [e.target.value] : []);
            this.handleSearch(selectedKeys, confirm, dataIndex)
          }}
          onBlur={() => this.setState(state => ({
            filterItem: { ...state.filterItem, [dataIndex]: false },
          }))}
          style={{ width: 188, marginBottom: 8, display: 'block' }}
        />
      </div>
    ),
    filterIcon: filtered => (
      <Icon type="search"
        onMouseDown={() => {
          this.setState(state => ({
            filterItem: { ...state.filterItem, [dataIndex]: true },
          }))
          setTimeout(() => this.searchInput.focus());
        }}
        style={{ color: filtered ? '#1890ff' : undefined }}
      />
    ),
    onFilter: (value, record) =>
      record[dataIndex]
        .toString()
        .toLowerCase()
        .includes(value.toLowerCase()),

    onFilterDropdownVisibleChange: (visible) => {
      if (visible) {
        setTimeout(() => this.searchInput.focus());
      }
    },
  })
  handleSearch = (selectedKeys, confirm, dataIndex) => {
    confirm();
  }
  keyChange = (key) => {
    this.valueObj = this.data[key]
    const { setFieldsValue } = this.props.form
    setFieldsValue({ value: undefined })
    this.setState({
      valueList: Object.keys(this.valueObj),
    })
  }
  handleQuery = (e) => {
    e && e.preventDefault()
    this.props.form.validateFields((err, values) => {
      if (!err) {
        const { value } = values
        if (value) {
          if (this.valueObj[value]) {
            this.setState({
              dataSource: this.valueObj[value].sort((a, b) => b.id - a.id),
            })
          } else {
            window.location.reload()
          }
        } else {
          const list = Object.keys(this.valueObj).reduce((acc, cur) => acc.concat(this.valueObj[cur]), [])
          this.setState({
            dataSource: list.sort((a, b) => b.id - a.id),
          })
        }
      }
    })
  }
  render() {
    const { dataSource, keys, keyList, valueList } = this.state
    const { getFieldDecorator } = this.props.form
    const columns = [
      { title: 'ID', align: 'center', dataIndex: 'id', sorter: (a, b) => a.id - b.id },
      {
        title: 'Rule ID',
        align: 'center',
        dataIndex: 'rule_id',
        render: ruleId => (<Link to={`/alerts_confirm/${ruleId}`}>{ruleId}</Link>),
      },
      { title: '报警值', align: 'center', dataIndex: 'value' },
      {
        title: '当前状态',
        align: 'center',
        dataIndex: 'status',
        filters: [{ text: '报警', value: 2 }, { text: '恢复', value: 0 }],
        onFilter: (value, record) => record.status === value,
        render: status => (
          <span>{status === 2 ? '报警' : '恢复'}</span>
        ),
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
        ...this.getColumnSearchProps('summary'),
        filterDropdownVisible: this.state.filterItem.summary,
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
        <Form layout="inline" onSubmit={this.handleQuery}>
          <Form.Item>
            {getFieldDecorator('key', {
              rules: [
                { required: true, message: '请输入key' },
              ],
            })(<Select style={{ width: 180 }} placeholder="请选择 key" onChange={this.keyChange}>
              {
                keyList && keyList.map(key => <Option value={key}>{key}</Option>)
              }
            </Select>)}
          </Form.Item>
          <Form.Item>
            {getFieldDecorator('value', {})(<Select style={{ width: 180 }} disable={!valueList.length} placeholder="请选择 value">
              {
                valueList && valueList.map(value => <Option value={value}>{value}</Option>)
              }
            </Select>)}
          </Form.Item>
          <Form.Item>
            <Button type="primary" htmlType="submit">查询</Button>
            <Button type="primary" style={{ marginLeft: 15 }} onClick={this.handleConfirm} disabled={!keys.length}>确认报警</Button>
          </Form.Item>
        </Form>
        <Table scroll={{ x: 1300 }} rowSelection={this.rowSelection} dataSource={dataSource} columns={columns} rowKey="id" />
        <AlertConfirmModal OnRef={c => this.onRef(c)} onSubmit={this.updatePromethus} />
      </div>
    )
  }
}
