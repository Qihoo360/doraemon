import React, { Component } from 'react'
import { Button, Table, message, Popconfirm } from 'antd'
import { getMaintain, addMaintain, updateMaintain, deleteMaintain, getHost } from '@apis/maintain'
import moment from 'moment'
import CreateEditMaintain from './maintain/create-edit-maintain'
import HostList from './maintain/host-list'

export default class Maintain extends Component {
  state = {
    dataSource: [],
  }
  columns = [
    { title: '编号', align: 'center', dataIndex: 'id' },
    { title: '机器列表', align: 'center', render: text => (<a onClick={() => this.showHost(text.id)}>查看列表</a>) },
    { title: '开始时间', align: 'center', dataIndex: 'time_start' },
    { title: '结束时间', align: 'center', dataIndex: 'time_end' },
    { title: '开始日期', align: 'center', dataIndex: 'day_start' },
    { title: '结束日期', align: 'center', dataIndex: 'day_end' },
    { title: '维护月', align: 'center', dataIndex: 'month' },
    {
      title: '有效期',
      align: 'center',
      dataIndex: 'valid',
      render: valid => (
        <span>{moment(valid).format('YYYY.MM.DD HH:mm:ss')}</span>
      ),
    },
    {
      title: '操作',
      align: 'center',
      key: 'action',
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
  componentDidMount() {
    this.getList()
  }
  getList() {
    getMaintain({}, (res) => {
      this.setState({
        dataSource: res.sort((a, b) => b.id - a.id),
      });
    });
  }

  showHost = (id) => {
    getHost({}, { id }, (res) => {
      this.hostList.updateValue(res)
    })
  }
  handleAdd = (e) => {
    this.createEditMaintain.updateValue()
  }
  handleEdit = (text, record) => {
    const { id } = record
    getHost({}, { id }, (res) => {
      this.createEditMaintain.updateValue({ hosts: (res || []).join('\r\n'), ...record })
    })
  }
  handleDelete = (record) => {
    const { id } = record
    deleteMaintain({}, { id }, (res) => {
      message.success(`删除 ${id} 成功`)
      this.getList()
    })
  }
  rulesUpdate = rule => new Promise((resolve) => {
    const { id, ...data } = rule
    if (id) {
      updateMaintain(data, { id }, (res) => {
        this.getList()
        resolve(true)
      })
      return
    }
    addMaintain(data, (res) => {
      resolve(true)
      this.getList()
    })
  })
  onRef(component) {
    this.createEditMaintain = component
  }
  onHostRef(component) {
    this.hostList = component
  }
  render() {
    const { dataSource } = this.state;
    return (
      <div>
        <div id="top-section">
          <Button type="primary" onClick={this.handleAdd}>添加</Button>
        </div>
        <CreateEditMaintain onRef={c => this.onRef(c)} onSubmit={this.rulesUpdate} />
        <Table dataSource={dataSource} columns={this.columns} rowKey="id" />
        <HostList onRef={c => this.onHostRef(c)} />
      </div>
    )
  }
}
