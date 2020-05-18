import React, { Component } from 'react'
import { Modal, message, Form, Input } from 'antd'
import { formItemLayout } from '@configs/const'

@Form.create({})
export default class CreateEditUser extends Component {
  state = {
    // labels: [0],
    id: 0,
    visiable: false,
  }
  componentDidMount() {
    this.props.onRef(this)
  }
  componentWillReceiveProps() {

  }
  updateValue(value) {
    const { form } = this.props
    form.resetFields();
    this.setState({
      id: value ? value.id : 0,
      visiable: true,
    })
    value && form.setFieldsValue(value)
  }

  handleOk = (e) => {
    this.props.form.validateFields(async (err, values) => {
      console.log(values)
      if (!err) {
        const { id } = this.state;
        const resultSuccess = await this.props.onSubmit({ id, ...values })
        if (resultSuccess) {
          if (id) {
            message.success('修改成功')
          } else {
            message.success('添加成功')
          }
          this.setState({
            visiable: false,
          })
        }
      }
    })
  }
  handleCancel = (e) => {
    const { form } = this.props
    form.resetFields()
    this.setState({
      visiable: false,
    })
  }
  // addLabel = (e) => {
  //   const { labels } = this.state
  //   const last = labels.length - 1
  //   const nextKeys = labels.concat(last + 1)
  //   this.setState({
  //     labels: nextKeys,
  //   })
  // }
  // removeLabel = (k) => {
  //   const { labels } = this.state
  //   if (labels.length === 1) { return }
  //   this.setState({
  //     labels: labels.filter(key => key !== k),
  //   })
  // }
  render() {
    const { getFieldDecorator } = this.props.form
    const { id, visiable } = this.state
    return (
      <Modal
        title={id ? '编辑报警组' : '添加本地用户'}
        visible={visiable}
        onOk={this.handleOk}
        onCancel={this.handleCancel}
        maskClosable={false}
      >
        <Form {...formItemLayout} layout="horizontal">
          <Form.Item label="用户名">
            {getFieldDecorator('name', {
              rules: [
                { required: true, message: '请输入用户名' },
              ],
            })(<Input />)}
          </Form.Item>
          {/* <Form.Item label="成员">
            {getFieldDecorator('user', {
              rules: [
                { required: true, message: '请输入成员' },
              ],
            })(<DInput />)}
          </Form.Item> */}
        </Form>
        <p align="center" style={{ color: '#FF0000' }} >新用户默认密码：123456</p>
      </Modal>
    )
  }
}
