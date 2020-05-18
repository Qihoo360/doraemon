import React, { Component } from 'react';
import { Button, Table, message, Popconfirm } from 'antd';
import { getUser, addUser, updatePassword, deleteUser } from '@apis/user';
import { getUserName } from '@apis/login';
import CreateEditUser from './user/create-delete-user';
import ChangeUserPassword from './user/change-user-password';

export default class Rules extends Component {
  state = {
    name: "",
    display: "none",
    dataSource: [],
  }
  columns = [
    { title: '编号', align: 'center', dataIndex: 'id' },
    { title: '用户名', align: 'center', dataIndex: 'name' },
    {
      title: '操作',
      align: 'center',
      key: 'action',
      render: (text, record, index) => (
        <span>
          {/* <a onClick={() => this.handleEdit(text, record)}>编辑</a> */}
          <Popconfirm
            title="确定要删除吗?"
            onConfirm={() => this.handleDelete(record)}
            okText="Yes"
            cancelText="No"
          >
            {/* {console.log(record)} */}
            {record.name != "admin" ? <a href="#">删除</a> : null}
          </Popconfirm>
        </span>
      ),
    },
  ];
  componentWillMount() {
    getUserName({}, (res) => {
      if (res == "admin") {
        this.setState({
          name: res,
          display: "block",
        });
      } else {
        this.setState({
          name: res,
        });
      }
    });
  }
  componentDidMount() {
    this.getList()
  }
  getList() {
    getUser({}, (res) => {
      this.setState({
        dataSource: res,
      });
    });
  }
  handleAdd = (e) => {
    this.createEditUser.updateValue()
  }
  handleChangePassword = (e) => {
    this.changeUserPassword.updateValue()
  }
  // handleEdit = (text, record) => {
  //   this.createEditUser.updateValue(record)
  // }
  handleDelete = (record) => {
    const { id,name } = record
    deleteUser({}, { id }, (res) => {
      message.success(`删除 ${name} 成功`)
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
    addUser(data, (res) => {
      resolve(true)
      this.getList()
    })
  })

  passwordUpdate = rule => new Promise((resolve) => {
    const { id, ...data } = rule
    data.name = this.state.name
    console.log(id, data)
    // if (id) {
    //   updateGroup(data, { id }, (res) => {
    //     this.getList()
    //     resolve(true)
    //   })
    //   return
    // }
    updatePassword(data, (res) => {
      resolve(true)
    })
  })

  onRef(component) {
    this.createEditUser = component
  }
  onRef2(component) {
    this.changeUserPassword = component
  }
  render() {
    const { dataSource } = this.state;
    return (
      <div>
        <div id="top-section">
          {
            this.state.display == "block" ? <Button type="primary" onClick={this.handleAdd}>添加用户</Button> : null
          }
          <Button type="primary" onClick={this.handleChangePassword}>修改密码</Button>
        </div>
        <div style={{ display: this.state.display }}>
          <CreateEditUser onRef={c => this.onRef(c)} onSubmit={this.rulesUpdate} />
          <ChangeUserPassword onRef={c => this.onRef2(c)} onSubmit={this.passwordUpdate} />
          <Table dataSource={dataSource} columns={this.columns} rowKey="id" />
        </div>
      </div>
    )
  }
}
