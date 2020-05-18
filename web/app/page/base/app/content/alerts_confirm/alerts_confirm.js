import React, { Component } from 'react'
import { Modal, message } from 'antd'
import { Form, Input, InputNumber, Select, Icon, Row, Col, Checkbox, Button, AutoComplete } from 'antd';
import { formItemLayout, formItemLayoutWithOutLabel } from '@configs/const';

@Form.create({})
export default class AlertConfirmModal extends Component {
  state = {
    visible: false,
  }
  constructor(props) {
    super(props)
    this.props.OnRef(this)
  }
  updateValue() {
    const { form } = this.props
    this.setState({
      visible: true,
    })
    form.resetFields();
  }
  handleOk = (e) => {
    this.props.form.validateFields(async (err, values) => {
      if (!err) {
        const resultSuccess = await this.props.onSubmit(values)
        if (resultSuccess) {
          message.success('确认成功')
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
    const { visible } = this.state

    return (
      <Modal
        title="报警确认"
        visible={visible}
        onOk={this.handleOk}
        onCancel={this.handleCancel}
        maskClosable={false}
      >
        <Form {...formItemLayout} layout="horizontal">
          <Form.Item label="确认时长(分)">
            {getFieldDecorator('duration', {
              rules: [
                { required: true, message: '请输入时长' },
              ],
            })(<InputNumber style={{ width: '100%' }} />)}
          </Form.Item>
        </Form>
      </Modal>
    )
  }
}
