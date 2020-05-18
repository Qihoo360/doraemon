import React, { Component } from 'react';
import { Button, Table, message, Popconfirm } from 'antd';
import { getGroup, addGroup, updateGroup, deleteGroup } from '@apis/group';
import CreateEditGroup from './group/create-edit-group';

export default class Rules extends Component {
  state = {
    dataSource: [],
  }
  columns = [
    { title: '编号', align: 'center', dataIndex: 'id' },
    { title: '组名', align: 'center', dataIndex: 'name' },
    { title: '成员', align: 'center', dataIndex: 'user' },
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
  ];
  componentDidMount() {
    this.getList()
  }
  getList() {
    getGroup({}, (res) => {
      this.setState({
        dataSource: res,
      });
    });
  }
  handleAdd = (e) => {
    this.createEditGroup.updateValue()
  }
  handleEdit = (text, record) => {
    this.createEditGroup.updateValue(record)
  }
  handleDelete = (record) => {
    const { id } = record
    deleteGroup({}, { id }, (res) => {
      message.success(`删除 ${id} 成功`)
      this.getList()
    })
  }
  rulesUpdate = rule => new Promise((resolve) => {
    const { id, ...data } = rule
    if (id) {
      updateGroup(data, { id }, (res) => {
        this.getList()
        resolve(true)
      })
      return
    }
    addGroup(data, (res) => {
      resolve(true)
      this.getList()
    })
  })
  onRef(component) {
    this.createEditGroup = component
  }
  render() {
    const { dataSource } = this.state;
    return (
      <div>
        <div id="top-section">
          <Button type="primary" onClick={this.handleAdd}>添加</Button>
        </div>
        <CreateEditGroup onRef={c => this.onRef(c)} onSubmit={this.rulesUpdate} />
        <Table dataSource={dataSource} columns={this.columns} rowKey="id" />
      </div>
    )
  }
}
