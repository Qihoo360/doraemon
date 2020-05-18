/* eslint-disable no-undef */
import React, { Component } from 'react'
import { Button, Table, message, Popconfirm, Input, Icon } from 'antd'
import { getRules, addRules, updateRules, deleteRules } from '@apis/rules'
import { connect } from 'react-redux'
import Highlighter from 'react-highlight-words'
import { withRouter } from 'react-router-dom'
import { parseQueryString } from '@configs/common'
import CreateEditRules from './rules/create-edit-rules'

@withRouter
@connect((state, prop) => ({
  promethusLink: state.promethus.link,
  strategyLink: state.strategy.link,
}))
export default class Rules extends Component {
  // search rule id
  searchId = undefined
  state = {
    dataSource: [],
    editModal: false,
    filterItem: {
      summary: false,
      description: false,
      prom_id: false,
      plan_id: false,
    },
  }
  paramsId = 0
  getColumnSearchProps = dataIndex => ({
    filterDropdown: ({ setSelectedKeys, selectedKeys, confirm }) => (
      <div style={{ padding: 8 }}>
        <Input
          ref={(node) => {
            this.searchInput = node;
          }}
          placeholder={`Search ${dataIndex}`}
          value={selectedKeys[0]}
          onInput={(e) => { setSelectedKeys(e.target.value ? [e.target.value] : []); this.handleSearch(selectedKeys, confirm, dataIndex) }}
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
          })); setTimeout(() => this.searchInput.focus());
        }}
        style={{ color: filtered ? '#1890ff' : undefined }}
      />
    ),
    onFilter: (value, record) => {
      let content
      // console.log(this.props.strategyLink[record[dataIndex]])
      switch (dataIndex) {
        // eslint-disable-next-line camelcase
        case 'prom_id': content = this.props.promethusLink ? (this.props.promethusLink[record[dataIndex]] || record[dataIndex]) : record[dataIndex]; break;
        // eslint-disable-next-line camelcase
        case 'plan_id': content = this.props.strategyLink ? (this.props.strategyLink[record[dataIndex]] || record[dataIndex]) : record[dataIndex]; break;
        default: content = record[dataIndex]
      }
      return content
        .toString()
        .toLowerCase()
        .includes(value.toLowerCase())
    },
    onFilterDropdownVisibleChange: (visible) => {
      if (visible) {
        setTimeout(() => this.searchInput.focus());
      }
    },
    render: text =>
      (this.state.searchedColumn === dataIndex ? (
        <Highlighter
          highlightStyle={{ backgroundColor: '#ffc069', padding: 0 }}
          searchWords={[this.state.searchText]}
          autoEscape
          textToHighlight={text.toString()}
        />
      ) : (
        text
      )),
  })
  handleSearch = (selectedKeys, confirm, dataIndex) => {
    confirm();
    this.setState({
      searchText: selectedKeys[0],
      searchedColumn: dataIndex,
    });
  }

  componentDidMount() {
    const { search } = this.props.location
    this.searchId = parseQueryString(search).id
    this.getList()
  }
  componentDidUpdate() {
    const { search } = this.props.location
    if (this.searchId !== parseQueryString(search).id) {
      this.searchId = parseQueryString(search).id
      this.getList()
    }
  }
  getList() {
    getRules({}, (res) => {
      this.setState({
        dataSource: res.sort((a, b) => b.id - a.id),
      });
    }, { params: { id: this.searchId } });
  }
  handleAdd = (e) => {
    this.setState({
      editModal: true,
    })
    this.createEditRules.updateValue()
  }
  handleEdit = (text, record) => {
    this.setState({
      editModal: true,
    })
    this.createEditRules.updateValue(record)
  }
  handleDelete = (record) => {
    const { id } = record
    deleteRules({}, { id }, (res) => {
      message.success(`删除 ${id} 成功`)
      this.getList()
    })
  }
  rulesUpdate = rule => new Promise((resolve) => {
    const { id, ...data } = rule
    if (id) {
      updateRules(data, { id }, (res) => {
        this.getList()
        this.setState({
          editModal: false,
        })
        resolve(true)
      })
      return
    }
    addRules(data, (res) => {
      resolve(true)
      this.getList()
      this.setState({
        editModal: false,
      })
    })
  })
  closeEditModal() {
    this.setState({
      editModal: false,
    })
  }
  onRefRule(component) {
    this.createEditRules = component
  }
  render() {
    const { editModal, dataSource } = this.state
    const columns = [
      { title: '编号', align: 'center', dataIndex: 'id' },
      {
        title: '表达式',
        align: 'center',
        width: '300px',
        render: (text, record) => (
          <code style={{ wordWrap: 'break-word' }}>{record.expr} {record.op} {record.value}</code>
        ),
      },
      { title: '持续时间', align: 'center', dataIndex: 'for' },
      {
        title: '标题',
        align: 'center',
        dataIndex: 'summary',
        ...this.getColumnSearchProps('summary'),
        filterDropdownVisible: this.state.filterItem.summary,
      },
      {
        title: '描述',
        align: 'center',
        dataIndex: 'description',
        ...this.getColumnSearchProps('description'),
        filterDropdownVisible: this.state.filterItem.description,
      },
      {
        title: '数据源',
        align: 'center',
        dataIndex: 'prom_id',
        ...this.getColumnSearchProps('prom_id'),
        filterDropdownVisible: this.state.filterItem.prom_id,
        render: promId => (
          <span>{this.props.promethusLink ? (this.props.promethusLink[promId] || promId) : promId}</span>
        ),
      },
      {
        title: '策略',
        align: 'center',
        dataIndex: 'plan_id',
        ...this.getColumnSearchProps('plan_id', true),
        filterDropdownVisible: this.state.filterItem.plan_id,
        render: planId => (
          <span>{this.props.strategyLink ? (this.props.strategyLink[planId] || planId) : planId}</span>
        ),
      },
      {
        title: '操作',
        align: 'center',
        key: 'action',
        width: '110px',
        render: (text, record, index) => (
          <span>
            <a onClick={() => this.handleEdit(text, record)}>编辑</a>
            <Popconfirm
              title="确定要删除吗?"
              onConfirm={() => this.handleDelete(record)}
              okText="Yes"
              cancelText="No"
            >
              <a href="#">删除</a>
            </Popconfirm>
          </span>
        ),
      },
    ]
    return (
      <div>
        <div id="top-section">
          <Button type="primary" onClick={this.handleAdd}>添加</Button>
        </div>
        <CreateEditRules onRef={c => this.onRefRule(c)} visiable={editModal} onClose={() => this.closeEditModal()} onSubmit={this.rulesUpdate} />
        <Table dataSource={dataSource} columns={columns} rowKey="id" />
      </div>
    )
  }
}
