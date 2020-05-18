import React, { Component } from 'react'
import { Modal, message } from 'antd'
import { Form, Input, InputNumber, Select, Icon, Row, Col, Checkbox, Button, AutoComplete } from 'antd';
import { formItemLayout, formItemLayoutWithOutLabel } from '@configs/const';
import { promethus } from '@actions/common'
import { connect } from 'react-redux'


@connect(() => ({}), dispatch => ({
  promethusAction: () => dispatch(promethus()),
}))
@Form.create({})
export default class CreateEditPromethus extends Component {
  state = {
    id: 0,
    visible: false,
  }
  constructor(props) {
    super(props)
    this.props.OnRef(this)
  }
  updateValue(value) {
    const { form } = this.props
    this.setState({
      id: value ? value.id : 0,
      visible: true,
    })
    if (!value) {
      form.resetFields();
      return
    }
    form.setFieldsValue(value)
  }
  handleOk = (e) => {
    this.props.form.validateFields(async (err, values) => {
      if (!err) {
        const { id } = this.state
        const resultSuccess = await this.props.onSubmit({ id, ...values })
        const { promethusAction } = this.props
        promethusAction()
        if (resultSuccess) {
          message.success(id ? '修改成功' : '添加成功')
          this.setState({
            visible: false,
          })
        }
      }
    })
  }
  handleCancel = (e) => {
    this.setState({
      visible: false,
    })
  }
  render() {
    const { getFieldDecorator } = this.props.form
    const { id, visible } = this.state

    return (
      <Modal
        title={id ? '修改数据源' : '添加数据源'}
        visible={visible}
        onOk={this.handleOk}
        onCancel={this.handleCancel}
        maskClosable={false}
      >
        <Form {...formItemLayout} layout="horizontal">
          <Form.Item label="名称">
            {getFieldDecorator('name', {
              rules: [
                { required: true, message: '请输入名称' },
              ],
            })(<Input />)}
          </Form.Item>
          <Form.Item label="URL">
            {getFieldDecorator('url', {
              rules: [
                { required: true, message: '请输入URL' },
              ],
            })(<Input />)}
          </Form.Item>
        </Form>
      </Modal>
    )
  }
}
