import React, { Component } from 'react'
import { Button, Table, message, Popconfirm } from 'antd'
import { getPromethus, addPromethus, updatePromethus, deletePromethus } from '@apis/promethus'
import CreateEditPromethus from './promethus/create-edit-promethus'

export default class Promethus extends Component {
  state = {
    dataSource: [],
  }
  componentDidMount() {
    this.getList()
  }
  getList() {
    getPromethus({}, (res) => {
      this.setState({
        dataSource: res,
      })
    })
  }
  columns = [
    { title: '编号', align: 'center', dataIndex: 'id', key: 'id' },
    { title: '名称', align: 'center', dataIndex: 'name', key: 'name' },
    { title: 'URL', align: 'center', dataIndex: 'url', key: 'url' },
    {
      title: '操作',
      align: 'center',
      key: 'action',
      render: (text, record, index) => (
        <span>
          <a onClick={() => this.handleEdit(text, record)}>编辑</a>
          <Popconfirm
            title="确定要删除吗?"
            onConfirm={() => this.confirmDelete(record)}
            okText="Yes"
            cancelText="No"
          >
            <a href="#">删除</a>
          </Popconfirm>
        </span>
      ),
    },
  ]
  confirmDelete(record) {
    const { id, name } = record
    deletePromethus({}, { id }, (res) => {
      message.success(`删除 ${name} 成功`)
      this.getList()
    })
  }
  handleAdd = (e) => {
    this.createEditPromethus.updateValue()
  }
  handleEdit(record) {
    this.createEditPromethus.updateValue(record)
  }
  onRefRro(component) {
    this.createEditPromethus = component
  }
  updatePromethus = values => new Promise((resolve) => {
    const { id, ...data } = values
    if (id) {
      // edit
      updatePromethus(data, { id }, (res) => {
        this.getList()
        resolve(true)
      })
      return
    }
    addPromethus(data, (res) => {
      this.getList()
      resolve(true)
    })
  })
  render() {
    const { dataSource } = this.state
    return (
      <div>
        <div id="top-section">
          <Button type="primary" onClick={this.handleAdd}>添加</Button>
        </div>
        <Table dataSource={dataSource} columns={this.columns} rowKey="id" />
        <CreateEditPromethus OnRef={c => this.onRefRro(c)} onSubmit={this.updatePromethus} />
      </div>
    )
  }
}
