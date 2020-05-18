import React, { Component } from 'react'
import { Modal, message, Form, Input, TimePicker, InputNumber, DatePicker, Select } from 'antd'
import { formItemLayout } from '@configs/const'
import moment from 'moment'

const { Option } = Select

@Form.create({})
export default class CreateEditMaintain extends Component {
  state = {
    labels: [0],
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
    value && form.setFieldsValue(this.unFormatValue(value))
  }
  dayEndValid = (rule, value, callback) => {
    const { form } = this.props
    const dayStart = form.getFieldValue('day_start')
    if (value && value < dayStart) {
      return callback('结束时间要晚于开始时间')
    }
    return callback()
  }
  formatValue(values) {
    const { ...value } = values;
    Object.keys(value).forEach((key) => {
      if (key === 'time_end' || key === 'time_start') {
        value[key] = moment(value[key]).format('HH:mm')
      }
      if (key === 'valid') {
        value[key] = moment(value[key]).format('YYYY-MM-DD HH:mm:ss')
      }
      if (key === 'month') {
        value[key] = (value[key] || []).join(',')
      }
    })
    return value
  }
  unFormatValue(values) {
    const { ...value } = values;
    Object.keys(value).forEach((key) => {
      if (key === 'time_end' || key === 'time_start') {
        value[key] = moment(value[key], 'HH:mm')
      }
      if (key === 'valid') {
        value[key] = moment(value[key], 'YYYY-MM-DD HH:mm:ss')
      }
      if (key === 'month') {
        value[key] = value[key].trim() === '' ? [] : value[key].split(',')
      }
    })
    return value
  }
  handleOk = (e) => {
    this.props.form.validateFields(async (err, values) => {
      console.log(values)
      if (!err) {
        const { id } = this.state;
        const resultSuccess = await this.props.onSubmit({ id, ...this.formatValue(values) })
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
  addLabel = (e) => {
    const { labels } = this.state
    const last = labels.length - 1
    const nextKeys = labels.concat(last + 1)
    this.setState({
      labels: nextKeys,
    })
  }
  removeLabel = (k) => {
    const { labels } = this.state
    if (labels.length === 1) { return }
    this.setState({
      labels: labels.filter(key => key !== k),
    })
  }
  render() {
    const { getFieldDecorator } = this.props.form
    const { id, visiable } = this.state
    return (
      <Modal
        title={id ? '编辑维护组' : '添加维护组'}
        visible={visiable}
        onOk={this.handleOk}
        onCancel={this.handleCancel}
        maskClosable={false}
      >
        <Form {...formItemLayout} layout="horizontal">
          <Form.Item label="维护时间段" required style={{ marginBottom: 0 }}>
            <Form.Item style={{ display: 'inline-block', width: 'calc(50% - 10px)' }}>
              {getFieldDecorator('time_start', {
                rules: [{ type: 'object', required: true, message: 'Please select time!' }],
              })(<TimePicker style={{ width: '100%' }} format="HH:mm" />)}
            </Form.Item>
            <span style={{ display: 'inline-block', width: '20px', textAlign: 'center' }}>~</span>
            <Form.Item style={{ display: 'inline-block', width: 'calc(50% - 10px)' }}>
              {getFieldDecorator('time_end', {
                rules: [
                  { type: 'object', required: true, message: 'Please select time!' },
                ],
              })(<TimePicker style={{ width: '100%' }} format="HH:mm" />)}
            </Form.Item>
          </Form.Item>
          <Form.Item label="维护月" required>
            {getFieldDecorator('month', {
              rules: [
                { type: 'array', required: true, message: 'Please select month' },
              ]
            })(<Select mode="multiple">
              { new Array(12).fill(1).map((item, index) => (<Option value={index + 1}>{index + 1}</Option>))}
            </Select>)}
          </Form.Item>
          <Form.Item label="维护日期" required style={{ marginBottom: 0 }}>
            <Form.Item style={{ display: 'inline-block', width: 'calc(50% - 10px)' }}>
              {getFieldDecorator('day_start', {
                rules: [{ required: true, message: 'Please input day!' }],
              })(<InputNumber style={{ width: '100%' }} format="HH:mm" />)}
            </Form.Item>
            <span style={{ display: 'inline-block', width: '20px', textAlign: 'center' }}>~</span>
            <Form.Item style={{ display: 'inline-block', width: 'calc(50% - 10px)' }}>
              {getFieldDecorator('day_end', {
                rules: [
                  { required: true, message: 'Please input day!' },
                  { validator: this.dayEndValid },
                ],
              })(<InputNumber style={{ width: '100%' }} format="HH:mm" />)}
            </Form.Item>
          </Form.Item>
          <Form.Item label="有效期">
            {getFieldDecorator('valid', {
              rules: [
                { required: true, message: '请填写有效期' },
              ],
            })(<DatePicker
              style={{ width: '100%' }}
              showTime={{ defaultValue: moment('00:00:00', 'HH:mm:ss') }}
              format="YYYY-MM-DD HH:mm:ss"
              placeholder={['请填写有效期']}
            />)}
          </Form.Item>
          <Form.Item label="机器列表">
            {getFieldDecorator('hosts', {
              rules: [
                { required: true, message: '请输入成员' },
              ],
            })(<Input.TextArea autoSize={{ minRows: 2 }} />)}
          </Form.Item>
        </Form>
      </Modal>
    )
  }
}
