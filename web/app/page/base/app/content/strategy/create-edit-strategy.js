import React, { Component } from 'react'
import { Modal, message, Form, Input } from 'antd'
import { formItemLayout } from '@configs/const'
import { strategy } from '@actions/common'
import { connect } from 'react-redux'


@connect(() => ({}), dispatch => ({
  strategyAction: () => dispatch(strategy()),
}))
@Form.create({})
export default class CreateEditStrategy extends Component {
  constructor(props) {
    super(props)
    this.props.OnRef(this)
  }
  state = {
    id: 0,
    visible: false,
  }
  updateValue(value) {
    const { form } = this.props
    this.setState({
      id: value ? value.id : 0,
      visible: true,
    })
    form.resetFields()
    value && form.setFieldsValue(value)
  }
  handleOk = () => {
    this.props.form.validateFields(async (err, values) => {
      if (!err) {
        const { id } = this.state
        const { strategyAction } = this.props
        const resultSuccess = await this.props.onSubmit({ id, ...values })
        strategyAction()
        if (resultSuccess) {
          message.success(id ? '编辑成功' : '添加成功')
          this.setState({
            visible: false,
          })
        }
      }
    })
  }
  handleCancel = () => {
    const { form } = this.props
    this.setState({
      visible: false,
    })
    form.resetFields()
  }
  render() {
    const { getFieldDecorator } = this.props.form
    const { id, visible } = this.state

    return (
      <Modal
        title={id ? '编辑报警计划' : '添加报警计划'}
        visible={visible}
        onOk={this.handleOk}
        onCancel={this.handleCancel}
        maskClosable={false}
      >
        <Form {...formItemLayout} layout="horizontal">
          <Form.Item label="名称">
            {getFieldDecorator('description', {
              rules: [
                { required: true, message: '请输入报警计划名称' },
              ],
            })(<Input />)}
          </Form.Item>
          <Form.Item label="描述">
            {getFieldDecorator('rule_labels', {})(<Input />)}
          </Form.Item>
        </Form>
      </Modal>
    )
  }
}
